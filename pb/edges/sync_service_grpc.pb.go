// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: edges/sync_service.proto

package edges

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
	SyncService_SetDeviceUpdated_FullMethodName    = "/edges.SyncService/SetDeviceUpdated"
	SyncService_GetDeviceUpdated_FullMethodName    = "/edges.SyncService/GetDeviceUpdated"
	SyncService_WaitDeviceUpdated_FullMethodName   = "/edges.SyncService/WaitDeviceUpdated"
	SyncService_SetOptionUpdated_FullMethodName    = "/edges.SyncService/SetOptionUpdated"
	SyncService_GetOptionUpdated_FullMethodName    = "/edges.SyncService/GetOptionUpdated"
	SyncService_SetSourceUpdated_FullMethodName    = "/edges.SyncService/SetSourceUpdated"
	SyncService_GetSourceUpdated_FullMethodName    = "/edges.SyncService/GetSourceUpdated"
	SyncService_SetTagUpdated_FullMethodName       = "/edges.SyncService/SetTagUpdated"
	SyncService_GetTagUpdated_FullMethodName       = "/edges.SyncService/GetTagUpdated"
	SyncService_SetConstUpdated_FullMethodName     = "/edges.SyncService/SetConstUpdated"
	SyncService_GetConstUpdated_FullMethodName     = "/edges.SyncService/GetConstUpdated"
	SyncService_SetClassUpdated_FullMethodName     = "/edges.SyncService/SetClassUpdated"
	SyncService_GetClassUpdated_FullMethodName     = "/edges.SyncService/GetClassUpdated"
	SyncService_SetAttrUpdated_FullMethodName      = "/edges.SyncService/SetAttrUpdated"
	SyncService_GetAttrUpdated_FullMethodName      = "/edges.SyncService/GetAttrUpdated"
	SyncService_SetLogicUpdated_FullMethodName     = "/edges.SyncService/SetLogicUpdated"
	SyncService_GetLogicUpdated_FullMethodName     = "/edges.SyncService/GetLogicUpdated"
	SyncService_SetFnUpdated_FullMethodName        = "/edges.SyncService/SetFnUpdated"
	SyncService_GetFnUpdated_FullMethodName        = "/edges.SyncService/GetFnUpdated"
	SyncService_SetTagValueUpdated_FullMethodName  = "/edges.SyncService/SetTagValueUpdated"
	SyncService_GetTagValueUpdated_FullMethodName  = "/edges.SyncService/GetTagValueUpdated"
	SyncService_WaitTagValueUpdated_FullMethodName = "/edges.SyncService/WaitTagValueUpdated"
)

// SyncServiceClient is the client API for SyncService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SyncServiceClient interface {
	SetDeviceUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error)
	GetDeviceUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error)
	WaitDeviceUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (SyncService_WaitDeviceUpdatedClient, error)
	SetOptionUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error)
	GetOptionUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error)
	SetSourceUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error)
	GetSourceUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error)
	SetTagUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error)
	GetTagUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error)
	SetConstUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error)
	GetConstUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error)
	SetClassUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error)
	GetClassUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error)
	SetAttrUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error)
	GetAttrUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error)
	SetLogicUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error)
	GetLogicUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error)
	SetFnUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error)
	GetFnUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error)
	SetTagValueUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error)
	GetTagValueUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error)
	WaitTagValueUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (SyncService_WaitTagValueUpdatedClient, error)
}

type syncServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSyncServiceClient(cc grpc.ClientConnInterface) SyncServiceClient {
	return &syncServiceClient{cc}
}

