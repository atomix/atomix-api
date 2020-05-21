// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/membership/membership.proto

package membership

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

// Cluster identifier
type ClusterId struct {
	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (m *ClusterId) Reset()         { *m = ClusterId{} }
func (m *ClusterId) String() string { return proto.CompactTextString(m) }
func (*ClusterId) ProtoMessage()    {}
func (*ClusterId) Descriptor() ([]byte, []int) {
	return fileDescriptor_907f8caf41e45230, []int{0}
}
func (m *ClusterId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClusterId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClusterId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClusterId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterId.Merge(m, src)
}
func (m *ClusterId) XXX_Size() int {
	return m.Size()
}
func (m *ClusterId) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterId.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterId proto.InternalMessageInfo

func (m *ClusterId) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClusterId) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

// Member identifier
type MemberId struct {
	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (m *MemberId) Reset()         { *m = MemberId{} }
func (m *MemberId) String() string { return proto.CompactTextString(m) }
func (*MemberId) ProtoMessage()    {}
func (*MemberId) Descriptor() ([]byte, []int) {
	return fileDescriptor_907f8caf41e45230, []int{1}
}
func (m *MemberId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MemberId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MemberId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MemberId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberId.Merge(m, src)
}
func (m *MemberId) XXX_Size() int {
	return m.Size()
}
func (m *MemberId) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberId.DiscardUnknown(m)
}

var xxx_messageInfo_MemberId proto.InternalMessageInfo

func (m *MemberId) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MemberId) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

// Member is a membership member
type Member struct {
	ID   MemberId `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Host string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Port int32    `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
}

func (m *Member) Reset()         { *m = Member{} }
func (m *Member) String() string { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()    {}
func (*Member) Descriptor() ([]byte, []int) {
	return fileDescriptor_907f8caf41e45230, []int{2}
}
func (m *Member) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Member) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Member.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Member) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Member.Merge(m, src)
}
func (m *Member) XXX_Size() int {
	return m.Size()
}
func (m *Member) XXX_DiscardUnknown() {
	xxx_messageInfo_Member.DiscardUnknown(m)
}

var xxx_messageInfo_Member proto.InternalMessageInfo

func (m *Member) GetID() MemberId {
	if m != nil {
		return m.ID
	}
	return MemberId{}
}

func (m *Member) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Member) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

// JoinClusterRequest is a request to join a membership
type JoinClusterRequest struct {
	Member    *Member   `protobuf:"bytes,1,opt,name=member,proto3" json:"member,omitempty"`
	ClusterID ClusterId `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id"`
}

