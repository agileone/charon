// Code generated by protoc-gen-go.
// source: auth.proto
// DO NOT EDIT!

/*
Package charonrpc is a generated protocol buffer package.

It is generated from these files:
	auth.proto
	user.proto
	group.proto
	permission.proto

It has these top-level messages:
	LoginRequest
	LogoutRequest
	IsAuthenticatedRequest
	IsGrantedRequest
	BelongsToRequest
	ActorResponse
	User
	CreateUserRequest
	CreateUserResponse
	GetUserRequest
	GetUserResponse
	ListUsersRequest
	ListUsersResponse
	DeleteUserRequest
	ModifyUserRequest
	ModifyUserResponse
	ListUserPermissionsRequest
	ListUserPermissionsResponse
	SetUserPermissionsRequest
	SetUserPermissionsResponse
	ListUserGroupsRequest
	ListUserGroupsResponse
	SetUserGroupsRequest
	SetUserGroupsResponse
	Group
	CreateGroupRequest
	CreateGroupResponse
	GetGroupRequest
	GetGroupResponse
	ListGroupsRequest
	ListGroupsResponse
	DeleteGroupRequest
	ModifyGroupRequest
	ModifyGroupResponse
	SetGroupPermissionsRequest
	SetGroupPermissionsResponse
	ListGroupPermissionsRequest
	ListGroupPermissionsResponse
	RegisterPermissionsRequest
	RegisterPermissionsResponse
	ListPermissionsRequest
	ListPermissionsResponse
	GetPermissionRequest
	GetPermissionResponse
*/
package charonrpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import google_protobuf1 "github.com/golang/protobuf/ptypes/wrappers"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LoginRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Client   string `protobuf:"bytes,3,opt,name=client" json:"client,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type LogoutRequest struct {
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
}

func (m *LogoutRequest) Reset()                    { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string            { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()               {}
func (*LogoutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type IsAuthenticatedRequest struct {
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
}

func (m *IsAuthenticatedRequest) Reset()                    { *m = IsAuthenticatedRequest{} }
func (m *IsAuthenticatedRequest) String() string            { return proto.CompactTextString(m) }
func (*IsAuthenticatedRequest) ProtoMessage()               {}
func (*IsAuthenticatedRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type IsGrantedRequest struct {
	UserId     int64  `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Permission string `protobuf:"bytes,2,opt,name=permission" json:"permission,omitempty"`
}

