// Code generated by protoc-gen-go. DO NOT EDIT.
// source: domains/rental/item.proto

package rental

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	user "github.com/ilhammhdd/kudaki-user-service/entities/domains/user"
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

type Item struct {
	Uuid                 string     `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name                 string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Amount               uint32     `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Unit                 string     `protobuf:"bytes,4,opt,name=unit,proto3" json:"unit,omitempty"`
	Price                uint32     `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	Description          string     `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Owner                *user.User `protobuf:"bytes,7,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a57166d64a283f1, []int{0}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetAmount() uint32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Item) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *Item) GetPrice() uint32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Item) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Item) GetOwner() *user.User {
	if m != nil {
		return m.Owner
	}
	return nil
}

func init() {
	proto.RegisterType((*Item)(nil), "domain.rental.Item")
}

func init() { proto.RegisterFile("domains/rental/item.proto", fileDescriptor_6a57166d64a283f1) }

var fileDescriptor_6a57166d64a283f1 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x89, 0xb6, 0x15, 0xb3, 0xec, 0xc1, 0x20, 0x1a, 0x3d, 0x15, 0x2f, 0xf6, 0xb2, 0x0d,
	0xe8, 0x1f, 0x90, 0xbd, 0x79, 0x2d, 0x78, 0xf1, 0x96, 0x6d, 0x06, 0x3b, 0xec, 0x26, 0x29, 0xc9,
	0x44, 0x7f, 0x9c, 0x7f, 0x4e, 0x9a, 0x54, 0xd0, 0x4b, 0x98, 0x79, 0xef, 0x7b, 0x90, 0x79, 0xfc,
	0xce, 0x78, 0xab, 0xd1, 0x45, 0x15, 0xc0, 0x91, 0x3e, 0x29, 0x24, 0xb0, 0xfd, 0x1c, 0x3c, 0x79,
	0xb1, 0x2d, 0x56, 0x5f, 0x9c, 0xfb, 0xdb, 0x5f, 0x32, 0x45, 0x08, 0xf9, 0x29, 0xdc, 0xc3, 0x37,
	0xe3, 0xd5, 0x2b, 0x81, 0x15, 0x82, 0x57, 0x29, 0xa1, 0x91, 0xac, 0x65, 0xdd, 0xe5, 0x90, 0xe7,
	0x45, 0x73, 0xda, 0x82, 0x3c, 0x2b, 0xda, 0x32, 0x8b, 0x1b, 0xde, 0x68, 0xeb, 0x93, 0x23, 0x79,
	0xde, 0xb2, 0x6e, 0x3b, 0xac, 0x5b, 0xce, 0x3b, 0x24, 0x59, 0xad, 0x79, 0x87, 0x24, 0xae, 0x79,
	0x3d, 0x07, 0x1c, 0x41, 0xd6, 0x19, 0x2d, 0x8b, 0x68, 0xf9, 0xc6, 0x40, 0x1c, 0x03, 0xce, 0x84,
	0xde, 0xc9, 0x26, 0x07, 0xfe, 0x4a, 0xe2, 0x91, 0xd7, 0xfe, 0xcb, 0x41, 0x90, 0x17, 0x2d, 0xeb,
	0x36, 0x4f, 0x57, 0xfd, 0x7a, 0x4c, 0xfe, 0xf7, 0x5b, 0x84, 0x30, 0x14, 0x7f, 0xbf, 0x7f, 0x7f,
	0xf9, 0x40, 0x9a, 0xd2, 0xa1, 0x1f, 0xbd, 0x55, 0x78, 0x9a, 0xb4, 0xb5, 0x93, 0x31, 0xea, 0x98,
	0x8c, 0x3e, 0xe2, 0x6e, 0xe1, 0x77, 0x11, 0xc2, 0x27, 0x8e, 0xa0, 0xc0, 0x11, 0x12, 0x42, 0x54,
	0xff, 0x4b, 0x3b, 0x34, 0xb9, 0x88, 0xe7, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xee, 0x33, 0xf9,
	0x26, 0x4d, 0x01, 0x00, 0x00,
}
