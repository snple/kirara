// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: nodes/device_service.proto

package nodes

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
	DeviceService_Login_FullMethodName           = "/nodes.DeviceService/Login"
	DeviceService_Update_FullMethodName          = "/nodes.DeviceService/Update"
	DeviceService_View_FullMethodName            = "/nodes.DeviceService/View"
	DeviceService_Link_FullMethodName            = "/nodes.DeviceService/Link"
	DeviceService_ViewWithDeleted_FullMethodName = "/nodes.DeviceService/ViewWithDeleted"
	DeviceService_Sync_FullMethodName            = "/nodes.DeviceService/Sync"
)

// DeviceServiceClient is the client API for DeviceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceServiceClient interface {
	Login(ctx context.Context, in *DeviceLoginRequest, opts ...grpc.CallOption) (*DeviceLoginReply, error)
	Update(ctx context.Context, in *pb.Device, opts ...grpc.CallOption) (*pb.Device, error)
	View(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*pb.Device, error)
	Link(ctx context.Context, in *DeviceLinkRequest, opts ...grpc.CallOption) (*pb.MyBool, error)
	ViewWithDeleted(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*pb.Device, error)
	Sync(ctx context.Context, in *pb.Device, opts ...grpc.CallOption) (*pb.MyBool, error)
}

type deviceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceServiceClient(cc grpc.ClientConnInterface) DeviceServiceClient {
	return &deviceServiceClient{cc}
}

func (c *deviceServiceClient) Login(ctx context.Context, in *DeviceLoginRequest, opts ...grpc.CallOption) (*DeviceLoginReply, error) {
	out := new(DeviceLoginReply)
	err := c.cc.Invoke(ctx, DeviceService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) Update(ctx context.Context, in *pb.Device, opts ...grpc.CallOption) (*pb.Device, error) {
	out := new(pb.Device)
	err := c.cc.Invoke(ctx, DeviceService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) View(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*pb.Device, error) {
	out := new(pb.Device)
	err := c.cc.Invoke(ctx, DeviceService_View_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) Link(ctx context.Context, in *DeviceLinkRequest, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, DeviceService_Link_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) ViewWithDeleted(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*pb.Device, error) {
	out := new(pb.Device)
	err := c.cc.Invoke(ctx, DeviceService_ViewWithDeleted_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) Sync(ctx context.Context, in *pb.Device, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, DeviceService_Sync_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceServiceServer is the server API for DeviceService service.
// All implementations must embed UnimplementedDeviceServiceServer
// for forward compatibility
type DeviceServiceServer interface {
	Login(context.Context, *DeviceLoginRequest) (*DeviceLoginReply, error)
	Update(context.Context, *pb.Device) (*pb.Device, error)
	View(context.Context, *pb.MyEmpty) (*pb.Device, error)
	Link(context.Context, *DeviceLinkRequest) (*pb.MyBool, error)
	ViewWithDeleted(context.Context, *pb.MyEmpty) (*pb.Device, error)
	Sync(context.Context, *pb.Device) (*pb.MyBool, error)
	mustEmbedUnimplementedDeviceServiceServer()
}

// UnimplementedDeviceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeviceServiceServer struct {
}

func (UnimplementedDeviceServiceServer) Login(context.Context, *DeviceLoginRequest) (*DeviceLoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedDeviceServiceServer) Update(context.Context, *pb.Device) (*pb.Device, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDeviceServiceServer) View(context.Context, *pb.MyEmpty) (*pb.Device, error) {
	return nil, status.Errorf(codes.Unimplemented, "method View not implemented")
}
func (UnimplementedDeviceServiceServer) Link(context.Context, *DeviceLinkRequest) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Link not implemented")
}
func (UnimplementedDeviceServiceServer) ViewWithDeleted(context.Context, *pb.MyEmpty) (*pb.Device, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewWithDeleted not implemented")
}
func (UnimplementedDeviceServiceServer) Sync(context.Context, *pb.Device) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedDeviceServiceServer) mustEmbedUnimplementedDeviceServiceServer() {}

// UnsafeDeviceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceServiceServer will
// result in compilation errors.
type UnsafeDeviceServiceServer interface {
	mustEmbedUnimplementedDeviceServiceServer()
}

func RegisterDeviceServiceServer(s grpc.ServiceRegistrar, srv DeviceServiceServer) {
	s.RegisterService(&DeviceService_ServiceDesc, srv)
}

func _DeviceService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).Login(ctx, req.(*DeviceLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Device)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).Update(ctx, req.(*pb.Device))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_View_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.MyEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).View(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceService_View_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).View(ctx, req.(*pb.MyEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_Link_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).Link(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceService_Link_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).Link(ctx, req.(*DeviceLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_ViewWithDeleted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.MyEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).ViewWithDeleted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceService_ViewWithDeleted_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).ViewWithDeleted(ctx, req.(*pb.MyEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Device)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceService_Sync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).Sync(ctx, req.(*pb.Device))
	}
	return interceptor(ctx, in, info, handler)
}

// DeviceService_ServiceDesc is the grpc.ServiceDesc for DeviceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeviceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nodes.DeviceService",
	HandlerType: (*DeviceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _DeviceService_Login_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DeviceService_Update_Handler,
		},
		{
			MethodName: "View",
			Handler:    _DeviceService_View_Handler,
		},
		{
			MethodName: "Link",
			Handler:    _DeviceService_Link_Handler,
		},
		{
			MethodName: "ViewWithDeleted",
			Handler:    _DeviceService_ViewWithDeleted_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _DeviceService_Sync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nodes/device_service.proto",
}
