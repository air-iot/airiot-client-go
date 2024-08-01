// Copyright 2015 The gRPC Authors
// protoc -I . --go_out=plugins=grpc:. ./warning.proto

// protoc -I ./ --go_out=. ./warning/warning.proto
// protoc -I ./ --go-grpc_out=. warning/warning.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: warning/warning.proto

package warning

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
	WarnService_Query_FullMethodName       = "/warning.WarnService/Query"
	WarnService_Get_FullMethodName         = "/warning.WarnService/Get"
	WarnService_Create_FullMethodName      = "/warning.WarnService/Create"
	WarnService_BatchCreate_FullMethodName = "/warning.WarnService/BatchCreate"
)

// WarnServiceClient is the client API for WarnService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WarnServiceClient interface {
	Query(ctx context.Context, in *QueryWarningRequest, opts ...grpc.CallOption) (*api.Response, error)
	Get(ctx context.Context, in *GetOrDeleteWarningRequest, opts ...grpc.CallOption) (*api.Response, error)
	Create(ctx context.Context, in *api.CreateRequest, opts ...grpc.CallOption) (*api.Response, error)
	BatchCreate(ctx context.Context, in *api.CreateRequest, opts ...grpc.CallOption) (*api.Response, error)
}

type warnServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWarnServiceClient(cc grpc.ClientConnInterface) WarnServiceClient {
	return &warnServiceClient{cc}
}

func (c *warnServiceClient) Query(ctx context.Context, in *QueryWarningRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, WarnService_Query_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warnServiceClient) Get(ctx context.Context, in *GetOrDeleteWarningRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, WarnService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warnServiceClient) Create(ctx context.Context, in *api.CreateRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, WarnService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warnServiceClient) BatchCreate(ctx context.Context, in *api.CreateRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, WarnService_BatchCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WarnServiceServer is the server API for WarnService service.
// All implementations must embed UnimplementedWarnServiceServer
// for forward compatibility
type WarnServiceServer interface {
	Query(context.Context, *QueryWarningRequest) (*api.Response, error)
	Get(context.Context, *GetOrDeleteWarningRequest) (*api.Response, error)
	Create(context.Context, *api.CreateRequest) (*api.Response, error)
	BatchCreate(context.Context, *api.CreateRequest) (*api.Response, error)
	mustEmbedUnimplementedWarnServiceServer()
}

// UnimplementedWarnServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWarnServiceServer struct {
}

func (UnimplementedWarnServiceServer) Query(context.Context, *QueryWarningRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedWarnServiceServer) Get(context.Context, *GetOrDeleteWarningRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedWarnServiceServer) Create(context.Context, *api.CreateRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedWarnServiceServer) BatchCreate(context.Context, *api.CreateRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchCreate not implemented")
}
func (UnimplementedWarnServiceServer) mustEmbedUnimplementedWarnServiceServer() {}

// UnsafeWarnServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WarnServiceServer will
// result in compilation errors.
type UnsafeWarnServiceServer interface {
	mustEmbedUnimplementedWarnServiceServer()
}

func RegisterWarnServiceServer(s grpc.ServiceRegistrar, srv WarnServiceServer) {
	s.RegisterService(&WarnService_ServiceDesc, srv)
}

