// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/rpc/rpc.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type RequestPing struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestPing) Reset()         { *m = RequestPing{} }
func (m *RequestPing) String() string { return proto.CompactTextString(m) }
func (*RequestPing) ProtoMessage()    {}
func (*RequestPing) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd1e66095a603e73, []int{0}
}

func (m *RequestPing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestPing.Unmarshal(m, b)
}
func (m *RequestPing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestPing.Marshal(b, m, deterministic)
}
func (m *RequestPing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestPing.Merge(m, src)
}
func (m *RequestPing) XXX_Size() int {
	return xxx_messageInfo_RequestPing.Size(m)
}
func (m *RequestPing) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestPing.DiscardUnknown(m)
}

var xxx_messageInfo_RequestPing proto.InternalMessageInfo

type ResponsePong struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponsePong) Reset()         { *m = ResponsePong{} }
func (m *ResponsePong) String() string { return proto.CompactTextString(m) }
func (*ResponsePong) ProtoMessage()    {}
func (*ResponsePong) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd1e66095a603e73, []int{1}
}

func (m *ResponsePong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponsePong.Unmarshal(m, b)
}
func (m *ResponsePong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponsePong.Marshal(b, m, deterministic)
}
func (m *ResponsePong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponsePong.Merge(m, src)
}
func (m *ResponsePong) XXX_Size() int {
	return xxx_messageInfo_ResponsePong.Size(m)
}
func (m *ResponsePong) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponsePong.DiscardUnknown(m)
}

var xxx_messageInfo_ResponsePong proto.InternalMessageInfo

type RequestMemberList struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestMemberList) Reset()         { *m = RequestMemberList{} }
func (m *RequestMemberList) String() string { return proto.CompactTextString(m) }
func (*RequestMemberList) ProtoMessage()    {}
func (*RequestMemberList) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd1e66095a603e73, []int{2}
}

func (m *RequestMemberList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestMemberList.Unmarshal(m, b)
}
func (m *RequestMemberList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestMemberList.Marshal(b, m, deterministic)
}
func (m *RequestMemberList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestMemberList.Merge(m, src)
}
func (m *RequestMemberList) XXX_Size() int {
	return xxx_messageInfo_RequestMemberList.Size(m)
}
func (m *RequestMemberList) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestMemberList.DiscardUnknown(m)
}

var xxx_messageInfo_RequestMemberList proto.InternalMessageInfo