func (c *syncServiceClient) SetDeviceUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, SyncService_SetDeviceUpdated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GetDeviceUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error) {
	out := new(SyncUpdated)
	err := c.cc.Invoke(ctx, SyncService_GetDeviceUpdated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) WaitDeviceUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (SyncService_WaitDeviceUpdatedClient, error) {
	stream, err := c.cc.NewStream(ctx, &SyncService_ServiceDesc.Streams[0], SyncService_WaitDeviceUpdated_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &syncServiceWaitDeviceUpdatedClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SyncService_WaitDeviceUpdatedClient interface {
	Recv() (*pb.MyBool, error)
	grpc.ClientStream
}

type syncServiceWaitDeviceUpdatedClient struct {
	grpc.ClientStream
}

func (x *syncServiceWaitDeviceUpdatedClient) Recv() (*pb.MyBool, error) {
	m := new(pb.MyBool)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *syncServiceClient) SetOptionUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, SyncService_SetOptionUpdated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GetOptionUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error) {
	out := new(SyncUpdated)
	err := c.cc.Invoke(ctx, SyncService_GetOptionUpdated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) SetSourceUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, SyncService_SetSourceUpdated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GetSourceUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error) {
	out := new(SyncUpdated)
	err := c.cc.Invoke(ctx, SyncService_GetSourceUpdated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) SetTagUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, SyncService_SetTagUpdated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GetTagUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error) {
	out := new(SyncUpdated)
	err := c.cc.Invoke(ctx, SyncService_GetTagUpdated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) SetConstUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, SyncService_SetConstUpdated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GetConstUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error) {
	out := new(SyncUpdated)
	err := c.cc.Invoke(ctx, SyncService_GetConstUpdated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) SetClassUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, SyncService_SetClassUpdated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GetClassUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error) {
	out := new(SyncUpdated)
	err := c.cc.Invoke(ctx, SyncService_GetClassUpdated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) SetAttrUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, SyncService_SetAttrUpdated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GetAttrUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error) {
	out := new(SyncUpdated)
	err := c.cc.Invoke(ctx, SyncService_GetAttrUpdated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) SetLogicUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, SyncService_SetLogicUpdated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GetLogicUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error) {
	out := new(SyncUpdated)
	err := c.cc.Invoke(ctx, SyncService_GetLogicUpdated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) SetFnUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, SyncService_SetFnUpdated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GetFnUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error) {
	out := new(SyncUpdated)
	err := c.cc.Invoke(ctx, SyncService_GetFnUpdated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) SetTagValueUpdated(ctx context.Context, in *SyncUpdated, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, SyncService_SetTagValueUpdated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GetTagValueUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (*SyncUpdated, error) {
	out := new(SyncUpdated)
	err := c.cc.Invoke(ctx, SyncService_GetTagValueUpdated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) WaitTagValueUpdated(ctx context.Context, in *pb.MyEmpty, opts ...grpc.CallOption) (SyncService_WaitTagValueUpdatedClient, error) {
	stream, err := c.cc.NewStream(ctx, &SyncService_ServiceDesc.Streams[1], SyncService_WaitTagValueUpdated_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &syncServiceWaitTagValueUpdatedClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SyncService_WaitTagValueUpdatedClient interface {
	Recv() (*pb.MyBool, error)
	grpc.ClientStream
}

type syncServiceWaitTagValueUpdatedClient struct {
	grpc.ClientStream
}

func (x *syncServiceWaitTagValueUpdatedClient) Recv() (*pb.MyBool, error) {
	m := new(pb.MyBool)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SyncServiceServer is the server API for SyncService service.
// All implementations must embed UnimplementedSyncServiceServer
// for forward compatibility
type SyncServiceServer interface {
	SetDeviceUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error)
	GetDeviceUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error)
	WaitDeviceUpdated(*pb.MyEmpty, SyncService_WaitDeviceUpdatedServer) error
	SetOptionUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error)
	GetOptionUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error)
	SetSourceUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error)
	GetSourceUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error)
	SetTagUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error)
	GetTagUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error)
	SetConstUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error)
	GetConstUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error)
	SetClassUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error)
	GetClassUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error)
	SetAttrUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error)
	GetAttrUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error)
	SetLogicUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error)
	GetLogicUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error)
	SetFnUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error)
	GetFnUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error)
	SetTagValueUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error)
	GetTagValueUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error)
	WaitTagValueUpdated(*pb.MyEmpty, SyncService_WaitTagValueUpdatedServer) error
	mustEmbedUnimplementedSyncServiceServer()
}

// UnimplementedSyncServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSyncServiceServer struct {
}

