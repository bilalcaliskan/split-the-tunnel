// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: routemanager.proto

package routemanager

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
	RouteManager_AddRoute_FullMethodName    = "/routemanager.RouteManager/AddRoute"
	RouteManager_RemoveRoute_FullMethodName = "/routemanager.RouteManager/RemoveRoute"
	RouteManager_ListRoutes_FullMethodName  = "/routemanager.RouteManager/ListRoutes"
)

// RouteManagerClient is the client API for RouteManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteManagerClient interface {
	AddRoute(ctx context.Context, in *AddRouteRequest, opts ...grpc.CallOption) (*AddRouteResponse, error)
	RemoveRoute(ctx context.Context, in *RemoveRouteRequest, opts ...grpc.CallOption) (*RemoveRouteResponse, error)
	ListRoutes(ctx context.Context, in *ListRoutesRequest, opts ...grpc.CallOption) (*ListRoutesResponse, error)
}

type routeManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteManagerClient(cc grpc.ClientConnInterface) RouteManagerClient {
	return &routeManagerClient{cc}
}

func (c *routeManagerClient) AddRoute(ctx context.Context, in *AddRouteRequest, opts ...grpc.CallOption) (*AddRouteResponse, error) {
	out := new(AddRouteResponse)
	err := c.cc.Invoke(ctx, RouteManager_AddRoute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeManagerClient) RemoveRoute(ctx context.Context, in *RemoveRouteRequest, opts ...grpc.CallOption) (*RemoveRouteResponse, error) {
	out := new(RemoveRouteResponse)
	err := c.cc.Invoke(ctx, RouteManager_RemoveRoute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeManagerClient) ListRoutes(ctx context.Context, in *ListRoutesRequest, opts ...grpc.CallOption) (*ListRoutesResponse, error) {
	out := new(ListRoutesResponse)
	err := c.cc.Invoke(ctx, RouteManager_ListRoutes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteManagerServer is the server API for RouteManager service.
// All implementations must embed UnimplementedRouteManagerServer
// for forward compatibility
type RouteManagerServer interface {
	AddRoute(context.Context, *AddRouteRequest) (*AddRouteResponse, error)
	RemoveRoute(context.Context, *RemoveRouteRequest) (*RemoveRouteResponse, error)
	ListRoutes(context.Context, *ListRoutesRequest) (*ListRoutesResponse, error)
	mustEmbedUnimplementedRouteManagerServer()
}

// UnimplementedRouteManagerServer must be embedded to have forward compatible implementations.
type UnimplementedRouteManagerServer struct {
}

func (UnimplementedRouteManagerServer) AddRoute(context.Context, *AddRouteRequest) (*AddRouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRoute not implemented")
}
func (UnimplementedRouteManagerServer) RemoveRoute(context.Context, *RemoveRouteRequest) (*RemoveRouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRoute not implemented")
}
func (UnimplementedRouteManagerServer) ListRoutes(context.Context, *ListRoutesRequest) (*ListRoutesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoutes not implemented")
}
func (UnimplementedRouteManagerServer) mustEmbedUnimplementedRouteManagerServer() {}

// UnsafeRouteManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouteManagerServer will
// result in compilation errors.
type UnsafeRouteManagerServer interface {
	mustEmbedUnimplementedRouteManagerServer()
}

func RegisterRouteManagerServer(s grpc.ServiceRegistrar, srv RouteManagerServer) {
	s.RegisterService(&RouteManager_ServiceDesc, srv)
}

func _RouteManager_AddRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteManagerServer).AddRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RouteManager_AddRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteManagerServer).AddRoute(ctx, req.(*AddRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteManager_RemoveRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteManagerServer).RemoveRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RouteManager_RemoveRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteManagerServer).RemoveRoute(ctx, req.(*RemoveRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteManager_ListRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteManagerServer).ListRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RouteManager_ListRoutes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteManagerServer).ListRoutes(ctx, req.(*ListRoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RouteManager_ServiceDesc is the grpc.ServiceDesc for RouteManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RouteManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "routemanager.RouteManager",
	HandlerType: (*RouteManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddRoute",
			Handler:    _RouteManager_AddRoute_Handler,
		},
		{
			MethodName: "RemoveRoute",
			Handler:    _RouteManager_RemoveRoute_Handler,
		},
		{
			MethodName: "ListRoutes",
			Handler:    _RouteManager_ListRoutes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "routemanager.proto",
}