type ResponseMemberList struct {
	Items                []*ClusterMember `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ResponseMemberList) Reset()         { *m = ResponseMemberList{} }
func (m *ResponseMemberList) String() string { return proto.CompactTextString(m) }
func (*ResponseMemberList) ProtoMessage()    {}
func (*ResponseMemberList) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd1e66095a603e73, []int{3}
}

func (m *ResponseMemberList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseMemberList.Unmarshal(m, b)
}
func (m *ResponseMemberList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseMemberList.Marshal(b, m, deterministic)
}
func (m *ResponseMemberList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseMemberList.Merge(m, src)
}
func (m *ResponseMemberList) XXX_Size() int {
	return xxx_messageInfo_ResponseMemberList.Size(m)
}
func (m *ResponseMemberList) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseMemberList.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseMemberList proto.InternalMessageInfo

func (m *ResponseMemberList) GetItems() []*ClusterMember {
	if m != nil {
		return m.Items
	}
	return nil
}

type ClusterMember struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterMember) Reset()         { *m = ClusterMember{} }
func (m *ClusterMember) String() string { return proto.CompactTextString(m) }
func (*ClusterMember) ProtoMessage()    {}
func (*ClusterMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd1e66095a603e73, []int{4}
}

func (m *ClusterMember) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterMember.Unmarshal(m, b)
}
func (m *ClusterMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterMember.Marshal(b, m, deterministic)
}
func (m *ClusterMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterMember.Merge(m, src)
}
func (m *ClusterMember) XXX_Size() int {
	return xxx_messageInfo_ClusterMember.Size(m)
}
func (m *ClusterMember) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterMember.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterMember proto.InternalMessageInfo

func (m *ClusterMember) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RequestUserList struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestUserList) Reset()         { *m = RequestUserList{} }
func (m *RequestUserList) String() string { return proto.CompactTextString(m) }
func (*RequestUserList) ProtoMessage()    {}
func (*RequestUserList) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd1e66095a603e73, []int{5}
}

func (m *RequestUserList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestUserList.Unmarshal(m, b)
}
func (m *RequestUserList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestUserList.Marshal(b, m, deterministic)
}
func (m *RequestUserList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestUserList.Merge(m, src)
}
func (m *RequestUserList) XXX_Size() int {
	return xxx_messageInfo_RequestUserList.Size(m)
}
func (m *RequestUserList) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestUserList.DiscardUnknown(m)
}

var xxx_messageInfo_RequestUserList proto.InternalMessageInfo

type ResponseUserList struct {
	Items                []*UserItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ResponseUserList) Reset()         { *m = ResponseUserList{} }
func (m *ResponseUserList) String() string { return proto.CompactTextString(m) }
func (*ResponseUserList) ProtoMessage()    {}
func (*ResponseUserList) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd1e66095a603e73, []int{6}
}

func (m *ResponseUserList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseUserList.Unmarshal(m, b)
}
func (m *ResponseUserList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseUserList.Marshal(b, m, deterministic)
}
func (m *ResponseUserList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseUserList.Merge(m, src)
}
func (m *ResponseUserList) XXX_Size() int {
	return xxx_messageInfo_ResponseUserList.Size(m)
}
func (m *ResponseUserList) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseUserList.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseUserList proto.InternalMessageInfo

func (m *ResponseUserList) GetItems() []*UserItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type UserItem struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Roles                []string `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserItem) Reset()         { *m = UserItem{} }
func (m *UserItem) String() string { return proto.CompactTextString(m) }
func (*UserItem) ProtoMessage()    {}
func (*UserItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd1e66095a603e73, []int{7}
}

func (m *UserItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserItem.Unmarshal(m, b)
}
func (m *UserItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserItem.Marshal(b, m, deterministic)
}
func (m *UserItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserItem.Merge(m, src)
}
func (m *UserItem) XXX_Size() int {
	return xxx_messageInfo_UserItem.Size(m)
}
func (m *UserItem) XXX_DiscardUnknown() {
	xxx_messageInfo_UserItem.DiscardUnknown(m)
}

var xxx_messageInfo_UserItem proto.InternalMessageInfo

func (m *UserItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserItem) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

type RequestUserAdd struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Role                 string   `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestUserAdd) Reset()         { *m = RequestUserAdd{} }
func (m *RequestUserAdd) String() string { return proto.CompactTextString(m) }
func (*RequestUserAdd) ProtoMessage()    {}
func (*RequestUserAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd1e66095a603e73, []int{8}
}

func (m *RequestUserAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestUserAdd.Unmarshal(m, b)
}
func (m *RequestUserAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestUserAdd.Marshal(b, m, deterministic)
}
func (m *RequestUserAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestUserAdd.Merge(m, src)
}
func (m *RequestUserAdd) XXX_Size() int {
	return xxx_messageInfo_RequestUserAdd.Size(m)
}
func (m *RequestUserAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestUserAdd.DiscardUnknown(m)
}

var xxx_messageInfo_RequestUserAdd proto.InternalMessageInfo

func (m *RequestUserAdd) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RequestUserAdd) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type ResponseUserAdd struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseUserAdd) Reset()         { *m = ResponseUserAdd{} }
func (m *ResponseUserAdd) String() string { return proto.CompactTextString(m) }
func (*ResponseUserAdd) ProtoMessage()    {}
func (*ResponseUserAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd1e66095a603e73, []int{9}
}

func (m *ResponseUserAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseUserAdd.Unmarshal(m, b)
}
func (m *ResponseUserAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseUserAdd.Marshal(b, m, deterministic)
}
func (m *ResponseUserAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseUserAdd.Merge(m, src)
}
func (m *ResponseUserAdd) XXX_Size() int {
	return xxx_messageInfo_ResponseUserAdd.Size(m)
}
func (m *ResponseUserAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseUserAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseUserAdd proto.InternalMessageInfo

func (m *ResponseUserAdd) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type ErrorUnauthorized struct {
	Endpoint             string   `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorUnauthorized) Reset()         { *m = ErrorUnauthorized{} }
