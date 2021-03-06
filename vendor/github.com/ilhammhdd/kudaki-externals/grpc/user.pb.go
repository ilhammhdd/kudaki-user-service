// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/user.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	events "github.com/ilhammhdd/kudaki-entities/events"
	user "github.com/ilhammhdd/kudaki-entities/user"
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

type UserAuthenticationRequested struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Jwt                  string   `protobuf:"bytes,2,opt,name=jwt,proto3" json:"jwt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserAuthenticationRequested) Reset()         { *m = UserAuthenticationRequested{} }
func (m *UserAuthenticationRequested) String() string { return proto.CompactTextString(m) }
func (*UserAuthenticationRequested) ProtoMessage()    {}
func (*UserAuthenticationRequested) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f3a799ad09b2336, []int{0}
}

func (m *UserAuthenticationRequested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAuthenticationRequested.Unmarshal(m, b)
}
func (m *UserAuthenticationRequested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAuthenticationRequested.Marshal(b, m, deterministic)
}
func (m *UserAuthenticationRequested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAuthenticationRequested.Merge(m, src)
}
func (m *UserAuthenticationRequested) XXX_Size() int {
	return xxx_messageInfo_UserAuthenticationRequested.Size(m)
}
func (m *UserAuthenticationRequested) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAuthenticationRequested.DiscardUnknown(m)
}

var xxx_messageInfo_UserAuthenticationRequested proto.InternalMessageInfo

func (m *UserAuthenticationRequested) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UserAuthenticationRequested) GetJwt() string {
	if m != nil {
		return m.Jwt
	}
	return ""
}

type UserAuthenticated struct {
	Uid                  string         `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	EventStatus          *events.Status `protobuf:"bytes,2,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	User                 *user.User     `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UserAuthenticated) Reset()         { *m = UserAuthenticated{} }
func (m *UserAuthenticated) String() string { return proto.CompactTextString(m) }
func (*UserAuthenticated) ProtoMessage()    {}
func (*UserAuthenticated) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f3a799ad09b2336, []int{1}
}

func (m *UserAuthenticated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAuthenticated.Unmarshal(m, b)
}
func (m *UserAuthenticated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAuthenticated.Marshal(b, m, deterministic)
}
func (m *UserAuthenticated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAuthenticated.Merge(m, src)
}
func (m *UserAuthenticated) XXX_Size() int {
	return xxx_messageInfo_UserAuthenticated.Size(m)
}
func (m *UserAuthenticated) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAuthenticated.DiscardUnknown(m)
}

var xxx_messageInfo_UserAuthenticated proto.InternalMessageInfo

func (m *UserAuthenticated) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UserAuthenticated) GetEventStatus() *events.Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *UserAuthenticated) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

type UserAuthorizationRequested struct {
	Uid                  string    `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Role                 user.Role `protobuf:"varint,2,opt,name=role,proto3,enum=entities.user.Role" json:"role,omitempty"`
	Jwt                  string    `protobuf:"bytes,3,opt,name=jwt,proto3" json:"jwt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UserAuthorizationRequested) Reset()         { *m = UserAuthorizationRequested{} }
func (m *UserAuthorizationRequested) String() string { return proto.CompactTextString(m) }
func (*UserAuthorizationRequested) ProtoMessage()    {}
func (*UserAuthorizationRequested) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f3a799ad09b2336, []int{2}
}

func (m *UserAuthorizationRequested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAuthorizationRequested.Unmarshal(m, b)
}
func (m *UserAuthorizationRequested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAuthorizationRequested.Marshal(b, m, deterministic)
}
func (m *UserAuthorizationRequested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAuthorizationRequested.Merge(m, src)
}
func (m *UserAuthorizationRequested) XXX_Size() int {
	return xxx_messageInfo_UserAuthorizationRequested.Size(m)
}
func (m *UserAuthorizationRequested) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAuthorizationRequested.DiscardUnknown(m)
}

var xxx_messageInfo_UserAuthorizationRequested proto.InternalMessageInfo

func (m *UserAuthorizationRequested) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UserAuthorizationRequested) GetRole() user.Role {
	if m != nil {
		return m.Role
	}
	return user.Role_USER
}

func (m *UserAuthorizationRequested) GetJwt() string {
	if m != nil {
		return m.Jwt
	}
	return ""
}

