//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package scaling

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/semi-technologies/weaviate/entities/backup"
	"github.com/semi-technologies/weaviate/usecases/sharding"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type ScaleOutManager struct {
	// the scaleOutManager needs to read and updated the sharding state of a
	// class. It can access it through the schemaManager
	schemaManager SchemaManager

	// get information about which nodes are in the cluster
	clusterState clusterState

	backerUpper BackerUpper

	nodes nodeClient

	persistenceRoot string

	logger logrus.FieldLogger
}

type clusterState interface {
	// AllNames() returns all the node names (not the hostnames!) including the
	// local one
	AllNames() []string
	LocalName() string
	NodeHostname(nodename string) (string, bool)
}

type BackerUpper interface {
	SingleShardBackup(
		ctx context.Context, bakID, className, shardName string,
	) (backup.ClassDescriptor, error)

	ShardsBackup(
		ctx context.Context, bakID, class string, shards []string,
	) (_ backup.ClassDescriptor, err error)
	ReleaseBackup(ctx context.Context, bakID, className string) error
}

func NewScaleOutManager(clusterState clusterState, backerUpper BackerUpper,
	nodeClient nodeClient, logger logrus.FieldLogger, persistenceRoot string,
) *ScaleOutManager {
	return &ScaleOutManager{
		clusterState:    clusterState,
		backerUpper:     backerUpper,
		nodes:           nodeClient,
		persistenceRoot: persistenceRoot,
	}
}

type SchemaManager interface {
	ShardingState(class string) *sharding.State
}

func (som *ScaleOutManager) SetSchemaManager(sm SchemaManager) {
	som.schemaManager = sm
}

// Scale returns the updated sharding state if successful. The caller must then
// make sure to broadcast that state to all nodes as part of the "update"
// transaction.
func (som *ScaleOutManager) Scale(ctx context.Context, className string,
	updated sharding.Config, prevReplFactor, newReplFactor int64,
) (*sharding.State, error) {
	// First identify what the sharding state was before this change. This is
	// mainly to be able to compare the diff later, so we know where we need to
	// make changes
	ssBefore := som.schemaManager.ShardingState(className)
	if ssBefore == nil {
		return nil, errors.Errorf("no sharding state for class %q", className)
	}
	if newReplFactor > prevReplFactor {
		return som.scaleOut(ctx, className, ssBefore, updated, newReplFactor)
	}

	if newReplFactor < prevReplFactor {
		return som.scaleIn(ctx, className, updated)
	}

	return nil, nil
}

