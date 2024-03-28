// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: cores/auth_service.proto

package cores

import (
	context "context"
	pb "github.com/snple/kirara/pb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AuthService_Login_FullMethodName           = "/cores.AuthService/Login"
	AuthService_ChangePass_FullMethodName      = "/cores.AuthService/ChangePass"
	AuthService_ForceChangePass_FullMethodName = "/cores.AuthService/ForceChangePass"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	ChangePass(ctx context.Context, in *ChangePassRequest, opts ...grpc.CallOption) (*pb.Message, error)
	ForceChangePass(ctx context.Context, in *ForceChangePassRequest, opts ...grpc.CallOption) (*pb.Message, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, AuthService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ChangePass(ctx context.Context, in *ChangePassRequest, opts ...grpc.CallOption) (*pb.Message, error) {
	out := new(pb.Message)
	err := c.cc.Invoke(ctx, AuthService_ChangePass_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ForceChangePass(ctx context.Context, in *ForceChangePassRequest, opts ...grpc.CallOption) (*pb.Message, error) {
	out := new(pb.Message)
	err := c.cc.Invoke(ctx, AuthService_ForceChangePass_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	ChangePass(context.Context, *ChangePassRequest) (*pb.Message, error)
	ForceChangePass(context.Context, *ForceChangePassRequest) (*pb.Message, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServiceServer) ChangePass(context.Context, *ChangePassRequest) (*pb.Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePass not implemented")
}
func (UnimplementedAuthServiceServer) ForceChangePass(context.Context, *ForceChangePassRequest) (*pb.Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForceChangePass not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ChangePass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ChangePass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ChangePass_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ChangePass(ctx, req.(*ChangePassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ForceChangePass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForceChangePassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ForceChangePass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ForceChangePass_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ForceChangePass(ctx, req.(*ForceChangePassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cores.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "ChangePass",
			Handler:    _AuthService_ChangePass_Handler,
		},
		{
			MethodName: "ForceChangePass",
			Handler:    _AuthService_ForceChangePass_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cores/auth_service.proto",
}
