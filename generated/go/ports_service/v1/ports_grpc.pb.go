// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: ports_service/v1/ports.proto

package ports_service

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
	PortsService_AddPorts_FullMethodName  = "/ports_service.v1.PortsService/AddPorts"
	PortsService_ListPorts_FullMethodName = "/ports_service.v1.PortsService/ListPorts"
)

// PortsServiceClient is the client API for PortsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortsServiceClient interface {
	AddPorts(ctx context.Context, in *AddPortsRequest, opts ...grpc.CallOption) (*AddPortsResponse, error)
	ListPorts(ctx context.Context, in *ListPortsRequest, opts ...grpc.CallOption) (*ListPortsResponse, error)
}

type portsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPortsServiceClient(cc grpc.ClientConnInterface) PortsServiceClient {
	return &portsServiceClient{cc}
}

func (c *portsServiceClient) AddPorts(ctx context.Context, in *AddPortsRequest, opts ...grpc.CallOption) (*AddPortsResponse, error) {
	out := new(AddPortsResponse)
	err := c.cc.Invoke(ctx, PortsService_AddPorts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portsServiceClient) ListPorts(ctx context.Context, in *ListPortsRequest, opts ...grpc.CallOption) (*ListPortsResponse, error) {
	out := new(ListPortsResponse)
	err := c.cc.Invoke(ctx, PortsService_ListPorts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortsServiceServer is the server API for PortsService service.
// All implementations must embed UnimplementedPortsServiceServer
// for forward compatibility
type PortsServiceServer interface {
	AddPorts(context.Context, *AddPortsRequest) (*AddPortsResponse, error)
	ListPorts(context.Context, *ListPortsRequest) (*ListPortsResponse, error)
	mustEmbedUnimplementedPortsServiceServer()
}

// UnimplementedPortsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPortsServiceServer struct {
}

func (UnimplementedPortsServiceServer) AddPorts(context.Context, *AddPortsRequest) (*AddPortsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPorts not implemented")
}
func (UnimplementedPortsServiceServer) ListPorts(context.Context, *ListPortsRequest) (*ListPortsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPorts not implemented")
}
func (UnimplementedPortsServiceServer) mustEmbedUnimplementedPortsServiceServer() {}

// UnsafePortsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortsServiceServer will
// result in compilation errors.
type UnsafePortsServiceServer interface {
	mustEmbedUnimplementedPortsServiceServer()
}

func RegisterPortsServiceServer(s grpc.ServiceRegistrar, srv PortsServiceServer) {
	s.RegisterService(&PortsService_ServiceDesc, srv)
}

func _PortsService_AddPorts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPortsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortsServiceServer).AddPorts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortsService_AddPorts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortsServiceServer).AddPorts(ctx, req.(*AddPortsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortsService_ListPorts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPortsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortsServiceServer).ListPorts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortsService_ListPorts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortsServiceServer).ListPorts(ctx, req.(*ListPortsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PortsService_ServiceDesc is the grpc.ServiceDesc for PortsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PortsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ports_service.v1.PortsService",
	HandlerType: (*PortsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPorts",
			Handler:    _PortsService_AddPorts_Handler,
		},
		{
			MethodName: "ListPorts",
			Handler:    _PortsService_ListPorts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ports_service/v1/ports.proto",
}
