// Code generated by protoc-gen-go. DO NOT EDIT.
// source: store/storefront.proto

package store

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

type Storefront struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	UserUuid             string   `protobuf:"bytes,2,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	TotalItem            int32    `protobuf:"varint,3,opt,name=total_item,json=totalItem,proto3" json:"total_item,omitempty"`
	Rating               float32  `protobuf:"fixed32,4,opt,name=rating,proto3" json:"rating,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Storefront) Reset()         { *m = Storefront{} }
func (m *Storefront) String() string { return proto.CompactTextString(m) }
func (*Storefront) ProtoMessage()    {}
func (*Storefront) Descriptor() ([]byte, []int) {
	return fileDescriptor_57a78b9ab3cd9ee2, []int{0}
}

func (m *Storefront) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Storefront.Unmarshal(m, b)
}
func (m *Storefront) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Storefront.Marshal(b, m, deterministic)
}
func (m *Storefront) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Storefront.Merge(m, src)
}
func (m *Storefront) XXX_Size() int {
	return xxx_messageInfo_Storefront.Size(m)
}
func (m *Storefront) XXX_DiscardUnknown() {
	xxx_messageInfo_Storefront.DiscardUnknown(m)
}

var xxx_messageInfo_Storefront proto.InternalMessageInfo

func (m *Storefront) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Storefront) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *Storefront) GetTotalItem() int32 {
	if m != nil {
		return m.TotalItem
	}
	return 0
}

func (m *Storefront) GetRating() float32 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func init() {
	proto.RegisterType((*Storefront)(nil), "entities.store.Storefront")
}

func init() { proto.RegisterFile("store/storefront.proto", fileDescriptor_57a78b9ab3cd9ee2) }

var fileDescriptor_57a78b9ab3cd9ee2 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8e, 0x3d, 0x0f, 0x82, 0x30,
	0x10, 0x40, 0x53, 0x44, 0x22, 0x37, 0x38, 0x74, 0x20, 0x24, 0xc6, 0x84, 0x38, 0x11, 0xa3, 0x30,
	0xf8, 0x0f, 0xdc, 0x5c, 0x31, 0x2e, 0x2e, 0xa4, 0xd8, 0x0a, 0x17, 0x28, 0x35, 0xe5, 0xfa, 0xff,
	0x0d, 0xf5, 0x63, 0xb9, 0xdc, 0xbd, 0x77, 0xc3, 0x83, 0x64, 0x22, 0x63, 0x55, 0xe9, 0xe7, 0xd3,
	0x9a, 0x91, 0x8a, 0x97, 0x35, 0x64, 0xf8, 0x5a, 0x8d, 0x84, 0x84, 0x6a, 0x2a, 0xbc, 0xda, 0x11,
	0xc0, 0xf5, 0xff, 0xc3, 0x39, 0x84, 0xce, 0xa1, 0x4c, 0x59, 0xc6, 0xf2, 0xb8, 0xf2, 0x3b, 0xdf,
	0x40, 0xec, 0x26, 0x65, 0x6b, 0x2f, 0x02, 0x2f, 0x56, 0x33, 0xb8, 0xcd, 0x72, 0x0b, 0x40, 0x86,
	0xc4, 0x50, 0x23, 0x29, 0x9d, 0x2e, 0x32, 0x96, 0x2f, 0xab, 0xd8, 0x93, 0x0b, 0x29, 0xcd, 0x13,
	0x88, 0xac, 0x20, 0x1c, 0xdb, 0x34, 0xcc, 0x58, 0x1e, 0x54, 0xdf, 0xeb, 0x7c, 0xb8, 0xef, 0x5b,
	0xa4, 0xce, 0x35, 0xc5, 0xc3, 0xe8, 0x12, 0x87, 0x4e, 0x68, 0xdd, 0x49, 0x59, 0xf6, 0x4e, 0x8a,
	0x1e, 0x8f, 0xbf, 0xc6, 0x4f, 0x7e, 0x13, 0xf9, 0xf4, 0xd3, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x12,
	0xf8, 0x4a, 0x9c, 0xd4, 0x00, 0x00, 0x00,
}