func (m *ErrorUnauthorized) String() string { return proto.CompactTextString(m) }
func (*ErrorUnauthorized) ProtoMessage()    {}
func (*ErrorUnauthorized) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd1e66095a603e73, []int{10}
}

func (m *ErrorUnauthorized) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorUnauthorized.Unmarshal(m, b)
}
func (m *ErrorUnauthorized) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorUnauthorized.Marshal(b, m, deterministic)
}
func (m *ErrorUnauthorized) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorUnauthorized.Merge(m, src)
}
func (m *ErrorUnauthorized) XXX_Size() int {
	return xxx_messageInfo_ErrorUnauthorized.Size(m)
}
func (m *ErrorUnauthorized) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorUnauthorized.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorUnauthorized proto.InternalMessageInfo

func (m *ErrorUnauthorized) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestPing)(nil), "proxy.rpc.RequestPing")
	proto.RegisterType((*ResponsePong)(nil), "proxy.rpc.ResponsePong")
	proto.RegisterType((*RequestMemberList)(nil), "proxy.rpc.RequestMemberList")
	proto.RegisterType((*ResponseMemberList)(nil), "proxy.rpc.ResponseMemberList")
	proto.RegisterType((*ClusterMember)(nil), "proxy.rpc.ClusterMember")
	proto.RegisterType((*RequestUserList)(nil), "proxy.rpc.RequestUserList")
	proto.RegisterType((*ResponseUserList)(nil), "proxy.rpc.ResponseUserList")
	proto.RegisterType((*UserItem)(nil), "proxy.rpc.UserItem")
	proto.RegisterType((*RequestUserAdd)(nil), "proxy.rpc.RequestUserAdd")
	proto.RegisterType((*ResponseUserAdd)(nil), "proxy.rpc.ResponseUserAdd")
	proto.RegisterType((*ErrorUnauthorized)(nil), "proxy.rpc.ErrorUnauthorized")
}

func init() { proto.RegisterFile("pkg/rpc/rpc.proto", fileDescriptor_bd1e66095a603e73) }

