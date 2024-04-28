// Copyright 2015 The gRPC Authors
// protoc -I . --go_out=plugins=grpc:. ./ds.proto

// protoc -I ./ --go_out=. ./dataService/ds.proto
// protoc -I ./ --go-grpc_out=. dataService/ds.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: dataService/ds.proto

package dataservice

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

const (
	DataGroupService_Query_FullMethodName      = "/dataService.DataGroupService/Query"
	DataGroupService_Get_FullMethodName        = "/dataService.DataGroupService/Get"
	DataGroupService_Delete_FullMethodName     = "/dataService.DataGroupService/Delete"
	DataGroupService_Update_FullMethodName     = "/dataService.DataGroupService/Update"
	DataGroupService_Replace_FullMethodName    = "/dataService.DataGroupService/Replace"
	DataGroupService_Create_FullMethodName     = "/dataService.DataGroupService/Create"
	DataGroupService_CreateMany_FullMethodName = "/dataService.DataGroupService/CreateMany"
	DataGroupService_DeleteMany_FullMethodName = "/dataService.DataGroupService/DeleteMany"
)

// DataGroupServiceClient is the client API for DataGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataGroupServiceClient interface {
	Query(ctx context.Context, in *api.QueryRequest, opts ...grpc.CallOption) (*api.Response, error)
	Get(ctx context.Context, in *api.GetOrDeleteRequest, opts ...grpc.CallOption) (*api.Response, error)
	Delete(ctx context.Context, in *api.GetOrDeleteRequest, opts ...grpc.CallOption) (*api.Response, error)
	Update(ctx context.Context, in *api.UpdateRequest, opts ...grpc.CallOption) (*api.Response, error)
	Replace(ctx context.Context, in *api.UpdateRequest, opts ...grpc.CallOption) (*api.Response, error)
	Create(ctx context.Context, in *api.CreateRequest, opts ...grpc.CallOption) (*api.Response, error)
	CreateMany(ctx context.Context, in *api.CreateRequest, opts ...grpc.CallOption) (*api.Response, error)
	DeleteMany(ctx context.Context, in *api.QueryRequest, opts ...grpc.CallOption) (*api.Response, error)
}

type dataGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataGroupServiceClient(cc grpc.ClientConnInterface) DataGroupServiceClient {
	return &dataGroupServiceClient{cc}
}

func (c *dataGroupServiceClient) Query(ctx context.Context, in *api.QueryRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, DataGroupService_Query_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataGroupServiceClient) Get(ctx context.Context, in *api.GetOrDeleteRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, DataGroupService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataGroupServiceClient) Delete(ctx context.Context, in *api.GetOrDeleteRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, DataGroupService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataGroupServiceClient) Update(ctx context.Context, in *api.UpdateRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, DataGroupService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataGroupServiceClient) Replace(ctx context.Context, in *api.UpdateRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, DataGroupService_Replace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataGroupServiceClient) Create(ctx context.Context, in *api.CreateRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, DataGroupService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataGroupServiceClient) CreateMany(ctx context.Context, in *api.CreateRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, DataGroupService_CreateMany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataGroupServiceClient) DeleteMany(ctx context.Context, in *api.QueryRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, DataGroupService_DeleteMany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataGroupServiceServer is the server API for DataGroupService service.
// All implementations must embed UnimplementedDataGroupServiceServer
// for forward compatibility
type DataGroupServiceServer interface {
	Query(context.Context, *api.QueryRequest) (*api.Response, error)
	Get(context.Context, *api.GetOrDeleteRequest) (*api.Response, error)
	Delete(context.Context, *api.GetOrDeleteRequest) (*api.Response, error)
	Update(context.Context, *api.UpdateRequest) (*api.Response, error)
	Replace(context.Context, *api.UpdateRequest) (*api.Response, error)
	Create(context.Context, *api.CreateRequest) (*api.Response, error)
	CreateMany(context.Context, *api.CreateRequest) (*api.Response, error)
	DeleteMany(context.Context, *api.QueryRequest) (*api.Response, error)
	mustEmbedUnimplementedDataGroupServiceServer()
}

// UnimplementedDataGroupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataGroupServiceServer struct {
}

func (UnimplementedDataGroupServiceServer) Query(context.Context, *api.QueryRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedDataGroupServiceServer) Get(context.Context, *api.GetOrDeleteRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDataGroupServiceServer) Delete(context.Context, *api.GetOrDeleteRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDataGroupServiceServer) Update(context.Context, *api.UpdateRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDataGroupServiceServer) Replace(context.Context, *api.UpdateRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Replace not implemented")
}
func (UnimplementedDataGroupServiceServer) Create(context.Context, *api.CreateRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDataGroupServiceServer) CreateMany(context.Context, *api.CreateRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMany not implemented")
}
func (UnimplementedDataGroupServiceServer) DeleteMany(context.Context, *api.QueryRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMany not implemented")
}
func (UnimplementedDataGroupServiceServer) mustEmbedUnimplementedDataGroupServiceServer() {}

// UnsafeDataGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataGroupServiceServer will
// result in compilation errors.
type UnsafeDataGroupServiceServer interface {
	mustEmbedUnimplementedDataGroupServiceServer()
}

func RegisterDataGroupServiceServer(s grpc.ServiceRegistrar, srv DataGroupServiceServer) {
	s.RegisterService(&DataGroupService_ServiceDesc, srv)
}

func _DataGroupService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataGroupServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataGroupService_Query_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataGroupServiceServer).Query(ctx, req.(*api.QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataGroupService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.GetOrDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataGroupServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataGroupService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataGroupServiceServer).Get(ctx, req.(*api.GetOrDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataGroupService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.GetOrDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataGroupServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataGroupService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataGroupServiceServer).Delete(ctx, req.(*api.GetOrDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataGroupService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataGroupServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataGroupService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataGroupServiceServer).Update(ctx, req.(*api.UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataGroupService_Replace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataGroupServiceServer).Replace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataGroupService_Replace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataGroupServiceServer).Replace(ctx, req.(*api.UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataGroupService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataGroupServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataGroupService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataGroupServiceServer).Create(ctx, req.(*api.CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataGroupService_CreateMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataGroupServiceServer).CreateMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataGroupService_CreateMany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataGroupServiceServer).CreateMany(ctx, req.(*api.CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataGroupService_DeleteMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataGroupServiceServer).DeleteMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataGroupService_DeleteMany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataGroupServiceServer).DeleteMany(ctx, req.(*api.QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataGroupService_ServiceDesc is the grpc.ServiceDesc for DataGroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataGroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dataService.DataGroupService",
	HandlerType: (*DataGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _DataGroupService_Query_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DataGroupService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DataGroupService_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DataGroupService_Update_Handler,
		},
		{
			MethodName: "Replace",
			Handler:    _DataGroupService_Replace_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _DataGroupService_Create_Handler,
		},
		{
			MethodName: "CreateMany",
			Handler:    _DataGroupService_CreateMany_Handler,
		},
		{
			MethodName: "DeleteMany",
			Handler:    _DataGroupService_DeleteMany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dataService/ds.proto",
}

const (
	DataInterfaceService_Query_FullMethodName      = "/dataService.DataInterfaceService/Query"
	DataInterfaceService_Get_FullMethodName        = "/dataService.DataInterfaceService/Get"
	DataInterfaceService_Delete_FullMethodName     = "/dataService.DataInterfaceService/Delete"
	DataInterfaceService_Update_FullMethodName     = "/dataService.DataInterfaceService/Update"
	DataInterfaceService_Replace_FullMethodName    = "/dataService.DataInterfaceService/Replace"
	DataInterfaceService_Create_FullMethodName     = "/dataService.DataInterfaceService/Create"
	DataInterfaceService_CreateMany_FullMethodName = "/dataService.DataInterfaceService/CreateMany"
	DataInterfaceService_DeleteMany_FullMethodName = "/dataService.DataInterfaceService/DeleteMany"
)

// DataInterfaceServiceClient is the client API for DataInterfaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataInterfaceServiceClient interface {
	Query(ctx context.Context, in *api.QueryRequest, opts ...grpc.CallOption) (*api.Response, error)
	Get(ctx context.Context, in *api.GetOrDeleteRequest, opts ...grpc.CallOption) (*api.Response, error)
	Delete(ctx context.Context, in *api.GetOrDeleteRequest, opts ...grpc.CallOption) (*api.Response, error)
	Update(ctx context.Context, in *api.UpdateRequest, opts ...grpc.CallOption) (*api.Response, error)
	Replace(ctx context.Context, in *api.UpdateRequest, opts ...grpc.CallOption) (*api.Response, error)
	Create(ctx context.Context, in *api.CreateRequest, opts ...grpc.CallOption) (*api.Response, error)
	CreateMany(ctx context.Context, in *api.CreateRequest, opts ...grpc.CallOption) (*api.Response, error)
	DeleteMany(ctx context.Context, in *api.QueryRequest, opts ...grpc.CallOption) (*api.Response, error)
}

type dataInterfaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataInterfaceServiceClient(cc grpc.ClientConnInterface) DataInterfaceServiceClient {
	return &dataInterfaceServiceClient{cc}
}

func (c *dataInterfaceServiceClient) Query(ctx context.Context, in *api.QueryRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, DataInterfaceService_Query_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataInterfaceServiceClient) Get(ctx context.Context, in *api.GetOrDeleteRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, DataInterfaceService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataInterfaceServiceClient) Delete(ctx context.Context, in *api.GetOrDeleteRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, DataInterfaceService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataInterfaceServiceClient) Update(ctx context.Context, in *api.UpdateRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, DataInterfaceService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataInterfaceServiceClient) Replace(ctx context.Context, in *api.UpdateRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, DataInterfaceService_Replace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataInterfaceServiceClient) Create(ctx context.Context, in *api.CreateRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, DataInterfaceService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataInterfaceServiceClient) CreateMany(ctx context.Context, in *api.CreateRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, DataInterfaceService_CreateMany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataInterfaceServiceClient) DeleteMany(ctx context.Context, in *api.QueryRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, DataInterfaceService_DeleteMany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataInterfaceServiceServer is the server API for DataInterfaceService service.
// All implementations must embed UnimplementedDataInterfaceServiceServer
// for forward compatibility
type DataInterfaceServiceServer interface {
	Query(context.Context, *api.QueryRequest) (*api.Response, error)
	Get(context.Context, *api.GetOrDeleteRequest) (*api.Response, error)
	Delete(context.Context, *api.GetOrDeleteRequest) (*api.Response, error)
	Update(context.Context, *api.UpdateRequest) (*api.Response, error)
	Replace(context.Context, *api.UpdateRequest) (*api.Response, error)
	Create(context.Context, *api.CreateRequest) (*api.Response, error)
	CreateMany(context.Context, *api.CreateRequest) (*api.Response, error)
	DeleteMany(context.Context, *api.QueryRequest) (*api.Response, error)
	mustEmbedUnimplementedDataInterfaceServiceServer()
}

// UnimplementedDataInterfaceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataInterfaceServiceServer struct {
}

func (UnimplementedDataInterfaceServiceServer) Query(context.Context, *api.QueryRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedDataInterfaceServiceServer) Get(context.Context, *api.GetOrDeleteRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDataInterfaceServiceServer) Delete(context.Context, *api.GetOrDeleteRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDataInterfaceServiceServer) Update(context.Context, *api.UpdateRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDataInterfaceServiceServer) Replace(context.Context, *api.UpdateRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Replace not implemented")
}
func (UnimplementedDataInterfaceServiceServer) Create(context.Context, *api.CreateRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDataInterfaceServiceServer) CreateMany(context.Context, *api.CreateRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMany not implemented")
}
func (UnimplementedDataInterfaceServiceServer) DeleteMany(context.Context, *api.QueryRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMany not implemented")
}
func (UnimplementedDataInterfaceServiceServer) mustEmbedUnimplementedDataInterfaceServiceServer() {}

// UnsafeDataInterfaceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataInterfaceServiceServer will
// result in compilation errors.
type UnsafeDataInterfaceServiceServer interface {
	mustEmbedUnimplementedDataInterfaceServiceServer()
}

func RegisterDataInterfaceServiceServer(s grpc.ServiceRegistrar, srv DataInterfaceServiceServer) {
	s.RegisterService(&DataInterfaceService_ServiceDesc, srv)
}

func _DataInterfaceService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataInterfaceServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataInterfaceService_Query_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataInterfaceServiceServer).Query(ctx, req.(*api.QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataInterfaceService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.GetOrDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataInterfaceServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataInterfaceService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataInterfaceServiceServer).Get(ctx, req.(*api.GetOrDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataInterfaceService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.GetOrDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataInterfaceServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataInterfaceService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataInterfaceServiceServer).Delete(ctx, req.(*api.GetOrDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataInterfaceService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataInterfaceServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataInterfaceService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataInterfaceServiceServer).Update(ctx, req.(*api.UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataInterfaceService_Replace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataInterfaceServiceServer).Replace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataInterfaceService_Replace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataInterfaceServiceServer).Replace(ctx, req.(*api.UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataInterfaceService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataInterfaceServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataInterfaceService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataInterfaceServiceServer).Create(ctx, req.(*api.CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataInterfaceService_CreateMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataInterfaceServiceServer).CreateMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataInterfaceService_CreateMany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataInterfaceServiceServer).CreateMany(ctx, req.(*api.CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataInterfaceService_DeleteMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataInterfaceServiceServer).DeleteMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataInterfaceService_DeleteMany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataInterfaceServiceServer).DeleteMany(ctx, req.(*api.QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataInterfaceService_ServiceDesc is the grpc.ServiceDesc for DataInterfaceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataInterfaceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dataService.DataInterfaceService",
	HandlerType: (*DataInterfaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _DataInterfaceService_Query_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DataInterfaceService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DataInterfaceService_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DataInterfaceService_Update_Handler,
		},
		{
			MethodName: "Replace",
			Handler:    _DataInterfaceService_Replace_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _DataInterfaceService_Create_Handler,
		},
		{
			MethodName: "CreateMany",
			Handler:    _DataInterfaceService_CreateMany_Handler,
		},
		{
			MethodName: "DeleteMany",
			Handler:    _DataInterfaceService_DeleteMany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dataService/ds.proto",
}

const (
	DataService_Proxy_FullMethodName = "/dataService.DataService/Proxy"
)

// DataServiceClient is the client API for DataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataServiceClient interface {
	Proxy(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ProxyResponse, error)
}

type dataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataServiceClient(cc grpc.ClientConnInterface) DataServiceClient {
	return &dataServiceClient{cc}
}

func (c *dataServiceClient) Proxy(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ProxyResponse, error) {
	out := new(ProxyResponse)
	err := c.cc.Invoke(ctx, DataService_Proxy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataServiceServer is the server API for DataService service.
// All implementations must embed UnimplementedDataServiceServer
// for forward compatibility
type DataServiceServer interface {
	Proxy(context.Context, *Request) (*ProxyResponse, error)
	mustEmbedUnimplementedDataServiceServer()
}

// UnimplementedDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataServiceServer struct {
}

func (UnimplementedDataServiceServer) Proxy(context.Context, *Request) (*ProxyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Proxy not implemented")
}
func (UnimplementedDataServiceServer) mustEmbedUnimplementedDataServiceServer() {}

// UnsafeDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataServiceServer will
// result in compilation errors.
type UnsafeDataServiceServer interface {
	mustEmbedUnimplementedDataServiceServer()
}

func RegisterDataServiceServer(s grpc.ServiceRegistrar, srv DataServiceServer) {
	s.RegisterService(&DataService_ServiceDesc, srv)
}

func _DataService_Proxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).Proxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataService_Proxy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).Proxy(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// DataService_ServiceDesc is the grpc.ServiceDesc for DataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dataService.DataService",
	HandlerType: (*DataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Proxy",
			Handler:    _DataService_Proxy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dataService/ds.proto",
}

const (
	DatasetViewService_Preview_FullMethodName = "/dataService.DatasetViewService/Preview"
)

// DatasetViewServiceClient is the client API for DatasetViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatasetViewServiceClient interface {
	Preview(ctx context.Context, in *ViewPreviewReq, opts ...grpc.CallOption) (*api.Response, error)
}

type datasetViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatasetViewServiceClient(cc grpc.ClientConnInterface) DatasetViewServiceClient {
	return &datasetViewServiceClient{cc}
}

func (c *datasetViewServiceClient) Preview(ctx context.Context, in *ViewPreviewReq, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, DatasetViewService_Preview_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatasetViewServiceServer is the server API for DatasetViewService service.
// All implementations must embed UnimplementedDatasetViewServiceServer
// for forward compatibility
type DatasetViewServiceServer interface {
	Preview(context.Context, *ViewPreviewReq) (*api.Response, error)
	mustEmbedUnimplementedDatasetViewServiceServer()
}

// UnimplementedDatasetViewServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDatasetViewServiceServer struct {
}

func (UnimplementedDatasetViewServiceServer) Preview(context.Context, *ViewPreviewReq) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Preview not implemented")
}
func (UnimplementedDatasetViewServiceServer) mustEmbedUnimplementedDatasetViewServiceServer() {}

// UnsafeDatasetViewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatasetViewServiceServer will
// result in compilation errors.
type UnsafeDatasetViewServiceServer interface {
	mustEmbedUnimplementedDatasetViewServiceServer()
}

func RegisterDatasetViewServiceServer(s grpc.ServiceRegistrar, srv DatasetViewServiceServer) {
	s.RegisterService(&DatasetViewService_ServiceDesc, srv)
}

func _DatasetViewService_Preview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewPreviewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetViewServiceServer).Preview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatasetViewService_Preview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetViewServiceServer).Preview(ctx, req.(*ViewPreviewReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DatasetViewService_ServiceDesc is the grpc.ServiceDesc for DatasetViewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DatasetViewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dataService.DatasetViewService",
	HandlerType: (*DatasetViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Preview",
			Handler:    _DatasetViewService_Preview_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dataService/ds.proto",
}
