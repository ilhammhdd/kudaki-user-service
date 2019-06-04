// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/rental.proto

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

func init() { proto.RegisterFile("rpc/rental.proto", fileDescriptor_d1c93b114b4e8bf1) }

var fileDescriptor_d1c93b114b4e8bf1 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8e, 0xc1, 0x8a, 0xc2, 0x30,
	0x10, 0x86, 0x59, 0x16, 0x7a, 0xc8, 0xee, 0x61, 0x37, 0x7a, 0x90, 0xfa, 0x02, 0x22, 0x98, 0x80,
	0x1e, 0xbd, 0x58, 0x7b, 0xf2, 0x5a, 0x6f, 0xde, 0xd2, 0xcc, 0x60, 0x42, 0x9b, 0x26, 0xa6, 0x13,
	0x5f, 0xc1, 0xd7, 0x16, 0x4a, 0xc5, 0x7a, 0x9c, 0xef, 0x9b, 0x0f, 0x7e, 0xf6, 0x17, 0x83, 0x96,
	0x11, 0x3b, 0x52, 0xad, 0x08, 0xd1, 0x93, 0xe7, 0xdf, 0x31, 0xe8, 0x7c, 0x86, 0x77, 0xec, 0xa8,
	0xff, 0x30, 0xdb, 0xc7, 0x17, 0xcb, 0xaa, 0x01, 0xf0, 0x3d, 0xfb, 0x3d, 0xa7, 0xda, 0x59, 0x1a,
	0xef, 0x85, 0x18, 0x02, 0x51, 0x1a, 0xd4, 0x8d, 0x4f, 0x54, 0xe1, 0x2d, 0x61, 0x4f, 0x08, 0xf9,
	0xff, 0xd4, 0x20, 0xf8, 0x44, 0xfc, 0xc0, 0x7e, 0x0a, 0x80, 0x52, 0x45, 0x3a, 0x11, 0x3a, 0xbe,
	0x1c, 0x3f, 0x26, 0xec, 0x9d, 0xcf, 0x5f, 0xf9, 0x68, 0x0a, 0x00, 0x84, 0xe3, 0xfa, 0xb2, 0xba,
	0x5a, 0x32, 0xa9, 0x16, 0xda, 0x3b, 0x69, 0x5b, 0xa3, 0x9c, 0x33, 0x00, 0xb2, 0x49, 0xa0, 0x1a,
	0xbb, 0xc1, 0x8e, 0x2c, 0x59, 0xec, 0x65, 0x0c, 0xba, 0xce, 0x86, 0xf1, 0xbb, 0x67, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xb9, 0xca, 0x1a, 0xc6, 0xea, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RentalClient is the client API for Rental service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RentalClient interface {
	SubmitRental(ctx context.Context, in *events.CheckoutRequested, opts ...grpc.CallOption) (*events.Checkedout, error)
	AddCartItem(ctx context.Context, in *events.AddCartItemRequested, opts ...grpc.CallOption) (*events.CartItemAdded, error)
}

type rentalClient struct {
	cc *grpc.ClientConn
}

func NewRentalClient(cc *grpc.ClientConn) RentalClient {
	return &rentalClient{cc}
}

func (c *rentalClient) SubmitRental(ctx context.Context, in *events.CheckoutRequested, opts ...grpc.CallOption) (*events.Checkedout, error) {
	out := new(events.Checkedout)
	err := c.cc.Invoke(ctx, "/rpc.Rental/SubmitRental", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rentalClient) AddCartItem(ctx context.Context, in *events.AddCartItemRequested, opts ...grpc.CallOption) (*events.CartItemAdded, error) {
	out := new(events.CartItemAdded)
	err := c.cc.Invoke(ctx, "/rpc.Rental/AddCartItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RentalServer is the server API for Rental service.
type RentalServer interface {
	SubmitRental(context.Context, *events.CheckoutRequested) (*events.Checkedout, error)
	AddCartItem(context.Context, *events.AddCartItemRequested) (*events.CartItemAdded, error)
}

// UnimplementedRentalServer can be embedded to have forward compatible implementations.
type UnimplementedRentalServer struct {
}

func (*UnimplementedRentalServer) SubmitRental(ctx context.Context, req *events.CheckoutRequested) (*events.Checkedout, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitRental not implemented")
}
func (*UnimplementedRentalServer) AddCartItem(ctx context.Context, req *events.AddCartItemRequested) (*events.CartItemAdded, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCartItem not implemented")
}

func RegisterRentalServer(s *grpc.Server, srv RentalServer) {
	s.RegisterService(&_Rental_serviceDesc, srv)
}

func _Rental_SubmitRental_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(events.CheckoutRequested)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentalServer).SubmitRental(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Rental/SubmitRental",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentalServer).SubmitRental(ctx, req.(*events.CheckoutRequested))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rental_AddCartItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(events.AddCartItemRequested)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentalServer).AddCartItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Rental/AddCartItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentalServer).AddCartItem(ctx, req.(*events.AddCartItemRequested))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rental_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Rental",
	HandlerType: (*RentalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitRental",
			Handler:    _Rental_SubmitRental_Handler,
		},
		{
			MethodName: "AddCartItem",
			Handler:    _Rental_AddCartItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/rental.proto",
}
