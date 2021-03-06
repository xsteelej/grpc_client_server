// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package portsDB

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// PortsDatabaseClient is the client API for PortsDatabase service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortsDatabaseClient interface {
	Write(ctx context.Context, in *Port, opts ...grpc.CallOption) (*WriteResponse, error)
	Read(ctx context.Context, in *PortRequest, opts ...grpc.CallOption) (*Port, error)
}

type portsDatabaseClient struct {
	cc grpc.ClientConnInterface
}

func NewPortsDatabaseClient(cc grpc.ClientConnInterface) PortsDatabaseClient {
	return &portsDatabaseClient{cc}
}

func (c *portsDatabaseClient) Write(ctx context.Context, in *Port, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := c.cc.Invoke(ctx, "/portsDB.PortsDatabase/write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portsDatabaseClient) Read(ctx context.Context, in *PortRequest, opts ...grpc.CallOption) (*Port, error) {
	out := new(Port)
	err := c.cc.Invoke(ctx, "/portsDB.PortsDatabase/read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortsDatabaseServer is the server API for PortsDatabase service.
// All implementations must embed UnimplementedPortsDatabaseServer
// for forward compatibility
type PortsDatabaseServer interface {
	Write(context.Context, *Port) (*WriteResponse, error)
	Read(context.Context, *PortRequest) (*Port, error)
	mustEmbedUnimplementedPortsDatabaseServer()
}

// UnimplementedPortsDatabaseServer must be embedded to have forward compatible implementations.
type UnimplementedPortsDatabaseServer struct {
}

func (UnimplementedPortsDatabaseServer) Write(context.Context, *Port) (*WriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedPortsDatabaseServer) Read(context.Context, *PortRequest) (*Port, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedPortsDatabaseServer) mustEmbedUnimplementedPortsDatabaseServer() {}

// UnsafePortsDatabaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortsDatabaseServer will
// result in compilation errors.
type UnsafePortsDatabaseServer interface {
	mustEmbedUnimplementedPortsDatabaseServer()
}

func RegisterPortsDatabaseServer(s grpc.ServiceRegistrar, srv PortsDatabaseServer) {
	s.RegisterService(&_PortsDatabase_serviceDesc, srv)
}

func _PortsDatabase_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Port)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortsDatabaseServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/portsDB.PortsDatabase/write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortsDatabaseServer).Write(ctx, req.(*Port))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortsDatabase_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortsDatabaseServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/portsDB.PortsDatabase/read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortsDatabaseServer).Read(ctx, req.(*PortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PortsDatabase_serviceDesc = grpc.ServiceDesc{
	ServiceName: "portsDB.PortsDatabase",
	HandlerType: (*PortsDatabaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "write",
			Handler:    _PortsDatabase_Write_Handler,
		},
		{
			MethodName: "read",
			Handler:    _PortsDatabase_Read_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ports.proto",
}
