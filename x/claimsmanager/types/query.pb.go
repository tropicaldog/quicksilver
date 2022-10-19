// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: quicksilver/claimsmanager/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	query "github.com/cosmos/cosmos-sdk/types/query"
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

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a524187a5a706bf7, []int{0}
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

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params defines the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a524187a5a706bf7, []int{1}
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

type QueryClaimsRequest struct {
	ChainId    string             `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty" yaml:"chain_id"`
	Address    string             `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Pagination *query.PageRequest `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryClaimsRequest) Reset()         { *m = QueryClaimsRequest{} }
func (m *QueryClaimsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryClaimsRequest) ProtoMessage()    {}
func (*QueryClaimsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a524187a5a706bf7, []int{2}
}
func (m *QueryClaimsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryClaimsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryClaimsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryClaimsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryClaimsRequest.Merge(m, src)
}
func (m *QueryClaimsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryClaimsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryClaimsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryClaimsRequest proto.InternalMessageInfo

func (m *QueryClaimsRequest) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *QueryClaimsRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *QueryClaimsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryClaimsResponse struct {
	Claims     []Claim             `protobuf:"bytes,1,rep,name=claims,proto3" json:"claims"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryClaimsResponse) Reset()         { *m = QueryClaimsResponse{} }
func (m *QueryClaimsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryClaimsResponse) ProtoMessage()    {}
func (*QueryClaimsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a524187a5a706bf7, []int{3}
}
func (m *QueryClaimsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryClaimsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryClaimsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryClaimsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryClaimsResponse.Merge(m, src)
}
func (m *QueryClaimsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryClaimsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryClaimsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryClaimsResponse proto.InternalMessageInfo

func (m *QueryClaimsResponse) GetClaims() []Claim {
	if m != nil {
		return m.Claims
	}
	return nil
}

func (m *QueryClaimsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "quicksilver.claimsmanager.v1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "quicksilver.claimsmanager.v1.QueryParamsResponse")
	proto.RegisterType((*QueryClaimsRequest)(nil), "quicksilver.claimsmanager.v1.QueryClaimsRequest")
	proto.RegisterType((*QueryClaimsResponse)(nil), "quicksilver.claimsmanager.v1.QueryClaimsResponse")
}

func init() {
	proto.RegisterFile("quicksilver/claimsmanager/v1/query.proto", fileDescriptor_a524187a5a706bf7)
}

var fileDescriptor_a524187a5a706bf7 = []byte{
	// 605 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x33, 0xa9, 0x26, 0x3a, 0x3d, 0x14, 0x26, 0x39, 0xc4, 0x50, 0xb6, 0x65, 0x2d, 0x35,
	0x88, 0xdd, 0xe9, 0xc6, 0x83, 0x20, 0x22, 0x36, 0x62, 0x45, 0x50, 0xac, 0x91, 0x0a, 0x7a, 0x09,
	0x93, 0xcd, 0xb0, 0x19, 0xcc, 0xee, 0x6c, 0x76, 0x66, 0x83, 0xa1, 0xf4, 0xe2, 0x27, 0x10, 0x3c,
	0x7a, 0x10, 0x2f, 0x7e, 0x02, 0x2f, 0x5e, 0xc4, 0x63, 0x8f, 0x55, 0x2f, 0x9e, 0x8a, 0x24, 0x7e,
	0x02, 0x3f, 0x81, 0x64, 0x66, 0x82, 0x59, 0x2b, 0x49, 0x7a, 0x51, 0x6f, 0x93, 0x79, 0xff, 0x97,
	0xf7, 0x7b, 0xff, 0xf7, 0x66, 0x61, 0xa5, 0x9b, 0x30, 0xef, 0xa9, 0x60, 0x9d, 0x1e, 0x8d, 0xb1,
	0xd7, 0x21, 0x2c, 0x10, 0x01, 0x09, 0x89, 0x4f, 0x63, 0xdc, 0x73, 0x71, 0x37, 0xa1, 0x71, 0xdf,
	0x89, 0x62, 0x2e, 0x39, 0x5a, 0x9e, 0x50, 0x3a, 0x29, 0xa5, 0xd3, 0x73, 0xcb, 0x45, 0x9f, 0xfb,
	0x5c, 0x09, 0xf1, 0xe8, 0xa4, 0x73, 0xca, 0xcb, 0x3e, 0xe7, 0x7e, 0x87, 0x62, 0x12, 0x31, 0x4c,
	0xc2, 0x90, 0x4b, 0x22, 0x19, 0x0f, 0x85, 0x89, 0x5e, 0xf4, 0xb8, 0x08, 0xb8, 0xc0, 0x4d, 0x22,
	0xa8, 0x2e, 0x85, 0x7b, 0x6e, 0x93, 0x4a, 0xe2, 0xe2, 0x88, 0xf8, 0x2c, 0x54, 0x62, 0xa3, 0x3d,
	0xa7, 0xb5, 0x0d, 0x5d, 0x42, 0xff, 0x30, 0xa1, 0xcd, 0xa9, 0x2d, 0xa4, 0x49, 0x55, 0x86, 0x5d,
	0x84, 0xe8, 0xc1, 0xa8, 0xdc, 0x0e, 0x89, 0x49, 0x20, 0xea, 0xb4, 0x9b, 0x50, 0x21, 0xed, 0xc7,
	0xb0, 0x90, 0xba, 0x15, 0x11, 0x0f, 0x05, 0x45, 0x35, 0x98, 0x8b, 0xd4, 0x4d, 0x09, 0xac, 0x82,
	0xca, 0x62, 0x75, 0xcd, 0x99, 0x66, 0x84, 0xa3, 0xb3, 0x6b, 0xa7, 0x0e, 0x8e, 0x56, 0x32, 0x75,
	0x93, 0x69, 0x7f, 0x04, 0xa6, 0xe2, 0x4d, 0x25, 0x37, 0x15, 0x91, 0x03, 0xcf, 0x78, 0x6d, 0xc2,
	0xc2, 0x06, 0x6b, 0xa9, 0x3f, 0x3f, 0x5b, 0x2b, 0xfc, 0x38, 0x5a, 0x59, 0xea, 0x93, 0xa0, 0x73,
	0xd5, 0x1e, 0x47, 0xec, 0x7a, 0x5e, 0x1d, 0xef, 0xb4, 0x50, 0x15, 0xe6, 0x49, 0xab, 0x15, 0x53,
	0x21, 0x4a, 0x59, 0x25, 0x2f, 0x7d, 0x7e, 0xb7, 0x51, 0x34, 0x66, 0x6c, 0xe9, 0xc8, 0x43, 0x19,
	0xb3, 0xd0, 0xaf, 0x8f, 0x85, 0x68, 0x1b, 0xc2, 0x5f, 0x66, 0x96, 0x16, 0x54, 0x0b, 0xeb, 0x8e,
	0xc9, 0x19, 0x39, 0xef, 0xe8, 0x21, 0x1b, 0xe7, 0x9d, 0x1d, 0xe2, 0x53, 0xc3, 0x57, 0x9f, 0xc8,
	0xb4, 0xdf, 0x00, 0x63, 0xcf, 0xb8, 0x05, 0x63, 0xcf, 0x16, 0xcc, 0x69, 0x0f, 0x4a, 0x60, 0x75,
	0xa1, 0xb2, 0x58, 0x3d, 0x3f, 0xdd, 0x1e, 0x95, 0x3d, 0x76, 0x47, 0x47, 0xd1, 0xed, 0x14, 0x62,
	0x56, 0x21, 0x5e, 0x98, 0x89, 0xa8, 0xeb, 0x4f, 0x32, 0x56, 0x5f, 0xe7, 0xe1, 0x69, 0xc5, 0x88,
	0x5e, 0x01, 0x98, 0xd3, 0x93, 0x40, 0x9b, 0xd3, 0x81, 0x8e, 0x2f, 0x42, 0xd9, 0x3d, 0x41, 0x86,
	0xa6, 0xb0, 0x2f, 0x3d, 0xff, 0xf2, 0xfd, 0x65, 0x76, 0x1d, 0xad, 0xe1, 0xa9, 0xcb, 0xa8, 0xd7,
	0x01, 0xbd, 0x05, 0x30, 0xa7, 0x6d, 0x9c, 0x8b, 0x2e, 0xb5, 0x34, 0x73, 0xd1, 0xa5, 0x67, 0x64,
	0x5f, 0x51, 0x74, 0x2e, 0xc2, 0x78, 0x8e, 0xa7, 0x82, 0xf7, 0xc6, 0x9b, 0xb7, 0x8f, 0x3e, 0x00,
	0xb8, 0x74, 0x97, 0x08, 0x79, 0x2b, 0xe2, 0x5e, 0xfb, 0x6f, 0x12, 0x6f, 0x2b, 0xe2, 0x1b, 0xe8,
	0xfa, 0x0c, 0x3f, 0x63, 0xda, 0x63, 0x3c, 0x11, 0x0d, 0x3a, 0x02, 0x6c, 0x1c, 0x6f, 0xe0, 0x3d,
	0x80, 0x70, 0x57, 0xd0, 0xf8, 0x3f, 0x64, 0x4f, 0x04, 0x8d, 0xf1, 0x9e, 0x79, 0xa6, 0xfb, 0x7f,
	0x30, 0xff, 0x13, 0x80, 0x85, 0x11, 0xfb, 0x3f, 0x19, 0xc0, 0xae, 0x6a, 0xe2, 0x3e, 0xba, 0x77,
	0xa2, 0x26, 0x66, 0xcd, 0xa3, 0xf6, 0xe8, 0x60, 0x60, 0x81, 0xc3, 0x81, 0x05, 0xbe, 0x0d, 0x2c,
	0xf0, 0x62, 0x68, 0x65, 0x0e, 0x87, 0x56, 0xe6, 0xeb, 0xd0, 0xca, 0x3c, 0xb9, 0xe6, 0x33, 0xd9,
	0x4e, 0x9a, 0x8e, 0xc7, 0x03, 0xcc, 0x42, 0x9f, 0x86, 0x09, 0x93, 0xfd, 0x8d, 0x66, 0xc2, 0x3a,
	0xad, 0x14, 0xc2, 0xb3, 0xdf, 0x20, 0x64, 0x3f, 0xa2, 0xa2, 0x99, 0x53, 0x1f, 0xf6, 0xcb, 0x3f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x9a, 0xcb, 0x61, 0xcf, 0x06, 0x00, 0x00,
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
	// Params returns the total set of participation rewards parameters.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	Claims(ctx context.Context, in *QueryClaimsRequest, opts ...grpc.CallOption) (*QueryClaimsResponse, error)
	LastEpochClaims(ctx context.Context, in *QueryClaimsRequest, opts ...grpc.CallOption) (*QueryClaimsResponse, error)
	UserClaims(ctx context.Context, in *QueryClaimsRequest, opts ...grpc.CallOption) (*QueryClaimsResponse, error)
	UserLastEpochClaims(ctx context.Context, in *QueryClaimsRequest, opts ...grpc.CallOption) (*QueryClaimsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/quicksilver.claimsmanager.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Claims(ctx context.Context, in *QueryClaimsRequest, opts ...grpc.CallOption) (*QueryClaimsResponse, error) {
	out := new(QueryClaimsResponse)
	err := c.cc.Invoke(ctx, "/quicksilver.claimsmanager.v1.Query/Claims", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) LastEpochClaims(ctx context.Context, in *QueryClaimsRequest, opts ...grpc.CallOption) (*QueryClaimsResponse, error) {
	out := new(QueryClaimsResponse)
	err := c.cc.Invoke(ctx, "/quicksilver.claimsmanager.v1.Query/LastEpochClaims", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) UserClaims(ctx context.Context, in *QueryClaimsRequest, opts ...grpc.CallOption) (*QueryClaimsResponse, error) {
	out := new(QueryClaimsResponse)
	err := c.cc.Invoke(ctx, "/quicksilver.claimsmanager.v1.Query/UserClaims", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) UserLastEpochClaims(ctx context.Context, in *QueryClaimsRequest, opts ...grpc.CallOption) (*QueryClaimsResponse, error) {
	out := new(QueryClaimsResponse)
	err := c.cc.Invoke(ctx, "/quicksilver.claimsmanager.v1.Query/UserLastEpochClaims", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params returns the total set of participation rewards parameters.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	Claims(context.Context, *QueryClaimsRequest) (*QueryClaimsResponse, error)
	LastEpochClaims(context.Context, *QueryClaimsRequest) (*QueryClaimsResponse, error)
	UserClaims(context.Context, *QueryClaimsRequest) (*QueryClaimsResponse, error)
	UserLastEpochClaims(context.Context, *QueryClaimsRequest) (*QueryClaimsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) Claims(ctx context.Context, req *QueryClaimsRequest) (*QueryClaimsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Claims not implemented")
}
func (*UnimplementedQueryServer) LastEpochClaims(ctx context.Context, req *QueryClaimsRequest) (*QueryClaimsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LastEpochClaims not implemented")
}
func (*UnimplementedQueryServer) UserClaims(ctx context.Context, req *QueryClaimsRequest) (*QueryClaimsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserClaims not implemented")
}
func (*UnimplementedQueryServer) UserLastEpochClaims(ctx context.Context, req *QueryClaimsRequest) (*QueryClaimsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLastEpochClaims not implemented")
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
		FullMethod: "/quicksilver.claimsmanager.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Claims_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryClaimsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Claims(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quicksilver.claimsmanager.v1.Query/Claims",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Claims(ctx, req.(*QueryClaimsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_LastEpochClaims_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryClaimsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LastEpochClaims(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quicksilver.claimsmanager.v1.Query/LastEpochClaims",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).LastEpochClaims(ctx, req.(*QueryClaimsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_UserClaims_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryClaimsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).UserClaims(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quicksilver.claimsmanager.v1.Query/UserClaims",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).UserClaims(ctx, req.(*QueryClaimsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_UserLastEpochClaims_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryClaimsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).UserLastEpochClaims(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quicksilver.claimsmanager.v1.Query/UserLastEpochClaims",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).UserLastEpochClaims(ctx, req.(*QueryClaimsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "quicksilver.claimsmanager.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Claims",
			Handler:    _Query_Claims_Handler,
		},
		{
			MethodName: "LastEpochClaims",
			Handler:    _Query_LastEpochClaims_Handler,
		},
		{
			MethodName: "UserClaims",
			Handler:    _Query_UserClaims_Handler,
		},
		{
			MethodName: "UserLastEpochClaims",
			Handler:    _Query_UserLastEpochClaims_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quicksilver/claimsmanager/v1/query.proto",
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

func (m *QueryClaimsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryClaimsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryClaimsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryClaimsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryClaimsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryClaimsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Claims) > 0 {
		for iNdEx := len(m.Claims) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Claims[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
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

func (m *QueryClaimsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryClaimsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Claims) > 0 {
		for _, e := range m.Claims {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
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
func (m *QueryClaimsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryClaimsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryClaimsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
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
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryClaimsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryClaimsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryClaimsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claims", wireType)
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
			m.Claims = append(m.Claims, Claim{})
			if err := m.Claims[len(m.Claims)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