func (m *IsGrantedRequest) Reset()                    { *m = IsGrantedRequest{} }
func (m *IsGrantedRequest) String() string            { return proto.CompactTextString(m) }
func (*IsGrantedRequest) ProtoMessage()               {}
func (*IsGrantedRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type BelongsToRequest struct {
	UserId  int64 `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	GroupId int64 `protobuf:"varint,2,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
}

func (m *BelongsToRequest) Reset()                    { *m = BelongsToRequest{} }
func (m *BelongsToRequest) String() string            { return proto.CompactTextString(m) }
func (*BelongsToRequest) ProtoMessage()               {}
func (*BelongsToRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ActorResponse struct {
	Id          int64    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Username    string   `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	FirstName   string   `protobuf:"bytes,3,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName    string   `protobuf:"bytes,4,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Permissions []string `protobuf:"bytes,5,rep,name=permissions" json:"permissions,omitempty"`
	IsSuperuser bool     `protobuf:"varint,6,opt,name=is_superuser,json=isSuperuser" json:"is_superuser,omitempty"`
	IsActive    bool     `protobuf:"varint,7,opt,name=is_active,json=isActive" json:"is_active,omitempty"`
	IsStuff     bool     `protobuf:"varint,8,opt,name=is_stuff,json=isStuff" json:"is_stuff,omitempty"`
	IsConfirmed bool     `protobuf:"varint,9,opt,name=is_confirmed,json=isConfirmed" json:"is_confirmed,omitempty"`
}

func (m *ActorResponse) Reset()                    { *m = ActorResponse{} }
func (m *ActorResponse) String() string            { return proto.CompactTextString(m) }
func (*ActorResponse) ProtoMessage()               {}
func (*ActorResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*LoginRequest)(nil), "charonrpc.LoginRequest")
	proto.RegisterType((*LogoutRequest)(nil), "charonrpc.LogoutRequest")
	proto.RegisterType((*IsAuthenticatedRequest)(nil), "charonrpc.IsAuthenticatedRequest")
	proto.RegisterType((*IsGrantedRequest)(nil), "charonrpc.IsGrantedRequest")
	proto.RegisterType((*BelongsToRequest)(nil), "charonrpc.BelongsToRequest")
	proto.RegisterType((*ActorResponse)(nil), "charonrpc.ActorResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Auth service

type AuthClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*google_protobuf1.StringValue, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	IsAuthenticated(ctx context.Context, in *IsAuthenticatedRequest, opts ...grpc.CallOption) (*google_protobuf1.BoolValue, error)
	Actor(ctx context.Context, in *google_protobuf1.StringValue, opts ...grpc.CallOption) (*ActorResponse, error)
	IsGranted(ctx context.Context, in *IsGrantedRequest, opts ...grpc.CallOption) (*google_protobuf1.BoolValue, error)
	BelongsTo(ctx context.Context, in *BelongsToRequest, opts ...grpc.CallOption) (*google_protobuf1.BoolValue, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*google_protobuf1.StringValue, error) {
	out := new(google_protobuf1.StringValue)
	err := grpc.Invoke(ctx, "/charonrpc.Auth/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/charonrpc.Auth/Logout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) IsAuthenticated(ctx context.Context, in *IsAuthenticatedRequest, opts ...grpc.CallOption) (*google_protobuf1.BoolValue, error) {
	out := new(google_protobuf1.BoolValue)
	err := grpc.Invoke(ctx, "/charonrpc.Auth/IsAuthenticated", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Actor(ctx context.Context, in *google_protobuf1.StringValue, opts ...grpc.CallOption) (*ActorResponse, error) {
	out := new(ActorResponse)
	err := grpc.Invoke(ctx, "/charonrpc.Auth/Actor", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) IsGranted(ctx context.Context, in *IsGrantedRequest, opts ...grpc.CallOption) (*google_protobuf1.BoolValue, error) {
	out := new(google_protobuf1.BoolValue)
	err := grpc.Invoke(ctx, "/charonrpc.Auth/IsGranted", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) BelongsTo(ctx context.Context, in *BelongsToRequest, opts ...grpc.CallOption) (*google_protobuf1.BoolValue, error) {
	out := new(google_protobuf1.BoolValue)
	err := grpc.Invoke(ctx, "/charonrpc.Auth/BelongsTo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthServer interface {
	Login(context.Context, *LoginRequest) (*google_protobuf1.StringValue, error)
	Logout(context.Context, *LogoutRequest) (*google_protobuf.Empty, error)
	IsAuthenticated(context.Context, *IsAuthenticatedRequest) (*google_protobuf1.BoolValue, error)
	Actor(context.Context, *google_protobuf1.StringValue) (*ActorResponse, error)
	IsGranted(context.Context, *IsGrantedRequest) (*google_protobuf1.BoolValue, error)
	BelongsTo(context.Context, *BelongsToRequest) (*google_protobuf1.BoolValue, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonrpc.Auth/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonrpc.Auth/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_IsAuthenticated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAuthenticatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).IsAuthenticated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonrpc.Auth/IsAuthenticated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).IsAuthenticated(ctx, req.(*IsAuthenticatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Actor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Actor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonrpc.Auth/Actor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Actor(ctx, req.(*google_protobuf1.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_IsGranted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsGrantedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).IsGranted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonrpc.Auth/IsGranted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).IsGranted(ctx, req.(*IsGrantedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_BelongsTo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BelongsToRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).BelongsTo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonrpc.Auth/BelongsTo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).BelongsTo(ctx, req.(*BelongsToRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "charonrpc.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Auth_Logout_Handler,
		},
		{
			MethodName: "IsAuthenticated",
			Handler:    _Auth_IsAuthenticated_Handler,
		},
		{
			MethodName: "Actor",
			Handler:    _Auth_Actor_Handler,
		},
		{
			MethodName: "IsGranted",
			Handler:    _Auth_IsGranted_Handler,
		},
		{
			MethodName: "BelongsTo",
			Handler:    _Auth_BelongsTo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x14, 0x54, 0x9d, 0x36, 0x89, 0x5f, 0x5a, 0xa8, 0xf6, 0x90, 0x1a, 0x17, 0x50, 0x9a, 0x53, 0x4f,
	0x8e, 0xd4, 0x9e, 0x00, 0x09, 0x94, 0x52, 0x40, 0x11, 0x15, 0x42, 0x49, 0xc5, 0x91, 0x68, 0x63,
	0x6f, 0x9c, 0x15, 0xce, 0xee, 0xb2, 0x6f, 0x4d, 0xd5, 0x5f, 0xc2, 0x99, 0x7f, 0x8a, 0x76, 0x1d,
	0x27, 0x76, 0x8a, 0x52, 0x71, 0x89, 0xf2, 0x66, 0x3c, 0xf3, 0x3e, 0x3c, 0x06, 0xa0, 0xb9, 0x59,
	0x44, 0x4a, 0x4b, 0x23, 0x89, 0x1f, 0x2f, 0xa8, 0x96, 0x42, 0xab, 0x38, 0xbc, 0x4c, 0xb9, 0x59,
	0xe4, 0xb3, 0x28, 0x96, 0xcb, 0x41, 0x2a, 0x33, 0x2a, 0xd2, 0x81, 0x7b, 0x66, 0x96, 0xcf, 0x07,
	0xca, 0xdc, 0x2b, 0x86, 0x03, 0xb6, 0x54, 0xe6, 0xbe, 0xf8, 0x2d, 0xf4, 0xe1, 0xab, 0xc7, 0x45,
	0x77, 0x9a, 0x2a, 0xc5, 0xf4, 0xe6, 0x4f, 0x21, 0xed, 0x7f, 0x87, 0xc3, 0x1b, 0x99, 0x72, 0x31,
	0x66, 0x3f, 0x73, 0x86, 0x86, 0x84, 0xd0, 0xce, 0x91, 0x69, 0x41, 0x97, 0x2c, 0xd8, 0xeb, 0xed,
	0x9d, 0xfb, 0xe3, 0x75, 0x6d, 0x39, 0x45, 0x11, 0xef, 0xa4, 0x4e, 0x02, 0xaf, 0xe0, 0xca, 0x9a,
	0x74, 0xa1, 0x19, 0x67, 0x9c, 0x09, 0x13, 0x34, 0x1c, 0xb3, 0xaa, 0xfa, 0x17, 0x70, 0x74, 0x23,
	0x53, 0x99, 0x9b, 0xb2, 0xc1, 0x19, 0x1c, 0xd2, 0x38, 0x66, 0x88, 0x53, 0x23, 0x7f, 0x30, 0xb1,
	0x6a, 0xd2, 0x29, 0xb0, 0x5b, 0x0b, 0xf5, 0xdf, 0x40, 0x77, 0x84, 0xc3, 0xdc, 0x2c, 0x98, 0x30,
	0x3c, 0xa6, 0x86, 0x25, 0xff, 0x21, 0xfe, 0x0c, 0xc7, 0x23, 0xfc, 0xa4, 0xa9, 0xa8, 0xc8, 0x4e,
	0xa0, 0x65, 0x97, 0x98, 0xf2, 0xc4, 0x29, 0x1a, 0xe3, 0xa6, 0x2d, 0x47, 0x09, 0x79, 0x09, 0xa0,
	0x98, 0x5e, 0x72, 0x44, 0x2e, 0xc5, 0x6a, 0xa7, 0x0a, 0xd2, 0xff, 0x08, 0xc7, 0x57, 0x2c, 0x93,
	0x22, 0xc5, 0x5b, 0xf9, 0xa8, 0xd9, 0x33, 0x68, 0xa7, 0x5a, 0xe6, 0xca, 0x32, 0x9e, 0x63, 0x5a,
	0xae, 0x1e, 0x25, 0xfd, 0xdf, 0x1e, 0x1c, 0x0d, 0x63, 0x23, 0xf5, 0x98, 0xa1, 0x92, 0x02, 0x19,
	0x79, 0x02, 0xde, 0xda, 0xc0, 0xe3, 0x49, 0xed, 0xee, 0xde, 0xd6, 0xdd, 0x5f, 0x00, 0xcc, 0xb9,
	0x46, 0x33, 0x75, 0x6c, 0x71, 0x5f, 0xdf, 0x21, 0x5f, 0x2c, 0x7d, 0x0a, 0x7e, 0x46, 0x4b, 0x76,
	0xbf, 0xd0, 0x5a, 0xc0, 0x91, 0x3d, 0xe8, 0x6c, 0xf6, 0xc1, 0xe0, 0xa0, 0xd7, 0xb0, 0x07, 0xab,
	0x40, 0xf6, 0xa6, 0x1c, 0xa7, 0x98, 0x2b, 0xa6, 0x6d, 0xc7, 0xa0, 0xd9, 0xdb, 0x3b, 0x6f, 0x8f,
	0x3b, 0x1c, 0x27, 0x25, 0x64, 0x3b, 0x70, 0x9c, 0xd2, 0xd8, 0xf0, 0x5f, 0x2c, 0x68, 0x39, 0xbe,
	0xcd, 0x71, 0xe8, 0x6a, 0xbb, 0xb6, 0xd5, 0x9b, 0x7c, 0x3e, 0x0f, 0xda, 0x8e, 0x6b, 0x71, 0x9c,
	0xd8, 0x72, 0x65, 0x1d, 0x4b, 0x31, 0xe7, 0x7a, 0xc9, 0x92, 0xc0, 0x2f, 0xad, 0xdf, 0x97, 0xd0,
	0xc5, 0x9f, 0x06, 0xec, 0xdb, 0x57, 0x4d, 0xde, 0xc2, 0x81, 0x0b, 0x22, 0x39, 0x89, 0xd6, 0x5f,
	0x43, 0x54, 0x8d, 0x66, 0xf8, 0x3c, 0x4a, 0xa5, 0x4c, 0x33, 0x16, 0x95, 0xd9, 0x8e, 0x26, 0x46,
	0x73, 0x91, 0x7e, 0xa3, 0x59, 0xce, 0xc8, 0x6b, 0x68, 0x16, 0x41, 0x23, 0x41, 0xdd, 0x60, 0x93,
	0xbd, 0xb0, 0xfb, 0xc0, 0xe1, 0x83, 0xfd, 0x8a, 0xc8, 0x57, 0x78, 0xba, 0x15, 0x38, 0x72, 0x56,
	0x31, 0xf9, 0x77, 0x18, 0xc3, 0xf0, 0x81, 0xdb, 0x95, 0x94, 0x59, 0x31, 0xcd, 0x3b, 0x38, 0x70,
	0xef, 0x9b, 0xec, 0x1c, 0x3a, 0xac, 0x8e, 0x5a, 0xcf, 0xc7, 0x35, 0xf8, 0xeb, 0x18, 0x93, 0xd3,
	0xda, 0x30, 0xf5, 0x70, 0xef, 0x1c, 0xe3, 0x1a, 0xfc, 0x75, 0x7e, 0x6b, 0x2e, 0xdb, 0xa9, 0xde,
	0xe5, 0x32, 0x6b, 0x3a, 0xec, 0xf2, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xec, 0x43, 0x58, 0xec,
	0xb3, 0x04, 0x00, 0x00,
}
