// Code generated by protoc-gen-go. DO NOT EDIT.
// source: n13t/idm/v0/idm.proto

package n13t_idm_v0

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN     HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING     HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING HealthCheckResponse_ServingStatus = 2
)

var HealthCheckResponse_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
}

var HealthCheckResponse_ServingStatus_value = map[string]int32{
	"UNKNOWN":     0,
	"SERVING":     1,
	"NOT_SERVING": 2,
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return proto.EnumName(HealthCheckResponse_ServingStatus_name, int32(x))
}

func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c4e677661a4128a4, []int{1, 0}
}

type HealthCheckRequest struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckRequest) Reset()         { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()    {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4e677661a4128a4, []int{0}
}

func (m *HealthCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckRequest.Unmarshal(m, b)
}
func (m *HealthCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckRequest.Marshal(b, m, deterministic)
}
func (m *HealthCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckRequest.Merge(m, src)
}
func (m *HealthCheckRequest) XXX_Size() int {
	return xxx_messageInfo_HealthCheckRequest.Size(m)
}
func (m *HealthCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckRequest proto.InternalMessageInfo

func (m *HealthCheckRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type HealthCheckResponse struct {
	Status               HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=n13t.idm.v0.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *HealthCheckResponse) Reset()         { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()    {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4e677661a4128a4, []int{1}
}

func (m *HealthCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckResponse.Unmarshal(m, b)
}
func (m *HealthCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckResponse.Marshal(b, m, deterministic)
}
func (m *HealthCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckResponse.Merge(m, src)
}
func (m *HealthCheckResponse) XXX_Size() int {
	return xxx_messageInfo_HealthCheckResponse.Size(m)
}
func (m *HealthCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckResponse proto.InternalMessageInfo

func (m *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if m != nil {
		return m.Status
	}
	return HealthCheckResponse_UNKNOWN
}

type ListUsersRequest struct {
	Search               string   `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersRequest) Reset()         { *m = ListUsersRequest{} }
func (m *ListUsersRequest) String() string { return proto.CompactTextString(m) }
func (*ListUsersRequest) ProtoMessage()    {}
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4e677661a4128a4, []int{2}
}

func (m *ListUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersRequest.Unmarshal(m, b)
}
func (m *ListUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersRequest.Marshal(b, m, deterministic)
}
func (m *ListUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersRequest.Merge(m, src)
}
func (m *ListUsersRequest) XXX_Size() int {
	return xxx_messageInfo_ListUsersRequest.Size(m)
}
func (m *ListUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersRequest proto.InternalMessageInfo

func (m *ListUsersRequest) GetSearch() string {
	if m != nil {
		return m.Search
	}
	return ""
}

type SearchRequest struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	PageNumber           int32    `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	ResultPerPage        int32    `protobuf:"varint,3,opt,name=result_per_page,json=resultPerPage,proto3" json:"result_per_page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4e677661a4128a4, []int{3}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchRequest) GetPageNumber() int32 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

func (m *SearchRequest) GetResultPerPage() int32 {
	if m != nil {
		return m.ResultPerPage
	}
	return 0
}

func init() {
	proto.RegisterEnum("n13t.idm.v0.HealthCheckResponse_ServingStatus", HealthCheckResponse_ServingStatus_name, HealthCheckResponse_ServingStatus_value)
	proto.RegisterType((*HealthCheckRequest)(nil), "n13t.idm.v0.HealthCheckRequest")
	proto.RegisterType((*HealthCheckResponse)(nil), "n13t.idm.v0.HealthCheckResponse")
	proto.RegisterType((*ListUsersRequest)(nil), "n13t.idm.v0.ListUsersRequest")
	proto.RegisterType((*SearchRequest)(nil), "n13t.idm.v0.SearchRequest")
}

func init() { proto.RegisterFile("n13t/idm/v0/idm.proto", fileDescriptor_c4e677661a4128a4) }

var fileDescriptor_c4e677661a4128a4 = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x71, 0x20, 0xa9, 0x3a, 0xa1, 0x4d, 0x32, 0x84, 0x10, 0x19, 0x50, 0x23, 0x1f, 0x50,
	0x95, 0x83, 0x1d, 0xda, 0x5b, 0xaf, 0xa5, 0x40, 0x45, 0xeb, 0x46, 0x09, 0xa1, 0x9c, 0x88, 0x36,
	0xc9, 0xe0, 0x58, 0xc4, 0x6b, 0x77, 0x77, 0x1d, 0x29, 0x42, 0x5c, 0x78, 0x04, 0x38, 0xf1, 0x5c,
	0xbc, 0x02, 0x0f, 0x82, 0x76, 0x6d, 0x97, 0xa6, 0x80, 0x10, 0x88, 0x93, 0x3d, 0xff, 0xcc, 0x7c,
	0x33, 0xb3, 0x33, 0x70, 0x97, 0x3f, 0xde, 0x57, 0x5e, 0x38, 0x8b, 0xbc, 0x65, 0x4f, 0x7f, 0xdc,
	0x44, 0xc4, 0x2a, 0xc6, 0xaa, 0x96, 0x5d, 0x6d, 0x2f, 0x7b, 0xf6, 0x83, 0x20, 0x8e, 0x83, 0x05,
	0x79, 0x2c, 0x09, 0x3d, 0xc6, 0x79, 0xac, 0x98, 0x0a, 0x63, 0x2e, 0xb3, 0x50, 0xfb, 0x7e, 0xee,
	0x35, 0xd6, 0x24, 0x7d, 0xeb, 0x51, 0x94, 0xa8, 0x55, 0xee, 0x6c, 0x5d, 0xc5, 0xa7, 0x92, 0x44,
	0xa6, 0x3b, 0x2e, 0xe0, 0x73, 0x62, 0x0b, 0x35, 0x3f, 0x9c, 0xd3, 0xf4, 0xdd, 0x80, 0x2e, 0x52,
	0x92, 0x0a, 0xdb, 0xb0, 0x21, 0x49, 0x2c, 0xc3, 0x29, 0xb5, 0xad, 0x8e, 0xb5, 0xbb, 0x39, 0x28,
	0x4c, 0xe7, 0x8b, 0x05, 0x77, 0xd6, 0x12, 0x64, 0x12, 0x73, 0x49, 0xf8, 0x14, 0x2a, 0x52, 0x31,
	0x95, 0x4a, 0x93, 0xb0, 0xbd, 0xe7, 0xba, 0x57, 0x1a, 0x77, 0x7f, 0x91, 0xe1, 0x0e, 0x35, 0x91,
	0x07, 0x43, 0x93, 0x35, 0xc8, 0xb3, 0x9d, 0x03, 0xd8, 0x5a, 0x73, 0x60, 0x15, 0x36, 0x46, 0xfe,
	0x0b, 0xff, 0xec, 0xdc, 0xaf, 0xdf, 0xd0, 0xc6, 0xf0, 0x68, 0xf0, 0xea, 0xd8, 0x7f, 0x56, 0xb7,
	0xb0, 0x06, 0x55, 0xff, 0xec, 0xe5, 0xb8, 0x10, 0x4a, 0x4e, 0x17, 0xea, 0x27, 0xa1, 0x54, 0x23,
	0x49, 0x42, 0x16, 0x93, 0xb4, 0xa0, 0x22, 0x89, 0x89, 0xe9, 0x3c, 0x1f, 0x24, 0xb7, 0x1c, 0xae,
	0xeb, 0xe8, 0xbf, 0x22, 0xb0, 0x09, 0xe5, 0x8b, 0x94, 0xc4, 0x2a, 0x8f, 0xcb, 0x0c, 0xdc, 0x81,
	0x6a, 0xc2, 0x02, 0x1a, 0xf3, 0x34, 0x9a, 0x90, 0x68, 0x97, 0x3a, 0xd6, 0x6e, 0x79, 0x00, 0x5a,
	0xf2, 0x8d, 0x82, 0x8f, 0xa0, 0x26, 0x48, 0xa6, 0x0b, 0x35, 0x4e, 0x48, 0x8c, 0xb5, 0xa3, 0x7d,
	0xd3, 0x04, 0x6d, 0x65, 0x72, 0x9f, 0x44, 0x9f, 0x05, 0xb4, 0xf7, 0xe9, 0x16, 0xe0, 0xf1, 0x8c,
	0xb8, 0x0a, 0xd5, 0xea, 0x94, 0x71, 0x16, 0x50, 0x44, 0x5c, 0xe1, 0x29, 0xc0, 0xa1, 0x20, 0xa6,
	0x48, 0x37, 0x8d, 0x8d, 0xb5, 0x47, 0xd3, 0x92, 0xdd, 0x72, 0xb3, 0xad, 0xba, 0xc5, 0x56, 0xdd,
	0x23, 0xbd, 0x55, 0xa7, 0xf9, 0xf1, 0xeb, 0xb7, 0xcf, 0xa5, 0x6d, 0x67, 0xb3, 0xd8, 0xa7, 0x3c,
	0xb0, 0xba, 0xf8, 0x1a, 0x6a, 0x97, 0x2f, 0x30, 0x54, 0x82, 0x58, 0x84, 0x0f, 0xd7, 0x98, 0xd7,
	0xdf, 0xc7, 0xfe, 0xb9, 0xa4, 0xd3, 0x30, 0xe8, 0x2a, 0xfe, 0x40, 0xf7, 0x2c, 0x1c, 0x02, 0x8c,
	0x92, 0xd9, 0x3f, 0x34, 0x6a, 0x1b, 0x5a, 0xd3, 0xae, 0x5d, 0xd2, 0xbc, 0xf7, 0x32, 0x9d, 0x7c,
	0xd0, 0xed, 0xf6, 0x01, 0x9e, 0xd0, 0x82, 0xfe, 0x1e, 0x7a, 0xcf, 0x40, 0x1b, 0xdd, 0xeb, 0x50,
	0x7c, 0x03, 0x65, 0x73, 0x65, 0xb8, 0xf3, 0xfb, 0xfb, 0xcb, 0x06, 0xef, 0xfc, 0xe9, 0x40, 0x1d,
	0x34, 0x45, 0x6e, 0x23, 0xe8, 0x22, 0x73, 0x13, 0x80, 0x27, 0x50, 0x3e, 0x67, 0x6a, 0x3a, 0xff,
	0x0f, 0xfc, 0x9e, 0x35, 0xa9, 0x98, 0xb1, 0xf6, 0xbf, 0x07, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x2b,
	0x00, 0x50, 0xfc, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IdentityManagementClient is the client API for IdentityManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IdentityManagementClient interface {
	CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*empty.Empty, error)
	/// the default rpc use for listing users
	ListUsersStream(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (IdentityManagement_ListUsersStreamClient, error)
	UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*empty.Empty, error)
	// https://github.com/grpc/grpc/blob/master/doc/health-checking.md
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	Watch(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (IdentityManagement_WatchClient, error)
}

type identityManagementClient struct {
	cc *grpc.ClientConn
}

func NewIdentityManagementClient(cc *grpc.ClientConn) IdentityManagementClient {
	return &identityManagementClient{cc}
}

func (c *identityManagementClient) CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/n13t.idm.v0.IdentityManagement/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityManagementClient) ListUsersStream(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (IdentityManagement_ListUsersStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_IdentityManagement_serviceDesc.Streams[0], "/n13t.idm.v0.IdentityManagement/ListUsersStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &identityManagementListUsersStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type IdentityManagement_ListUsersStreamClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type identityManagementListUsersStreamClient struct {
	grpc.ClientStream
}

func (x *identityManagementListUsersStreamClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *identityManagementClient) UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/n13t.idm.v0.IdentityManagement/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityManagementClient) DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/n13t.idm.v0.IdentityManagement/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityManagementClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/n13t.idm.v0.IdentityManagement/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityManagementClient) Watch(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (IdentityManagement_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_IdentityManagement_serviceDesc.Streams[1], "/n13t.idm.v0.IdentityManagement/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &identityManagementWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type IdentityManagement_WatchClient interface {
	Recv() (*HealthCheckResponse, error)
	grpc.ClientStream
}

type identityManagementWatchClient struct {
	grpc.ClientStream
}

func (x *identityManagementWatchClient) Recv() (*HealthCheckResponse, error) {
	m := new(HealthCheckResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// IdentityManagementServer is the server API for IdentityManagement service.
type IdentityManagementServer interface {
	CreateUser(context.Context, *User) (*empty.Empty, error)
	/// the default rpc use for listing users
	ListUsersStream(*ListUsersRequest, IdentityManagement_ListUsersStreamServer) error
	UpdateUser(context.Context, *User) (*empty.Empty, error)
	DeleteUser(context.Context, *User) (*empty.Empty, error)
	// https://github.com/grpc/grpc/blob/master/doc/health-checking.md
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	Watch(*HealthCheckRequest, IdentityManagement_WatchServer) error
}

// UnimplementedIdentityManagementServer can be embedded to have forward compatible implementations.
type UnimplementedIdentityManagementServer struct {
}

func (*UnimplementedIdentityManagementServer) CreateUser(ctx context.Context, req *User) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedIdentityManagementServer) ListUsersStream(req *ListUsersRequest, srv IdentityManagement_ListUsersStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ListUsersStream not implemented")
}
func (*UnimplementedIdentityManagementServer) UpdateUser(ctx context.Context, req *User) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedIdentityManagementServer) DeleteUser(ctx context.Context, req *User) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (*UnimplementedIdentityManagementServer) Check(ctx context.Context, req *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (*UnimplementedIdentityManagementServer) Watch(req *HealthCheckRequest, srv IdentityManagement_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}

func RegisterIdentityManagementServer(s *grpc.Server, srv IdentityManagementServer) {
	s.RegisterService(&_IdentityManagement_serviceDesc, srv)
}

func _IdentityManagement_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityManagementServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n13t.idm.v0.IdentityManagement/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityManagementServer).CreateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityManagement_ListUsersStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListUsersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IdentityManagementServer).ListUsersStream(m, &identityManagementListUsersStreamServer{stream})
}

type IdentityManagement_ListUsersStreamServer interface {
	Send(*User) error
	grpc.ServerStream
}

type identityManagementListUsersStreamServer struct {
	grpc.ServerStream
}

func (x *identityManagementListUsersStreamServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func _IdentityManagement_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityManagementServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n13t.idm.v0.IdentityManagement/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityManagementServer).UpdateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityManagement_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityManagementServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n13t.idm.v0.IdentityManagement/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityManagementServer).DeleteUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityManagement_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityManagementServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n13t.idm.v0.IdentityManagement/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityManagementServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityManagement_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HealthCheckRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IdentityManagementServer).Watch(m, &identityManagementWatchServer{stream})
}

type IdentityManagement_WatchServer interface {
	Send(*HealthCheckResponse) error
	grpc.ServerStream
}

type identityManagementWatchServer struct {
	grpc.ServerStream
}

func (x *identityManagementWatchServer) Send(m *HealthCheckResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _IdentityManagement_serviceDesc = grpc.ServiceDesc{
	ServiceName: "n13t.idm.v0.IdentityManagement",
	HandlerType: (*IdentityManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _IdentityManagement_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _IdentityManagement_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _IdentityManagement_DeleteUser_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _IdentityManagement_Check_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListUsersStream",
			Handler:       _IdentityManagement_ListUsersStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Watch",
			Handler:       _IdentityManagement_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "n13t/idm/v0/idm.proto",
}
