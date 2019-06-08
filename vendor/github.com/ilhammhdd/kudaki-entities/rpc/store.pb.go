// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/store.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	events "github.com/ilhammhdd/kudaki-entities/events"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xbf, 0x4a, 0xc7, 0x30,
	0x14, 0x85, 0x41, 0xd1, 0x21, 0x0e, 0x62, 0xf8, 0x4d, 0xd5, 0x45, 0x70, 0x10, 0xc1, 0x04, 0xf4,
	0x05, 0xac, 0xb8, 0xb8, 0xb6, 0xb8, 0x74, 0x4b, 0x73, 0xaf, 0x26, 0xb4, 0x69, 0x62, 0x72, 0xeb,
	0x9b, 0xfa, 0x3e, 0xd2, 0xb4, 0x52, 0xff, 0x54, 0xdc, 0x92, 0xf3, 0x9d, 0xf3, 0x2d, 0x97, 0x1d,
	0xc7, 0xa0, 0x65, 0x22, 0x1f, 0x51, 0x84, 0xe8, 0xc9, 0xf3, 0xfd, 0x18, 0x74, 0xc1, 0xf1, 0x0d,
	0x07, 0x4a, 0x5f, 0xc1, 0xcd, 0xfb, 0x1e, 0x3b, 0xa8, 0xa7, 0x3f, 0xaf, 0xd8, 0x49, 0x09, 0x90,
	0xdf, 0xcf, 0xd1, 0x0f, 0xf4, 0x48, 0xe8, 0xf8, 0xb9, 0xc8, 0x1b, 0xf1, 0x8b, 0x54, 0xf8, 0x3a,
	0x62, 0x22, 0x84, 0xa2, 0x58, 0x2a, 0xdf, 0x79, 0x09, 0x80, 0xc0, 0x1b, 0xb6, 0x7b, 0xc0, 0x1e,
	0x09, 0x7f, 0x68, 0x2f, 0x96, 0xcd, 0x16, 0x5c, 0xcd, 0x67, 0x9b, 0xe6, 0x79, 0x93, 0xdd, 0x4f,
	0x01, 0xd4, 0x9f, 0xee, 0x2d, 0xf8, 0x9f, 0x7b, 0xde, 0x00, 0xbf, 0x63, 0x47, 0x35, 0xaa, 0xa8,
	0xcd, 0x14, 0x26, 0x7e, 0xfa, 0x59, 0x5e, 0xb3, 0xd5, 0xb4, 0x5b, 0x60, 0x8e, 0xe7, 0x06, 0xc2,
	0xfd, 0x55, 0x73, 0xf9, 0x62, 0xc9, 0x8c, 0xad, 0xd0, 0xde, 0x49, 0xdb, 0x1b, 0xe5, 0x9c, 0x01,
	0x90, 0xdd, 0x08, 0xaa, 0xb3, 0xd7, 0x38, 0x90, 0x25, 0x8b, 0x49, 0xc6, 0xa0, 0xdb, 0xc3, 0x7c,
	0x8a, 0xdb, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x37, 0xec, 0x03, 0x27, 0xb6, 0x01, 0x00, 0x00,
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
	AddStorefrontItem(ctx context.Context, in *events.AddStorefrontItemRequested, opts ...grpc.CallOption) (*events.StorefrontItemAdded, error)
	DeleteStorefrontItem(ctx context.Context, in *events.DeleteStorefrontItemRequested, opts ...grpc.CallOption) (*events.StorefrontItemDeleted, error)
	UpdateStorefrontItem(ctx context.Context, in *events.UpdateStorefrontItemRequested, opts ...grpc.CallOption) (*events.StorefrontItemUpdated, error)
	SearchItems(ctx context.Context, in *events.SearchItemsRequested, opts ...grpc.CallOption) (*events.ItemsSearched, error)
}

type storeClient struct {
	cc *grpc.ClientConn
}

func NewStoreClient(cc *grpc.ClientConn) StoreClient {
	return &storeClient{cc}
}