func (UnimplementedSyncServiceServer) SetDeviceUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDeviceUpdated not implemented")
}
func (UnimplementedSyncServiceServer) GetDeviceUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceUpdated not implemented")
}
func (UnimplementedSyncServiceServer) WaitDeviceUpdated(*pb.MyEmpty, SyncService_WaitDeviceUpdatedServer) error {
	return status.Errorf(codes.Unimplemented, "method WaitDeviceUpdated not implemented")
}
func (UnimplementedSyncServiceServer) SetOptionUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetOptionUpdated not implemented")
}
func (UnimplementedSyncServiceServer) GetOptionUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOptionUpdated not implemented")
}
func (UnimplementedSyncServiceServer) SetSourceUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSourceUpdated not implemented")
}
func (UnimplementedSyncServiceServer) GetSourceUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSourceUpdated not implemented")
}
func (UnimplementedSyncServiceServer) SetTagUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTagUpdated not implemented")
}
func (UnimplementedSyncServiceServer) GetTagUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTagUpdated not implemented")
}
func (UnimplementedSyncServiceServer) SetConstUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConstUpdated not implemented")
}
func (UnimplementedSyncServiceServer) GetConstUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConstUpdated not implemented")
}
func (UnimplementedSyncServiceServer) SetClassUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetClassUpdated not implemented")
}
func (UnimplementedSyncServiceServer) GetClassUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClassUpdated not implemented")
}
func (UnimplementedSyncServiceServer) SetAttrUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAttrUpdated not implemented")
}
func (UnimplementedSyncServiceServer) GetAttrUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttrUpdated not implemented")
}
func (UnimplementedSyncServiceServer) SetLogicUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLogicUpdated not implemented")
}
func (UnimplementedSyncServiceServer) GetLogicUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogicUpdated not implemented")
}
func (UnimplementedSyncServiceServer) SetFnUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFnUpdated not implemented")
}
func (UnimplementedSyncServiceServer) GetFnUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFnUpdated not implemented")
}
func (UnimplementedSyncServiceServer) SetTagValueUpdated(context.Context, *SyncUpdated) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTagValueUpdated not implemented")
}
func (UnimplementedSyncServiceServer) GetTagValueUpdated(context.Context, *pb.MyEmpty) (*SyncUpdated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTagValueUpdated not implemented")
}
func (UnimplementedSyncServiceServer) WaitTagValueUpdated(*pb.MyEmpty, SyncService_WaitTagValueUpdatedServer) error {
	return status.Errorf(codes.Unimplemented, "method WaitTagValueUpdated not implemented")
}
func (UnimplementedSyncServiceServer) mustEmbedUnimplementedSyncServiceServer() {}

// UnsafeSyncServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SyncServiceServer will
// result in compilation errors.
type UnsafeSyncServiceServer interface {
	mustEmbedUnimplementedSyncServiceServer()
}

func RegisterSyncServiceServer(s grpc.ServiceRegistrar, srv SyncServiceServer) {
	s.RegisterService(&SyncService_ServiceDesc, srv)
}

