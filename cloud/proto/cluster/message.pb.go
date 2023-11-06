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

package cluster

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Command_Type int32

const (
	Command_TYPE_UNSPECIFIED         Command_Type = 0
	Command_TYPE_ADD_CLASS           Command_Type = 1
	Command_TYPE_UPDATE_CLASS        Command_Type = 2
	Command_TYPE_DELETE_CLASS        Command_Type = 3
	Command_TYPE_RESTORE_CLASS       Command_Type = 4
	Command_TYPE_ADD_PROPERTY        Command_Type = 5
	Command_TYPE_UPDATE_SHARD_STATUS Command_Type = 10
	Command_TYPE_ADD_TENANT          Command_Type = 16
	Command_TYPE_UPDATE_TENANT       Command_Type = 17
	Command_TYPE_DELETE_TENANT       Command_Type = 18
)

// Enum value maps for Command_Type.
var (
	Command_Type_name = map[int32]string{
		0:  "TYPE_UNSPECIFIED",
		1:  "TYPE_ADD_CLASS",
		2:  "TYPE_UPDATE_CLASS",
		3:  "TYPE_DELETE_CLASS",
		4:  "TYPE_RESTORE_CLASS",
		5:  "TYPE_ADD_PROPERTY",
		10: "TYPE_UPDATE_SHARD_STATUS",
		16: "TYPE_ADD_TENANT",
		17: "TYPE_UPDATE_TENANT",
		18: "TYPE_DELETE_TENANT",
	}
	Command_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED":         0,
		"TYPE_ADD_CLASS":           1,
		"TYPE_UPDATE_CLASS":        2,
		"TYPE_DELETE_CLASS":        3,
		"TYPE_RESTORE_CLASS":       4,
		"TYPE_ADD_PROPERTY":        5,
		"TYPE_UPDATE_SHARD_STATUS": 10,
		"TYPE_ADD_TENANT":          16,
		"TYPE_UPDATE_TENANT":       17,
		"TYPE_DELETE_TENANT":       18,
	}
)

func (x Command_Type) Enum() *Command_Type {
	p := new(Command_Type)
	*p = x
	return p
}

func (x Command_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Command_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_cluster_message_proto_enumTypes[0].Descriptor()
}

func (Command_Type) Type() protoreflect.EnumType {
	return &file_cluster_message_proto_enumTypes[0]
}

func (x Command_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Command_Type.Descriptor instead.
func (Command_Type) EnumDescriptor() ([]byte, []int) {
	return file_cluster_message_proto_rawDescGZIP(), []int{4, 0}
}

type JoinPeerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Voter   bool   `protobuf:"varint,3,opt,name=voter,proto3" json:"voter,omitempty"`
}

func (x *JoinPeerRequest) Reset() {
	*x = JoinPeerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinPeerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinPeerRequest) ProtoMessage() {}

func (x *JoinPeerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinPeerRequest.ProtoReflect.Descriptor instead.
func (*JoinPeerRequest) Descriptor() ([]byte, []int) {
	return file_cluster_message_proto_rawDescGZIP(), []int{0}
}

func (x *JoinPeerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *JoinPeerRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *JoinPeerRequest) GetVoter() bool {
	if x != nil {
		return x.Voter
	}
	return false
}

type RemovePeerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemovePeerRequest) Reset() {
	*x = RemovePeerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePeerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePeerRequest) ProtoMessage() {}

