// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/controller/partition.proto

package controller

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Partition identifier
type PartitionId struct {
	Partition            int32             `protobuf:"varint,1,opt,name=partition,proto3" json:"partition,omitempty"`
	Group                *PartitionGroupId `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PartitionId) Reset()         { *m = PartitionId{} }
func (m *PartitionId) String() string { return proto.CompactTextString(m) }
func (*PartitionId) ProtoMessage()    {}
func (*PartitionId) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fbf59d47cbc2727, []int{0}
}
func (m *PartitionId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionId.Unmarshal(m, b)
}
func (m *PartitionId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionId.Marshal(b, m, deterministic)
}
func (m *PartitionId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionId.Merge(m, src)
}
func (m *PartitionId) XXX_Size() int {
	return xxx_messageInfo_PartitionId.Size(m)
}
func (m *PartitionId) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionId.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionId proto.InternalMessageInfo

func (m *PartitionId) GetPartition() int32 {
	if m != nil {
		return m.Partition
	}
	return 0
}

func (m *PartitionId) GetGroup() *PartitionGroupId {
	if m != nil {
		return m.Group
	}
	return nil
}

// Partition group name
type PartitionGroupId struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PartitionGroupId) Reset()         { *m = PartitionGroupId{} }
func (m *PartitionGroupId) String() string { return proto.CompactTextString(m) }
func (*PartitionGroupId) ProtoMessage()    {}
func (*PartitionGroupId) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fbf59d47cbc2727, []int{1}
}
func (m *PartitionGroupId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionGroupId.Unmarshal(m, b)
}
func (m *PartitionGroupId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionGroupId.Marshal(b, m, deterministic)
}
func (m *PartitionGroupId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionGroupId.Merge(m, src)
}
func (m *PartitionGroupId) XXX_Size() int {
	return xxx_messageInfo_PartitionGroupId.Size(m)
}
func (m *PartitionGroupId) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionGroupId.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionGroupId proto.InternalMessageInfo

func (m *PartitionGroupId) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PartitionGroupId) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

// Partition configuration
type PartitionConfig struct {
	Partition            *PartitionId  `protobuf:"bytes,1,opt,name=partition,proto3" json:"partition,omitempty"`
	Controller           *NodeConfig   `protobuf:"bytes,2,opt,name=controller,proto3" json:"controller,omitempty"`
	Members              []*NodeConfig `protobuf:"bytes,3,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PartitionConfig) Reset()         { *m = PartitionConfig{} }
func (m *PartitionConfig) String() string { return proto.CompactTextString(m) }
func (*PartitionConfig) ProtoMessage()    {}
func (*PartitionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fbf59d47cbc2727, []int{2}
}
func (m *PartitionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionConfig.Unmarshal(m, b)
}
func (m *PartitionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionConfig.Marshal(b, m, deterministic)
}
func (m *PartitionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionConfig.Merge(m, src)
}
func (m *PartitionConfig) XXX_Size() int {
	return xxx_messageInfo_PartitionConfig.Size(m)
}
func (m *PartitionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionConfig proto.InternalMessageInfo

func (m *PartitionConfig) GetPartition() *PartitionId {
	if m != nil {
		return m.Partition
	}
	return nil
}

func (m *PartitionConfig) GetController() *NodeConfig {
	if m != nil {
		return m.Controller
	}
	return nil
}

func (m *PartitionConfig) GetMembers() []*NodeConfig {
	if m != nil {
		return m.Members
	}
	return nil
}

// Node configuration
type NodeConfig struct {
	ID                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Host                 string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Port                 int32    `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	APIPort              int32    `protobuf:"varint,4,opt,name=api_port,json=apiPort,proto3" json:"api_port,omitempty"`
	ProtocolPort         int32    `protobuf:"varint,5,opt,name=protocol_port,json=protocolPort,proto3" json:"protocol_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeConfig) Reset()         { *m = NodeConfig{} }
func (m *NodeConfig) String() string { return proto.CompactTextString(m) }
func (*NodeConfig) ProtoMessage()    {}
func (*NodeConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fbf59d47cbc2727, []int{3}
}
func (m *NodeConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeConfig.Unmarshal(m, b)
}
func (m *NodeConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeConfig.Marshal(b, m, deterministic)
}
func (m *NodeConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeConfig.Merge(m, src)
}
func (m *NodeConfig) XXX_Size() int {
	return xxx_messageInfo_NodeConfig.Size(m)
}
func (m *NodeConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NodeConfig proto.InternalMessageInfo

func (m *NodeConfig) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *NodeConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *NodeConfig) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *NodeConfig) GetAPIPort() int32 {
	if m != nil {
		return m.APIPort
	}
	return 0
}

func (m *NodeConfig) GetProtocolPort() int32 {
	if m != nil {
		return m.ProtocolPort
	}
	return 0
}

// Partition group info
type PartitionGroup struct {
	ID                   *PartitionGroupId   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Spec                 *PartitionGroupSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Partitions           []*Partition        `protobuf:"bytes,3,rep,name=partitions,proto3" json:"partitions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PartitionGroup) Reset()         { *m = PartitionGroup{} }
func (m *PartitionGroup) String() string { return proto.CompactTextString(m) }
func (*PartitionGroup) ProtoMessage()    {}
func (*PartitionGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fbf59d47cbc2727, []int{4}
}
func (m *PartitionGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionGroup.Unmarshal(m, b)
}
func (m *PartitionGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionGroup.Marshal(b, m, deterministic)
}
func (m *PartitionGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionGroup.Merge(m, src)
}
func (m *PartitionGroup) XXX_Size() int {
	return xxx_messageInfo_PartitionGroup.Size(m)
}
func (m *PartitionGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionGroup.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionGroup proto.InternalMessageInfo

func (m *PartitionGroup) GetID() *PartitionGroupId {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *PartitionGroup) GetSpec() *PartitionGroupSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *PartitionGroup) GetPartitions() []*Partition {
	if m != nil {
		return m.Partitions
	}
	return nil
}

// Partition info
type Partition struct {
	PartitionID          int32                `protobuf:"varint,1,opt,name=partition_id,json=partitionId,proto3" json:"partition_id,omitempty"`
	Endpoints            []*PartitionEndpoint `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Partition) Reset()         { *m = Partition{} }
func (m *Partition) String() string { return proto.CompactTextString(m) }
func (*Partition) ProtoMessage()    {}
func (*Partition) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fbf59d47cbc2727, []int{5}
}
func (m *Partition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Partition.Unmarshal(m, b)
}
func (m *Partition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Partition.Marshal(b, m, deterministic)
}
func (m *Partition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Partition.Merge(m, src)
}
func (m *Partition) XXX_Size() int {
	return xxx_messageInfo_Partition.Size(m)
}
func (m *Partition) XXX_DiscardUnknown() {
	xxx_messageInfo_Partition.DiscardUnknown(m)
}

var xxx_messageInfo_Partition proto.InternalMessageInfo

func (m *Partition) GetPartitionID() int32 {
	if m != nil {
		return m.PartitionID
	}
	return 0
}

func (m *Partition) GetEndpoints() []*PartitionEndpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

// Partition endpoint
type PartitionEndpoint struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PartitionEndpoint) Reset()         { *m = PartitionEndpoint{} }
func (m *PartitionEndpoint) String() string { return proto.CompactTextString(m) }
func (*PartitionEndpoint) ProtoMessage()    {}
func (*PartitionEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fbf59d47cbc2727, []int{6}
}
func (m *PartitionEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionEndpoint.Unmarshal(m, b)
}
func (m *PartitionEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionEndpoint.Marshal(b, m, deterministic)
}
func (m *PartitionEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionEndpoint.Merge(m, src)
}
func (m *PartitionEndpoint) XXX_Size() int {
	return xxx_messageInfo_PartitionEndpoint.Size(m)
}
func (m *PartitionEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionEndpoint proto.InternalMessageInfo

func (m *PartitionEndpoint) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *PartitionEndpoint) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