func _SyncService_SetDeviceUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncUpdated)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).SetDeviceUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_SetDeviceUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).SetDeviceUpdated(ctx, req.(*SyncUpdated))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GetDeviceUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.MyEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetDeviceUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_GetDeviceUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetDeviceUpdated(ctx, req.(*pb.MyEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_WaitDeviceUpdated_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(pb.MyEmpty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SyncServiceServer).WaitDeviceUpdated(m, &syncServiceWaitDeviceUpdatedServer{stream})
}

type SyncService_WaitDeviceUpdatedServer interface {
	Send(*pb.MyBool) error
	grpc.ServerStream
}

type syncServiceWaitDeviceUpdatedServer struct {
	grpc.ServerStream
}

func (x *syncServiceWaitDeviceUpdatedServer) Send(m *pb.MyBool) error {
	return x.ServerStream.SendMsg(m)
}

func _SyncService_SetOptionUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncUpdated)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).SetOptionUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_SetOptionUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).SetOptionUpdated(ctx, req.(*SyncUpdated))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GetOptionUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.MyEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetOptionUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_GetOptionUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetOptionUpdated(ctx, req.(*pb.MyEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_SetSourceUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncUpdated)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).SetSourceUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_SetSourceUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).SetSourceUpdated(ctx, req.(*SyncUpdated))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GetSourceUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.MyEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetSourceUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_GetSourceUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetSourceUpdated(ctx, req.(*pb.MyEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_SetTagUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncUpdated)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).SetTagUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_SetTagUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).SetTagUpdated(ctx, req.(*SyncUpdated))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GetTagUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.MyEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetTagUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_GetTagUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetTagUpdated(ctx, req.(*pb.MyEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_SetConstUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncUpdated)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).SetConstUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_SetConstUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).SetConstUpdated(ctx, req.(*SyncUpdated))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GetConstUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.MyEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetConstUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_GetConstUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetConstUpdated(ctx, req.(*pb.MyEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_SetClassUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncUpdated)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).SetClassUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_SetClassUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).SetClassUpdated(ctx, req.(*SyncUpdated))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GetClassUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.MyEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetClassUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_GetClassUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetClassUpdated(ctx, req.(*pb.MyEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_SetAttrUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncUpdated)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).SetAttrUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_SetAttrUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).SetAttrUpdated(ctx, req.(*SyncUpdated))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GetAttrUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.MyEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetAttrUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_GetAttrUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetAttrUpdated(ctx, req.(*pb.MyEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_SetLogicUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncUpdated)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).SetLogicUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_SetLogicUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).SetLogicUpdated(ctx, req.(*SyncUpdated))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GetLogicUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.MyEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetLogicUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_GetLogicUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetLogicUpdated(ctx, req.(*pb.MyEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_SetFnUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncUpdated)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).SetFnUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_SetFnUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).SetFnUpdated(ctx, req.(*SyncUpdated))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GetFnUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.MyEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetFnUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_GetFnUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetFnUpdated(ctx, req.(*pb.MyEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_SetTagValueUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncUpdated)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).SetTagValueUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_SetTagValueUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).SetTagValueUpdated(ctx, req.(*SyncUpdated))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GetTagValueUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.MyEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetTagValueUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_GetTagValueUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetTagValueUpdated(ctx, req.(*pb.MyEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_WaitTagValueUpdated_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(pb.MyEmpty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SyncServiceServer).WaitTagValueUpdated(m, &syncServiceWaitTagValueUpdatedServer{stream})
}

type SyncService_WaitTagValueUpdatedServer interface {
	Send(*pb.MyBool) error
	grpc.ServerStream
}

type syncServiceWaitTagValueUpdatedServer struct {
	grpc.ServerStream
}

func (x *syncServiceWaitTagValueUpdatedServer) Send(m *pb.MyBool) error {
	return x.ServerStream.SendMsg(m)
}

// SyncService_ServiceDesc is the grpc.ServiceDesc for SyncService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SyncService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "edges.SyncService",
	HandlerType: (*SyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetDeviceUpdated",
			Handler:    _SyncService_SetDeviceUpdated_Handler,
		},
		{
			MethodName: "GetDeviceUpdated",
			Handler:    _SyncService_GetDeviceUpdated_Handler,
		},
		{
			MethodName: "SetOptionUpdated",
			Handler:    _SyncService_SetOptionUpdated_Handler,
		},
		{
			MethodName: "GetOptionUpdated",
			Handler:    _SyncService_GetOptionUpdated_Handler,
		},
		{
			MethodName: "SetSourceUpdated",
			Handler:    _SyncService_SetSourceUpdated_Handler,
		},
		{
			MethodName: "GetSourceUpdated",
			Handler:    _SyncService_GetSourceUpdated_Handler,
		},
		{
			MethodName: "SetTagUpdated",
			Handler:    _SyncService_SetTagUpdated_Handler,
		},
		{
			MethodName: "GetTagUpdated",
			Handler:    _SyncService_GetTagUpdated_Handler,
		},
		{
			MethodName: "SetConstUpdated",
			Handler:    _SyncService_SetConstUpdated_Handler,
		},
		{
			MethodName: "GetConstUpdated",
			Handler:    _SyncService_GetConstUpdated_Handler,
		},
		{
			MethodName: "SetClassUpdated",
			Handler:    _SyncService_SetClassUpdated_Handler,
		},
		{
			MethodName: "GetClassUpdated",
			Handler:    _SyncService_GetClassUpdated_Handler,
		},
		{
			MethodName: "SetAttrUpdated",
			Handler:    _SyncService_SetAttrUpdated_Handler,
		},
		{
			MethodName: "GetAttrUpdated",
			Handler:    _SyncService_GetAttrUpdated_Handler,
		},
		{
			MethodName: "SetLogicUpdated",
			Handler:    _SyncService_SetLogicUpdated_Handler,
		},
		{
			MethodName: "GetLogicUpdated",
			Handler:    _SyncService_GetLogicUpdated_Handler,
		},
		{
			MethodName: "SetFnUpdated",
			Handler:    _SyncService_SetFnUpdated_Handler,
		},
		{
			MethodName: "GetFnUpdated",
			Handler:    _SyncService_GetFnUpdated_Handler,
		},
		{
			MethodName: "SetTagValueUpdated",
			Handler:    _SyncService_SetTagValueUpdated_Handler,
		},
		{
			MethodName: "GetTagValueUpdated",
			Handler:    _SyncService_GetTagValueUpdated_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WaitDeviceUpdated",
			Handler:       _SyncService_WaitDeviceUpdated_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WaitTagValueUpdated",
			Handler:       _SyncService_WaitTagValueUpdated_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "edges/sync_service.proto",
}