func _WarnService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryWarningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarnServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarnService_Query_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarnServiceServer).Query(ctx, req.(*QueryWarningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarnService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrDeleteWarningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarnServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarnService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarnServiceServer).Get(ctx, req.(*GetOrDeleteWarningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarnService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarnServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarnService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarnServiceServer).Create(ctx, req.(*api.CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarnService_BatchCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarnServiceServer).BatchCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarnService_BatchCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarnServiceServer).BatchCreate(ctx, req.(*api.CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WarnService_ServiceDesc is the grpc.ServiceDesc for WarnService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WarnService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "warning.WarnService",
	HandlerType: (*WarnServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _WarnService_Query_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _WarnService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _WarnService_Create_Handler,
		},
		{
			MethodName: "BatchCreate",
			Handler:    _WarnService_BatchCreate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "warning/warning.proto",
}

const (
	RuleService_Query_FullMethodName       = "/warning.RuleService/Query"
	RuleService_BatchCreate_FullMethodName = "/warning.RuleService/BatchCreate"
	RuleService_Update_FullMethodName      = "/warning.RuleService/Update"
	RuleService_Delete_FullMethodName      = "/warning.RuleService/Delete"
	RuleService_Replace_FullMethodName     = "/warning.RuleService/Replace"
)

// RuleServiceClient is the client API for RuleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RuleServiceClient interface {
	Query(ctx context.Context, in *api.QueryRequest, opts ...grpc.CallOption) (*api.Response, error)
	BatchCreate(ctx context.Context, in *api.CreateRequest, opts ...grpc.CallOption) (*api.Response, error)
	Update(ctx context.Context, in *api.UpdateRequest, opts ...grpc.CallOption) (*api.Response, error)
	Delete(ctx context.Context, in *api.GetOrDeleteRequest, opts ...grpc.CallOption) (*api.Response, error)
	Replace(ctx context.Context, in *api.UpdateRequest, opts ...grpc.CallOption) (*api.Response, error)
}

type ruleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRuleServiceClient(cc grpc.ClientConnInterface) RuleServiceClient {
	return &ruleServiceClient{cc}
}

func (c *ruleServiceClient) Query(ctx context.Context, in *api.QueryRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, RuleService_Query_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleServiceClient) BatchCreate(ctx context.Context, in *api.CreateRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, RuleService_BatchCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleServiceClient) Update(ctx context.Context, in *api.UpdateRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, RuleService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleServiceClient) Delete(ctx context.Context, in *api.GetOrDeleteRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, RuleService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleServiceClient) Replace(ctx context.Context, in *api.UpdateRequest, opts ...grpc.CallOption) (*api.Response, error) {
	out := new(api.Response)
	err := c.cc.Invoke(ctx, RuleService_Replace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RuleServiceServer is the server API for RuleService service.
// All implementations must embed UnimplementedRuleServiceServer
// for forward compatibility
type RuleServiceServer interface {
	Query(context.Context, *api.QueryRequest) (*api.Response, error)
	BatchCreate(context.Context, *api.CreateRequest) (*api.Response, error)
	Update(context.Context, *api.UpdateRequest) (*api.Response, error)
	Delete(context.Context, *api.GetOrDeleteRequest) (*api.Response, error)
	Replace(context.Context, *api.UpdateRequest) (*api.Response, error)
	mustEmbedUnimplementedRuleServiceServer()
}

// UnimplementedRuleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRuleServiceServer struct {
}

func (UnimplementedRuleServiceServer) Query(context.Context, *api.QueryRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedRuleServiceServer) BatchCreate(context.Context, *api.CreateRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchCreate not implemented")
}
func (UnimplementedRuleServiceServer) Update(context.Context, *api.UpdateRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRuleServiceServer) Delete(context.Context, *api.GetOrDeleteRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRuleServiceServer) Replace(context.Context, *api.UpdateRequest) (*api.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Replace not implemented")
}
func (UnimplementedRuleServiceServer) mustEmbedUnimplementedRuleServiceServer() {}

// UnsafeRuleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RuleServiceServer will
// result in compilation errors.
type UnsafeRuleServiceServer interface {
	mustEmbedUnimplementedRuleServiceServer()
}

func RegisterRuleServiceServer(s grpc.ServiceRegistrar, srv RuleServiceServer) {
	s.RegisterService(&RuleService_ServiceDesc, srv)
}

func _RuleService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuleService_Query_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleServiceServer).Query(ctx, req.(*api.QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleService_BatchCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleServiceServer).BatchCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuleService_BatchCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleServiceServer).BatchCreate(ctx, req.(*api.CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuleService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleServiceServer).Update(ctx, req.(*api.UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.GetOrDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuleService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleServiceServer).Delete(ctx, req.(*api.GetOrDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleService_Replace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleServiceServer).Replace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuleService_Replace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleServiceServer).Replace(ctx, req.(*api.UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RuleService_ServiceDesc is the grpc.ServiceDesc for RuleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RuleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "warning.RuleService",
	HandlerType: (*RuleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _RuleService_Query_Handler,
		},
		{
			MethodName: "BatchCreate",
			Handler:    _RuleService_BatchCreate_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _RuleService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RuleService_Delete_Handler,
		},
		{
			MethodName: "Replace",
			Handler:    _RuleService_Replace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "warning/warning.proto",
}