func (x *RemovePeerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePeerRequest.ProtoReflect.Descriptor instead.
func (*RemovePeerRequest) Descriptor() ([]byte, []int) {
	return file_cluster_message_proto_rawDescGZIP(), []int{1}
}

func (x *RemovePeerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RemovePeerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemovePeerResponse) Reset() {
	*x = RemovePeerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePeerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePeerResponse) ProtoMessage() {}

func (x *RemovePeerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePeerResponse.ProtoReflect.Descriptor instead.
func (*RemovePeerResponse) Descriptor() ([]byte, []int) {
	return file_cluster_message_proto_rawDescGZIP(), []int{2}
}

type JoinPeerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *JoinPeerResponse) Reset() {
	*x = JoinPeerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinPeerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinPeerResponse) ProtoMessage() {}

func (x *JoinPeerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinPeerResponse.ProtoReflect.Descriptor instead.
func (*JoinPeerResponse) Descriptor() ([]byte, []int) {
	return file_cluster_message_proto_rawDescGZIP(), []int{3}
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       Command_Type `protobuf:"varint,1,opt,name=type,proto3,enum=weaviate.cloud.internal.cluster.Command_Type" json:"type,omitempty"`
	Class      string       `protobuf:"bytes,2,opt,name=class,proto3" json:"class,omitempty"`
	SubCommand []byte       `protobuf:"bytes,3,opt,name=sub_command,json=subCommand,proto3" json:"sub_command,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_cluster_message_proto_rawDescGZIP(), []int{4}
}

func (x *Command) GetType() Command_Type {
	if x != nil {
		return x.Type
	}
	return Command_TYPE_UNSPECIFIED
}

func (x *Command) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

func (x *Command) GetSubCommand() []byte {
	if x != nil {
		return x.SubCommand
	}
	return nil
}

type AddTenantsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenants []*Tenant `protobuf:"bytes,1,rep,name=tenants,proto3" json:"tenants,omitempty"`
}

func (x *AddTenantsRequest) Reset() {
	*x = AddTenantsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTenantsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTenantsRequest) ProtoMessage() {}

func (x *AddTenantsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTenantsRequest.ProtoReflect.Descriptor instead.
func (*AddTenantsRequest) Descriptor() ([]byte, []int) {
	return file_cluster_message_proto_rawDescGZIP(), []int{5}
}

func (x *AddTenantsRequest) GetTenants() []*Tenant {
	if x != nil {
		return x.Tenants
	}
	return nil
}

type UpdateTenantsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenants []*Tenant `protobuf:"bytes,1,rep,name=tenants,proto3" json:"tenants,omitempty"`
}

func (x *UpdateTenantsRequest) Reset() {
	*x = UpdateTenantsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTenantsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTenantsRequest) ProtoMessage() {}

func (x *UpdateTenantsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTenantsRequest.ProtoReflect.Descriptor instead.
func (*UpdateTenantsRequest) Descriptor() ([]byte, []int) {
	return file_cluster_message_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateTenantsRequest) GetTenants() []*Tenant {
	if x != nil {
		return x.Tenants
	}
	return nil
}

type DeleteTenantsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenants []string `protobuf:"bytes,1,rep,name=tenants,proto3" json:"tenants,omitempty"`
}

func (x *DeleteTenantsRequest) Reset() {
	*x = DeleteTenantsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTenantsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTenantsRequest) ProtoMessage() {}

func (x *DeleteTenantsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTenantsRequest.ProtoReflect.Descriptor instead.
func (*DeleteTenantsRequest) Descriptor() ([]byte, []int) {
	return file_cluster_message_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteTenantsRequest) GetTenants() []string {
	if x != nil {
		return x.Tenants
	}
	return nil
}

type Tenant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Nodes  []string `protobuf:"bytes,3,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *Tenant) Reset() {
	*x = Tenant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_message_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tenant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tenant) ProtoMessage() {}

func (x *Tenant) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_message_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tenant.ProtoReflect.Descriptor instead.
func (*Tenant) Descriptor() ([]byte, []int) {
	return file_cluster_message_proto_rawDescGZIP(), []int{8}
}

func (x *Tenant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tenant) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Tenant) GetNodes() []string {
	if x != nil {
		return x.Nodes
	}
	return nil
}

var File_cluster_message_proto protoreflect.FileDescriptor

var file_cluster_message_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x51, 0x0a, 0x0f, 0x4a, 0x6f, 0x69, 0x6e,
	0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x22, 0x23, 0x0a, 0x11, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x14, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x4a, 0x6f, 0x69, 0x6e, 0x50, 0x65,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xf6, 0x02, 0x0a, 0x07, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x41, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x22, 0xf0, 0x01, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x12, 0x0a, 0x0e, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x43, 0x4c, 0x41, 0x53,
	0x53, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41,
	0x54, 0x45, 0x5f, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x10,
	0x03, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x54, 0x4f, 0x52,
	0x45, 0x5f, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52, 0x54, 0x59, 0x10, 0x05,
	0x12, 0x1c, 0x0a, 0x18, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f,
	0x53, 0x48, 0x41, 0x52, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x0a, 0x12, 0x13,
	0x0a, 0x0f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x54, 0x45, 0x4e, 0x41, 0x4e,
	0x54, 0x10, 0x10, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41,
	0x54, 0x45, 0x5f, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x54, 0x10, 0x11, 0x12, 0x16, 0x0a, 0x12, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x54, 0x45, 0x4e, 0x41, 0x4e,
	0x54, 0x10, 0x12, 0x22, 0x56, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x07, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x77, 0x65, 0x61, 0x76,
	0x69, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x52, 0x07, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x22, 0x59, 0x0a, 0x14, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x07, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x07, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x22, 0x30, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x22, 0x4a, 0x0a, 0x06, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x32, 0xfc, 0x01, 0x0a, 0x0e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x77, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x50, 0x65, 0x65, 0x72, 0x12, 0x32, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x77, 0x65, 0x61, 0x76,
	0x69, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x71, 0x0a, 0x08, 0x4a, 0x6f, 0x69, 0x6e, 0x50, 0x65, 0x65, 0x72, 0x12, 0x30, 0x2e, 0x77,
	0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4a,
	0x6f, 0x69, 0x6e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31,
	0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x85, 0x02, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x65, 0x61, 0x76,
	0x69, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x42, 0x0c, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65,
	0x2f, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0xa2, 0x02, 0x04,
	0x57, 0x43, 0x49, 0x43, 0xaa, 0x02, 0x1f, 0x57, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0xca, 0x02, 0x1f, 0x57, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74,
	0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0xe2, 0x02, 0x2b, 0x57, 0x65, 0x61, 0x76, 0x69,
	0x61, 0x74, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x5c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x22, 0x57, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74,
	0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x3a, 0x3a, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cluster_message_proto_rawDescOnce sync.Once
	file_cluster_message_proto_rawDescData = file_cluster_message_proto_rawDesc
)

