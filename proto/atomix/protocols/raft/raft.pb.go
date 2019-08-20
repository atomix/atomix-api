// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/protocols/raft/raft.proto

package raft

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Storage level
type StorageLevel int32

const (
	StorageLevel_DISK   StorageLevel = 0
	StorageLevel_MAPPED StorageLevel = 1
)

var StorageLevel_name = map[int32]string{
	0: "DISK",
	1: "MAPPED",
}

var StorageLevel_value = map[string]int32{
	"DISK":   0,
	"MAPPED": 1,
}

func (x StorageLevel) String() string {
	return proto.EnumName(StorageLevel_name, int32(x))
}

func (StorageLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a5bc54f311a0a43e, []int{0}
}

// Member group strategy
type MemberGroupStrategy int32

const (
	MemberGroupStrategy_HOST_AWARE MemberGroupStrategy = 0
	MemberGroupStrategy_RACK_AWARE MemberGroupStrategy = 1
	MemberGroupStrategy_ZONE_AWARE MemberGroupStrategy = 2
)

var MemberGroupStrategy_name = map[int32]string{
	0: "HOST_AWARE",
	1: "RACK_AWARE",
	2: "ZONE_AWARE",
}

var MemberGroupStrategy_value = map[string]int32{
	"HOST_AWARE": 0,
	"RACK_AWARE": 1,
	"ZONE_AWARE": 2,
}

func (x MemberGroupStrategy) String() string {
	return proto.EnumName(MemberGroupStrategy_name, int32(x))
}

func (MemberGroupStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a5bc54f311a0a43e, []int{1}
}

// Raft protocol configuration
type RaftProtocol struct {
	ElectionTimeout   *time.Duration  `protobuf:"bytes,1,opt,name=election_timeout,json=electionTimeout,proto3,stdduration" json:"election_timeout,omitempty"`
	HeartbeatInterval *time.Duration  `protobuf:"bytes,2,opt,name=heartbeat_interval,json=heartbeatInterval,proto3,stdduration" json:"heartbeat_interval,omitempty"`
	Storage           *StorageSpec    `protobuf:"bytes,3,opt,name=storage,proto3" json:"storage,omitempty"`
	Compaction        *CompactionSpec `protobuf:"bytes,4,opt,name=compaction,proto3" json:"compaction,omitempty"`
}

func (m *RaftProtocol) Reset()         { *m = RaftProtocol{} }
func (m *RaftProtocol) String() string { return proto.CompactTextString(m) }
func (*RaftProtocol) ProtoMessage()    {}
func (*RaftProtocol) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5bc54f311a0a43e, []int{0}
}
func (m *RaftProtocol) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RaftProtocol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RaftProtocol.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RaftProtocol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RaftProtocol.Merge(m, src)
}
func (m *RaftProtocol) XXX_Size() int {
	return m.Size()
}
func (m *RaftProtocol) XXX_DiscardUnknown() {
	xxx_messageInfo_RaftProtocol.DiscardUnknown(m)
}

var xxx_messageInfo_RaftProtocol proto.InternalMessageInfo

func (m *RaftProtocol) GetElectionTimeout() *time.Duration {
	if m != nil {
		return m.ElectionTimeout
	}
	return nil
}

func (m *RaftProtocol) GetHeartbeatInterval() *time.Duration {
	if m != nil {
		return m.HeartbeatInterval
	}
	return nil
}

func (m *RaftProtocol) GetStorage() *StorageSpec {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *RaftProtocol) GetCompaction() *CompactionSpec {
	if m != nil {
		return m.Compaction
	}
	return nil
}

// Partition group storage configuration
type StorageSpec struct {
	MaxEntrySize  uint32       `protobuf:"varint,1,opt,name=max_entry_size,json=maxEntrySize,proto3" json:"max_entry_size,omitempty"`
	SegmentSize   uint32       `protobuf:"varint,2,opt,name=segment_size,json=segmentSize,proto3" json:"segment_size,omitempty"`
	Level         StorageLevel `protobuf:"varint,3,opt,name=level,proto3,enum=atomix.protocols.raft.StorageLevel" json:"level,omitempty"`
	FlushOnCommit bool         `protobuf:"varint,4,opt,name=flush_on_commit,json=flushOnCommit,proto3" json:"flush_on_commit,omitempty"`
}

func (m *StorageSpec) Reset()         { *m = StorageSpec{} }
func (m *StorageSpec) String() string { return proto.CompactTextString(m) }
func (*StorageSpec) ProtoMessage()    {}
func (*StorageSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5bc54f311a0a43e, []int{1}
}
func (m *StorageSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StorageSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StorageSpec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StorageSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageSpec.Merge(m, src)
}
func (m *StorageSpec) XXX_Size() int {
	return m.Size()
}
func (m *StorageSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageSpec.DiscardUnknown(m)
}