// Partition group specification
type PartitionGroupSpec struct {
	Partitions           uint32     `protobuf:"varint,1,opt,name=partitions,proto3" json:"partitions,omitempty"`
	PartitionSize        uint32     `protobuf:"varint,2,opt,name=partition_size,json=partitionSize,proto3" json:"partition_size,omitempty"`
	Protocol             *types.Any `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PartitionGroupSpec) Reset()         { *m = PartitionGroupSpec{} }
func (m *PartitionGroupSpec) String() string { return proto.CompactTextString(m) }
func (*PartitionGroupSpec) ProtoMessage()    {}
func (*PartitionGroupSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fbf59d47cbc2727, []int{7}
}
func (m *PartitionGroupSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionGroupSpec.Unmarshal(m, b)
}
func (m *PartitionGroupSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionGroupSpec.Marshal(b, m, deterministic)
}
func (m *PartitionGroupSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionGroupSpec.Merge(m, src)
}
func (m *PartitionGroupSpec) XXX_Size() int {
	return xxx_messageInfo_PartitionGroupSpec.Size(m)
}
func (m *PartitionGroupSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionGroupSpec.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionGroupSpec proto.InternalMessageInfo

func (m *PartitionGroupSpec) GetPartitions() uint32 {
	if m != nil {
		return m.Partitions
	}
	return 0
}

func (m *PartitionGroupSpec) GetPartitionSize() uint32 {
	if m != nil {
		return m.PartitionSize
	}
	return 0
}

func (m *PartitionGroupSpec) GetProtocol() *types.Any {
	if m != nil {
		return m.Protocol
	}
	return nil
}

func init() {
	proto.RegisterType((*PartitionId)(nil), "atomix.controller.PartitionId")
	proto.RegisterType((*PartitionGroupId)(nil), "atomix.controller.PartitionGroupId")
	proto.RegisterType((*PartitionConfig)(nil), "atomix.controller.PartitionConfig")
	proto.RegisterType((*NodeConfig)(nil), "atomix.controller.NodeConfig")
	proto.RegisterType((*PartitionGroup)(nil), "atomix.controller.PartitionGroup")
	proto.RegisterType((*Partition)(nil), "atomix.controller.Partition")
	proto.RegisterType((*PartitionEndpoint)(nil), "atomix.controller.PartitionEndpoint")
	proto.RegisterType((*PartitionGroupSpec)(nil), "atomix.controller.PartitionGroupSpec")
}

func init() { proto.RegisterFile("atomix/controller/partition.proto", fileDescriptor_6fbf59d47cbc2727) }

var fileDescriptor_6fbf59d47cbc2727 = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x4e, 0xdb, 0x40,
	0x14, 0x96, 0x4d, 0x42, 0xf0, 0x33, 0x81, 0x32, 0x42, 0x95, 0x8b, 0x28, 0x49, 0x4d, 0xa9, 0x58,
	0xd9, 0x55, 0xba, 0xa8, 0x10, 0x74, 0x41, 0x9a, 0xaa, 0xf2, 0xa6, 0x8a, 0x86, 0x03, 0x20, 0xc7,
	0x9e, 0xb8, 0x23, 0x25, 0x7e, 0x23, 0xdb, 0x48, 0x85, 0x6d, 0xd7, 0x3d, 0x40, 0x2f, 0xd4, 0xde,
	0x22, 0x8b, 0x9c, 0xa4, 0xf2, 0xd8, 0x1e, 0x3b, 0x04, 0x85, 0xae, 0xf2, 0xf4, 0xbd, 0xef, 0xfd,
	0x7d, 0x5f, 0xc6, 0xf0, 0xc6, 0xcf, 0x70, 0xce, 0x7f, 0xb8, 0x01, 0xc6, 0x59, 0x82, 0xb3, 0x19,
	0x4b, 0x5c, 0xe1, 0x27, 0x19, 0xcf, 0x38, 0xc6, 0x8e, 0x48, 0x30, 0x43, 0x72, 0x50, 0x50, 0x9c,
	0x9a, 0x72, 0xf4, 0x2a, 0x42, 0x8c, 0x66, 0xcc, 0x95, 0x84, 0xc9, 0xdd, 0xd4, 0xf5, 0xe3, 0xfb,
	0x82, 0x7d, 0x74, 0x18, 0x61, 0x84, 0x32, 0x74, 0xf3, 0xa8, 0x40, 0xed, 0x29, 0x98, 0xe3, 0xaa,
	0xad, 0x17, 0x92, 0x63, 0x30, 0xd4, 0x14, 0x4b, 0xeb, 0x6b, 0xe7, 0x6d, 0x5a, 0x03, 0xe4, 0x02,
	0xda, 0x51, 0x82, 0x77, 0xc2, 0xd2, 0xfb, 0xda, 0xb9, 0x39, 0x38, 0x75, 0xd6, 0x16, 0x70, 0x54,
	0xb3, 0xaf, 0x39, 0xd1, 0x0b, 0x69, 0x51, 0x61, 0x8f, 0xe0, 0xc5, 0xe3, 0x14, 0x21, 0xd0, 0x8a,
	0xfd, 0x39, 0x93, 0x73, 0x0c, 0x2a, 0xe3, 0x7c, 0x81, 0xfc, 0x37, 0x15, 0x7e, 0xc0, 0xe4, 0x18,
	0x83, 0xd6, 0x80, 0xfd, 0x57, 0x83, 0x7d, 0xd5, 0xe6, 0x33, 0xc6, 0x53, 0x1e, 0x91, 0xab, 0xc7,
	0x2b, 0x9b, 0x83, 0x93, 0x4d, 0x8b, 0x79, 0x61, 0xf3, 0xa4, 0x4f, 0x00, 0x35, 0xa9, 0xbc, 0xeb,
	0xf5, 0x13, 0xe5, 0xdf, 0x30, 0x64, 0xc5, 0x40, 0xda, 0x28, 0x20, 0x1f, 0xa1, 0x33, 0x67, 0xf3,
	0x09, 0x4b, 0x52, 0x6b, 0xab, 0xbf, 0xf5, 0x7c, 0x6d, 0xc5, 0xb6, 0x7f, 0x6b, 0x00, 0x35, 0x4e,
	0x5e, 0x82, 0xce, 0xc3, 0x42, 0x88, 0xe1, 0xf6, 0x72, 0xd1, 0xd3, 0xbd, 0x11, 0xd5, 0xb9, 0x94,
	0xe8, 0x3b, 0xa6, 0x59, 0xa9, 0x84, 0x8c, 0x73, 0x4c, 0x60, 0x92, 0x59, 0x5b, 0xd2, 0x1e, 0x19,
	0x93, 0x77, 0xb0, 0xe3, 0x0b, 0x7e, 0x2b, 0xf1, 0x56, 0x8e, 0x0f, 0xcd, 0xe5, 0xa2, 0xd7, 0xb9,
	0x1e, 0x7b, 0x63, 0x4c, 0x32, 0xda, 0xf1, 0x05, 0xcf, 0x03, 0x72, 0x0a, 0x5d, 0xe9, 0x7b, 0x80,
	0xb3, 0x82, 0xdc, 0x96, 0x4d, 0x76, 0x2b, 0x30, 0x27, 0xd9, 0x7f, 0x34, 0xd8, 0x5b, 0x35, 0x8b,
	0x5c, 0xaa, 0xfd, 0xfe, 0xcf, 0xf6, 0x95, 0x23, 0x2e, 0xa0, 0x95, 0x0a, 0x16, 0x94, 0xea, 0x9e,
	0x3d, 0x5b, 0x7e, 0x23, 0x58, 0x40, 0x65, 0x09, 0xb9, 0x02, 0x50, 0x5e, 0x55, 0x12, 0x1f, 0x6f,
	0x6a, 0x40, 0x1b, 0x7c, 0xfb, 0xa7, 0x06, 0x86, 0xca, 0x90, 0x01, 0xec, 0xaa, 0xdc, 0x6d, 0x79,
	0x4d, 0x7b, 0xb8, 0xbf, 0x5c, 0xf4, 0x1a, 0x4f, 0x60, 0x44, 0x4d, 0xd1, 0x78, 0x0f, 0x43, 0x30,
	0x58, 0x1c, 0x0a, 0xe4, 0x71, 0x96, 0x5a, 0xba, 0x1c, 0xff, 0x76, 0xd3, 0xf8, 0x2f, 0x25, 0x99,
	0xd6, 0x65, 0xf6, 0x25, 0x1c, 0xac, 0xe5, 0x95, 0xb1, 0xda, 0x13, 0xc6, 0xea, 0xb5, 0xb1, 0xf6,
	0x2f, 0x0d, 0xc8, 0xba, 0x3a, 0xe4, 0x64, 0x45, 0x97, 0xbc, 0x49, 0xb7, 0x79, 0x39, 0x39, 0x83,
	0xbd, 0xfa, 0xd6, 0x94, 0x3f, 0x14, 0x6f, 0xa9, 0x4b, 0xbb, 0x0a, 0xbd, 0xe1, 0x0f, 0x8c, 0xbc,
	0x87, 0x9d, 0xca, 0x79, 0xf9, 0x77, 0x32, 0x07, 0x87, 0x4e, 0xf1, 0x05, 0x71, 0xaa, 0x2f, 0x88,
	0x73, 0x1d, 0xdf, 0x53, 0xc5, 0x9a, 0x6c, 0xcb, 0xe8, 0xc3, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x1f, 0xdf, 0xea, 0x12, 0x9f, 0x04, 0x00, 0x00,
}
