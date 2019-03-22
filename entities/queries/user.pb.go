// Code generated by protoc-gen-go. DO NOT EDIT.
// source: queries/user.proto

package queries

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	domains "github.com/ilhammhdd/kudaki-user-service/entities/domains"
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

type RetrieveUsers struct {
	Uuid                 string          `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Filter               *domains.Filter `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RetrieveUsers) Reset()         { *m = RetrieveUsers{} }
func (m *RetrieveUsers) String() string { return proto.CompactTextString(m) }
func (*RetrieveUsers) ProtoMessage()    {}
func (*RetrieveUsers) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea3496c9576abf8f, []int{0}
}

func (m *RetrieveUsers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveUsers.Unmarshal(m, b)
}
func (m *RetrieveUsers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveUsers.Marshal(b, m, deterministic)
}
func (m *RetrieveUsers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveUsers.Merge(m, src)
}
func (m *RetrieveUsers) XXX_Size() int {
	return xxx_messageInfo_RetrieveUsers.Size(m)
}
func (m *RetrieveUsers) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveUsers.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveUsers proto.InternalMessageInfo

func (m *RetrieveUsers) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *RetrieveUsers) GetFilter() *domains.Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type RetrieveProfile struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	ProfileUuid          string   `protobuf:"bytes,2,opt,name=profile_uuid,json=profileUuid,proto3" json:"profile_uuid,omitempty"`
	UserUuid             string   `protobuf:"bytes,3,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetrieveProfile) Reset()         { *m = RetrieveProfile{} }
func (m *RetrieveProfile) String() string { return proto.CompactTextString(m) }
func (*RetrieveProfile) ProtoMessage()    {}
func (*RetrieveProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea3496c9576abf8f, []int{1}
}

func (m *RetrieveProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveProfile.Unmarshal(m, b)
}
func (m *RetrieveProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveProfile.Marshal(b, m, deterministic)
}
func (m *RetrieveProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveProfile.Merge(m, src)
}
func (m *RetrieveProfile) XXX_Size() int {
	return xxx_messageInfo_RetrieveProfile.Size(m)
}
func (m *RetrieveProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveProfile.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveProfile proto.InternalMessageInfo

func (m *RetrieveProfile) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *RetrieveProfile) GetProfileUuid() string {
	if m != nil {
		return m.ProfileUuid
	}
	return ""
}

func (m *RetrieveProfile) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func init() {
	proto.RegisterType((*RetrieveUsers)(nil), "query.RetrieveUsers")
	proto.RegisterType((*RetrieveProfile)(nil), "query.RetrieveProfile")
}

func init() { proto.RegisterFile("queries/user.proto", fileDescriptor_ea3496c9576abf8f) }

var fileDescriptor_ea3496c9576abf8f = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd0, 0xbf, 0x4b, 0xc4, 0x30,
	0x14, 0x07, 0x70, 0x7a, 0xea, 0xe1, 0xe5, 0xfc, 0x01, 0xc1, 0xe1, 0xd0, 0xe5, 0xec, 0x20, 0x5d,
	0x9a, 0x80, 0x4e, 0xe2, 0xe6, 0xe0, 0xe2, 0x22, 0x85, 0x2e, 0x2e, 0xd2, 0x36, 0xaf, 0xf6, 0xd1,
	0xa6, 0xa9, 0xf9, 0x51, 0xf0, 0xbf, 0x97, 0x24, 0xed, 0xe6, 0x16, 0xbe, 0x9f, 0x97, 0x2f, 0xbc,
	0x47, 0xe8, 0x8f, 0x03, 0x8d, 0x60, 0xb8, 0x33, 0xa0, 0xd9, 0xa4, 0x95, 0x55, 0xf4, 0xcc, 0x67,
	0xbf, 0xb7, 0x37, 0x42, 0xc9, 0x0a, 0x47, 0xc3, 0x5b, 0x1c, 0xec, 0x8a, 0xe9, 0x3b, 0xb9, 0x2c,
	0xc0, 0x6a, 0x84, 0x19, 0x4a, 0x03, 0xda, 0x50, 0x4a, 0x4e, 0x9d, 0x43, 0x71, 0x48, 0x8e, 0x49,
	0xb6, 0x2b, 0xc2, 0x9b, 0x3e, 0x90, 0x6d, 0xfc, 0x74, 0xd8, 0x1c, 0x93, 0x6c, 0xff, 0x78, 0xc5,
	0x62, 0x17, 0x7b, 0x0b, 0x69, 0xb1, 0x68, 0x0a, 0xe4, 0x7a, 0x2d, 0xfb, 0xd0, 0xaa, 0xc5, 0x01,
	0xfe, 0xad, 0xbb, 0x27, 0x17, 0x53, 0xe4, 0xaf, 0x60, 0x9b, 0x60, 0xfb, 0x25, 0x2b, 0xfd, 0xc8,
	0x1d, 0xd9, 0xf9, 0x0d, 0xa2, 0x9f, 0x04, 0x3f, 0xf7, 0x81, 0xc7, 0xd7, 0x97, 0xcf, 0xe7, 0x6f,
	0xb4, 0x9d, 0xab, 0x59, 0xa3, 0x24, 0xc7, 0xa1, 0xab, 0xa4, 0xec, 0x84, 0xe0, 0xbd, 0x13, 0x55,
	0x8f, 0xb9, 0x9f, 0xcb, 0x0d, 0xe8, 0x19, 0x1b, 0xe0, 0x30, 0x5a, 0xb4, 0xfe, 0x20, 0xcb, 0x61,
	0xea, 0x6d, 0xd8, 0xfb, 0xe9, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x53, 0x06, 0x27, 0x2a, 0x01,
	0x00, 0x00,
}