var xxx_messageInfo_StorageSpec proto.InternalMessageInfo

func (m *StorageSpec) GetMaxEntrySize() uint32 {
	if m != nil {
		return m.MaxEntrySize
	}
	return 0
}

func (m *StorageSpec) GetSegmentSize() uint32 {
	if m != nil {
		return m.SegmentSize
	}
	return 0
}

func (m *StorageSpec) GetLevel() StorageLevel {
	if m != nil {
		return m.Level
	}
	return StorageLevel_DISK
}

func (m *StorageSpec) GetFlushOnCommit() bool {
	if m != nil {
		return m.FlushOnCommit
	}
	return false
}

// Partition group compaction configuration
type CompactionSpec struct {
	Dynamic        bool    `protobuf:"varint,1,opt,name=dynamic,proto3" json:"dynamic,omitempty"`
	FreeDiskBuffer float64 `protobuf:"fixed64,2,opt,name=free_disk_buffer,json=freeDiskBuffer,proto3" json:"free_disk_buffer,omitempty"`
}

func (m *CompactionSpec) Reset()         { *m = CompactionSpec{} }
func (m *CompactionSpec) String() string { return proto.CompactTextString(m) }
func (*CompactionSpec) ProtoMessage()    {}
func (*CompactionSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5bc54f311a0a43e, []int{2}
}
func (m *CompactionSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CompactionSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CompactionSpec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CompactionSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompactionSpec.Merge(m, src)
}
func (m *CompactionSpec) XXX_Size() int {
	return m.Size()
}
func (m *CompactionSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_CompactionSpec.DiscardUnknown(m)
}

var xxx_messageInfo_CompactionSpec proto.InternalMessageInfo

func (m *CompactionSpec) GetDynamic() bool {
	if m != nil {
		return m.Dynamic
	}
	return false
}

func (m *CompactionSpec) GetFreeDiskBuffer() float64 {
	if m != nil {
		return m.FreeDiskBuffer
	}
	return 0
}

func init() {
	proto.RegisterEnum("atomix.protocols.raft.StorageLevel", StorageLevel_name, StorageLevel_value)
	proto.RegisterEnum("atomix.protocols.raft.MemberGroupStrategy", MemberGroupStrategy_name, MemberGroupStrategy_value)
	proto.RegisterType((*RaftProtocol)(nil), "atomix.protocols.raft.RaftProtocol")
	proto.RegisterType((*StorageSpec)(nil), "atomix.protocols.raft.StorageSpec")
	proto.RegisterType((*CompactionSpec)(nil), "atomix.protocols.raft.CompactionSpec")
}

func init() { proto.RegisterFile("atomix/protocols/raft/raft.proto", fileDescriptor_a5bc54f311a0a43e) }