var fileDescriptor_bd1e66095a603e73 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x5f, 0xaf, 0xd2, 0x30,
	0x14, 0x77, 0xbb, 0x17, 0x2f, 0x9c, 0x2b, 0xe8, 0x8a, 0xd1, 0x39, 0x35, 0x62, 0x9f, 0xf0, 0xc1,
	0x0d, 0xd0, 0xc4, 0x27, 0x13, 0x27, 0xfa, 0x60, 0xd4, 0x84, 0x2c, 0xc1, 0x07, 0xdf, 0x06, 0xad,
	0xa3, 0x81, 0xb5, 0xb3, 0xed, 0x8c, 0xfa, 0x05, 0xfd, 0x5a, 0xa6, 0xfb, 0x03, 0x83, 0x71, 0x1f,
	0x48, 0xe8, 0xef, 0xfc, 0xfe, 0x9c, 0x9d, 0xf6, 0x80, 0x93, 0x6d, 0x93, 0x40, 0x66, 0x6b, 0xf3,
	0xf3, 0x33, 0x29, 0xb4, 0x40, 0xbd, 0x4c, 0x8a, 0xdf, 0x7f, 0x7c, 0x99, 0xad, 0x71, 0x1f, 0xae,
	0x23, 0xfa, 0x33, 0xa7, 0x4a, 0x2f, 0x18, 0x4f, 0xf0, 0x00, 0xee, 0x44, 0x54, 0x65, 0x82, 0x2b,
	0xba, 0x10, 0x3c, 0xc1, 0x43, 0x70, 0xaa, 0xf2, 0x57, 0x9a, 0xae, 0xa8, 0xfc, 0xc2, 0x94, 0xc6,
	0x1f, 0x00, 0xd5, 0xa4, 0x03, 0x8a, 0x7c, 0xe8, 0x30, 0x4d, 0x53, 0xe5, 0x5a, 0xa3, 0x8b, 0xf1,
	0xf5, 0xcc, 0xf5, 0xf7, 0x21, 0xfe, 0x7c, 0x97, 0x2b, 0x4d, 0x65, 0x49, 0x8e, 0x4a, 0x1a, 0x7e,
	0x06, 0xfd, 0x23, 0x1c, 0x0d, 0xc0, 0x66, 0xc4, 0xb5, 0x46, 0xd6, 0xb8, 0x17, 0xd9, 0x8c, 0x60,
	0x07, 0xee, 0x56, 0xd9, 0x4b, 0x55, 0x25, 0xbf, 0x85, 0x7b, 0x75, 0x72, 0x8d, 0xa1, 0x17, 0xc7,
	0xb9, 0xc3, 0x46, 0xae, 0xe1, 0x7c, 0xd2, 0x34, 0xad, 0x23, 0x27, 0xd0, 0xad, 0xa1, 0xd3, 0x34,
	0x74, 0x1f, 0x3a, 0x52, 0xec, 0xa8, 0x72, 0xed, 0xd1, 0xc5, 0xb8, 0x17, 0x95, 0x07, 0xfc, 0x1a,
	0x06, 0x8d, 0x1e, 0x42, 0x42, 0x5a, 0x3a, 0x04, 0x97, 0x86, 0xea, 0xda, 0x05, 0x52, 0xfc, 0xc7,
	0xcf, 0x4d, 0xe7, 0x87, 0x36, 0x2b, 0x99, 0xd8, 0x16, 0xb2, 0x6e, 0x64, 0x8b, 0x2d, 0x0e, 0xc0,
	0xf9, 0x28, 0xa5, 0x90, 0x4b, 0x1e, 0xe7, 0x7a, 0x23, 0x24, 0xfb, 0x4b, 0x09, 0xf2, 0xa0, 0x4b,
	0x39, 0xc9, 0x04, 0xe3, 0xba, 0x4a, 0xd8, 0x9f, 0x67, 0xdf, 0xe0, 0xaa, 0x1a, 0x17, 0xfa, 0x0c,
	0xd0, 0x98, 0xfb, 0x93, 0xc6, 0x07, 0xb7, 0xee, 0xca, 0x7b, 0x7a, 0x54, 0x3d, 0xbd, 0x34, 0x7c,
	0x6b, 0xf6, 0xcf, 0x82, 0x4e, 0x48, 0x52, 0xc6, 0xd1, 0x1b, 0xb8, 0x34, 0x6f, 0x00, 0x3d, 0x68,
	0x1b, 0x1a, 0xdc, 0x7b, 0x78, 0xc6, 0xca, 0x3c, 0x12, 0x34, 0x2f, 0xc7, 0x5a, 0x74, 0xe3, 0xb5,
	0xc5, 0x75, 0xcd, 0x7b, 0x7c, 0xc6, 0x60, 0x2f, 0x7c, 0x07, 0x57, 0xf5, 0xac, 0x1e, 0x9d, 0xf7,
	0x08, 0x09, 0xf1, 0xbc, 0x1b, 0x2c, 0x42, 0x42, 0xde, 0x4f, 0xbf, 0x07, 0x09, 0xd3, 0x9b, 0x7c,
	0xe5, 0xaf, 0x45, 0x1a, 0xfc, 0x98, 0x4e, 0x27, 0xc1, 0x2e, 0x4e, 0x64, 0xcc, 0x13, 0x16, 0xf3,
	0x97, 0x85, 0x2e, 0x30, 0xbb, 0xa0, 0xa8, 0xfc, 0x45, 0xa5, 0x59, 0x87, 0xd5, 0xed, 0x62, 0x1f,
	0x5e, 0xfd, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x82, 0x5a, 0xd1, 0x24, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClusterClient is the client API for Cluster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClusterClient interface {
	MemberList(ctx context.Context, in *RequestMemberList, opts ...grpc.CallOption) (*ResponseMemberList, error)
}

type clusterClient struct {
	cc *grpc.ClientConn
}

