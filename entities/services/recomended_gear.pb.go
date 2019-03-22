// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/recomended_gear.proto

package services

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/ilhammhdd/kudaki-user-service/entities/commands"
	_ "github.com/ilhammhdd/kudaki-user-service/entities/events"
	_ "github.com/ilhammhdd/kudaki-user-service/entities/queries"
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

func init() { proto.RegisterFile("services/recomended_gear.proto", fileDescriptor_41629076709a089e) }

var fileDescriptor_41629076709a089e = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x4f, 0x4b, 0xfb, 0x30,
	0x18, 0xc7, 0x4f, 0xbf, 0x9f, 0x90, 0x83, 0x60, 0x98, 0x1b, 0x64, 0xae, 0x7b, 0x05, 0x4b, 0x41,
	0x6f, 0xe2, 0x65, 0x32, 0xf1, 0x20, 0x78, 0xe8, 0x50, 0x71, 0x17, 0xc9, 0xfa, 0x3c, 0xac, 0x61,
	0x4b, 0x32, 0xd3, 0xb4, 0xe2, 0x6b, 0xf3, 0xcd, 0x49, 0xff, 0x6d, 0xd8, 0xa6, 0xbb, 0x3e, 0xdf,
	0x4f, 0x3f, 0x49, 0xf3, 0xf0, 0x25, 0x41, 0x8a, 0x36, 0x97, 0x31, 0xa6, 0xa1, 0xc5, 0xd8, 0x28,
	0xd4, 0x80, 0xf0, 0xb1, 0x41, 0x61, 0xf9, 0xde, 0x1a, 0x67, 0xe8, 0x59, 0x9d, 0xb3, 0x51, 0x6c,
	0x94, 0x12, 0x1a, 0xd2, 0x50, 0x99, 0x4c, 0x3b, 0x21, 0x75, 0x45, 0xb0, 0x4b, 0xcc, 0x51, 0xbb,
	0xce, 0x78, 0xf8, 0x99, 0xa1, 0x95, 0xd8, 0x9e, 0x5f, 0xff, 0xfc, 0x23, 0xe7, 0xd1, 0xe1, 0xa8,
	0x47, 0x14, 0x96, 0xae, 0xc8, 0x28, 0x42, 0x67, 0x25, 0xe6, 0xf8, 0x37, 0x49, 0x69, 0xc0, 0x0b,
	0xcd, 0x37, 0xef, 0xc9, 0xd9, 0x94, 0x97, 0xa7, 0xf3, 0xd6, 0xbc, 0xc1, 0x81, 0xbe, 0x91, 0xa1,
	0xff, 0x5b, 0x3a, 0x39, 0xa9, 0x66, 0x81, 0xd7, 0x7c, 0x14, 0x3f, 0x91, 0x8b, 0x39, 0x40, 0xcb,
	0xc9, 0x78, 0xfd, 0x4a, 0xbc, 0x93, 0x31, 0xe6, 0x15, 0xce, 0x01, 0x10, 0xe8, 0x33, 0xa1, 0x0f,
	0x20, 0x5d, 0xcb, 0x36, 0x3e, 0xd8, 0xba, 0x21, 0x1b, 0x7b, 0x75, 0x05, 0x88, 0x40, 0x97, 0x64,
	0xb0, 0xc0, 0x1d, 0xba, 0xee, 0x3f, 0x37, 0x46, 0x5f, 0xcc, 0xae, 0xbc, 0xce, 0x0a, 0x2d, 0xa5,
	0x2f, 0xfb, 0x57, 0x73, 0x42, 0xea, 0x8b, 0x7b, 0xa4, 0x15, 0x0a, 0xf4, 0x9d, 0x0c, 0x17, 0xe6,
	0x4b, 0x7b, 0xb4, 0xd3, 0xe3, 0x5d, 0xbd, 0x40, 0xcf, 0x86, 0x1a, 0x18, 0x68, 0x44, 0x06, 0x4b,
	0x14, 0x36, 0x4e, 0x3a, 0xcf, 0x5a, 0x2d, 0xde, 0x17, 0xb2, 0x89, 0x57, 0x5a, 0xa1, 0x08, 0xf7,
	0x77, 0xab, 0xdb, 0x8d, 0x74, 0x49, 0xb6, 0x2e, 0x2e, 0x17, 0xca, 0x5d, 0x22, 0x94, 0x4a, 0x00,
	0xc2, 0x6d, 0x06, 0x62, 0x2b, 0x67, 0x59, 0x8a, 0x76, 0x56, 0x37, 0x26, 0x44, 0xed, 0xa4, 0x2b,
	0x1a, 0xd0, 0x54, 0x6c, 0xfd, 0xbf, 0xac, 0xc0, 0xcd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7e,
	0x27, 0xc7, 0x4e, 0x75, 0x03, 0x00, 0x00,
}