var fileDescriptor_a5bc54f311a0a43e = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xcd, 0x6e, 0xd3, 0x4e,
	0x14, 0xc5, 0xe3, 0xfc, 0xf3, 0x6f, 0xa3, 0x9b, 0x34, 0x0d, 0x03, 0x48, 0xa6, 0x0b, 0x53, 0x42,
	0x41, 0x55, 0x17, 0x8e, 0x54, 0x56, 0x48, 0x6c, 0xf2, 0x61, 0x41, 0x29, 0x6d, 0x22, 0x3b, 0x12,
	0x12, 0x1b, 0x6b, 0xec, 0x5c, 0xbb, 0x56, 0x3d, 0x9e, 0x68, 0x3c, 0xae, 0x92, 0x3e, 0x05, 0x4b,
	0xde, 0x82, 0x35, 0x6f, 0xc0, 0xb2, 0x4b, 0x76, 0xa0, 0xe4, 0x45, 0x90, 0xc7, 0x4e, 0x14, 0x24,
	0x40, 0x6c, 0xac, 0xb9, 0xc7, 0xe7, 0x77, 0xee, 0x07, 0x1c, 0x52, 0xc9, 0x59, 0x34, 0xef, 0xce,
	0x04, 0x97, 0xdc, 0xe7, 0x71, 0xda, 0x15, 0x34, 0x90, 0xea, 0x63, 0x2a, 0x8d, 0x3c, 0x2c, 0x1c,
	0xe6, 0xc6, 0x61, 0xe6, 0x3f, 0x0f, 0x8c, 0x90, 0xf3, 0x30, 0xc6, 0x02, 0xf4, 0xb2, 0xa0, 0x3b,
	0xcd, 0x04, 0x95, 0x11, 0x4f, 0x0a, 0xe3, 0xc1, 0x83, 0x90, 0x87, 0x5c, 0x3d, 0xbb, 0xf9, 0xab,
	0x50, 0x3b, 0x9f, 0xab, 0xd0, 0xb4, 0x69, 0x20, 0xc7, 0x65, 0x18, 0x79, 0x0b, 0x6d, 0x8c, 0xd1,
	0xcf, 0x41, 0x57, 0x46, 0x0c, 0x79, 0x26, 0x75, 0xed, 0x50, 0x3b, 0x6e, 0x9c, 0x3e, 0x32, 0x8b,
	0x0e, 0xe6, 0xba, 0x83, 0x39, 0x2c, 0x3b, 0xf4, 0x6b, 0x9f, 0xbe, 0x3f, 0xd6, 0xec, 0xfd, 0x35,
	0x38, 0x29, 0x38, 0x72, 0x09, 0xe4, 0x0a, 0xa9, 0x90, 0x1e, 0x52, 0xe9, 0x46, 0x89, 0x44, 0x71,
	0x43, 0x63, 0xbd, 0xfa, 0x6f, 0x69, 0xf7, 0x36, 0xe8, 0x59, 0x49, 0x92, 0x57, 0xb0, 0x9b, 0x4a,
	0x2e, 0x68, 0x88, 0xfa, 0x7f, 0x2a, 0xa4, 0x63, 0xfe, 0xf6, 0x16, 0xa6, 0x53, 0xb8, 0x9c, 0x19,
	0xfa, 0xf6, 0x1a, 0x21, 0x16, 0x80, 0xcf, 0xd9, 0x8c, 0xaa, 0x11, 0xf5, 0x9a, 0x0a, 0x78, 0xf6,
	0x87, 0x80, 0xc1, 0xc6, 0xa8, 0x32, 0xb6, 0xc0, 0xce, 0x17, 0x0d, 0x1a, 0x5b, 0xf9, 0xe4, 0x08,
	0x5a, 0x8c, 0xce, 0x5d, 0x4c, 0xa4, 0x58, 0xb8, 0x69, 0x74, 0x8b, 0xea, 0x5c, 0x7b, 0x76, 0x93,
	0xd1, 0xb9, 0x95, 0x8b, 0x4e, 0x74, 0x8b, 0xe4, 0x09, 0x34, 0x53, 0x0c, 0x19, 0x26, 0xb2, 0xf0,
	0x54, 0x95, 0xa7, 0x51, 0x6a, 0xca, 0xf2, 0x12, 0xfe, 0x8f, 0xf1, 0x06, 0x63, 0xb5, 0x5b, 0xeb,
	0xf4, 0xe9, 0xdf, 0x77, 0x7b, 0x97, 0x5b, 0xed, 0x82, 0x20, 0xcf, 0x61, 0x3f, 0x88, 0xb3, 0xf4,
	0xca, 0xe5, 0x89, 0xeb, 0x73, 0xc6, 0x22, 0xa9, 0xf6, 0xab, 0xdb, 0x7b, 0x4a, 0x1e, 0x25, 0x03,
	0x25, 0x76, 0x26, 0xd0, 0xfa, 0x75, 0x33, 0xa2, 0xc3, 0xee, 0x74, 0x91, 0x50, 0x16, 0xf9, 0x6a,
	0xec, 0xba, 0xbd, 0x2e, 0xc9, 0x31, 0xb4, 0x03, 0x81, 0xe8, 0x4e, 0xa3, 0xf4, 0xda, 0xf5, 0xb2,
	0x20, 0x40, 0xa1, 0xa6, 0xd6, 0xec, 0x56, 0xae, 0x0f, 0xa3, 0xf4, 0xba, 0xaf, 0xd4, 0x93, 0x23,
	0x68, 0x6e, 0x0f, 0x45, 0xea, 0x50, 0x1b, 0x9e, 0x39, 0xe7, 0xed, 0x0a, 0x01, 0xd8, 0xb9, 0xe8,
	0x8d, 0xc7, 0xd6, 0xb0, 0xad, 0x9d, 0x58, 0x70, 0xff, 0x02, 0x99, 0x87, 0xe2, 0xb5, 0xe0, 0xd9,
	0xcc, 0x91, 0x82, 0x4a, 0x0c, 0x17, 0xa4, 0x05, 0xf0, 0x66, 0xe4, 0x4c, 0xdc, 0xde, 0xfb, 0x9e,
	0x6d, 0xb5, 0x2b, 0x79, 0x6d, 0xf7, 0x06, 0xe7, 0x65, 0xad, 0xe5, 0xf5, 0x87, 0xd1, 0xa5, 0x55,
	0xd6, 0xd5, 0xbe, 0xfe, 0x75, 0x69, 0x68, 0x77, 0x4b, 0x43, 0xfb, 0xb1, 0x34, 0xb4, 0x8f, 0x2b,
	0xa3, 0x72, 0xb7, 0x32, 0x2a, 0xdf, 0x56, 0x46, 0xc5, 0xdb, 0x51, 0x87, 0x7a, 0xf1, 0x33, 0x00,
	0x00, 0xff, 0xff, 0xfd, 0xb0, 0x8e, 0x5b, 0x42, 0x03, 0x00, 0x00,
}

