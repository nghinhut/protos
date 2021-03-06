// Code generated by protoc-gen-go. DO NOT EDIT.
// source: n13t/idm/v0/search.proto

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

type SearchRequest struct {
	Index                string   `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Query                string   `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	Limit                int32    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32    `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_81f6c60b662e27bb, []int{0}
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

func (m *SearchRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *SearchRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *SearchRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type SearchResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_81f6c60b662e27bb, []int{1}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *SearchResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*SearchRequest)(nil), "n13t.idm.v0.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "n13t.idm.v0.SearchResponse")
}

func init() {
	proto.RegisterFile("n13t/idm/v0/search.proto", fileDescriptor_81f6c60b662e27bb)
}

var fileDescriptor_81f6c60b662e27bb = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xd1, 0x6a, 0x83, 0x30,
	0x14, 0x86, 0xc9, 0x9c, 0xc2, 0x22, 0x1b, 0x5b, 0x18, 0x12, 0x76, 0x25, 0xde, 0xcc, 0xab, 0xe8,
	0xe6, 0x63, 0xf4, 0x2e, 0xa5, 0x0f, 0xd0, 0xd6, 0x23, 0x0d, 0x54, 0xa3, 0x39, 0x51, 0xda, 0xb7,
	0x2f, 0x49, 0x2c, 0x78, 0x97, 0xef, 0x3b, 0x7f, 0xf8, 0x7f, 0xca, 0x87, 0xbf, 0xc6, 0x56, 0xaa,
	0xed, 0xab, 0xa5, 0xae, 0x10, 0x8e, 0xe6, 0x7c, 0x11, 0xa3, 0xd1, 0x56, 0xb3, 0xd4, 0x5d, 0x84,
	0x6a, 0x7b, 0xb1, 0xd4, 0x3f, 0xd9, 0x36, 0x36, 0x23, 0x98, 0x10, 0x2a, 0x14, 0x7d, 0xdf, 0xfb,
	0x4f, 0x12, 0xa6, 0x19, 0xd0, 0xb2, 0x6f, 0x1a, 0xab, 0xa1, 0x85, 0x1b, 0x27, 0x39, 0x29, 0xdf,
	0x64, 0x00, 0x67, 0xa7, 0x19, 0xcc, 0x9d, 0xbf, 0x04, 0xeb, 0xc1, 0xd9, 0xab, 0xea, 0x95, 0xe5,
	0x51, 0x4e, 0xca, 0x58, 0x06, 0x60, 0x19, 0x4d, 0x74, 0xd7, 0x21, 0x58, 0xfe, 0xea, 0xf5, 0x4a,
	0xc5, 0x8e, 0x7e, 0x3c, 0xab, 0x70, 0xd4, 0x03, 0x02, 0xfb, 0xa5, 0xb1, 0x9b, 0x82, 0x9c, 0xe4,
	0x51, 0x99, 0xfe, 0x7f, 0x89, 0xcd, 0x62, 0x71, 0x40, 0x30, 0x32, 0xdc, 0xd9, 0x27, 0x8d, 0xc0,
	0x98, 0xb5, 0xdc, 0x3d, 0x4f, 0x89, 0x9f, 0xdf, 0x3c, 0x02, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xe5,
	0x4d, 0xb8, 0xff, 0x00, 0x00, 0x00,
}