func file_cluster_message_proto_rawDescGZIP() []byte {
	file_cluster_message_proto_rawDescOnce.Do(func() {
		file_cluster_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_cluster_message_proto_rawDescData)
	})
	return file_cluster_message_proto_rawDescData
}

var file_cluster_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cluster_message_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_cluster_message_proto_goTypes = []interface{}{
	(Command_Type)(0),            // 0: weaviate.cloud.internal.cluster.Command.Type
	(*JoinPeerRequest)(nil),      // 1: weaviate.cloud.internal.cluster.JoinPeerRequest
	(*RemovePeerRequest)(nil),    // 2: weaviate.cloud.internal.cluster.RemovePeerRequest
	(*RemovePeerResponse)(nil),   // 3: weaviate.cloud.internal.cluster.RemovePeerResponse
	(*JoinPeerResponse)(nil),     // 4: weaviate.cloud.internal.cluster.JoinPeerResponse
	(*Command)(nil),              // 5: weaviate.cloud.internal.cluster.Command
	(*AddTenantsRequest)(nil),    // 6: weaviate.cloud.internal.cluster.AddTenantsRequest
	(*UpdateTenantsRequest)(nil), // 7: weaviate.cloud.internal.cluster.UpdateTenantsRequest
	(*DeleteTenantsRequest)(nil), // 8: weaviate.cloud.internal.cluster.DeleteTenantsRequest
	(*Tenant)(nil),               // 9: weaviate.cloud.internal.cluster.Tenant
}
var file_cluster_message_proto_depIdxs = []int32{
	0, // 0: weaviate.cloud.internal.cluster.Command.type:type_name -> weaviate.cloud.internal.cluster.Command.Type
	9, // 1: weaviate.cloud.internal.cluster.AddTenantsRequest.tenants:type_name -> weaviate.cloud.internal.cluster.Tenant
	9, // 2: weaviate.cloud.internal.cluster.UpdateTenantsRequest.tenants:type_name -> weaviate.cloud.internal.cluster.Tenant
	2, // 3: weaviate.cloud.internal.cluster.ClusterService.RemovePeer:input_type -> weaviate.cloud.internal.cluster.RemovePeerRequest
	1, // 4: weaviate.cloud.internal.cluster.ClusterService.JoinPeer:input_type -> weaviate.cloud.internal.cluster.JoinPeerRequest
	3, // 5: weaviate.cloud.internal.cluster.ClusterService.RemovePeer:output_type -> weaviate.cloud.internal.cluster.RemovePeerResponse
	4, // 6: weaviate.cloud.internal.cluster.ClusterService.JoinPeer:output_type -> weaviate.cloud.internal.cluster.JoinPeerResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_cluster_message_proto_init() }
func file_cluster_message_proto_init() {
	if File_cluster_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cluster_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinPeerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cluster_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePeerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cluster_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePeerResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cluster_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinPeerResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cluster_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cluster_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTenantsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cluster_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTenantsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cluster_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTenantsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cluster_message_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tenant); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cluster_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cluster_message_proto_goTypes,
		DependencyIndexes: file_cluster_message_proto_depIdxs,
		EnumInfos:         file_cluster_message_proto_enumTypes,
		MessageInfos:      file_cluster_message_proto_msgTypes,
	}.Build()
	File_cluster_message_proto = out.File
	file_cluster_message_proto_rawDesc = nil
	file_cluster_message_proto_goTypes = nil
	file_cluster_message_proto_depIdxs = nil
}