func (c *storeClient) AddStorefrontItem(ctx context.Context, in *events.AddStorefrontItemRequested, opts ...grpc.CallOption) (*events.StorefrontItemAdded, error) {
	out := new(events.StorefrontItemAdded)
	err := c.cc.Invoke(ctx, "/rpc.Store/AddStorefrontItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) DeleteStorefrontItem(ctx context.Context, in *events.DeleteStorefrontItemRequested, opts ...grpc.CallOption) (*events.StorefrontItemDeleted, error) {
	out := new(events.StorefrontItemDeleted)
	err := c.cc.Invoke(ctx, "/rpc.Store/DeleteStorefrontItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) UpdateStorefrontItem(ctx context.Context, in *events.UpdateStorefrontItemRequested, opts ...grpc.CallOption) (*events.StorefrontItemUpdated, error) {
	out := new(events.StorefrontItemUpdated)
	err := c.cc.Invoke(ctx, "/rpc.Store/UpdateStorefrontItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) SearchItems(ctx context.Context, in *events.SearchItemsRequested, opts ...grpc.CallOption) (*events.ItemsSearched, error) {
	out := new(events.ItemsSearched)
	err := c.cc.Invoke(ctx, "/rpc.Store/SearchItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServer is the server API for Store service.
type StoreServer interface {
	AddStorefrontItem(context.Context, *events.AddStorefrontItemRequested) (*events.StorefrontItemAdded, error)
	DeleteStorefrontItem(context.Context, *events.DeleteStorefrontItemRequested) (*events.StorefrontItemDeleted, error)
	UpdateStorefrontItem(context.Context, *events.UpdateStorefrontItemRequested) (*events.StorefrontItemUpdated, error)
	SearchItems(context.Context, *events.SearchItemsRequested) (*events.ItemsSearched, error)
}

// UnimplementedStoreServer can be embedded to have forward compatible implementations.
type UnimplementedStoreServer struct {
}

func (*UnimplementedStoreServer) AddStorefrontItem(ctx context.Context, req *events.AddStorefrontItemRequested) (*events.StorefrontItemAdded, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStorefrontItem not implemented")
}
func (*UnimplementedStoreServer) DeleteStorefrontItem(ctx context.Context, req *events.DeleteStorefrontItemRequested) (*events.StorefrontItemDeleted, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStorefrontItem not implemented")
}
func (*UnimplementedStoreServer) UpdateStorefrontItem(ctx context.Context, req *events.UpdateStorefrontItemRequested) (*events.StorefrontItemUpdated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStorefrontItem not implemented")
}
func (*UnimplementedStoreServer) SearchItems(ctx context.Context, req *events.SearchItemsRequested) (*events.ItemsSearched, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchItems not implemented")
}

func RegisterStoreServer(s *grpc.Server, srv StoreServer) {
	s.RegisterService(&_Store_serviceDesc, srv)
}

func _Store_AddStorefrontItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(events.AddStorefrontItemRequested)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).AddStorefrontItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Store/AddStorefrontItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).AddStorefrontItem(ctx, req.(*events.AddStorefrontItemRequested))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_DeleteStorefrontItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(events.DeleteStorefrontItemRequested)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).DeleteStorefrontItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Store/DeleteStorefrontItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).DeleteStorefrontItem(ctx, req.(*events.DeleteStorefrontItemRequested))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_UpdateStorefrontItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(events.UpdateStorefrontItemRequested)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).UpdateStorefrontItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Store/UpdateStorefrontItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).UpdateStorefrontItem(ctx, req.(*events.UpdateStorefrontItemRequested))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_SearchItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(events.SearchItemsRequested)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).SearchItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Store/SearchItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).SearchItems(ctx, req.(*events.SearchItemsRequested))
	}
	return interceptor(ctx, in, info, handler)
}

var _Store_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Store",
	HandlerType: (*StoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddStorefrontItem",
			Handler:    _Store_AddStorefrontItem_Handler,
		},
		{
			MethodName: "DeleteStorefrontItem",
			Handler:    _Store_DeleteStorefrontItem_Handler,
		},
		{
			MethodName: "UpdateStorefrontItem",
			Handler:    _Store_UpdateStorefrontItem_Handler,
		},
		{
			MethodName: "SearchItems",
			Handler:    _Store_SearchItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/store.proto",
}
