// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/storefront.proto

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

func init() { proto.RegisterFile("services/storefront.proto", fileDescriptor_9590186f48f18dfd) }

var fileDescriptor_9590186f48f18dfd = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4a, 0x03, 0x31,
	0x14, 0x45, 0xc1, 0x85, 0x42, 0x96, 0xaf, 0x15, 0x31, 0xad, 0xfa, 0x05, 0x4d, 0x40, 0x77, 0xe2,
	0xa6, 0xa2, 0x82, 0x1b, 0x41, 0xc7, 0x55, 0x77, 0xd3, 0xbe, 0xa7, 0x13, 0x3a, 0x49, 0x34, 0x79,
	0x33, 0xe0, 0xdf, 0xfa, 0x29, 0x92, 0x76, 0x3a, 0x92, 0x3a, 0x6e, 0xef, 0xb9, 0x39, 0x99, 0xc9,
	0x15, 0xa7, 0x91, 0x42, 0x6b, 0x56, 0x14, 0x75, 0x64, 0x1f, 0xe8, 0x2d, 0x78, 0xc7, 0xea, 0x23,
	0x78, 0xf6, 0x70, 0xd4, 0x21, 0x79, 0xbc, 0xf2, 0xd6, 0x96, 0x0e, 0xa3, 0x0e, 0xe4, 0xb8, 0xac,
	0xb7, 0x5c, 0x8e, 0x3f, 0x1b, 0x0a, 0x86, 0xf6, 0xd2, 0x11, 0xb5, 0xe4, 0x38, 0x0f, 0x2f, 0xbf,
	0x0f, 0x84, 0x28, 0x92, 0xff, 0x21, 0xf9, 0xe1, 0x59, 0x8c, 0xe6, 0x88, 0x8f, 0x4c, 0xf6, 0xd5,
	0x17, 0xfd, 0xb5, 0x30, 0x55, 0xdd, 0x45, 0x6a, 0x80, 0xca, 0xa9, 0xda, 0x98, 0x55, 0x02, 0x73,
	0x44, 0xc2, 0xec, 0xec, 0x93, 0x80, 0x7b, 0x34, 0xfc, 0x9b, 0xa4, 0x1a, 0x4c, 0x7a, 0xe3, 0x5f,
	0x28, 0x27, 0x9d, 0x30, 0x8f, 0x53, 0x91, 0x10, 0x16, 0xe2, 0xe4, 0x85, 0x38, 0x18, 0x6a, 0x29,
	0xe7, 0x11, 0xce, 0x55, 0xfa, 0xf1, 0x2f, 0xf5, 0x0f, 0x97, 0x17, 0x83, 0xde, 0xb8, 0xab, 0x23,
	0x14, 0x62, 0x7c, 0x47, 0x35, 0xf1, 0xde, 0x49, 0x38, 0xeb, 0xbf, 0x76, 0x08, 0xf7, 0x0f, 0x90,
	0xc7, 0xdb, 0x2a, 0xde, 0xde, 0x2c, 0xae, 0xdf, 0x0d, 0x57, 0xcd, 0x32, 0x49, 0xb4, 0xa9, 0xab,
	0xd2, 0xda, 0x0a, 0x51, 0xaf, 0x1b, 0x2c, 0xd7, 0x66, 0xd6, 0x44, 0x0a, 0xb3, 0x6e, 0x50, 0x4d,
	0x8e, 0x0d, 0xa7, 0xe5, 0x76, 0xe3, 0x2f, 0x0f, 0x37, 0x3b, 0x5d, 0xfd, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xb7, 0xd2, 0x47, 0xcd, 0x0f, 0x02, 0x00, 0x00,
}
