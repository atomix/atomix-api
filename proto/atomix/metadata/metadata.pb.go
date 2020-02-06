// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/metadata/metadata.proto

package metadata

import (
	primitive "atomix/primitive"
	context "context"
	fmt "fmt"
	headers "github.com/atomix/api/proto/atomix/headers"
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

// GetPrimitivesRequest is a request for primitive metadata
type GetPrimitivesRequest struct {
	Header    *headers.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Namespace string                 `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Type      string                 `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (m *GetPrimitivesRequest) Reset()         { *m = GetPrimitivesRequest{} }
func (m *GetPrimitivesRequest) String() string { return proto.CompactTextString(m) }
func (*GetPrimitivesRequest) ProtoMessage()    {}
func (*GetPrimitivesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_878692e84ef30791, []int{0}
}
func (m *GetPrimitivesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetPrimitivesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetPrimitivesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetPrimitivesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPrimitivesRequest.Merge(m, src)
}
func (m *GetPrimitivesRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetPrimitivesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPrimitivesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPrimitivesRequest proto.InternalMessageInfo

func (m *GetPrimitivesRequest) GetHeader() *headers.RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *GetPrimitivesRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GetPrimitivesRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

// GetPrimitivesResponse is a response containing primitive metadata
type GetPrimitivesResponse struct {
	Primitives []*PrimitiveMetadata `protobuf:"bytes,1,rep,name=primitives,proto3" json:"primitives,omitempty"`
}

func (m *GetPrimitivesResponse) Reset()         { *m = GetPrimitivesResponse{} }
func (m *GetPrimitivesResponse) String() string { return proto.CompactTextString(m) }
func (*GetPrimitivesResponse) ProtoMessage()    {}
func (*GetPrimitivesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_878692e84ef30791, []int{1}
}
func (m *GetPrimitivesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetPrimitivesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetPrimitivesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetPrimitivesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPrimitivesResponse.Merge(m, src)
}
func (m *GetPrimitivesResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetPrimitivesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPrimitivesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPrimitivesResponse proto.InternalMessageInfo

func (m *GetPrimitivesResponse) GetPrimitives() []*PrimitiveMetadata {
	if m != nil {
		return m.Primitives
	}
	return nil
}

// PrimitiveMetadata indicates the type and name of a primitive
type PrimitiveMetadata struct {
	Type string          `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Name *primitive.Name `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *PrimitiveMetadata) Reset()         { *m = PrimitiveMetadata{} }
func (m *PrimitiveMetadata) String() string { return proto.CompactTextString(m) }
func (*PrimitiveMetadata) ProtoMessage()    {}
func (*PrimitiveMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_878692e84ef30791, []int{2}
}
func (m *PrimitiveMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PrimitiveMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PrimitiveMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PrimitiveMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimitiveMetadata.Merge(m, src)
}
func (m *PrimitiveMetadata) XXX_Size() int {
	return m.Size()
}
func (m *PrimitiveMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimitiveMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_PrimitiveMetadata proto.InternalMessageInfo

func (m *PrimitiveMetadata) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PrimitiveMetadata) GetName() *primitive.Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func init() {
	proto.RegisterType((*GetPrimitivesRequest)(nil), "atomix.metadata.GetPrimitivesRequest")
	proto.RegisterType((*GetPrimitivesResponse)(nil), "atomix.metadata.GetPrimitivesResponse")
	proto.RegisterType((*PrimitiveMetadata)(nil), "atomix.metadata.PrimitiveMetadata")
}

func init() { proto.RegisterFile("atomix/metadata/metadata.proto", fileDescriptor_878692e84ef30791) }

var fileDescriptor_878692e84ef30791 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xb3, 0xb6, 0x14, 0x3a, 0x41, 0x8a, 0x8b, 0x4a, 0x08, 0x75, 0x09, 0x01, 0x25, 0x78,
	0x88, 0x10, 0xf1, 0x05, 0x7a, 0xd1, 0x8b, 0x22, 0xe9, 0x51, 0x10, 0xd6, 0x76, 0xc0, 0x1c, 0xd2,
	0xa4, 0xd9, 0xb5, 0xe8, 0xc9, 0x57, 0xf0, 0xb1, 0x3c, 0xf6, 0xe8, 0x51, 0x92, 0x17, 0x11, 0xf7,
	0x4f, 0xa2, 0xad, 0xe0, 0x29, 0x93, 0x6f, 0xbe, 0x99, 0xf9, 0xcd, 0x2c, 0x30, 0x2e, 0x8b, 0x3c,
	0x7b, 0x3e, 0xcb, 0x51, 0xf2, 0x39, 0x97, 0xbc, 0x0d, 0xe2, 0xb2, 0x2a, 0x64, 0x41, 0x47, 0x3a,
	0x1f, 0x5b, 0xd9, 0x1f, 0x9b, 0x82, 0x47, 0xe4, 0x73, 0xac, 0x84, 0xfd, 0x6a, 0xbb, 0x1f, 0x98,
	0x6c, 0x59, 0x65, 0x79, 0x26, 0xb3, 0x15, 0x76, 0x91, 0x76, 0x84, 0xaf, 0xb0, 0x7f, 0x89, 0xf2,
	0xd6, 0xaa, 0x22, 0xc5, 0xe5, 0x13, 0x0a, 0x49, 0x2f, 0x60, 0xa0, 0x5b, 0x79, 0x24, 0x20, 0x91,
	0x9b, 0x1c, 0xc5, 0x66, 0xb2, 0x1d, 0x60, 0x8c, 0x57, 0xea, 0x37, 0x35, 0x66, 0x3a, 0x86, 0xe1,
	0x82, 0xe7, 0x28, 0x4a, 0x3e, 0x43, 0x6f, 0x27, 0x20, 0xd1, 0x30, 0xed, 0x04, 0x4a, 0xa1, 0x2f,
	0x5f, 0x4a, 0xf4, 0x7a, 0x2a, 0xa1, 0xe2, 0xf0, 0x0e, 0x0e, 0x36, 0x00, 0x44, 0x59, 0x2c, 0x04,
	0xd2, 0x09, 0x40, 0x0b, 0x2b, 0x3c, 0x12, 0xf4, 0x22, 0x37, 0x09, 0xe3, 0x8d, 0xfd, 0xe3, 0xb6,
	0xf0, 0xda, 0x28, 0xe9, 0x8f, 0xaa, 0x70, 0x0a, 0x7b, 0x5b, 0x86, 0x96, 0x82, 0x74, 0x14, 0xf4,
	0x14, 0xfa, 0xdf, 0x98, 0x0a, 0xd9, 0x4d, 0x0e, 0xed, 0x98, 0xee, 0x5a, 0x37, 0x3c, 0xc7, 0x54,
	0x79, 0x92, 0x25, 0x8c, 0x6c, 0xaf, 0x29, 0x56, 0xab, 0x6c, 0x86, 0xf4, 0x1e, 0x76, 0x7f, 0x2d,
	0x41, 0x8f, 0xb7, 0x40, 0xff, 0xba, 0xb2, 0x7f, 0xf2, 0x9f, 0x4d, 0xdf, 0x62, 0xe2, 0xbd, 0xd7,
	0x8c, 0xac, 0x6b, 0x46, 0x3e, 0x6b, 0x46, 0xde, 0x1a, 0xe6, 0xac, 0x1b, 0xe6, 0x7c, 0x34, 0xcc,
	0x79, 0x18, 0xa8, 0x67, 0x3c, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x23, 0x53, 0x2a, 0xb7, 0x39,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetadataServiceClient is the client API for MetadataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetadataServiceClient interface {
	// GetPrimitives returns a list of primitives in a partition
	GetPrimitives(ctx context.Context, in *GetPrimitivesRequest, opts ...grpc.CallOption) (*GetPrimitivesResponse, error)
}

type metadataServiceClient struct {
	cc *grpc.ClientConn
}

func NewMetadataServiceClient(cc *grpc.ClientConn) MetadataServiceClient {
	return &metadataServiceClient{cc}
}

func (c *metadataServiceClient) GetPrimitives(ctx context.Context, in *GetPrimitivesRequest, opts ...grpc.CallOption) (*GetPrimitivesResponse, error) {
	out := new(GetPrimitivesResponse)
	err := c.cc.Invoke(ctx, "/atomix.metadata.MetadataService/GetPrimitives", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetadataServiceServer is the server API for MetadataService service.
type MetadataServiceServer interface {
	// GetPrimitives returns a list of primitives in a partition
	GetPrimitives(context.Context, *GetPrimitivesRequest) (*GetPrimitivesResponse, error)
}

// UnimplementedMetadataServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMetadataServiceServer struct {
}

func (*UnimplementedMetadataServiceServer) GetPrimitives(ctx context.Context, req *GetPrimitivesRequest) (*GetPrimitivesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrimitives not implemented")
}

func RegisterMetadataServiceServer(s *grpc.Server, srv MetadataServiceServer) {
	s.RegisterService(&_MetadataService_serviceDesc, srv)
}

func _MetadataService_GetPrimitives_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrimitivesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).GetPrimitives(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.metadata.MetadataService/GetPrimitives",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).GetPrimitives(ctx, req.(*GetPrimitivesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetadataService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "atomix.metadata.MetadataService",
	HandlerType: (*MetadataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPrimitives",
			Handler:    _MetadataService_GetPrimitives_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "atomix/metadata/metadata.proto",
}

func (m *GetPrimitivesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetPrimitivesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetPrimitivesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0x12
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMetadata(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetPrimitivesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetPrimitivesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetPrimitivesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Primitives) > 0 {
		for iNdEx := len(m.Primitives) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Primitives[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMetadata(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *PrimitiveMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrimitiveMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PrimitiveMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Name != nil {
		{
			size, err := m.Name.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMetadata(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMetadata(dAtA []byte, offset int, v uint64) int {
	offset -= sovMetadata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetPrimitivesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	return n
}

func (m *GetPrimitivesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Primitives) > 0 {
		for _, e := range m.Primitives {
			l = e.Size()
			n += 1 + l + sovMetadata(uint64(l))
		}
	}
	return n
}

func (m *PrimitiveMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	if m.Name != nil {
		l = m.Name.Size()
		n += 1 + l + sovMetadata(uint64(l))
	}
	return n
}

func sovMetadata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMetadata(x uint64) (n int) {
	return sovMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetPrimitivesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
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
			return fmt.Errorf("proto: GetPrimitivesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetPrimitivesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &headers.RequestHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetadata
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMetadata
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
func (m *GetPrimitivesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
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
			return fmt.Errorf("proto: GetPrimitivesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetPrimitivesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Primitives", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Primitives = append(m.Primitives, &PrimitiveMetadata{})
			if err := m.Primitives[len(m.Primitives)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetadata
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMetadata
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
func (m *PrimitiveMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
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
			return fmt.Errorf("proto: PrimitiveMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrimitiveMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Name == nil {
				m.Name = &primitive.Name{}
			}
			if err := m.Name.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetadata
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMetadata
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
func skipMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetadata
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
					return 0, ErrIntOverflowMetadata
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
					return 0, ErrIntOverflowMetadata
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
				return 0, ErrInvalidLengthMetadata
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMetadata
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMetadata
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMetadata        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetadata          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMetadata = fmt.Errorf("proto: unexpected end of group")
)
