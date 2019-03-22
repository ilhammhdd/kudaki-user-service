// Code generated by protoc-gen-go. DO NOT EDIT.
// source: domains/status.proto

package domains

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Status int32

const (
	Status_ON_PROGRESS  Status = 0
	Status_SUCCESS      Status = 1
	Status_FAILED       Status = 2
	Status_CANCELED     Status = 3
	Status_PENDING      Status = 4
	Status_INSTALLED    Status = 5
	Status_PAID_OFF     Status = 6
	Status_NOT_PAID_YET Status = 7
	Status_AVAILABLE    Status = 8
	Status_UNAVAILABLE  Status = 9
	Status_OUT_OF_STOCK Status = 10
)

var Status_name = map[int32]string{
	0:  "ON_PROGRESS",
	1:  "SUCCESS",
	2:  "FAILED",
	3:  "CANCELED",
	4:  "PENDING",
	5:  "INSTALLED",
	6:  "PAID_OFF",
	7:  "NOT_PAID_YET",
	8:  "AVAILABLE",
	9:  "UNAVAILABLE",
	10: "OUT_OF_STOCK",
}

var Status_value = map[string]int32{
	"ON_PROGRESS":  0,
	"SUCCESS":      1,
	"FAILED":       2,
	"CANCELED":     3,
	"PENDING":      4,
	"INSTALLED":    5,
	"PAID_OFF":     6,
	"NOT_PAID_YET": 7,
	"AVAILABLE":    8,
	"UNAVAILABLE":  9,
	"OUT_OF_STOCK": 10,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cf8a23022947bfa2, []int{0}
}

type EventStatus struct {
	Code                 int32                `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Messages             []string             `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
	Data                 []byte               `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EventStatus) Reset()         { *m = EventStatus{} }
func (m *EventStatus) String() string { return proto.CompactTextString(m) }
func (*EventStatus) ProtoMessage()    {}
func (*EventStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf8a23022947bfa2, []int{0}
}

func (m *EventStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventStatus.Unmarshal(m, b)
}
func (m *EventStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventStatus.Marshal(b, m, deterministic)
}
func (m *EventStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventStatus.Merge(m, src)
}
func (m *EventStatus) XXX_Size() int {
	return xxx_messageInfo_EventStatus.Size(m)
}
func (m *EventStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_EventStatus.DiscardUnknown(m)
}

var xxx_messageInfo_EventStatus proto.InternalMessageInfo

func (m *EventStatus) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *EventStatus) GetMessages() []string {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *EventStatus) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *EventStatus) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterEnum("domain.Status", Status_name, Status_value)
	proto.RegisterType((*EventStatus)(nil), "domain.EventStatus")
}

func init() { proto.RegisterFile("domains/status.proto", fileDescriptor_cf8a23022947bfa2) }

var fileDescriptor_cf8a23022947bfa2 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x91, 0x4d, 0x6b, 0x83, 0x30,
	0x18, 0xc7, 0x67, 0x5f, 0x6c, 0x8d, 0x1d, 0x0b, 0x61, 0x07, 0xe9, 0x65, 0xb2, 0x93, 0x0c, 0xaa,
	0xb0, 0x5d, 0x36, 0x76, 0xb2, 0x56, 0x8b, 0x4c, 0xb4, 0xa8, 0x1d, 0x6c, 0x17, 0x49, 0x6b, 0x66,
	0x43, 0x6b, 0x2d, 0x4d, 0xec, 0x97, 0xd8, 0x77, 0xd9, 0x67, 0x1c, 0xb1, 0x6f, 0xb7, 0xe7, 0xff,
	0xcf, 0xef, 0x21, 0x3f, 0x12, 0x70, 0x9f, 0x57, 0x25, 0xa6, 0x5b, 0x66, 0x31, 0x8e, 0x79, 0xcd,
	0xcc, 0xdd, 0xbe, 0xe2, 0x15, 0x92, 0x8f, 0xed, 0xf0, 0xa1, 0xa8, 0xaa, 0x62, 0x43, 0xac, 0xa6,
	0x5d, 0xd4, 0x3f, 0x16, 0xa7, 0x25, 0x61, 0x1c, 0x97, 0xbb, 0x23, 0xf8, 0xf8, 0x2b, 0x01, 0xd5,
	0x3d, 0x90, 0x2d, 0x4f, 0x9a, 0x75, 0x84, 0x40, 0x67, 0x59, 0xe5, 0x44, 0x93, 0x74, 0xc9, 0xe8,
	0xc6, 0xcd, 0x8c, 0x86, 0xa0, 0x5f, 0x12, 0xc6, 0x70, 0x41, 0x98, 0xd6, 0xd2, 0xdb, 0x86, 0x12,
	0x5f, 0xb2, 0xe0, 0x73, 0xcc, 0xb1, 0xd6, 0xd6, 0x25, 0x63, 0x10, 0x37, 0x33, 0x7a, 0x05, 0xca,
	0xe5, 0x1a, 0xad, 0xa3, 0x4b, 0x86, 0xfa, 0x3c, 0x34, 0x8f, 0x22, 0xe6, 0x59, 0xc4, 0x4c, 0xcf,
	0x44, 0x7c, 0x85, 0x9f, 0xfe, 0x24, 0x20, 0x9f, 0x44, 0xee, 0x80, 0x1a, 0x85, 0xd9, 0x2c, 0x8e,
	0xa6, 0xb1, 0x9b, 0x24, 0xf0, 0x06, 0xa9, 0xa0, 0x97, 0xcc, 0x1d, 0x47, 0x04, 0x09, 0x01, 0x20,
	0x7b, 0xb6, 0x1f, 0xb8, 0x13, 0xd8, 0x42, 0x03, 0xd0, 0x77, 0xec, 0xd0, 0x71, 0x45, 0x6a, 0x0b,
	0x6c, 0xe6, 0x86, 0x13, 0x3f, 0x9c, 0xc2, 0x0e, 0xba, 0x05, 0x8a, 0x1f, 0x26, 0xa9, 0x1d, 0x88,
	0xb3, 0xae, 0x20, 0x67, 0xb6, 0x3f, 0xc9, 0x22, 0xcf, 0x83, 0x32, 0x82, 0x60, 0x10, 0x46, 0x69,
	0xd6, 0x34, 0x5f, 0x6e, 0x0a, 0x7b, 0x02, 0xb7, 0x3f, 0x6d, 0x3f, 0xb0, 0xc7, 0x81, 0x0b, 0xfb,
	0x42, 0x61, 0x1e, 0x5e, 0x0b, 0x45, 0x6c, 0x44, 0xf3, 0x34, 0x8b, 0xbc, 0x2c, 0x49, 0x23, 0xe7,
	0x03, 0x82, 0xf1, 0xfb, 0xf7, 0x5b, 0x41, 0xf9, 0xaa, 0x5e, 0x98, 0xcb, 0xaa, 0xb4, 0xe8, 0x66,
	0x85, 0xcb, 0x72, 0x95, 0xe7, 0xd6, 0xba, 0xce, 0xf1, 0x9a, 0x8e, 0x6a, 0x46, 0xf6, 0x23, 0x46,
	0xf6, 0x07, 0xba, 0x24, 0x16, 0xd9, 0x72, 0xca, 0x29, 0x61, 0xd6, 0xe9, 0xc7, 0x16, 0x72, 0xf3,
	0x18, 0x2f, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa6, 0xcf, 0x2f, 0x59, 0xc3, 0x01, 0x00, 0x00,
}
