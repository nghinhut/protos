// Code generated by protoc-gen-go. DO NOT EDIT.
// source: n13t/idm/v0/list.proto

package n13t_idm_v0

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ListUsersRequest struct {
	// Optional. The maximum number of Users to return in the response.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. A pagination token returned from a previous call to `ListUsers`
	// that indicates where this listing should continue from.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersRequest) Reset()         { *m = ListUsersRequest{} }
func (m *ListUsersRequest) String() string { return proto.CompactTextString(m) }
func (*ListUsersRequest) ProtoMessage()    {}
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_42dabac6d506bfca, []int{0}
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

func (m *ListUsersRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListUsersRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListUsersResponse struct {
	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	// A pagination token returned from a previous call to `ListUsers`
	// that indicates from where listing should continue.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	Err                  string   `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersResponse) Reset()         { *m = ListUsersResponse{} }
func (m *ListUsersResponse) String() string { return proto.CompactTextString(m) }
func (*ListUsersResponse) ProtoMessage()    {}
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_42dabac6d506bfca, []int{1}
}

func (m *ListUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersResponse.Unmarshal(m, b)
}
func (m *ListUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersResponse.Marshal(b, m, deterministic)
}
func (m *ListUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersResponse.Merge(m, src)
}
func (m *ListUsersResponse) XXX_Size() int {
	return xxx_messageInfo_ListUsersResponse.Size(m)
}
func (m *ListUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersResponse proto.InternalMessageInfo

func (m *ListUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *ListUsersResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func (m *ListUsersResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*ListUsersRequest)(nil), "n13t.idm.v0.ListUsersRequest")
	proto.RegisterType((*ListUsersResponse)(nil), "n13t.idm.v0.ListUsersResponse")
}

func init() {
	proto.RegisterFile("n13t/idm/v0/list.proto", fileDescriptor_42dabac6d506bfca)
}

var fileDescriptor_42dabac6d506bfca = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xd1, 0x4b, 0x85, 0x30,
	0x14, 0x87, 0x51, 0x31, 0xf2, 0x48, 0xa4, 0x7b, 0x08, 0x29, 0x02, 0xf1, 0xa1, 0x7c, 0x9a, 0x96,
	0xff, 0x46, 0x44, 0xac, 0x7a, 0x96, 0xc2, 0x43, 0x8c, 0x72, 0xb3, 0x9d, 0x29, 0xe1, 0x5f, 0x1f,
	0x9b, 0xf7, 0x82, 0xf7, 0x6d, 0x7c, 0xbf, 0x8f, 0x7d, 0x1b, 0x5c, 0xa9, 0x87, 0xce, 0x36, 0x72,
	0x18, 0x9b, 0xa5, 0x6d, 0x7e, 0x24, 0x59, 0x3e, 0x19, 0x6d, 0x35, 0x4b, 0x1d, 0xe7, 0x72, 0x18,
	0xf9, 0xd2, 0x5e, 0x9f, 0x48, 0x33, 0xa1, 0xd9, 0xa4, 0xea, 0x19, 0xb2, 0x27, 0x49, 0xf6, 0x9d,
	0xd0, 0x90, 0xc0, 0xdf, 0x19, 0xc9, 0xb2, 0x1b, 0x48, 0xa6, 0x8f, 0x2f, 0xec, 0x49, 0xae, 0x58,
	0x84, 0x65, 0x50, 0xc7, 0xe2, 0xdc, 0x81, 0x57, 0xb9, 0x22, 0xbb, 0x05, 0xf0, 0xa3, 0xd5, 0xdf,
	0xa8, 0x8a, 0xa8, 0x0c, 0xea, 0x44, 0x78, 0xfd, 0xcd, 0x81, 0x6a, 0x81, 0x7c, 0x77, 0x1f, 0x4d,
	0x5a, 0x11, 0xb2, 0x7b, 0x88, 0x5d, 0x92, 0x8a, 0xa0, 0x8c, 0xea, 0xf4, 0x31, 0xe7, 0xbb, 0x97,
	0x71, 0xa7, 0x8a, 0x6d, 0x67, 0x77, 0x70, 0xa9, 0xf0, 0xcf, 0xf6, 0xbb, 0x42, 0xe8, 0x0b, 0x17,
	0x0e, 0xbf, 0x1c, 0x2b, 0x2c, 0x83, 0x08, 0x8d, 0x39, 0xd4, 0xdd, 0xf1, 0xf3, 0xcc, 0x7f, 0xa7,
	0xfb, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x97, 0xc4, 0x17, 0x47, 0x0d, 0x01, 0x00, 0x00,
}
