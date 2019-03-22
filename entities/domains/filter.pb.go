// Code generated by protoc-gen-go. DO NOT EDIT.
// source: domains/filter.proto

package domains

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

type Range struct {
	Min                  uint32   `protobuf:"varint,1,opt,name=min,proto3" json:"min,omitempty"`
	Max                  uint32   `protobuf:"varint,2,opt,name=max,proto3" json:"max,omitempty"`
	Unit                 string   `protobuf:"bytes,3,opt,name=unit,proto3" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Range) Reset()         { *m = Range{} }
func (m *Range) String() string { return proto.CompactTextString(m) }
func (*Range) ProtoMessage()    {}
func (*Range) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ca831ad796a6446, []int{0}
}

func (m *Range) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Range.Unmarshal(m, b)
}
func (m *Range) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Range.Marshal(b, m, deterministic)
}
func (m *Range) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Range.Merge(m, src)
}
func (m *Range) XXX_Size() int {
	return xxx_messageInfo_Range.Size(m)
}
func (m *Range) XXX_DiscardUnknown() {
	xxx_messageInfo_Range.DiscardUnknown(m)
}

var xxx_messageInfo_Range proto.InternalMessageInfo

func (m *Range) GetMin() uint32 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *Range) GetMax() uint32 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *Range) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

type Category struct {
	Name                 string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SubCategories        []*Category `protobuf:"bytes,2,rep,name=sub_categories,json=subCategories,proto3" json:"sub_categories,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Category) Reset()         { *m = Category{} }
func (m *Category) String() string { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()    {}
func (*Category) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ca831ad796a6446, []int{1}
}

func (m *Category) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Category.Unmarshal(m, b)
}
func (m *Category) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Category.Marshal(b, m, deterministic)
}
func (m *Category) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Category.Merge(m, src)
}
func (m *Category) XXX_Size() int {
	return xxx_messageInfo_Category.Size(m)
}
func (m *Category) XXX_DiscardUnknown() {
	xxx_messageInfo_Category.DiscardUnknown(m)
}

var xxx_messageInfo_Category proto.InternalMessageInfo

func (m *Category) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Category) GetSubCategories() []*Category {
	if m != nil {
		return m.SubCategories
	}
	return nil
}

type Filter struct {
	Range                *Range      `protobuf:"bytes,1,opt,name=range,proto3" json:"range,omitempty"`
	Categories           []*Category `protobuf:"bytes,2,rep,name=categories,proto3" json:"categories,omitempty"`
	Status               []Status    `protobuf:"varint,3,rep,packed,name=status,proto3,enum=domain.Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ca831ad796a6446, []int{2}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetRange() *Range {
	if m != nil {
		return m.Range
	}
	return nil
}

func (m *Filter) GetCategories() []*Category {
	if m != nil {
		return m.Categories
	}
	return nil
}

func (m *Filter) GetStatus() []Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*Range)(nil), "domain.Range")
	proto.RegisterType((*Category)(nil), "domain.Category")
	proto.RegisterType((*Filter)(nil), "domain.Filter")
}

func init() { proto.RegisterFile("domains/filter.proto", fileDescriptor_4ca831ad796a6446) }

var fileDescriptor_4ca831ad796a6446 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x49, 0x63, 0x83, 0x9d, 0xd2, 0x52, 0x16, 0x0f, 0xc1, 0x53, 0x88, 0x20, 0xb9, 0x34,
	0x91, 0x78, 0x10, 0xf1, 0x20, 0x58, 0xf0, 0x03, 0xac, 0x07, 0xc1, 0x8b, 0x6c, 0x92, 0x35, 0x19,
	0xda, 0xdd, 0xc8, 0xfe, 0x91, 0xfa, 0x09, 0xfc, 0xda, 0x92, 0xdd, 0xa4, 0x78, 0xf4, 0x36, 0xbc,
	0xdf, 0x9b, 0xd9, 0xb7, 0x0f, 0x2e, 0x9a, 0x5e, 0x30, 0x94, 0xba, 0xf8, 0xc0, 0x83, 0xe1, 0x2a,
	0xff, 0x54, 0xbd, 0xe9, 0x49, 0xe4, 0xd5, 0xcb, 0x13, 0xd5, 0x86, 0x19, 0xab, 0x3d, 0x4d, 0x1f,
	0x61, 0x4e, 0x99, 0x6c, 0x39, 0xd9, 0x40, 0x28, 0x50, 0xc6, 0x41, 0x12, 0x64, 0x2b, 0x3a, 0x8c,
	0x4e, 0x61, 0xc7, 0x78, 0x36, 0x2a, 0xec, 0x48, 0x08, 0x9c, 0x59, 0x89, 0x26, 0x0e, 0x93, 0x20,
	0x5b, 0x50, 0x37, 0xa7, 0xaf, 0x70, 0xbe, 0x63, 0x86, 0xb7, 0xbd, 0xfa, 0x1e, 0xb8, 0x64, 0x82,
	0xbb, 0x23, 0x0b, 0xea, 0x66, 0x72, 0x07, 0x6b, 0x6d, 0xab, 0xf7, 0xda, 0x7b, 0x90, 0xeb, 0x78,
	0x96, 0x84, 0xd9, 0xb2, 0xdc, 0xe4, 0x3e, 0x4f, 0x3e, 0x6d, 0xd3, 0x95, 0xb6, 0xd5, 0xee, 0x64,
	0x4b, 0x7f, 0x02, 0x88, 0x9e, 0xdd, 0x47, 0xc8, 0x15, 0xcc, 0xd5, 0x10, 0xd2, 0x1d, 0x5e, 0x96,
	0xab, 0x69, 0xd5, 0x25, 0xa7, 0x9e, 0x91, 0x1b, 0x80, 0x7f, 0x3c, 0xf2, 0xc7, 0x43, 0xae, 0x21,
	0xf2, 0x5d, 0xc4, 0x61, 0x12, 0x66, 0xeb, 0x72, 0x3d, 0xb9, 0x5f, 0x9c, 0x4a, 0x47, 0xfa, 0xf4,
	0xf0, 0x76, 0xdf, 0xa2, 0xe9, 0x6c, 0x95, 0xd7, 0xbd, 0x28, 0xf0, 0xd0, 0x31, 0x21, 0xba, 0xa6,
	0x29, 0xf6, 0xb6, 0x61, 0x7b, 0xdc, 0x5a, 0xcd, 0xd5, 0x56, 0x73, 0xf5, 0x85, 0x35, 0x2f, 0xb8,
	0x34, 0x68, 0x90, 0xeb, 0x62, 0x6c, 0xbb, 0x8a, 0x5c, 0xcf, 0xb7, 0xbf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x17, 0x81, 0xf9, 0xc7, 0x9d, 0x01, 0x00, 0x00,
}
