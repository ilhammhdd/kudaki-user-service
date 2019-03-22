// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/profile.proto

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

func init() { proto.RegisterFile("services/profile.proto", fileDescriptor_acd4ea30cd9ea06e) }

var fileDescriptor_acd4ea30cd9ea06e = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4d, 0x6a, 0xc3, 0x30,
	0x10, 0x85, 0xe9, 0xa6, 0x06, 0x41, 0x29, 0x55, 0xa9, 0x0b, 0xba, 0x83, 0x25, 0x68, 0x77, 0xa5,
	0x64, 0xe1, 0x13, 0x84, 0x40, 0x36, 0xd9, 0xc9, 0xd6, 0x24, 0x1e, 0x6c, 0x59, 0x8e, 0x7e, 0x0c,
	0x39, 0x44, 0xee, 0x1c, 0x64, 0xcb, 0x60, 0x67, 0xa9, 0xef, 0xc1, 0xa7, 0xf7, 0x86, 0xe4, 0x0e,
	0xec, 0x88, 0x35, 0x38, 0x31, 0x58, 0x73, 0xc6, 0x0e, 0xf8, 0x60, 0x8d, 0x37, 0x34, 0x4b, 0x9c,
	0x7d, 0xd6, 0x46, 0x6b, 0xd9, 0x2b, 0x27, 0x82, 0x03, 0x3b, 0xa7, 0x8c, 0x5e, 0x03, 0x58, 0x84,
	0x0d, 0xfb, 0x80, 0x11, 0x7a, 0xbf, 0x46, 0x3f, 0xf7, 0x17, 0x92, 0xed, 0x67, 0x2d, 0xdd, 0x91,
	0xb7, 0xe3, 0xa0, 0xa4, 0x87, 0x05, 0xe4, 0x3c, 0x99, 0xf9, 0x86, 0xb3, 0x2f, 0x3e, 0x89, 0x78,
	0x7a, 0xcf, 0xa1, 0xa2, 0x25, 0x79, 0x3f, 0x80, 0xb7, 0x08, 0xe3, 0xca, 0x10, 0x6b, 0xdc, 0xf8,
	0x13, 0x67, 0xdf, 0x5b, 0xc3, 0x12, 0xab, 0xf2, 0xff, 0xf4, 0x77, 0x41, 0xdf, 0x84, 0x2a, 0x7e,
	0x2d, 0xb0, 0x6b, 0xa4, 0xd6, 0x8d, 0x52, 0xa2, 0x0d, 0x4a, 0xb6, 0x58, 0xc4, 0xe6, 0x45, 0xda,
	0x2d, 0xa0, 0xf7, 0xe8, 0xe3, 0xc4, 0xe5, 0x40, 0xd5, 0xeb, 0x34, 0xea, 0xf7, 0x11, 0x00, 0x00,
	0xff, 0xff, 0xf7, 0x3f, 0x90, 0xd5, 0x33, 0x01, 0x00, 0x00,
}