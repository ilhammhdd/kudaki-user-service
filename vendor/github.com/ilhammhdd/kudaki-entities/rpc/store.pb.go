// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/store.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	events "github.com/ilhammhdd/kudaki-entities/events"
	grpc "google.golang.org/grpc"
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

func init() { proto.RegisterFile("rpc/store.proto", fileDescriptor_4e215feb7542d463) }

var fileDescriptor_4e215feb7542d463 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8e, 0xbd, 0x0a, 0x02, 0x31,
	0x10, 0x84, 0x0b, 0xd1, 0x22, 0x8d, 0x90, 0x32, 0x9d, 0x9d, 0x08, 0x26, 0xa0, 0x4f, 0x70, 0x76,
	0xb6, 0xda, 0x69, 0x75, 0x97, 0x5d, 0x4d, 0x38, 0xf3, 0xe3, 0x66, 0xe3, 0xf3, 0x8b, 0xc1, 0x42,
	0x8b, 0xeb, 0x66, 0xf8, 0xbe, 0x81, 0x11, 0x4b, 0xca, 0xd6, 0x14, 0x4e, 0x84, 0x3a, 0x53, 0xe2,
	0x24, 0x67, 0x94, 0xad, 0x92, 0xf8, 0xc2, 0xc8, 0xe5, 0x17, 0xec, 0x40, 0xcc, 0xcf, 0x9f, 0x2a,
	0xaf, 0x42, 0x75, 0x00, 0x2d, 0xdf, 0x28, 0x45, 0x3e, 0x32, 0x86, 0x13, 0x3e, 0x2b, 0x16, 0x46,
	0x90, 0x2b, 0xdd, 0xb6, 0x7a, 0x5a, 0x51, 0xea, 0xab, 0xfc, 0xf3, 0x0e, 0x00, 0xe1, 0xb0, 0xb9,
	0xac, 0xef, 0x9e, 0x5d, 0x1d, 0xb4, 0x4d, 0xc1, 0xf8, 0x87, 0xeb, 0x43, 0x70, 0x00, 0x66, 0xac,
	0xd0, 0x8f, 0x7e, 0x8b, 0x91, 0x3d, 0x7b, 0x2c, 0x86, 0xb2, 0x1d, 0x16, 0xed, 0xd8, 0xfe, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0x29, 0x20, 0x97, 0x05, 0xc4, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StoreClient is the client API for Store service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoreClient interface {
	AddStorefrontItemRequested(ctx context.Context, in *events.AddStorefrontItemRequested, opts ...grpc.CallOption) (*events.StorefrontItemAdded, error)
}

type storeClient struct {
	cc *grpc.ClientConn
}

func NewStoreClient(cc *grpc.ClientConn) StoreClient {
	return &storeClient{cc}
}

func (c *storeClient) AddStorefrontItemRequested(ctx context.Context, in *events.AddStorefrontItemRequested, opts ...grpc.CallOption) (*events.StorefrontItemAdded, error) {
	out := new(events.StorefrontItemAdded)
	err := c.cc.Invoke(ctx, "/rpc.Store/AddStorefrontItemRequested", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServer is the server API for Store service.
type StoreServer interface {
	AddStorefrontItemRequested(context.Context, *events.AddStorefrontItemRequested) (*events.StorefrontItemAdded, error)
}

func RegisterStoreServer(s *grpc.Server, srv StoreServer) {
	s.RegisterService(&_Store_serviceDesc, srv)
}

func _Store_AddStorefrontItemRequested_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(events.AddStorefrontItemRequested)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).AddStorefrontItemRequested(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Store/AddStorefrontItemRequested",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).AddStorefrontItemRequested(ctx, req.(*events.AddStorefrontItemRequested))
	}
	return interceptor(ctx, in, info, handler)
}

var _Store_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Store",
	HandlerType: (*StoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddStorefrontItemRequested",
			Handler:    _Store_AddStorefrontItemRequested_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/store.proto",
}
