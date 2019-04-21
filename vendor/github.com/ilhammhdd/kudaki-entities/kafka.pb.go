// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kafka.proto

package kudaki_entities

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

type Topics int32

const (
	Topics_SIGNED_UP                    Topics = 0
	Topics_USER_VERIFICATION_EMAIL_SENT Topics = 1
	Topics_SIGN_UP_REQUESTED            Topics = 2
	Topics_VERIFY_USER_REQUESTED        Topics = 3
)

var Topics_name = map[int32]string{
	0: "SIGNED_UP",
	1: "USER_VERIFICATION_EMAIL_SENT",
	2: "SIGN_UP_REQUESTED",
	3: "VERIFY_USER_REQUESTED",
}

var Topics_value = map[string]int32{
	"SIGNED_UP":                    0,
	"USER_VERIFICATION_EMAIL_SENT": 1,
	"SIGN_UP_REQUESTED":            2,
	"VERIFY_USER_REQUESTED":        3,
}

func (x Topics) String() string {
	return proto.EnumName(Topics_name, int32(x))
}

func (Topics) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_68928ed13de9fb92, []int{0}
}

func init() {
	proto.RegisterEnum("entities.Topics", Topics_name, Topics_value)
}

func init() { proto.RegisterFile("kafka.proto", fileDescriptor_68928ed13de9fb92) }

var fileDescriptor_68928ed13de9fb92 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x4e, 0x4c, 0xcb,
	0x4e, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0xcd, 0x2b, 0xc9, 0x2c, 0xc9, 0x4c,
	0x2d, 0xd6, 0xca, 0xe6, 0x62, 0x0b, 0xc9, 0x2f, 0xc8, 0x4c, 0x2e, 0x16, 0xe2, 0xe5, 0xe2, 0x0c,
	0xf6, 0x74, 0xf7, 0x73, 0x75, 0x89, 0x0f, 0x0d, 0x10, 0x60, 0x10, 0x52, 0xe0, 0x92, 0x09, 0x0d,
	0x76, 0x0d, 0x8a, 0x0f, 0x73, 0x0d, 0xf2, 0x74, 0xf3, 0x74, 0x76, 0x0c, 0xf1, 0xf4, 0xf7, 0x8b,
	0x77, 0xf5, 0x75, 0xf4, 0xf4, 0x89, 0x0f, 0x76, 0xf5, 0x0b, 0x11, 0x60, 0x14, 0x12, 0xe5, 0x12,
	0x04, 0x69, 0x88, 0x0f, 0x0d, 0x88, 0x0f, 0x72, 0x0d, 0x0c, 0x75, 0x0d, 0x0e, 0x71, 0x75, 0x11,
	0x60, 0x12, 0x92, 0xe4, 0x12, 0x05, 0xeb, 0x89, 0x8c, 0x07, 0xeb, 0x47, 0x48, 0x31, 0x3b, 0xa9,
	0x45, 0xa9, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x67, 0xe6, 0x64,
	0x24, 0xe6, 0xe6, 0x66, 0xa4, 0xa4, 0xe8, 0x67, 0x97, 0xa6, 0x24, 0x66, 0x67, 0xea, 0xc2, 0x1c,
	0x95, 0xc4, 0x06, 0x76, 0xa5, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x1c, 0xaf, 0xfd, 0xb4,
	0x00, 0x00, 0x00,
}