func NewClusterClient(cc *grpc.ClientConn) ClusterClient {
	return &clusterClient{cc}
}

func (c *clusterClient) MemberList(ctx context.Context, in *RequestMemberList, opts ...grpc.CallOption) (*ResponseMemberList, error) {
	out := new(ResponseMemberList)
	err := c.cc.Invoke(ctx, "/proxy.rpc.Cluster/MemberList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterServer is the server API for Cluster service.
type ClusterServer interface {
	MemberList(context.Context, *RequestMemberList) (*ResponseMemberList, error)
}

// UnimplementedClusterServer can be embedded to have forward compatible implementations.
type UnimplementedClusterServer struct {
}

func (*UnimplementedClusterServer) MemberList(ctx context.Context, req *RequestMemberList) (*ResponseMemberList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberList not implemented")
}

func RegisterClusterServer(s *grpc.Server, srv ClusterServer) {
	s.RegisterService(&_Cluster_serviceDesc, srv)
}

func _Cluster_MemberList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMemberList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).MemberList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proxy.rpc.Cluster/MemberList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).MemberList(ctx, req.(*RequestMemberList))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cluster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proxy.rpc.Cluster",
	HandlerType: (*ClusterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MemberList",
			Handler:    _Cluster_MemberList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/rpc/rpc.proto",
}

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdminClient interface {
	Ping(ctx context.Context, in *RequestPing, opts ...grpc.CallOption) (*ResponsePong, error)
	UserList(ctx context.Context, in *RequestUserList, opts ...grpc.CallOption) (*ResponseUserList, error)
	UserAdd(ctx context.Context, in *RequestUserAdd, opts ...grpc.CallOption) (*ResponseUserAdd, error)
}

type adminClient struct {
	cc *grpc.ClientConn
}

func NewAdminClient(cc *grpc.ClientConn) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) Ping(ctx context.Context, in *RequestPing, opts ...grpc.CallOption) (*ResponsePong, error) {
	out := new(ResponsePong)
	err := c.cc.Invoke(ctx, "/proxy.rpc.Admin/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) UserList(ctx context.Context, in *RequestUserList, opts ...grpc.CallOption) (*ResponseUserList, error) {
	out := new(ResponseUserList)
	err := c.cc.Invoke(ctx, "/proxy.rpc.Admin/UserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) UserAdd(ctx context.Context, in *RequestUserAdd, opts ...grpc.CallOption) (*ResponseUserAdd, error) {
	out := new(ResponseUserAdd)
	err := c.cc.Invoke(ctx, "/proxy.rpc.Admin/UserAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
type AdminServer interface {
	Ping(context.Context, *RequestPing) (*ResponsePong, error)
	UserList(context.Context, *RequestUserList) (*ResponseUserList, error)
	UserAdd(context.Context, *RequestUserAdd) (*ResponseUserAdd, error)
}

// UnimplementedAdminServer can be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (*UnimplementedAdminServer) Ping(ctx context.Context, req *RequestPing) (*ResponsePong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedAdminServer) UserList(ctx context.Context, req *RequestUserList) (*ResponseUserList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserList not implemented")
}
func (*UnimplementedAdminServer) UserAdd(ctx context.Context, req *RequestUserAdd) (*ResponseUserAdd, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAdd not implemented")
}

func RegisterAdminServer(s *grpc.Server, srv AdminServer) {
	s.RegisterService(&_Admin_serviceDesc, srv)
}

func _Admin_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestPing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proxy.rpc.Admin/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).Ping(ctx, req.(*RequestPing))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_UserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUserList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).UserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proxy.rpc.Admin/UserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).UserList(ctx, req.(*RequestUserList))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_UserAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUserAdd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).UserAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proxy.rpc.Admin/UserAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).UserAdd(ctx, req.(*RequestUserAdd))
	}
	return interceptor(ctx, in, info, handler)
}

var _Admin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proxy.rpc.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Admin_Ping_Handler,
		},
		{
			MethodName: "UserList",
			Handler:    _Admin_UserList_Handler,
		},
		{
			MethodName: "UserAdd",
			Handler:    _Admin_UserAdd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/rpc/rpc.proto",
}
