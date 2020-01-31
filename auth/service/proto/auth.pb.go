// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth/service/proto/auth.proto

package go_micro_auth

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

type ServiceAccount struct {
	Parent               *Resource         `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	Token                string            `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Created              int64             `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Expiry               int64             `protobuf:"varint,4,opt,name=expiry,proto3" json:"expiry,omitempty"`
	Roles                []*Role           `protobuf:"bytes,5,rep,name=roles,proto3" json:"roles,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ServiceAccount) Reset()         { *m = ServiceAccount{} }
func (m *ServiceAccount) String() string { return proto.CompactTextString(m) }
func (*ServiceAccount) ProtoMessage()    {}
func (*ServiceAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{0}
}

func (m *ServiceAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceAccount.Unmarshal(m, b)
}
func (m *ServiceAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceAccount.Marshal(b, m, deterministic)
}
func (m *ServiceAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceAccount.Merge(m, src)
}
func (m *ServiceAccount) XXX_Size() int {
	return xxx_messageInfo_ServiceAccount.Size(m)
}
func (m *ServiceAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceAccount.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceAccount proto.InternalMessageInfo

func (m *ServiceAccount) GetParent() *Resource {
	if m != nil {
		return m.Parent
	}
	return nil
}

func (m *ServiceAccount) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ServiceAccount) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *ServiceAccount) GetExpiry() int64 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

func (m *ServiceAccount) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *ServiceAccount) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Role struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Resource             *Resource `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{1}
}

func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (m *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(m, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Role) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

type Resource struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{2}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Resource) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type GenerateRequest struct {
	ServiceAccount       *ServiceAccount `protobuf:"bytes,1,opt,name=service_account,json=serviceAccount,proto3" json:"service_account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GenerateRequest) Reset()         { *m = GenerateRequest{} }
func (m *GenerateRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateRequest) ProtoMessage()    {}
func (*GenerateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{3}
}