// scaleOut is a relatively primitive implementation that takes a shard and
// copies it onto another node. Then it returns the new and updated sharding
// state where the node-association is updated to point to all nodes where the
// shard can be found.
//
// This implementation is meant as a temporary one and should probably be
// replaced with something more sophisticated. The main issues are:
//
//   - Everything is synchronous. This blocks the users request until all data
//     is copied which is not great.
//   - Everything is sequential. A lot of the things could probably happen in
//     parallel
//
// Follow the in-line comments to see how this implementation achieves scalign
// out
func (som *ScaleOutManager) scaleOut(ctx context.Context, className string, ssBefore *sharding.State,
	updated sharding.Config, replFactor int64,
) (*sharding.State, error) {
	// Create a deep copy of the old sharding state, so we can start building the
	// updated state. Because this is a deep copy we don't risk leaking our
	// changes to anyone else. We can return the changes in the end where the
	// caller can then make sure to broadcast the new state to the cluster.
	ssAfter := ssBefore.DeepCopy()
	ssAfter.Config = updated

	// Identify all shards of the class and adjust the replicas. After this is
	// done, the affected shards now belong to more nodes than they did before.
	for name, shard := range ssAfter.Physical {
		shard.AdjustReplicas(int(replFactor), som.clusterState)
		ssAfter.Physical[name] = shard
	}
	// resolve hosts beforehand
	localShards := make([]string, 0, len(ssBefore.Physical))
	remoteShards := make(map[string]string)
	for name := range ssBefore.Physical {
		if ssBefore.IsShardLocal(name) {
			localShards = append(localShards, name)
		} else {
			belongsTo := ssBefore.Physical[name].BelongsToNode()
			host, ok := som.clusterState.NodeHostname(belongsTo)
			if !ok {
				return nil, fmt.Errorf("cannot resolve hostname for node %q", belongsTo)
			}
			remoteShards[name] = host
		}
	}
	// However, so far we have only updated config, now we also need to actually
	// copy files.
	var g errgroup.Group
	for shard, host := range remoteShards {
		shard, host := shard, host
		g.Go(func() error {
			err := som.nodes.IncreaseReplicationFactor(ctx, host, className, ssBefore, &ssAfter)
			if err != nil {
				belongsTo := ssBefore.Physical[shard].BelongsToNode()
				return fmt.Errorf("increase replication factor for class %q on node %q: %w", className, belongsTo, err)
			}
			return nil
		})
	}
	g.Go(func() error { // TODO: get all shards backups otherwise the second goroutine will be just waiting
		if err := som.localScaleOut(ctx, className, localShards, ssBefore, &ssAfter); err != nil {
			return fmt.Errorf("increase local replication factor: %w", err)
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		return nil, err
	}

	// Finally, return sharding state back to schema manager. The schema manager
	// will then broadcast this updated state to the cluster. This is essentially
	// what will take the new replication shards live: On the new nodes, if
	// traffic is incoming, IsShardLocal() would have returned false before. But
	// now that a copy of the local shard is present it will return true and
	// serve the traffic.
	return &ssAfter, nil
}

func (som *ScaleOutManager) LocalScaleOut(ctx context.Context,
	className string, ssBefore, ssAfter *sharding.State,
) error {
	// Next, we will iterate over every shard that this class has. This is the
	// meat&bones of this implemenation. For each shard, we're roughly doing the
	// following:
	// - Create a single shard backup, so the shard is safe to copy
	// - Figure out the copy targets (i.e. each node that is part of the after
	//   state, but wasn't part of the before state yet)
	// - Create an empty shard on the target node
	// - Copy over all files from the backup
	// - Reinit the shard to recognize the copied files
	// - Release the single-shard backup
	ssBefore.SetLocalName(som.clusterState.LocalName())
	ssAfter.SetLocalName(som.clusterState.LocalName())
	localShards := make([]string, 0, len(ssBefore.Physical))
	for shardName := range ssBefore.Physical {
		if ssBefore.IsShardLocal(shardName) {
			localShards = append(localShards, shardName)
		}
	}
	if err := som.localScaleOut(ctx, className, localShards, ssBefore, ssAfter); err != nil {
		return err
	}

	return nil
}

func (som *ScaleOutManager) localScaleOut(ctx context.Context,
	className string, shards []string, ssBefore, ssAfter *sharding.State,
) error {
	if len(shards) < 1 {
		return nil
	}
	// Create backup of the single shard
	bakID := fmt.Sprintf("_internal_scaleout_%s", uuid.New().String()) // todo better name
	bak, err := som.backerUpper.ShardsBackup(ctx, bakID, className, shards)
	if err != nil {
		return errors.Wrap(err, "create snapshot")
	}

	defer func() {
		err := som.backerUpper.ReleaseBackup(context.Background(), bakID, className)
		if err != nil {
			som.logger.WithField("scaler", "releaseBackup").WithField("class", className).Error(err)
		}
	}()

	for _, desc := range bak.Shards {
		shardName := desc.Name
		// Figure out which nodes are new by diffing the before and after state
		// TODO: This manual diffing is ugly, refactor!
		newNodes := ssAfter.Physical[shardName].BelongsToNodes
		previousNodes := ssBefore.Physical[shardName].BelongsToNodes
		// This relies on the convention that new nodes are always appended at the end
		additions := newNodes[len(previousNodes):]

		// Iterate over the new target nodes and copy files
		for _, targetNode := range additions {
			// bakShard := bak.Shards[0]
			// if bakShard.Name != shardName {
			// 	// this sanity check is only needed because of the Shards[0] above. If this
			// 	// supports multi-shard, we need a better logic anyway
			// 	return fmt.Errorf("shard name mismatch in backup: %q vs %q", bakShard.Name,
			// 		shardName)
			// }

			// Create an empty shard on the remote node. This is a requirement to
			// copy files in the next step. If the empty shard didn't exist, we'd
			// have no copy target which could receive the files.
			//
			if err := som.CreateShard(ctx, targetNode, className, shardName); err != nil {
				return fmt.Errorf("create new shard on remote node: %w", err)
			}

			// Transfer each file that's part of the backup.
			for _, file := range desc.Files {
				err := som.PutFile(ctx, file, targetNode, className, shardName)
				if err != nil {
					return fmt.Errorf("copy files to remote node: %w", err)
				}
			}

			// Transfer shard metadata files
			err := som.PutFile(ctx, desc.ShardVersionPath, targetNode, className, shardName)
			if err != nil {
				return fmt.Errorf("copy shard version to remote node: %w", err)
			}

			err = som.PutFile(ctx, desc.DocIDCounterPath, targetNode, className, shardName)
			if err != nil {
				return fmt.Errorf("copy index counter to remote node: %w", err)
			}

			err = som.PutFile(ctx, desc.PropLengthTrackerPath, targetNode, className, shardName)
			if err != nil {
				return fmt.Errorf("copy prop length tracker to remote node: %w", err)
			}

			// Now that all files are on the remote node's new shard, the shard needs
			// to be reinitialized. Otherwise, it would not recognize the files when
			// serving traffic later.
			if err := som.ReInitShard(ctx, targetNode, className, shardName); err != nil {
				return fmt.Errorf("create new shard on remote node: %w", err)
			}

			// Clean up after ourselves and prevent blocking future backups.
			if err := som.backerUpper.ReleaseBackup(ctx, bakID, className); err != nil {
				return fmt.Errorf("release shard backup: %w", err)
			}
		}
	}

	return nil
}

func (som *ScaleOutManager) scaleIn(ctx context.Context, className string,
	updated sharding.Config,
) (*sharding.State, error) {
	return nil, errors.Errorf("scaling in (reducing replica count) not supported yet")
}

func (som *ScaleOutManager) PutFile(ctx context.Context, sourceFileName string,
	targetNode, className, shardName string,
) error {
	absPath := filepath.Join(som.persistenceRoot, sourceFileName)

	hostname, ok := som.clusterState.NodeHostname(targetNode)
	if !ok {
		return fmt.Errorf("resolve hostname for node %q", targetNode)
	}

	f, err := os.Open(absPath)
	if err != nil {
		return fmt.Errorf("open file %q for reading: %w", absPath, err)
	}

	return som.nodes.PutFile(ctx, hostname, className, shardName, sourceFileName, f)
}

func (som *ScaleOutManager) CreateShard(ctx context.Context,
	targetNode, className, shardName string,
) error {
	hostname, ok := som.clusterState.NodeHostname(targetNode)
	if !ok {
		return fmt.Errorf("resolve hostname for node %q", targetNode)
	}

	return som.nodes.CreateShard(ctx, hostname, className, shardName)
}

func (som *ScaleOutManager) ReInitShard(ctx context.Context,
	targetNode, className, shardName string,
) error {
	hostname, ok := som.clusterState.NodeHostname(targetNode)
	if !ok {
		return fmt.Errorf("resolve hostname for node %q", targetNode)
	}

	return som.nodes.ReInitShard(ctx, hostname, className, shardName)
}

type nodeClient interface {
	PutFile(ctx context.Context, hostName, indexName,
		shardName, fileName string, payload io.ReadSeekCloser) error
	CreateShard(ctx context.Context,
		hostName, indexName, shardName string) error
	ReInitShard(ctx context.Context,
		hostName, indexName, shardName string) error
	IncreaseReplicationFactor(ctx context.Context, hostName, indexName string,
		ssBefore, ssAfter *sharding.State) error
}