type UserAuthorized struct {
	Uid                  string         `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	EventStatus          *events.Status `protobuf:"bytes,2,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	User                 *user.User     `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UserAuthorized) Reset()         { *m = UserAuthorized{} }
func (m *UserAuthorized) String() string { return proto.CompactTextString(m) }
func (*UserAuthorized) ProtoMessage()    {}
func (*UserAuthorized) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f3a799ad09b2336, []int{3}
}

func (m *UserAuthorized) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAuthorized.Unmarshal(m, b)
}
func (m *UserAuthorized) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAuthorized.Marshal(b, m, deterministic)
}
func (m *UserAuthorized) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAuthorized.Merge(m, src)
}
func (m *UserAuthorized) XXX_Size() int {
	return xxx_messageInfo_UserAuthorized.Size(m)
}
func (m *UserAuthorized) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAuthorized.DiscardUnknown(m)
}

var xxx_messageInfo_UserAuthorized proto.InternalMessageInfo

func (m *UserAuthorized) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UserAuthorized) GetEventStatus() *events.Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *UserAuthorized) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*UserAuthenticationRequested)(nil), "rpc.UserAuthenticationRequested")
	proto.RegisterType((*UserAuthenticated)(nil), "rpc.UserAuthenticated")
	proto.RegisterType((*UserAuthorizationRequested)(nil), "rpc.UserAuthorizationRequested")
	proto.RegisterType((*UserAuthorized)(nil), "rpc.UserAuthorized")
}

func init() { proto.RegisterFile("grpc/user.proto", fileDescriptor_3f3a799ad09b2336) }

