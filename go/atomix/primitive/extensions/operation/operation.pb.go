// Code generated by protoc-gen-go. DO NOT EDIT.
// source: atomix/primitive/operation.proto

package operation

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// OperationType is an enum for specifying the type of operation
type OperationType int32

const (
	OperationType_COMMAND OperationType = 0
	OperationType_QUERY   OperationType = 1
)

var OperationType_name = map[int32]string{
	0: "COMMAND",
	1: "QUERY",
}

var OperationType_value = map[string]int32{
	"COMMAND": 0,
	"QUERY":   1,
}

func (x OperationType) String() string {
	return proto.EnumName(OperationType_name, int32(x))
}

func (OperationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7578e8285e8099c5, []int{0}
}

// AggregateStrategy is an enum for indicating the strategy used to aggregate a field
type AggregateStrategy int32

const (
	AggregateStrategy_CHOOSE_FIRST AggregateStrategy = 0
	AggregateStrategy_APPEND       AggregateStrategy = 1
	AggregateStrategy_SUM          AggregateStrategy = 2
)

var AggregateStrategy_name = map[int32]string{
	0: "CHOOSE_FIRST",
	1: "APPEND",
	2: "SUM",
}

var AggregateStrategy_value = map[string]int32{
	"CHOOSE_FIRST": 0,
	"APPEND":       1,
	"SUM":          2,
}

func (x AggregateStrategy) String() string {
	return proto.EnumName(AggregateStrategy_name, int32(x))
}

func (AggregateStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7578e8285e8099c5, []int{1}
}

var E_Name = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         61000,
	Name:          "atomix.primitive.operation.name",
	Tag:           "bytes,61000,opt,name=name",
	Filename:      "atomix/primitive/operation.proto",
}

var E_Type = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*OperationType)(nil),
	Field:         61001,
	Name:          "atomix.primitive.operation.type",
	Tag:           "varint,61001,opt,name=type,enum=atomix.primitive.operation.OperationType",
	Filename:      "atomix/primitive/operation.proto",
}

var E_Async = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         61002,
	Name:          "atomix.primitive.operation.async",
	Tag:           "varint,61002,opt,name=async",
	Filename:      "atomix/primitive/operation.proto",
}

var E_Headers = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         62000,
	Name:          "atomix.primitive.operation.headers",
	Tag:           "varint,62000,opt,name=headers",
	Filename:      "atomix/primitive/operation.proto",
}

var E_Aggregate = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*AggregateStrategy)(nil),
	Field:         62001,
	Name:          "atomix.primitive.operation.aggregate",
	Tag:           "varint,62001,opt,name=aggregate,enum=atomix.primitive.operation.AggregateStrategy",
	Filename:      "atomix/primitive/operation.proto",
}

func init() {
	proto.RegisterEnum("atomix.primitive.operation.OperationType", OperationType_name, OperationType_value)
	proto.RegisterEnum("atomix.primitive.operation.AggregateStrategy", AggregateStrategy_name, AggregateStrategy_value)
	proto.RegisterExtension(E_Name)
	proto.RegisterExtension(E_Type)
	proto.RegisterExtension(E_Async)
	proto.RegisterExtension(E_Headers)
	proto.RegisterExtension(E_Aggregate)
}

func init() { proto.RegisterFile("atomix/primitive/operation.proto", fileDescriptor_7578e8285e8099c5) }

var fileDescriptor_7578e8285e8099c5 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcb, 0x4e, 0xc2, 0x40,
	0x14, 0x86, 0x41, 0x6e, 0x72, 0xbc, 0xa4, 0xce, 0xca, 0x90, 0x68, 0xd8, 0xa9, 0x24, 0x4e, 0x13,
	0x35, 0x26, 0x76, 0x47, 0xb8, 0x44, 0x17, 0xa5, 0xd8, 0xc2, 0xc2, 0x15, 0x19, 0xe0, 0x58, 0x26,
	0x01, 0xa6, 0x99, 0x8e, 0x46, 0xde, 0x4a, 0xdf, 0x42, 0x7d, 0x06, 0x97, 0x3c, 0x88, 0x61, 0x26,
	0xc5, 0x18, 0xa2, 0xdd, 0x7f, 0xdf, 0xfc, 0xff, 0x7f, 0x06, 0xaa, 0x4c, 0x89, 0x19, 0x7f, 0xb1,
	0x23, 0xc9, 0x67, 0x5c, 0xf1, 0x67, 0xb4, 0x45, 0x84, 0x92, 0x29, 0x2e, 0xe6, 0x34, 0x92, 0x42,
	0x09, 0x52, 0x31, 0x04, 0x5d, 0x13, 0x74, 0x4d, 0x54, 0xaa, 0xa1, 0x10, 0xe1, 0x14, 0x6d, 0x4d,
	0x0e, 0x9f, 0x1e, 0xed, 0x31, 0xc6, 0x23, 0xc9, 0x23, 0x25, 0xa4, 0xb1, 0x6b, 0x27, 0xb0, 0xe7,
	0x25, 0x78, 0x6f, 0x11, 0x21, 0xd9, 0x81, 0x52, 0xc3, 0x73, 0xdd, 0x7a, 0xa7, 0x69, 0x65, 0x48,
	0x19, 0x0a, 0xf7, 0xfd, 0x96, 0xff, 0x60, 0x65, 0x6b, 0x0e, 0x1c, 0xd4, 0xc3, 0x50, 0x62, 0xc8,
	0x14, 0x06, 0x4a, 0x32, 0x85, 0xe1, 0x82, 0x58, 0xb0, 0xdb, 0xb8, 0xf5, 0xbc, 0xa0, 0x35, 0x68,
	0xdf, 0xf9, 0x41, 0xcf, 0xca, 0x10, 0x80, 0x62, 0xbd, 0xdb, 0x6d, 0x75, 0x9a, 0x56, 0x96, 0x94,
	0x20, 0x17, 0xf4, 0x5d, 0x6b, 0xcb, 0xb9, 0x82, 0xfc, 0x9c, 0xcd, 0x90, 0x1c, 0x53, 0xd3, 0x87,
	0x26, 0x7d, 0xa8, 0x8b, 0x6a, 0x22, 0xc6, 0x5e, 0xb4, 0x8a, 0x8f, 0x0f, 0xdf, 0xbf, 0x72, 0xd5,
	0xec, 0x69, 0xd9, 0xd7, 0xb4, 0x33, 0x80, 0xbc, 0x5a, 0x35, 0x4a, 0xb3, 0x3e, 0xb4, 0xb5, 0x7f,
	0x71, 0x46, 0xff, 0xbe, 0x04, 0xfd, 0x35, 0xd2, 0xd7, 0x0f, 0x3b, 0xd7, 0x50, 0x60, 0xf1, 0x62,
	0x3e, 0x4a, 0x4d, 0xf8, 0xd4, 0x09, 0xdb, 0xbe, 0xc1, 0x9d, 0x1b, 0x28, 0x4d, 0x90, 0x8d, 0x51,
	0xc6, 0xe4, 0x68, 0xc3, 0x6c, 0x73, 0x9c, 0xae, 0xc5, 0xd7, 0xa5, 0x11, 0x13, 0xde, 0x99, 0x42,
	0x99, 0x25, 0x57, 0x4c, 0x93, 0xdf, 0x96, 0x66, 0xd7, 0xf9, 0x7f, 0xbb, 0x36, 0xfe, 0xc4, 0xff,
	0x09, 0x18, 0x16, 0xf5, 0xc3, 0x97, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x20, 0xf9, 0x54,
	0x45, 0x02, 0x00, 0x00,
}