// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.5.1
// source: sync/sync.proto

package sync

import (
	context "context"
	api "github.com/air-iot/api-client-go/v4/api"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SyncServiceClient is the client API for SyncService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SyncServiceClient interface {
	Query(ctx context.Context, in *api.QueryRequest, opts ...grpc.CallOption) (*api.Response, error)
	Get(ctx context.Context, in *api.GetOrDeleteRequest, opts ...grpc.CallOption) (*api.Response, error)
	Update(ctx context.Context, in *api.UpdateRequest, opts ...grpc.CallOption) (*api.Response, error)
	Delete(ctx context.Context, in *api.GetOrDeleteRequest, opts ...grpc.CallOption) (*api.Response, error)
	Replace(ctx context.Context, in *api.UpdateRequest, opts ...grpc.CallOption) (*api.Response, error)
}

type syncServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSyncServiceClient(cc grpc.ClientConnInterface) SyncServiceClient {
	return &syncServiceClient{cc}
}

func (c *syncServiceClient) Query(ctx context.Context, in *api.QueryRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, "/sync.SyncService/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) Get(ctx context.Context, in *api.GetOrDeleteRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, "/sync.SyncService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) Update(ctx context.Context, in *api.UpdateRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, "/sync.SyncService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) Delete(ctx context.Context, in *api.GetOrDeleteRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, "/sync.SyncService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) Replace(ctx context.Context, in *api.UpdateRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, "/sync.SyncService/Replace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyncServiceServer is the server API for SyncService service.
// All implementations must embed UnimplementedSyncServiceServer
// for forward compatibility
type SyncServiceServer interface {
	Query(context.Context, *api.QueryRequest) (*api.Response, error)
	Get(context.Context, *api.GetOrDeleteRequest) (*api.Response, error)
	Update(context.Context, *api.UpdateRequest) (*api.Response, error)
	Delete(context.Context, *api.GetOrDeleteRequest) (*api.Response, error)
	Replace(context.Context, *api.UpdateRequest) (*api.Response, error)
	mustEmbedUnimplementedSyncServiceServer()
}

// UnimplementedSyncServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSyncServiceServer struct {
}

func (UnimplementedSyncServiceServer) Query(context.Context, *api.QueryRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedSyncServiceServer) Get(context.Context, *api.GetOrDeleteRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSyncServiceServer) Update(context.Context, *api.UpdateRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSyncServiceServer) Delete(context.Context, *api.GetOrDeleteRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSyncServiceServer) Replace(context.Context, *api.UpdateRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Replace not implemented")
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

func _SyncService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sync.SyncService/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).Query(ctx, req.(*api.QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.GetOrDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sync.SyncService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).Get(ctx, req.(*api.GetOrDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sync.SyncService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).Update(ctx, req.(*api.UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.GetOrDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sync.SyncService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).Delete(ctx, req.(*api.GetOrDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_Replace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).Replace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sync.SyncService/Replace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).Replace(ctx, req.(*api.UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SyncService_ServiceDesc is the grpc.ServiceDesc for SyncService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SyncService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sync.SyncService",
	HandlerType: (*SyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _SyncService_Query_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _SyncService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SyncService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SyncService_Delete_Handler,
		},
		{
			MethodName: "Replace",
			Handler:    _SyncService_Replace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sync/sync.proto",
}
