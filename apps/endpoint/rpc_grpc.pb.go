// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.3
// source: apps/endpoint/pb/rpc.proto

package endpoint

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Rpc_CreateEndpoint_FullMethodName = "/mcenter.endpoint.Rpc/CreateEndpoint"
	Rpc_QueryEndpoint_FullMethodName  = "/mcenter.endpoint.Rpc/QueryEndpoint"
	Rpc_DeleteEndpoint_FullMethodName = "/mcenter.endpoint.Rpc/DeleteEndpoint"
)

// RpcClient is the client API for Rpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcClient interface {
	CreateEndpoint(ctx context.Context, in *EndpointSpecSet, opts ...grpc.CallOption) (*EndpointSet, error)
	QueryEndpoint(ctx context.Context, in *QueryEndpointRequest, opts ...grpc.CallOption) (*EndpointSet, error)
	DeleteEndpoint(ctx context.Context, in *DeleteEndpointRequest, opts ...grpc.CallOption) (*EndpointSet, error)
}

type rpcClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcClient(cc grpc.ClientConnInterface) RpcClient {
	return &rpcClient{cc}
}

func (c *rpcClient) CreateEndpoint(ctx context.Context, in *EndpointSpecSet, opts ...grpc.CallOption) (*EndpointSet, error) {
	out := new(EndpointSet)
	err := c.cc.Invoke(ctx, Rpc_CreateEndpoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) QueryEndpoint(ctx context.Context, in *QueryEndpointRequest, opts ...grpc.CallOption) (*EndpointSet, error) {
	out := new(EndpointSet)
	err := c.cc.Invoke(ctx, Rpc_QueryEndpoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) DeleteEndpoint(ctx context.Context, in *DeleteEndpointRequest, opts ...grpc.CallOption) (*EndpointSet, error) {
	out := new(EndpointSet)
	err := c.cc.Invoke(ctx, Rpc_DeleteEndpoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcServer is the server API for Rpc service.
// All implementations must embed UnimplementedRpcServer
// for forward compatibility
type RpcServer interface {
	CreateEndpoint(context.Context, *EndpointSpecSet) (*EndpointSet, error)
	QueryEndpoint(context.Context, *QueryEndpointRequest) (*EndpointSet, error)
	DeleteEndpoint(context.Context, *DeleteEndpointRequest) (*EndpointSet, error)
	mustEmbedUnimplementedRpcServer()
}

// UnimplementedRpcServer must be embedded to have forward compatible implementations.
type UnimplementedRpcServer struct {
}

func (UnimplementedRpcServer) CreateEndpoint(context.Context, *EndpointSpecSet) (*EndpointSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEndpoint not implemented")
}
func (UnimplementedRpcServer) QueryEndpoint(context.Context, *QueryEndpointRequest) (*EndpointSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryEndpoint not implemented")
}
func (UnimplementedRpcServer) DeleteEndpoint(context.Context, *DeleteEndpointRequest) (*EndpointSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEndpoint not implemented")
}
func (UnimplementedRpcServer) mustEmbedUnimplementedRpcServer() {}

// UnsafeRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcServer will
// result in compilation errors.
type UnsafeRpcServer interface {
	mustEmbedUnimplementedRpcServer()
}

func RegisterRpcServer(s grpc.ServiceRegistrar, srv RpcServer) {
	s.RegisterService(&Rpc_ServiceDesc, srv)
}

func _Rpc_CreateEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndpointSpecSet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).CreateEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_CreateEndpoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).CreateEndpoint(ctx, req.(*EndpointSpecSet))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_QueryEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).QueryEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_QueryEndpoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).QueryEndpoint(ctx, req.(*QueryEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_DeleteEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).DeleteEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_DeleteEndpoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).DeleteEndpoint(ctx, req.(*DeleteEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Rpc_ServiceDesc is the grpc.ServiceDesc for Rpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mcenter.endpoint.Rpc",
	HandlerType: (*RpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEndpoint",
			Handler:    _Rpc_CreateEndpoint_Handler,
		},
		{
			MethodName: "QueryEndpoint",
			Handler:    _Rpc_QueryEndpoint_Handler,
		},
		{
			MethodName: "DeleteEndpoint",
			Handler:    _Rpc_DeleteEndpoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/endpoint/pb/rpc.proto",
}
