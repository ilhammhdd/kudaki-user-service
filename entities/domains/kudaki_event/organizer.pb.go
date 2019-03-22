// Code generated by protoc-gen-go. DO NOT EDIT.
// source: domains/kudaki_event/organizer.proto

package kudaki_event

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

type Organizer struct {
	Uuid                 string     `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	User                 *user.User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Name                 string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Phone                string     `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	PhotoPath            string     `protobuf:"bytes,5,opt,name=photo_path,json=photoPath,proto3" json:"photo_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Organizer) Reset()         { *m = Organizer{} }
func (m *Organizer) String() string { return proto.CompactTextString(m) }
func (*Organizer) ProtoMessage()    {}
func (*Organizer) Descriptor() ([]byte, []int) {
	return fileDescriptor_19b2457f4f8d435f, []int{0}
}

func (m *Organizer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Organizer.Unmarshal(m, b)
}
func (m *Organizer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Organizer.Marshal(b, m, deterministic)
}
func (m *Organizer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Organizer.Merge(m, src)
}
func (m *Organizer) XXX_Size() int {
	return xxx_messageInfo_Organizer.Size(m)
}
func (m *Organizer) XXX_DiscardUnknown() {
	xxx_messageInfo_Organizer.DiscardUnknown(m)
}

var xxx_messageInfo_Organizer proto.InternalMessageInfo

func (m *Organizer) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Organizer) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Organizer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Organizer) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Organizer) GetPhotoPath() string {
	if m != nil {
		return m.PhotoPath
	}
	return ""
}

func init() {
	proto.RegisterType((*Organizer)(nil), "domain.kudaki_event.Organizer")
}

func init() {
	proto.RegisterFile("domains/kudaki_event/organizer.proto", fileDescriptor_19b2457f4f8d435f)
}

var fileDescriptor_19b2457f4f8d435f = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x59, 0xdd, 0x0a, 0x1b, 0x4f, 0x46, 0xc1, 0x20, 0x08, 0x45, 0x14, 0x7a, 0xe9, 0x06,
	0xf4, 0x1f, 0x78, 0x10, 0x6f, 0x4a, 0xc1, 0x8b, 0x97, 0x92, 0x36, 0x43, 0x33, 0xd4, 0x24, 0x4b,
	0x32, 0xe9, 0xc1, 0x3f, 0xe1, 0x5f, 0x96, 0x9d, 0x74, 0xc1, 0x83, 0x97, 0x30, 0x99, 0xf7, 0xbd,
	0x17, 0x5e, 0xc4, 0xbd, 0x8d, 0xde, 0x60, 0xc8, 0x7a, 0x5f, 0xac, 0xd9, 0xe3, 0x1a, 0x0e, 0x10,
	0x48, 0xc7, 0xb4, 0x33, 0x01, 0xbf, 0x21, 0xf5, 0x43, 0x8a, 0x14, 0xe5, 0x65, 0xa5, 0xfa, 0xbf,
	0xd0, 0xcd, 0xf5, 0x64, 0x2d, 0x19, 0x12, 0x1f, 0x95, 0xbe, 0xfb, 0x69, 0x44, 0xf7, 0x36, 0x25,
	0x48, 0x29, 0xda, 0x52, 0xd0, 0xaa, 0x66, 0xde, 0x2c, 0xba, 0x15, 0xcf, 0xf2, 0x41, 0xb4, 0x23,
	0xaf, 0x4e, 0xe6, 0xcd, 0xe2, 0xfc, 0xf1, 0xa2, 0x3f, 0xc6, 0x73, 0xc6, 0x47, 0x86, 0xb4, 0x62,
	0x79, 0xb4, 0x06, 0xe3, 0x41, 0x9d, 0x56, 0xeb, 0x38, 0xcb, 0x2b, 0x31, 0x1b, 0x5c, 0x0c, 0xa0,
	0x5a, 0x5e, 0xd6, 0x8b, 0xbc, 0x15, 0x62, 0x70, 0x91, 0xe2, 0x7a, 0x30, 0xe4, 0xd4, 0x8c, 0xa5,
	0x8e, 0x37, 0xef, 0x86, 0xdc, 0xf3, 0xeb, 0xe7, 0xcb, 0x0e, 0xc9, 0x95, 0x4d, 0xbf, 0x8d, 0x5e,
	0xe3, 0x97, 0x33, 0xde, 0x3b, 0x6b, 0x8f, 0xa5, 0x97, 0xe3, 0x63, 0xcb, 0x0c, 0xe9, 0x80, 0x5b,
	0xd0, 0x10, 0x08, 0x09, 0x21, 0xeb, 0xff, 0x7e, 0x66, 0x73, 0xc6, 0x15, 0x9f, 0x7e, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xc0, 0x15, 0xb8, 0x7e, 0x38, 0x01, 0x00, 0x00,
}