//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package schema

import (
	"context"
	"fmt"
	"strings"

	"github.com/weaviate/weaviate/entities/schema"

	"github.com/pkg/errors"
	"github.com/weaviate/weaviate/entities/models"
)

// AddClassProperty to an existing Class
func (m *Manager) AddClassProperty(ctx context.Context, principal *models.Principal,
	class string, property *models.Property,
) error {
	err := m.Authorizer.Authorize(principal, "update", "schema/objects")
	if err != nil {
		return err
	}

	return m.addClassProperty(ctx, class, property)
}

func (m *Manager) addClassProperty(ctx context.Context,
	className string, prop *models.Property,
) error {
	m.Lock()
	defer m.Unlock()

	class, err := schema.GetClassByName(m.state.ObjectSchema, className)
	if err != nil {
		return err
	}
	prop.Name = schema.LowerCaseFirstLetter(prop.Name)

	if err := m.setNewPropDefaults(class, prop); err != nil {
		return err
	}

	existingPropertyNames := map[string]bool{}
	for _, existingProperty := range class.Properties {
		existingPropertyNames[strings.ToLower(existingProperty.Name)] = true
	}
	if err := m.validateProperty(prop, className, existingPropertyNames, false); err != nil {
		return err
	}

	tx, err := m.cluster.BeginTransaction(ctx, AddProperty,
		AddPropertyPayload{className, prop}, DefaultTxTTL)
	if err != nil {
		// possible causes for errors could be nodes down (we expect every node to
		// the up for a schema transaction) or concurrent transactions from other
		// nodes
		return errors.Wrap(err, "open cluster-wide transaction")
	}

	if err := m.cluster.CommitWriteTransaction(ctx, tx); err != nil {
		// Only log the commit error, but do not abort the changes locally. Once
		// we've told others to commit, we also need to commit ourselves!
		//
		// The idea is that if we abort our changes we are guaranteed to create an
		// inconsistency as soon as any other node honored the commit. This would
		// for example be the case in a 3-node cluster where node 1 is the
		// coordinator, node 2 honored the commit and node 3 died during the commit
		// phase.
		//
		// In this scenario it is far more desirable to make sure that node 1 and
		// node 2 stay in sync, as node 3 - who may or may not have missed the
		// update - can use a local WAL from the first TX phase to replay any
		// missing changes once it's back.
		m.logger.WithError(err).Errorf("not every node was able to commit")
	}

	return m.addClassPropertyApplyChanges(ctx, className, prop)
}

func (m *Manager) setNewPropDefaults(class *models.Class, prop *models.Property) error {
	m.setPropertyDefaults(prop)
	if err := validateUserProp(class, prop); err != nil {
		return err
	}
	m.moduleConfig.SetSinglePropertyDefaults(class, prop)
	return nil
}

func validateUserProp(class *models.Class, prop *models.Property) error {
	if prop.ModuleConfig == nil {
		return nil
	} else {
		modconfig, ok := prop.ModuleConfig.(map[string]interface{})
		if !ok {
			return fmt.Errorf("%v property config invalid", prop.Name)
		}
		vectorizerConfig, ok := modconfig[class.Vectorizer]
		if !ok {
			return fmt.Errorf("%v vectorizer module not part of the property", class.Vectorizer)
		}
		_, ok = vectorizerConfig.(map[string]interface{})
		if !ok {
			return fmt.Errorf("vectorizer config for vectorizer %v, not of type map[string]interface{}", class.Vectorizer)
		}
	}
	return nil
}

func (m *Manager) addClassPropertyApplyChanges(ctx context.Context,
	className string, prop *models.Property,
) error {
	class, err := schema.GetClassByName(m.state.ObjectSchema, className)
	if err != nil {
		return err
	}

	class.Properties = append(class.Properties, prop)
	err = m.saveSchema(ctx)
	if err != nil {
		return err
	}

	return m.migrator.AddProperty(ctx, className, prop)
}
