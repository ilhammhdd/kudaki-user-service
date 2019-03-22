// Code generated by protoc-gen-go. DO NOT EDIT.
// source: domains/rental/cart.proto

package rental

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	domains "github.com/ilhammhdd/kudaki-user-service/entities/domains"
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

type Cart struct {
	Uuid                 string         `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	User                 *user.User     `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Items                []*Item        `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Total                float64        `protobuf:"fixed64,4,opt,name=total,proto3" json:"total,omitempty"`
	Status               domains.Status `protobuf:"varint,5,opt,name=status,proto3,enum=domain.Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Cart) Reset()         { *m = Cart{} }
func (m *Cart) String() string { return proto.CompactTextString(m) }
func (*Cart) ProtoMessage()    {}
func (*Cart) Descriptor() ([]byte, []int) {
	return fileDescriptor_017500e750829d52, []int{0}
}

func (m *Cart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cart.Unmarshal(m, b)
}
func (m *Cart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cart.Marshal(b, m, deterministic)
}
func (m *Cart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cart.Merge(m, src)
}
func (m *Cart) XXX_Size() int {
	return xxx_messageInfo_Cart.Size(m)
}
func (m *Cart) XXX_DiscardUnknown() {
	xxx_messageInfo_Cart.DiscardUnknown(m)
}

var xxx_messageInfo_Cart proto.InternalMessageInfo

func (m *Cart) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Cart) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Cart) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Cart) GetTotal() float64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Cart) GetStatus() domains.Status {
	if m != nil {
		return m.Status
	}
	return domains.Status_ON_PROGRESS
}

func init() {
	proto.RegisterType((*Cart)(nil), "domain.rental.Cart")
}

func init() { proto.RegisterFile("domains/rental/cart.proto", fileDescriptor_017500e750829d52) }

var fileDescriptor_017500e750829d52 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4e, 0x03, 0x21,
	0x18, 0x84, 0x83, 0xdd, 0x36, 0x91, 0xc6, 0x26, 0x62, 0x13, 0xd7, 0x9e, 0x88, 0x89, 0x06, 0x0f,
	0x85, 0x64, 0x7d, 0x01, 0x53, 0x4f, 0x5e, 0x31, 0x5e, 0xbc, 0xd1, 0x85, 0xb8, 0xa4, 0xcb, 0x62,
	0xe0, 0xc7, 0x87, 0xf2, 0x29, 0xcd, 0xc2, 0xee, 0xa1, 0x5e, 0x08, 0xfc, 0xdf, 0x64, 0x66, 0xf8,
	0xf1, 0x9d, 0xf6, 0x4e, 0xd9, 0x21, 0x8a, 0x60, 0x06, 0x50, 0xbd, 0x68, 0x55, 0x00, 0xfe, 0x1d,
	0x3c, 0x78, 0x72, 0x55, 0x10, 0x2f, 0x64, 0x77, 0x3b, 0x2b, 0x53, 0x34, 0x21, 0x1f, 0x45, 0xb7,
	0xfb, 0x6f, 0x61, 0xc1, 0xb8, 0x09, 0x6d, 0x67, 0x14, 0x41, 0x41, 0x8a, 0x65, 0x7a, 0xff, 0x8b,
	0x70, 0xf5, 0xaa, 0x02, 0x10, 0x82, 0xab, 0x94, 0xac, 0xae, 0x11, 0x45, 0xec, 0x52, 0xe6, 0x3b,
	0x79, 0xc0, 0xd5, 0xe8, 0x5d, 0x5f, 0x50, 0xc4, 0xd6, 0xcd, 0x35, 0x9f, 0x4a, 0xe4, 0xbc, 0x8f,
	0x68, 0x82, 0xcc, 0x98, 0x3c, 0xe1, 0xe5, 0x98, 0x13, 0xeb, 0x05, 0x5d, 0xb0, 0x75, 0x73, 0xc3,
	0xcf, 0xca, 0xf2, 0x37, 0x30, 0x4e, 0x16, 0x05, 0xd9, 0xe2, 0x25, 0x78, 0x50, 0x7d, 0x5d, 0x51,
	0xc4, 0x90, 0x2c, 0x0f, 0xf2, 0x88, 0x57, 0xa5, 0x54, 0xbd, 0xa4, 0x88, 0x6d, 0x9a, 0xcd, 0xec,
	0xf0, 0x9e, 0xa7, 0x72, 0xa2, 0x87, 0xc3, 0xe7, 0xcb, 0x97, 0x85, 0x2e, 0x1d, 0x79, 0xeb, 0x9d,
	0xb0, 0x7d, 0xa7, 0x9c, 0xeb, 0xb4, 0x16, 0xa7, 0xa4, 0xd5, 0xc9, 0xee, 0xc7, 0x32, 0xfb, 0x68,
	0xc2, 0x8f, 0x6d, 0x8d, 0x30, 0x03, 0x58, 0xb0, 0x26, 0x8a, 0xf3, 0x8d, 0x1c, 0x57, 0xf9, 0xdf,
	0xcf, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6d, 0xce, 0x21, 0x01, 0x6d, 0x01, 0x00, 0x00,
}