func (m *JoinClusterRequest) Reset()         { *m = JoinClusterRequest{} }
func (m *JoinClusterRequest) String() string { return proto.CompactTextString(m) }
func (*JoinClusterRequest) ProtoMessage()    {}
func (*JoinClusterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_907f8caf41e45230, []int{3}
}
func (m *JoinClusterRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JoinClusterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JoinClusterRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JoinClusterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinClusterRequest.Merge(m, src)
}
func (m *JoinClusterRequest) XXX_Size() int {
	return m.Size()
}
func (m *JoinClusterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinClusterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinClusterRequest proto.InternalMessageInfo

func (m *JoinClusterRequest) GetMember() *Member {
	if m != nil {
		return m.Member
	}
	return nil
}

func (m *JoinClusterRequest) GetClusterID() ClusterId {
	if m != nil {
		return m.ClusterID
	}
	return ClusterId{}
}

// JoinClusterResponse is a response to joining a membership
type JoinClusterResponse struct {
	ClusterID ClusterId `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id"`
	Members   []Member  `protobuf:"bytes,2,rep,name=members,proto3" json:"members"`
}

func (m *JoinClusterResponse) Reset()         { *m = JoinClusterResponse{} }
func (m *JoinClusterResponse) String() string { return proto.CompactTextString(m) }
func (*JoinClusterResponse) ProtoMessage()    {}
func (*JoinClusterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_907f8caf41e45230, []int{4}
}
func (m *JoinClusterResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JoinClusterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JoinClusterResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JoinClusterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinClusterResponse.Merge(m, src)
}
func (m *JoinClusterResponse) XXX_Size() int {
	return m.Size()
}
func (m *JoinClusterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinClusterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JoinClusterResponse proto.InternalMessageInfo

func (m *JoinClusterResponse) GetClusterID() ClusterId {
	if m != nil {
		return m.ClusterID
	}
	return ClusterId{}
}

func (m *JoinClusterResponse) GetMembers() []Member {
	if m != nil {
		return m.Members
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterId)(nil), "atomix.membership.ClusterId")
	proto.RegisterType((*MemberId)(nil), "atomix.membership.MemberId")
	proto.RegisterType((*Member)(nil), "atomix.membership.Member")
	proto.RegisterType((*JoinClusterRequest)(nil), "atomix.membership.JoinClusterRequest")
	proto.RegisterType((*JoinClusterResponse)(nil), "atomix.membership.JoinClusterResponse")
}

func init() { proto.RegisterFile("atomix/membership/membership.proto", fileDescriptor_907f8caf41e45230) }

var fileDescriptor_907f8caf41e45230 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xc1, 0x4e, 0xfa, 0x40,
	0x10, 0xc6, 0xbb, 0x85, 0x3f, 0x7f, 0x3b, 0x9c, 0x58, 0x3d, 0x54, 0x24, 0x85, 0x34, 0xd1, 0x70,
	0x2a, 0x8a, 0x07, 0x63, 0xa2, 0x97, 0xca, 0x05, 0x13, 0x3c, 0xac, 0x77, 0x0d, 0x94, 0x0d, 0x6c,
	0x62, 0xd9, 0xda, 0x2d, 0xc6, 0xc7, 0xf0, 0xe6, 0xd5, 0xc7, 0xe1, 0xc8, 0xd1, 0x13, 0x31, 0xe5,
	0x45, 0x4c, 0xb7, 0x4b, 0x41, 0x41, 0x4d, 0xf4, 0xd4, 0x2f, 0x3b, 0xdf, 0x37, 0xbf, 0x99, 0x49,
	0xc1, 0xee, 0x46, 0xdc, 0x67, 0x8f, 0x0d, 0x9f, 0xfa, 0x3d, 0x1a, 0x8a, 0x21, 0x0b, 0x56, 0xa4,
	0x13, 0x84, 0x3c, 0xe2, 0xb8, 0x94, 0x7a, 0x9c, 0x65, 0xa1, 0xbc, 0x33, 0xe0, 0x03, 0x2e, 0xab,
	0x8d, 0x44, 0xa5, 0x46, 0xfb, 0x1c, 0x8c, 0x8b, 0xbb, 0xb1, 0x88, 0x68, 0xd8, 0xee, 0x63, 0x0c,
	0xf9, 0x51, 0xd7, 0xa7, 0x26, 0xaa, 0xa1, 0xba, 0x41, 0xa4, 0xc6, 0x15, 0x30, 0x92, 0xaf, 0x08,
	0xba, 0x1e, 0x35, 0x75, 0x59, 0x58, 0x3e, 0xd8, 0x67, 0xb0, 0xd5, 0x91, 0x88, 0x5f, 0xa5, 0x19,
	0x14, 0xd2, 0x34, 0x3e, 0x01, 0x9d, 0xf5, 0x65, 0xb2, 0xd8, 0xdc, 0x73, 0xd6, 0x86, 0x77, 0x16,
	0x10, 0x17, 0x26, 0xb3, 0xaa, 0x16, 0xcf, 0xaa, 0x7a, 0xbb, 0x45, 0x74, 0x26, 0xa1, 0x43, 0x2e,
	0x22, 0xd5, 0x5b, 0xea, 0xe4, 0x2d, 0xe0, 0x61, 0x64, 0xe6, 0x6a, 0xa8, 0xfe, 0x8f, 0x48, 0x6d,
	0x3f, 0x23, 0xc0, 0x97, 0x9c, 0x8d, 0xd4, 0xb2, 0x84, 0xde, 0x8f, 0xa9, 0x88, 0xf0, 0x11, 0x14,
	0x52, 0x8a, 0x62, 0xef, 0x7e, 0xc9, 0x26, 0xca, 0x88, 0xaf, 0x00, 0xbc, 0xb4, 0xc9, 0x2d, 0xeb,
	0x4b, 0x6e, 0xb1, 0x59, 0xd9, 0x10, 0xcb, 0xce, 0xea, 0x96, 0xd4, 0xcc, 0xd9, 0xa5, 0x5b, 0xc4,
	0xf0, 0x16, 0x55, 0xfb, 0x05, 0xc1, 0xf6, 0x87, 0xc9, 0x44, 0xc0, 0x47, 0x82, 0x7e, 0xe2, 0xa0,
	0xbf, 0x72, 0xf0, 0x29, 0xfc, 0x57, 0x29, 0x53, 0xaf, 0xe5, 0xbe, 0xdd, 0xd5, 0xcd, 0x27, 0x9d,
	0xc8, 0xc2, 0xdf, 0x14, 0x50, 0xea, 0x64, 0x9e, 0x6b, 0x1a, 0x3e, 0x30, 0x8f, 0xe2, 0x1b, 0x28,
	0xae, 0x8c, 0x8d, 0xf7, 0x37, 0x74, 0x5b, 0x3f, 0x78, 0xf9, 0xe0, 0x27, 0x5b, 0xba, 0xfd, 0x21,
	0x72, 0xcd, 0x49, 0x6c, 0xa1, 0x69, 0x6c, 0xa1, 0xb7, 0xd8, 0x42, 0x4f, 0x73, 0x4b, 0x9b, 0xce,
	0x2d, 0xed, 0x75, 0x6e, 0x69, 0xbd, 0x82, 0xfc, 0x75, 0x8f, 0xdf, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x01, 0x68, 0xec, 0x20, 0x09, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MembershipServiceClient is the client API for MembershipService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MembershipServiceClient interface {
	// Joins a member to a cluster
	JoinCluster(ctx context.Context, in *JoinClusterRequest, opts ...grpc.CallOption) (MembershipService_JoinClusterClient, error)
}

type membershipServiceClient struct {
	cc *grpc.ClientConn
}

func NewMembershipServiceClient(cc *grpc.ClientConn) MembershipServiceClient {
	return &membershipServiceClient{cc}
}

func (c *membershipServiceClient) JoinCluster(ctx context.Context, in *JoinClusterRequest, opts ...grpc.CallOption) (MembershipService_JoinClusterClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MembershipService_serviceDesc.Streams[0], "/atomix.membership.MembershipService/JoinCluster", opts...)
	if err != nil {
		return nil, err
	}
	x := &membershipServiceJoinClusterClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MembershipService_JoinClusterClient interface {
	Recv() (*JoinClusterResponse, error)
	grpc.ClientStream
}

type membershipServiceJoinClusterClient struct {
	grpc.ClientStream
}

func (x *membershipServiceJoinClusterClient) Recv() (*JoinClusterResponse, error) {
	m := new(JoinClusterResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MembershipServiceServer is the server API for MembershipService service.
type MembershipServiceServer interface {
	// Joins a member to a cluster
	JoinCluster(*JoinClusterRequest, MembershipService_JoinClusterServer) error
}

// UnimplementedMembershipServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMembershipServiceServer struct {
}

func (*UnimplementedMembershipServiceServer) JoinCluster(req *JoinClusterRequest, srv MembershipService_JoinClusterServer) error {
	return status.Errorf(codes.Unimplemented, "method JoinCluster not implemented")
}

func RegisterMembershipServiceServer(s *grpc.Server, srv MembershipServiceServer) {
	s.RegisterService(&_MembershipService_serviceDesc, srv)
}

func _MembershipService_JoinCluster_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(JoinClusterRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MembershipServiceServer).JoinCluster(m, &membershipServiceJoinClusterServer{stream})
}

type MembershipService_JoinClusterServer interface {
	Send(*JoinClusterResponse) error
	grpc.ServerStream
}

type membershipServiceJoinClusterServer struct {
	grpc.ServerStream
}

func (x *membershipServiceJoinClusterServer) Send(m *JoinClusterResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _MembershipService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "atomix.membership.MembershipService",
	HandlerType: (*MembershipServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "JoinCluster",
			Handler:       _MembershipService_JoinCluster_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "atomix/membership/membership.proto",
}

func (m *ClusterId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterId) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClusterId) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintMembership(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintMembership(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MemberId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MemberId) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MemberId) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintMembership(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintMembership(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Member) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Member) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Member) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Port != 0 {
		i = encodeVarintMembership(dAtA, i, uint64(m.Port))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Host) > 0 {
		i -= len(m.Host)
		copy(dAtA[i:], m.Host)
		i = encodeVarintMembership(dAtA, i, uint64(len(m.Host)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMembership(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *JoinClusterRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JoinClusterRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *JoinClusterRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ClusterID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMembership(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Member != nil {
		{
			size, err := m.Member.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMembership(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *JoinClusterResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JoinClusterResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *JoinClusterResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Members) > 0 {
		for iNdEx := len(m.Members) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Members[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMembership(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.ClusterID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMembership(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintMembership(dAtA []byte, offset int, v uint64) int {
	offset -= sovMembership(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClusterId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMembership(uint64(l))
	}
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovMembership(uint64(l))
	}
	return n
}

func (m *MemberId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMembership(uint64(l))
	}
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovMembership(uint64(l))
	}
	return n
}

func (m *Member) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovMembership(uint64(l))
	l = len(m.Host)
	if l > 0 {
		n += 1 + l + sovMembership(uint64(l))
	}
	if m.Port != 0 {
		n += 1 + sovMembership(uint64(m.Port))
	}
	return n
}

func (m *JoinClusterRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Member != nil {
		l = m.Member.Size()
		n += 1 + l + sovMembership(uint64(l))
	}
	l = m.ClusterID.Size()
	n += 1 + l + sovMembership(uint64(l))
	return n
}

func (m *JoinClusterResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ClusterID.Size()
	n += 1 + l + sovMembership(uint64(l))
	if len(m.Members) > 0 {
		for _, e := range m.Members {
			l = e.Size()
			n += 1 + l + sovMembership(uint64(l))
		}
	}
	return n
}

func sovMembership(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMembership(x uint64) (n int) {
	return sovMembership(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClusterId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMembership
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
			return fmt.Errorf("proto: ClusterId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMembership(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMembership
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMembership
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
func (m *MemberId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMembership
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
			return fmt.Errorf("proto: MemberId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemberId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMembership(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMembership
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMembership
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
func (m *Member) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMembership
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
			return fmt.Errorf("proto: Member: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Member: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
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
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Host", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Host = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMembership(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMembership
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMembership
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
func (m *JoinClusterRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMembership
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
			return fmt.Errorf("proto: JoinClusterRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JoinClusterRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Member", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
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
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Member == nil {
				m.Member = &Member{}
			}
			if err := m.Member.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
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
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClusterID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMembership(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMembership
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMembership
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
func (m *JoinClusterResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMembership
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
			return fmt.Errorf("proto: JoinClusterResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JoinClusterResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
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
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClusterID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMembership
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
				return ErrInvalidLengthMembership
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMembership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, Member{})
			if err := m.Members[len(m.Members)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMembership(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMembership
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMembership
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
func skipMembership(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMembership
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
					return 0, ErrIntOverflowMembership
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMembership
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
				return 0, ErrInvalidLengthMembership
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMembership
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMembership
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMembership        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMembership          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMembership = fmt.Errorf("proto: unexpected end of group")
)