func (m *RaftProtocol) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RaftProtocol) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RaftProtocol) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Compaction != nil {
		{
			size, err := m.Compaction.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaft(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Storage != nil {
		{
			size, err := m.Storage.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaft(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.HeartbeatInterval != nil {
		n3, err3 := github_com_gogo_protobuf_types.StdDurationMarshalTo(*m.HeartbeatInterval, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(*m.HeartbeatInterval):])
		if err3 != nil {
			return 0, err3
		}
		i -= n3
		i = encodeVarintRaft(dAtA, i, uint64(n3))
		i--
		dAtA[i] = 0x12
	}
	if m.ElectionTimeout != nil {
		n4, err4 := github_com_gogo_protobuf_types.StdDurationMarshalTo(*m.ElectionTimeout, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(*m.ElectionTimeout):])
		if err4 != nil {
			return 0, err4
		}
		i -= n4
		i = encodeVarintRaft(dAtA, i, uint64(n4))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StorageSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StorageSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StorageSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FlushOnCommit {
		i--
		if m.FlushOnCommit {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Level != 0 {
		i = encodeVarintRaft(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x18
	}
	if m.SegmentSize != 0 {
		i = encodeVarintRaft(dAtA, i, uint64(m.SegmentSize))
		i--
		dAtA[i] = 0x10
	}
	if m.MaxEntrySize != 0 {
		i = encodeVarintRaft(dAtA, i, uint64(m.MaxEntrySize))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CompactionSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CompactionSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CompactionSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FreeDiskBuffer != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.FreeDiskBuffer))))
		i--
		dAtA[i] = 0x11
	}
	if m.Dynamic {
		i--
		if m.Dynamic {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRaft(dAtA []byte, offset int, v uint64) int {
	offset -= sovRaft(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RaftProtocol) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ElectionTimeout != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdDuration(*m.ElectionTimeout)
		n += 1 + l + sovRaft(uint64(l))
	}
	if m.HeartbeatInterval != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdDuration(*m.HeartbeatInterval)
		n += 1 + l + sovRaft(uint64(l))
	}
	if m.Storage != nil {
		l = m.Storage.Size()
		n += 1 + l + sovRaft(uint64(l))
	}
	if m.Compaction != nil {
		l = m.Compaction.Size()
		n += 1 + l + sovRaft(uint64(l))
	}
	return n
}

func (m *StorageSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxEntrySize != 0 {
		n += 1 + sovRaft(uint64(m.MaxEntrySize))
	}
	if m.SegmentSize != 0 {
		n += 1 + sovRaft(uint64(m.SegmentSize))
	}
	if m.Level != 0 {
		n += 1 + sovRaft(uint64(m.Level))
	}
	if m.FlushOnCommit {
		n += 2
	}
	return n
}

func (m *CompactionSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Dynamic {
		n += 2
	}
	if m.FreeDiskBuffer != 0 {
		n += 9
	}
	return n
}

func sovRaft(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRaft(x uint64) (n int) {
	return sovRaft(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RaftProtocol) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RaftProtocol: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaftProtocol: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ElectionTimeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRaft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ElectionTimeout == nil {
				m.ElectionTimeout = new(time.Duration)
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(m.ElectionTimeout, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeartbeatInterval", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRaft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HeartbeatInterval == nil {
				m.HeartbeatInterval = new(time.Duration)
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(m.HeartbeatInterval, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Storage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRaft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Storage == nil {
				m.Storage = &StorageSpec{}
			}
			if err := m.Storage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Compaction", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRaft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Compaction == nil {
				m.Compaction = &CompactionSpec{}
			}
			if err := m.Compaction.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRaft
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StorageSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StorageSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StorageSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxEntrySize", wireType)
			}
			m.MaxEntrySize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxEntrySize |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SegmentSize", wireType)
			}
			m.SegmentSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SegmentSize |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			m.Level = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Level |= StorageLevel(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FlushOnCommit", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FlushOnCommit = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRaft
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CompactionSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CompactionSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CompactionSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dynamic", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Dynamic = bool(v != 0)
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field FreeDiskBuffer", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.FreeDiskBuffer = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRaft
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRaft(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRaft
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthRaft
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthRaft
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRaft
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRaft(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthRaft
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRaft = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRaft   = fmt.Errorf("proto: integer overflow")
)