func (m *GenerateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateRequest.Unmarshal(m, b)
}
func (m *GenerateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateRequest.Marshal(b, m, deterministic)
}
func (m *GenerateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateRequest.Merge(m, src)
}
func (m *GenerateRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateRequest.Size(m)
}
func (m *GenerateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateRequest proto.InternalMessageInfo

func (m *GenerateRequest) GetServiceAccount() *ServiceAccount {
	if m != nil {
		return m.ServiceAccount
	}
	return nil
}

type GenerateResponse struct {
	ServiceAccount       *ServiceAccount `protobuf:"bytes,1,opt,name=service_account,json=serviceAccount,proto3" json:"service_account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GenerateResponse) Reset()         { *m = GenerateResponse{} }
func (m *GenerateResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateResponse) ProtoMessage()    {}
func (*GenerateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{4}
}

func (m *GenerateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateResponse.Unmarshal(m, b)
}
func (m *GenerateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateResponse.Marshal(b, m, deterministic)
}
func (m *GenerateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateResponse.Merge(m, src)
}
func (m *GenerateResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateResponse.Size(m)
}
func (m *GenerateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateResponse proto.InternalMessageInfo

func (m *GenerateResponse) GetServiceAccount() *ServiceAccount {
	if m != nil {
		return m.ServiceAccount
	}
	return nil
}

type ValidateRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateRequest) Reset()         { *m = ValidateRequest{} }
func (m *ValidateRequest) String() string { return proto.CompactTextString(m) }
func (*ValidateRequest) ProtoMessage()    {}
func (*ValidateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{5}
}

func (m *ValidateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateRequest.Unmarshal(m, b)
}
func (m *ValidateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateRequest.Marshal(b, m, deterministic)
}
func (m *ValidateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateRequest.Merge(m, src)
}
func (m *ValidateRequest) XXX_Size() int {
	return xxx_messageInfo_ValidateRequest.Size(m)
}
func (m *ValidateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateRequest proto.InternalMessageInfo

func (m *ValidateRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ValidateResponse struct {
	ServiceAccount       *ServiceAccount `protobuf:"bytes,1,opt,name=service_account,json=serviceAccount,proto3" json:"service_account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ValidateResponse) Reset()         { *m = ValidateResponse{} }
func (m *ValidateResponse) String() string { return proto.CompactTextString(m) }
func (*ValidateResponse) ProtoMessage()    {}
func (*ValidateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{6}
}

func (m *ValidateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateResponse.Unmarshal(m, b)
}
func (m *ValidateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateResponse.Marshal(b, m, deterministic)
}
func (m *ValidateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateResponse.Merge(m, src)
}
func (m *ValidateResponse) XXX_Size() int {
	return xxx_messageInfo_ValidateResponse.Size(m)
}
func (m *ValidateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateResponse proto.InternalMessageInfo

func (m *ValidateResponse) GetServiceAccount() *ServiceAccount {
	if m != nil {
		return m.ServiceAccount
	}
	return nil
}

type RevokeRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RevokeRequest) Reset()         { *m = RevokeRequest{} }
func (m *RevokeRequest) String() string { return proto.CompactTextString(m) }
func (*RevokeRequest) ProtoMessage()    {}
func (*RevokeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{7}
}

func (m *RevokeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokeRequest.Unmarshal(m, b)
}
func (m *RevokeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokeRequest.Marshal(b, m, deterministic)
}
func (m *RevokeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokeRequest.Merge(m, src)
}
func (m *RevokeRequest) XXX_Size() int {
	return xxx_messageInfo_RevokeRequest.Size(m)
}
func (m *RevokeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RevokeRequest proto.InternalMessageInfo

func (m *RevokeRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RevokeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RevokeResponse) Reset()         { *m = RevokeResponse{} }
func (m *RevokeResponse) String() string { return proto.CompactTextString(m) }
func (*RevokeResponse) ProtoMessage()    {}
func (*RevokeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{8}
}

func (m *RevokeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokeResponse.Unmarshal(m, b)
}
func (m *RevokeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokeResponse.Marshal(b, m, deterministic)
}
func (m *RevokeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokeResponse.Merge(m, src)
}
func (m *RevokeResponse) XXX_Size() int {
	return xxx_messageInfo_RevokeResponse.Size(m)
}
func (m *RevokeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RevokeResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ServiceAccount)(nil), "go.micro.auth.ServiceAccount")
	proto.RegisterMapType((map[string]string)(nil), "go.micro.auth.ServiceAccount.MetadataEntry")
	proto.RegisterType((*Role)(nil), "go.micro.auth.Role")
	proto.RegisterType((*Resource)(nil), "go.micro.auth.Resource")
	proto.RegisterType((*GenerateRequest)(nil), "go.micro.auth.GenerateRequest")
	proto.RegisterType((*GenerateResponse)(nil), "go.micro.auth.GenerateResponse")
	proto.RegisterType((*ValidateRequest)(nil), "go.micro.auth.ValidateRequest")
	proto.RegisterType((*ValidateResponse)(nil), "go.micro.auth.ValidateResponse")
	proto.RegisterType((*RevokeRequest)(nil), "go.micro.auth.RevokeRequest")
	proto.RegisterType((*RevokeResponse)(nil), "go.micro.auth.RevokeResponse")
}

func init() { proto.RegisterFile("auth/service/proto/auth.proto", fileDescriptor_21300bfacc51fc2a) }

var fileDescriptor_21300bfacc51fc2a = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdf, 0x8b, 0xd3, 0x40,
	0x10, 0xbe, 0xa4, 0x6d, 0xec, 0xcd, 0xd1, 0x1f, 0x8c, 0xa2, 0xa1, 0x58, 0x2d, 0x01, 0xb1, 0x22,
	0xa4, 0xd0, 0x7b, 0x11, 0x7d, 0xba, 0x07, 0xed, 0xd3, 0x21, 0xac, 0x20, 0xe8, 0x8b, 0xac, 0xc9,
	0xe0, 0x85, 0xb6, 0xd9, 0xb8, 0xd9, 0x14, 0xfb, 0xaf, 0xf9, 0x3f, 0xf9, 0x3f, 0xc8, 0x6e, 0x76,
	0x7b, 0x6d, 0xca, 0xdd, 0xbd, 0xf4, 0x6d, 0x66, 0xe7, 0x9b, 0x6f, 0xbf, 0x6f, 0x76, 0x58, 0x18,
	0xf3, 0x4a, 0xdd, 0xcc, 0x4a, 0x92, 0x9b, 0x2c, 0xa1, 0x59, 0x21, 0x85, 0x12, 0x33, 0x7d, 0x14,
	0x9b, 0x10, 0x7b, 0xbf, 0x44, 0xbc, 0xce, 0x12, 0x29, 0x62, 0x7d, 0x18, 0xfd, 0xf5, 0xa1, 0xff,
	0xa5, 0xc6, 0x5e, 0x25, 0x89, 0xa8, 0x72, 0x85, 0x33, 0x08, 0x0a, 0x2e, 0x29, 0x57, 0xa1, 0x37,
	0xf1, 0xa6, 0x17, 0xf3, 0x67, 0xf1, 0x41, 0x4b, 0xcc, 0xa8, 0x14, 0x95, 0x4c, 0x88, 0x59, 0x18,
	0x3e, 0x81, 0x8e, 0x12, 0x4b, 0xca, 0x43, 0x7f, 0xe2, 0x4d, 0xcf, 0x59, 0x9d, 0x60, 0x08, 0x8f,
	0x12, 0x49, 0x5c, 0x51, 0x1a, 0xb6, 0x26, 0xde, 0xb4, 0xc5, 0x5c, 0x8a, 0x4f, 0x21, 0xa0, 0x3f,
	0x45, 0x26, 0xb7, 0x61, 0xdb, 0x14, 0x6c, 0x86, 0x6f, 0xa0, 0x23, 0xc5, 0x8a, 0xca, 0xb0, 0x33,
	0x69, 0x4d, 0x2f, 0xe6, 0x8f, 0x9b, 0xf7, 0x8a, 0x15, 0xb1, 0x1a, 0x81, 0x0b, 0xe8, 0xae, 0x49,
	0xf1, 0x94, 0x2b, 0x1e, 0x06, 0x06, 0xfd, 0xb6, 0x81, 0x3e, 0x34, 0x15, 0x5f, 0x5b, 0xf4, 0xc7,
	0x5c, 0xc9, 0x2d, 0xdb, 0x35, 0x8f, 0x3e, 0x40, 0xef, 0xa0, 0x84, 0x43, 0x68, 0x2d, 0x69, 0x6b,
	0xac, 0x9f, 0x33, 0x1d, 0x6a, 0x7b, 0x1b, 0xbe, 0xaa, 0xc8, 0xd9, 0x33, 0xc9, 0x7b, 0xff, 0x9d,
	0x17, 0x7d, 0x86, 0xb6, 0x16, 0x85, 0x08, 0xed, 0x9c, 0xaf, 0xc9, 0x36, 0x99, 0x18, 0x2f, 0xa1,
	0x2b, 0xed, 0xa0, 0x4c, 0xe3, 0x3d, 0x73, 0xdc, 0x01, 0xa3, 0x18, 0xba, 0xee, 0x14, 0xfb, 0xe0,
	0x67, 0xa9, 0xa5, 0xf4, 0xb3, 0x54, 0x5f, 0xa2, 0xb6, 0x85, 0x53, 0x61, 0xe2, 0xe8, 0x1b, 0x0c,
	0x16, 0x94, 0x93, 0xe4, 0x8a, 0x18, 0xfd, 0xae, 0xa8, 0x54, 0xf8, 0x09, 0x06, 0xf6, 0xed, 0x7f,
	0xf0, 0xda, 0xbb, 0x7d, 0xc6, 0xf1, 0xbd, 0x03, 0x62, 0xfd, 0xf2, 0x20, 0x8f, 0xbe, 0xc3, 0xf0,
	0x96, 0xba, 0x2c, 0x44, 0x5e, 0xd2, 0xc9, 0xb8, 0x5f, 0xc3, 0xe0, 0x2b, 0x5f, 0x65, 0xe9, 0x9e,
	0xec, 0xdd, 0x0e, 0x79, 0x7b, 0x3b, 0xa4, 0x45, 0xdc, 0x02, 0x4f, 0x2c, 0xe2, 0x15, 0xf4, 0x18,
	0x6d, 0xc4, 0xf2, 0x01, 0x09, 0x43, 0xe8, 0x3b, 0x58, 0x2d, 0x60, 0xfe, 0xcf, 0x83, 0xf6, 0x55,
	0xa5, 0x6e, 0xf0, 0x1a, 0xba, 0x6e, 0x44, 0xf8, 0xa2, 0x71, 0x79, 0xe3, 0x59, 0x46, 0x2f, 0xef,
	0xac, 0xd7, 0xac, 0xd1, 0x99, 0xa6, 0x73, 0x66, 0x8f, 0xe8, 0x1a, 0xe3, 0x3a, 0xa2, 0x6b, 0x4e,
	0x29, 0x3a, 0xc3, 0x05, 0x04, 0xb5, 0x70, 0x7c, 0x7e, 0xb4, 0x78, 0x7b, 0xb6, 0x47, 0xe3, 0x3b,
	0xaa, 0x8e, 0xe8, 0x67, 0x60, 0x3e, 0x8e, 0xcb, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2c, 0xe3,
	0x2d, 0xb8, 0x59, 0x04, 0x00, 0x00,
}