var fileDescriptor_3f3a799ad09b2336 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0xa5, 0xa6, 0x2e, 0x78, 0xd7, 0x6d, 0xdd, 0x5b, 0x5c, 0x4a, 0x44, 0x5d, 0xfa, 0xb2, 0x22,
	0x9a, 0xe8, 0xfa, 0xec, 0xc3, 0x2a, 0x82, 0xc2, 0x2a, 0x92, 0xa2, 0x0f, 0xbe, 0x48, 0x36, 0x73,
	0x4d, 0x66, 0x9b, 0xce, 0xc4, 0xf9, 0xb0, 0x2a, 0x08, 0xfe, 0x51, 0xff, 0x8b, 0xcc, 0x24, 0x6d,
	0x52, 0x53, 0xf4, 0xcd, 0xb7, 0xde, 0x73, 0xce, 0x9c, 0x39, 0x87, 0xb9, 0x0d, 0x8c, 0x73, 0x55,
	0x65, 0xb1, 0xd5, 0xa4, 0xa2, 0x4a, 0x49, 0x23, 0x31, 0x50, 0x55, 0x16, 0x1e, 0xd2, 0x17, 0x12,
	0x46, 0x77, 0xf0, 0x70, 0xd2, 0x40, 0xda, 0xa4, 0xc6, 0xea, 0x06, 0x1c, 0x3b, 0x41, 0x47, 0x35,
	0x3b, 0x83, 0x5b, 0xef, 0x34, 0xa9, 0x33, 0x6b, 0x0a, 0x12, 0x86, 0x67, 0xa9, 0xe1, 0x52, 0x24,
	0xf4, 0xd9, 0x92, 0x36, 0xc4, 0xf0, 0x06, 0x04, 0x96, 0xb3, 0xe9, 0xe0, 0x78, 0x70, 0xef, 0x5a,
	0xe2, 0x7e, 0x3a, 0xe4, 0x72, 0x65, 0xa6, 0x57, 0x6a, 0xe4, 0x72, 0x65, 0x66, 0x3f, 0x07, 0x70,
	0xf8, 0x87, 0xc7, 0xce, 0x93, 0x8f, 0xe0, 0xba, 0x8f, 0xf4, 0xb1, 0x4e, 0xe4, 0x2d, 0xf6, 0x4f,
	0x0f, 0x22, 0x0f, 0x46, 0x73, 0x0f, 0x26, 0xfb, 0x7e, 0xaa, 0x07, 0x3c, 0x81, 0xa1, 0x8b, 0x3a,
	0x0d, 0xbc, 0x72, 0x12, 0x39, 0x77, 0xc3, 0x49, 0x47, 0xbe, 0x80, 0xbb, 0x33, 0xf1, 0x82, 0x19,
	0x87, 0x70, 0x9d, 0x40, 0x2a, 0xfe, 0xfd, 0x9f, 0x25, 0x4e, 0x60, 0xa8, 0x64, 0x49, 0x3e, 0xc2,
	0xa8, 0x67, 0x9c, 0xc8, 0x92, 0x12, 0x2f, 0x58, 0xb7, 0x0d, 0xda, 0xb6, 0x3f, 0x60, 0xd4, 0xbd,
	0xea, 0x3f, 0x37, 0x3d, 0xfd, 0x15, 0xc0, 0xd0, 0x8d, 0xf8, 0x18, 0xf6, 0xe6, 0x3c, 0x17, 0xb6,
	0xc2, 0xa3, 0xb5, 0xaf, 0x1f, 0x37, 0xb5, 0xc3, 0x71, 0x07, 0x27, 0x66, 0x2b, 0x7c, 0x0a, 0xf0,
	0x9e, 0x14, 0xff, 0xf4, 0xcd, 0x1b, 0x84, 0x0d, 0xdd, 0x42, 0xed, 0xd1, 0x49, 0xc3, 0x39, 0xd4,
	0xf3, 0x9c, 0x18, 0xc6, 0x70, 0xf5, 0x5c, 0xe6, 0x5c, 0xe0, 0xcd, 0x86, 0xf5, 0x53, 0xff, 0xbe,
	0x73, 0x99, 0xe7, 0xc4, 0xb8, 0xc0, 0x37, 0x80, 0xfd, 0xdd, 0xc2, 0xe3, 0x48, 0x55, 0x59, 0xf4,
	0x97, 0xa5, 0x0b, 0x8f, 0x76, 0x29, 0x88, 0xe1, 0xab, 0x76, 0xcf, 0x36, 0xaf, 0x8c, 0x77, 0xb7,
	0xc4, 0xfd, 0xd7, 0x0f, 0x27, 0x3d, 0x01, 0x31, 0x7c, 0x09, 0xa3, 0xe7, 0x45, 0x2a, 0x72, 0x7a,
	0x9b, 0x6a, 0xbd, 0x92, 0x8a, 0xe1, 0x9d, 0x26, 0xfd, 0x36, 0xdc, 0x0d, 0x55, 0xf3, 0x6b, 0xa6,
	0xd6, 0x31, 0x7c, 0x0d, 0x07, 0x09, 0x69, 0x32, 0x1b, 0xa3, 0xdb, 0x8d, 0x70, 0x0b, 0x6d, 0x7d,
	0x76, 0xd2, 0x2f, 0x96, 0x29, 0x2f, 0xe7, 0x24, 0xcc, 0xb3, 0x07, 0x1f, 0xee, 0xe7, 0xdc, 0x14,
	0xf6, 0x22, 0xca, 0xe4, 0x32, 0xe6, 0x65, 0x91, 0x2e, 0x97, 0x05, 0x63, 0xf1, 0xc2, 0xb2, 0x74,
	0xc1, 0x1f, 0xd2, 0x57, 0x43, 0x4a, 0xa4, 0xa5, 0x8e, 0xdd, 0x67, 0xe0, 0x62, 0xcf, 0xff, 0x89,
	0x9f, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xda, 0x95, 0x25, 0x1c, 0x15, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	Signup(ctx context.Context, in *events.SignupRequested, opts ...grpc.CallOption) (*events.Signedup, error)
	VerifyUser(ctx context.Context, in *events.VerifyUserRequested, opts ...grpc.CallOption) (*events.UserVerified, error)
	Login(ctx context.Context, in *events.LoginRequested, opts ...grpc.CallOption) (*events.Loggedin, error)
	UserAuthentication(ctx context.Context, in *UserAuthenticationRequested, opts ...grpc.CallOption) (*UserAuthenticated, error)
	UserAuthorization(ctx context.Context, in *UserAuthorizationRequested, opts ...grpc.CallOption) (*UserAuthorized, error)
	ChangePassword(ctx context.Context, in *events.ChangePasswordRequested, opts ...grpc.CallOption) (*events.PasswordChanged, error)
	ResetPassword(ctx context.Context, in *events.ResetPasswordRequested, opts ...grpc.CallOption) (*events.ResetPasswordEmailSent, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) Signup(ctx context.Context, in *events.SignupRequested, opts ...grpc.CallOption) (*events.Signedup, error) {
	out := new(events.Signedup)
	err := c.cc.Invoke(ctx, "/rpc.User/Signup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) VerifyUser(ctx context.Context, in *events.VerifyUserRequested, opts ...grpc.CallOption) (*events.UserVerified, error) {
	out := new(events.UserVerified)
	err := c.cc.Invoke(ctx, "/rpc.User/VerifyUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Login(ctx context.Context, in *events.LoginRequested, opts ...grpc.CallOption) (*events.Loggedin, error) {
	out := new(events.Loggedin)
	err := c.cc.Invoke(ctx, "/rpc.User/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserAuthentication(ctx context.Context, in *UserAuthenticationRequested, opts ...grpc.CallOption) (*UserAuthenticated, error) {
	out := new(UserAuthenticated)
	err := c.cc.Invoke(ctx, "/rpc.User/UserAuthentication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserAuthorization(ctx context.Context, in *UserAuthorizationRequested, opts ...grpc.CallOption) (*UserAuthorized, error) {
	out := new(UserAuthorized)
	err := c.cc.Invoke(ctx, "/rpc.User/UserAuthorization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ChangePassword(ctx context.Context, in *events.ChangePasswordRequested, opts ...grpc.CallOption) (*events.PasswordChanged, error) {
	out := new(events.PasswordChanged)
	err := c.cc.Invoke(ctx, "/rpc.User/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ResetPassword(ctx context.Context, in *events.ResetPasswordRequested, opts ...grpc.CallOption) (*events.ResetPasswordEmailSent, error) {
	out := new(events.ResetPasswordEmailSent)
	err := c.cc.Invoke(ctx, "/rpc.User/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	Signup(context.Context, *events.SignupRequested) (*events.Signedup, error)
	VerifyUser(context.Context, *events.VerifyUserRequested) (*events.UserVerified, error)
	Login(context.Context, *events.LoginRequested) (*events.Loggedin, error)
	UserAuthentication(context.Context, *UserAuthenticationRequested) (*UserAuthenticated, error)
	UserAuthorization(context.Context, *UserAuthorizationRequested) (*UserAuthorized, error)
	ChangePassword(context.Context, *events.ChangePasswordRequested) (*events.PasswordChanged, error)
	ResetPassword(context.Context, *events.ResetPasswordRequested) (*events.ResetPasswordEmailSent, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) Signup(ctx context.Context, req *events.SignupRequested) (*events.Signedup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (*UnimplementedUserServer) VerifyUser(ctx context.Context, req *events.VerifyUserRequested) (*events.UserVerified, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyUser not implemented")
}
func (*UnimplementedUserServer) Login(ctx context.Context, req *events.LoginRequested) (*events.Loggedin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedUserServer) UserAuthentication(ctx context.Context, req *UserAuthenticationRequested) (*UserAuthenticated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAuthentication not implemented")
}
func (*UnimplementedUserServer) UserAuthorization(ctx context.Context, req *UserAuthorizationRequested) (*UserAuthorized, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAuthorization not implemented")
}
func (*UnimplementedUserServer) ChangePassword(ctx context.Context, req *events.ChangePasswordRequested) (*events.PasswordChanged, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (*UnimplementedUserServer) ResetPassword(ctx context.Context, req *events.ResetPasswordRequested) (*events.ResetPasswordEmailSent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(events.SignupRequested)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.User/Signup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Signup(ctx, req.(*events.SignupRequested))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_VerifyUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(events.VerifyUserRequested)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).VerifyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.User/VerifyUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).VerifyUser(ctx, req.(*events.VerifyUserRequested))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(events.LoginRequested)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.User/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*events.LoginRequested))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserAuthentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAuthenticationRequested)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserAuthentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.User/UserAuthentication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserAuthentication(ctx, req.(*UserAuthenticationRequested))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserAuthorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAuthorizationRequested)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserAuthorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.User/UserAuthorization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserAuthorization(ctx, req.(*UserAuthorizationRequested))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(events.ChangePasswordRequested)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.User/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ChangePassword(ctx, req.(*events.ChangePasswordRequested))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(events.ResetPasswordRequested)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.User/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ResetPassword(ctx, req.(*events.ResetPasswordRequested))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _User_Signup_Handler,
		},
		{
			MethodName: "VerifyUser",
			Handler:    _User_VerifyUser_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "UserAuthentication",
			Handler:    _User_UserAuthentication_Handler,
		},
		{
			MethodName: "UserAuthorization",
			Handler:    _User_UserAuthorization_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _User_ChangePassword_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _User_ResetPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/user.proto",
}
