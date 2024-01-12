// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: layer/bridge/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e48df680904493de, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e48df680904493de, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

//	message QueryBridgeValidatorsResponse {
//	  repeated BridgeValidator bridge_validators = 1;
//	  pagination.PageResponse page_response = 2;
//	}
type QueryGetEvmValidatorsRequest struct {
}

func (m *QueryGetEvmValidatorsRequest) Reset()         { *m = QueryGetEvmValidatorsRequest{} }
func (m *QueryGetEvmValidatorsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetEvmValidatorsRequest) ProtoMessage()    {}
func (*QueryGetEvmValidatorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e48df680904493de, []int{2}
}
func (m *QueryGetEvmValidatorsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetEvmValidatorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetEvmValidatorsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetEvmValidatorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetEvmValidatorsRequest.Merge(m, src)
}
func (m *QueryGetEvmValidatorsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetEvmValidatorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetEvmValidatorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetEvmValidatorsRequest proto.InternalMessageInfo

type QueryGetEvmValidatorsResponse struct {
	EthAddress []string `protobuf:"bytes,1,rep,name=ethAddress,proto3" json:"ethAddress,omitempty"`
}

func (m *QueryGetEvmValidatorsResponse) Reset()         { *m = QueryGetEvmValidatorsResponse{} }
func (m *QueryGetEvmValidatorsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetEvmValidatorsResponse) ProtoMessage()    {}
func (*QueryGetEvmValidatorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e48df680904493de, []int{3}
}
func (m *QueryGetEvmValidatorsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetEvmValidatorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetEvmValidatorsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetEvmValidatorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetEvmValidatorsResponse.Merge(m, src)
}
func (m *QueryGetEvmValidatorsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetEvmValidatorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetEvmValidatorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetEvmValidatorsResponse proto.InternalMessageInfo

func (m *QueryGetEvmValidatorsResponse) GetEthAddress() []string {
	if m != nil {
		return m.EthAddress
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "layer.bridge.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "layer.bridge.QueryParamsResponse")
	proto.RegisterType((*QueryGetEvmValidatorsRequest)(nil), "layer.bridge.QueryGetEvmValidatorsRequest")
	proto.RegisterType((*QueryGetEvmValidatorsResponse)(nil), "layer.bridge.QueryGetEvmValidatorsResponse")
}

func init() { proto.RegisterFile("layer/bridge/query.proto", fileDescriptor_e48df680904493de) }

var fileDescriptor_e48df680904493de = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4e, 0xe2, 0x40,
	0x1c, 0xc6, 0x5b, 0x76, 0x97, 0x64, 0x67, 0xf7, 0xb0, 0x99, 0xe5, 0xc0, 0x36, 0xec, 0x2c, 0xdb,
	0x83, 0x41, 0xd4, 0x4e, 0xa8, 0x0f, 0x60, 0xc4, 0x18, 0xe3, 0x4d, 0x39, 0x78, 0xf0, 0x42, 0xa6,
	0x74, 0x32, 0x34, 0x69, 0x3b, 0x65, 0x66, 0xda, 0xc8, 0xd5, 0x27, 0x30, 0xf1, 0x01, 0xf4, 0x71,
	0x38, 0x92, 0x78, 0xf1, 0x64, 0x0c, 0xf8, 0x20, 0x86, 0x69, 0x25, 0x20, 0x60, 0xbc, 0x35, 0xff,
	0xef, 0xfb, 0x7f, 0xdf, 0x6f, 0xfe, 0x29, 0xa8, 0x86, 0x64, 0x48, 0x05, 0xf6, 0x44, 0xe0, 0x33,
	0x8a, 0x07, 0x29, 0x15, 0x43, 0x27, 0x11, 0x5c, 0x71, 0xf8, 0x53, 0x2b, 0x4e, 0xae, 0x58, 0x15,
	0xc6, 0x19, 0xd7, 0x02, 0x9e, 0x7d, 0xe5, 0x1e, 0xab, 0xc6, 0x38, 0x67, 0x21, 0xc5, 0x24, 0x09,
	0x30, 0x89, 0x63, 0xae, 0x88, 0x0a, 0x78, 0x2c, 0x0b, 0xb5, 0xd9, 0xe3, 0x32, 0xe2, 0x12, 0x7b,
	0x44, 0x16, 0xd1, 0x38, 0x6b, 0x79, 0x54, 0x91, 0x16, 0x4e, 0x08, 0x0b, 0x62, 0x6d, 0x2e, 0xbc,
	0x7f, 0x96, 0x38, 0x12, 0x22, 0x48, 0x54, 0xc4, 0xd8, 0x15, 0x00, 0xcf, 0x67, 0xcb, 0x67, 0x7a,
	0xd8, 0xa1, 0x83, 0x94, 0x4a, 0x65, 0x9f, 0x82, 0xdf, 0x4b, 0x53, 0x99, 0xf0, 0x58, 0x52, 0xe8,
	0x82, 0x72, 0xbe, 0x5c, 0x35, 0xeb, 0x66, 0xe3, 0x87, 0x5b, 0x71, 0x16, 0x9f, 0xe1, 0xe4, 0xee,
	0xf6, 0xd7, 0xd1, 0xd3, 0x3f, 0xa3, 0x53, 0x38, 0x6d, 0x04, 0x6a, 0x3a, 0xea, 0x84, 0xaa, 0xe3,
	0x2c, 0xba, 0x20, 0x61, 0xe0, 0x13, 0xc5, 0xc5, 0xbc, 0xea, 0x00, 0xfc, 0xdd, 0xa0, 0x17, 0xa5,
	0x08, 0x00, 0xaa, 0xfa, 0x87, 0xbe, 0x2f, 0xa8, 0x9c, 0x15, 0x7f, 0x69, 0x7c, 0xef, 0x2c, 0x4c,
	0xdc, 0xfb, 0x12, 0xf8, 0xa6, 0x13, 0x60, 0x0a, 0xca, 0x39, 0x02, 0xac, 0x2f, 0x83, 0xad, 0xbe,
	0xd0, 0xfa, 0xff, 0x81, 0x23, 0x2f, 0xb6, 0xb7, 0xae, 0x1f, 0x5e, 0x6e, 0x4b, 0x75, 0x88, 0xb0,
	0xa2, 0x61, 0xc8, 0xc5, 0x5e, 0xc0, 0xf1, 0x9a, 0x43, 0xc2, 0x3b, 0x13, 0xfc, 0x7a, 0x4f, 0x0f,
	0x9b, 0x6b, 0xf2, 0x37, 0x9c, 0xc0, 0xda, 0xf9, 0x94, 0xb7, 0xa0, 0x72, 0x35, 0xd5, 0x2e, 0x6c,
	0x6e, 0xa2, 0x62, 0x54, 0x75, 0x69, 0x16, 0x75, 0xb3, 0xf9, 0x6e, 0xfb, 0x68, 0x34, 0x41, 0xe6,
	0x78, 0x82, 0xcc, 0xe7, 0x09, 0x32, 0x6f, 0xa6, 0xc8, 0x18, 0x4f, 0x91, 0xf1, 0x38, 0x45, 0xc6,
	0xe5, 0x36, 0x0b, 0x54, 0x3f, 0xf5, 0x9c, 0x1e, 0x8f, 0x56, 0xf2, 0xae, 0xde, 0x12, 0xd5, 0x30,
	0xa1, 0xd2, 0x2b, 0xeb, 0x1f, 0x66, 0xff, 0x35, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xd0, 0x39, 0x22,
	0xd5, 0x02, 0x00, 0x00,
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
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of GetEvmValidators items.
	GetEvmValidators(ctx context.Context, in *QueryGetEvmValidatorsRequest, opts ...grpc.CallOption) (*QueryGetEvmValidatorsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/layer.bridge.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetEvmValidators(ctx context.Context, in *QueryGetEvmValidatorsRequest, opts ...grpc.CallOption) (*QueryGetEvmValidatorsResponse, error) {
	out := new(QueryGetEvmValidatorsResponse)
	err := c.cc.Invoke(ctx, "/layer.bridge.Query/GetEvmValidators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of GetEvmValidators items.
	GetEvmValidators(context.Context, *QueryGetEvmValidatorsRequest) (*QueryGetEvmValidatorsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) GetEvmValidators(ctx context.Context, req *QueryGetEvmValidatorsRequest) (*QueryGetEvmValidatorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvmValidators not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/layer.bridge.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetEvmValidators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetEvmValidatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetEvmValidators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/layer.bridge.Query/GetEvmValidators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetEvmValidators(ctx, req.(*QueryGetEvmValidatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "layer.bridge.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "GetEvmValidators",
			Handler:    _Query_GetEvmValidators_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "layer/bridge/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetEvmValidatorsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetEvmValidatorsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetEvmValidatorsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryGetEvmValidatorsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetEvmValidatorsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetEvmValidatorsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EthAddress) > 0 {
		for iNdEx := len(m.EthAddress) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.EthAddress[iNdEx])
			copy(dAtA[i:], m.EthAddress[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.EthAddress[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
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
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetEvmValidatorsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryGetEvmValidatorsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.EthAddress) > 0 {
		for _, s := range m.EthAddress {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryGetEvmValidatorsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetEvmValidatorsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetEvmValidatorsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *QueryGetEvmValidatorsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetEvmValidatorsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetEvmValidatorsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthAddress", wireType)
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
			m.EthAddress = append(m.EthAddress, string(dAtA[iNdEx:postIndex]))
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
