// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: teritori/airdrop/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type QueryAllocationRequest struct {
	// address is the address to query allocation for.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryAllocationRequest) Reset()         { *m = QueryAllocationRequest{} }
func (m *QueryAllocationRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllocationRequest) ProtoMessage()    {}
func (*QueryAllocationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_efe03a7078585dc1, []int{0}
}
func (m *QueryAllocationRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllocationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllocationRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllocationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllocationRequest.Merge(m, src)
}
func (m *QueryAllocationRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllocationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllocationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllocationRequest proto.InternalMessageInfo

type QueryAllocationResponse struct {
	Allocation *AirdropAllocation `protobuf:"bytes,1,opt,name=allocation,proto3" json:"allocation,omitempty"`
}

func (m *QueryAllocationResponse) Reset()         { *m = QueryAllocationResponse{} }
func (m *QueryAllocationResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllocationResponse) ProtoMessage()    {}
func (*QueryAllocationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_efe03a7078585dc1, []int{1}
}
func (m *QueryAllocationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllocationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllocationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllocationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllocationResponse.Merge(m, src)
}
func (m *QueryAllocationResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllocationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllocationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllocationResponse proto.InternalMessageInfo

func (m *QueryAllocationResponse) GetAllocation() *AirdropAllocation {
	if m != nil {
		return m.Allocation
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryAllocationRequest)(nil), "teritori.airdrop.v1beta1.QueryAllocationRequest")
	proto.RegisterType((*QueryAllocationResponse)(nil), "teritori.airdrop.v1beta1.QueryAllocationResponse")
}

func init() {
	proto.RegisterFile("teritori/airdrop/v1beta1/query.proto", fileDescriptor_efe03a7078585dc1)
}

var fileDescriptor_efe03a7078585dc1 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x29, 0x49, 0x2d, 0xca,
	0x2c, 0xc9, 0x2f, 0xca, 0xd4, 0x4f, 0xcc, 0x2c, 0x4a, 0x29, 0xca, 0x2f, 0xd0, 0x2f, 0x33, 0x4c,
	0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x92, 0x80, 0xa9, 0xd2, 0x83, 0xaa, 0xd2, 0x83, 0xaa, 0x92, 0x12, 0x49, 0xcf, 0x4f, 0xcf,
	0x07, 0x2b, 0xd2, 0x07, 0xb1, 0x20, 0xea, 0xa5, 0x64, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5,
	0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0xa1,
	0xb2, 0x72, 0xc9, 0xf9, 0xc5, 0xb9, 0xf9, 0xc5, 0xfa, 0x49, 0x89, 0xc5, 0xa9, 0x70, 0xeb, 0x92,
	0xf3, 0x33, 0xf3, 0xa0, 0xf2, 0x9a, 0x38, 0xdd, 0x94, 0x98, 0x93, 0x93, 0x9f, 0x0c, 0x36, 0x0b,
	0xa2, 0x54, 0xc9, 0x86, 0x4b, 0x2c, 0x10, 0xe4, 0x4e, 0x47, 0xb8, 0x44, 0x50, 0x6a, 0x61, 0x69,
	0x6a, 0x71, 0x89, 0x90, 0x04, 0x17, 0x7b, 0x62, 0x4a, 0x4a, 0x51, 0x6a, 0x71, 0xb1, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0x6b, 0xc5, 0xd1, 0xb1, 0x40, 0x9e, 0xe1, 0xc5, 0x02, 0x79,
	0x06, 0xa5, 0x34, 0x2e, 0x71, 0x0c, 0xdd, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0xde, 0x5c,
	0x5c, 0x08, 0xcb, 0xc0, 0x26, 0x70, 0x1b, 0x69, 0xeb, 0xe1, 0x0a, 0x06, 0x3d, 0x47, 0x08, 0x1f,
	0xc9, 0x20, 0x24, 0xed, 0x46, 0x9b, 0x19, 0xb9, 0x58, 0xc1, 0x16, 0x09, 0xad, 0x64, 0xe4, 0xe2,
	0x42, 0x28, 0x12, 0x32, 0xc0, 0x6d, 0x22, 0x76, 0x6f, 0x49, 0x19, 0x92, 0xa0, 0x03, 0xe2, 0x15,
	0x25, 0xb3, 0xa6, 0xcb, 0x4f, 0x26, 0x33, 0x19, 0x08, 0xe9, 0xe9, 0x13, 0x11, 0xae, 0xfa, 0xd5,
	0xd0, 0x60, 0xaa, 0x75, 0xf2, 0x3a, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f,
	0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28,
	0x83, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0xbf, 0x88, 0x90, 0x00,
	0xff, 0x00, 0xb8, 0xd1, 0xba, 0xc9, 0x19, 0x89, 0x99, 0x79, 0xfa, 0x15, 0x70, 0x2b, 0x4a, 0x2a,
	0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0xd1, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x82,
	0xd2, 0x08, 0x6f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	Allocation(ctx context.Context, in *QueryAllocationRequest, opts ...grpc.CallOption) (*QueryAllocationResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Allocation(ctx context.Context, in *QueryAllocationRequest, opts ...grpc.CallOption) (*QueryAllocationResponse, error) {
	out := new(QueryAllocationResponse)
	err := c.cc.Invoke(ctx, "/teritori.airdrop.v1beta1.Query/Allocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Allocation(context.Context, *QueryAllocationRequest) (*QueryAllocationResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Allocation(ctx context.Context, req *QueryAllocationRequest) (*QueryAllocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Allocation not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Allocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Allocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/teritori.airdrop.v1beta1.Query/Allocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Allocation(ctx, req.(*QueryAllocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "teritori.airdrop.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Allocation",
			Handler:    _Query_Allocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teritori/airdrop/v1beta1/query.proto",
}

func (m *QueryAllocationRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllocationRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllocationRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllocationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllocationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllocationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Allocation != nil {
		{
			size, err := m.Allocation.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryAllocationRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllocationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Allocation != nil {
		l = m.Allocation.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryAllocationRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllocationRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllocationRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllocationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllocationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllocationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allocation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Allocation == nil {
				m.Allocation = &AirdropAllocation{}
			}
			if err := m.Allocation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
