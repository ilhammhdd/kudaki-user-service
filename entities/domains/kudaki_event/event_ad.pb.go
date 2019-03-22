// Code generated by protoc-gen-go. DO NOT EDIT.
// source: domains/kudaki_event/event_ad.proto

package kudaki_event

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	domains "github.com/ilhammhdd/kudaki-user-service/entities/domains"
	mountain "github.com/ilhammhdd/kudaki-user-service/entities/domains/mountain"
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

type EventAd struct {
	Uuid                 string               `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Start                *timestamp.Timestamp `protobuf:"bytes,3,opt,name=start,proto3" json:"start,omitempty"`
	End                  *timestamp.Timestamp `protobuf:"bytes,4,opt,name=end,proto3" json:"end,omitempty"`
	Mountain             *mountain.Mountain   `protobuf:"bytes,5,opt,name=mountain,proto3" json:"mountain,omitempty"`
	Organizer            *Organizer           `protobuf:"bytes,6,opt,name=organizer,proto3" json:"organizer,omitempty"`
	Description          string               `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Status               domains.Status       `protobuf:"varint,8,opt,name=status,proto3,enum=domain.Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EventAd) Reset()         { *m = EventAd{} }
func (m *EventAd) String() string { return proto.CompactTextString(m) }
func (*EventAd) ProtoMessage()    {}
func (*EventAd) Descriptor() ([]byte, []int) {
	return fileDescriptor_71ffc53d3b504a23, []int{0}
}

func (m *EventAd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventAd.Unmarshal(m, b)
}
func (m *EventAd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventAd.Marshal(b, m, deterministic)
}
func (m *EventAd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAd.Merge(m, src)
}
func (m *EventAd) XXX_Size() int {
	return xxx_messageInfo_EventAd.Size(m)
}
func (m *EventAd) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAd.DiscardUnknown(m)
}

var xxx_messageInfo_EventAd proto.InternalMessageInfo

func (m *EventAd) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *EventAd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EventAd) GetStart() *timestamp.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *EventAd) GetEnd() *timestamp.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *EventAd) GetMountain() *mountain.Mountain {
	if m != nil {
		return m.Mountain
	}
	return nil
}

func (m *EventAd) GetOrganizer() *Organizer {
	if m != nil {
		return m.Organizer
	}
	return nil
}

func (m *EventAd) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EventAd) GetStatus() domains.Status {
	if m != nil {
		return m.Status
	}
	return domains.Status_ON_PROGRESS
}

func init() {
	proto.RegisterType((*EventAd)(nil), "domain.kudaki_event.EventAd")
}

func init() {
	proto.RegisterFile("domains/kudaki_event/event_ad.proto", fileDescriptor_71ffc53d3b504a23)
}

var fileDescriptor_71ffc53d3b504a23 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x4f, 0x4b, 0x23, 0x31,
	0x00, 0xc5, 0x99, 0xfe, 0x6f, 0x0a, 0x3d, 0x64, 0xf7, 0x90, 0xed, 0x61, 0x3b, 0xec, 0x8a, 0xf4,
	0x60, 0x13, 0xa9, 0x78, 0xf3, 0xa2, 0xa0, 0x78, 0x11, 0x61, 0xf4, 0xe4, 0xa5, 0xa4, 0x4d, 0x9c,
	0x86, 0x36, 0x49, 0xc9, 0x9f, 0x1e, 0xfc, 0x94, 0x7e, 0x24, 0x99, 0x64, 0x32, 0xf6, 0x50, 0xf0,
	0x32, 0x3c, 0xde, 0xfc, 0x5e, 0xc2, 0x7b, 0x01, 0xff, 0x99, 0x96, 0x54, 0x28, 0x4b, 0xb6, 0x9e,
	0xd1, 0xad, 0x58, 0xf2, 0x03, 0x57, 0x8e, 0x84, 0xef, 0x92, 0x32, 0xbc, 0x37, 0xda, 0x69, 0xf8,
	0x2b, 0x42, 0xf8, 0x98, 0x99, 0x4c, 0x4b, 0xad, 0xcb, 0x1d, 0x27, 0x01, 0x59, 0xf9, 0x77, 0xe2,
	0x84, 0xe4, 0xd6, 0x51, 0xb9, 0x8f, 0xa9, 0xc9, 0x34, 0x1d, 0x2d, 0xb5, 0x57, 0x8e, 0x0a, 0xd5,
	0x88, 0x1a, 0x38, 0x3b, 0x79, 0xb7, 0x36, 0x25, 0x55, 0xe2, 0x83, 0x9b, 0x9a, 0xfa, 0x9d, 0x28,
	0xeb, 0xa8, 0xf3, 0x36, 0xba, 0xff, 0x3e, 0x5b, 0xa0, 0x7f, 0x5f, 0xf1, 0xb7, 0x0c, 0x42, 0xd0,
	0xf1, 0x5e, 0x30, 0x94, 0xe5, 0xd9, 0x6c, 0x58, 0x04, 0x5d, 0x79, 0x8a, 0x4a, 0x8e, 0x5a, 0xd1,
	0xab, 0x34, 0xbc, 0x04, 0x5d, 0xeb, 0xa8, 0x71, 0xa8, 0x9d, 0x67, 0xb3, 0xd1, 0x62, 0x82, 0x63,
	0x03, 0x9c, 0x1a, 0xe0, 0xd7, 0xd4, 0xa0, 0x88, 0x20, 0xbc, 0x00, 0x6d, 0xae, 0x18, 0xea, 0xfc,
	0xc8, 0x57, 0x18, 0xbc, 0x06, 0x83, 0xd4, 0x10, 0x75, 0x43, 0xe4, 0x0f, 0xae, 0x97, 0x6b, 0x9a,
	0x3f, 0xd5, 0xa2, 0x68, 0x50, 0x78, 0x03, 0x86, 0x4d, 0x67, 0xd4, 0x0b, 0xb9, 0xbf, 0xf8, 0xc4,
	0xe2, 0xf8, 0x39, 0x51, 0xc5, 0x77, 0x00, 0xe6, 0x60, 0xc4, 0xb8, 0x5d, 0x1b, 0xb1, 0x77, 0x42,
	0x2b, 0xd4, 0x0f, 0x7d, 0x8f, 0x2d, 0x78, 0x0e, 0x7a, 0x71, 0x3a, 0x34, 0xc8, 0xb3, 0xd9, 0x78,
	0x31, 0x4e, 0x87, 0xbf, 0x04, 0xb7, 0xa8, 0xff, 0xde, 0x3d, 0xbe, 0x3d, 0x94, 0xc2, 0x6d, 0xfc,
	0x0a, 0xaf, 0xb5, 0x24, 0x62, 0xb7, 0xa1, 0x52, 0x6e, 0x18, 0xab, 0x5f, 0x67, 0xee, 0x2d, 0x37,
	0x73, 0xcb, 0xcd, 0x41, 0xac, 0x39, 0xe1, 0xca, 0x09, 0x27, 0xb8, 0x25, 0xa7, 0x9e, 0x70, 0xd5,
	0x0b, 0x0b, 0x5d, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0x20, 0xa8, 0x22, 0xd2, 0x5d, 0x02, 0x00,
	0x00,
}