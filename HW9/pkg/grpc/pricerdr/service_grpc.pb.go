// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: service.proto

package pricerdr

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

// ListServiceClient is the client API for ListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ListServiceClient interface {
	Create(ctx context.Context, in *List, opts ...grpc.CallOption) (*ListId, error)
	Read(ctx context.Context, in *ListId, opts ...grpc.CallOption) (*List, error)
	Update(ctx context.Context, in *List, opts ...grpc.CallOption) (*ListId, error)
	Delete(ctx context.Context, in *ListId, opts ...grpc.CallOption) (*ListId, error)
}

type listServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewListServiceClient(cc grpc.ClientConnInterface) ListServiceClient {
	return &listServiceClient{cc}
}

func (c *listServiceClient) Create(ctx context.Context, in *List, opts ...grpc.CallOption) (*ListId, error) {
	out := new(ListId)
	err := c.cc.Invoke(ctx, "/ListService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listServiceClient) Read(ctx context.Context, in *ListId, opts ...grpc.CallOption) (*List, error) {
	out := new(List)
	err := c.cc.Invoke(ctx, "/ListService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listServiceClient) Update(ctx context.Context, in *List, opts ...grpc.CallOption) (*ListId, error) {
	out := new(ListId)
	err := c.cc.Invoke(ctx, "/ListService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listServiceClient) Delete(ctx context.Context, in *ListId, opts ...grpc.CallOption) (*ListId, error) {
	out := new(ListId)
	err := c.cc.Invoke(ctx, "/ListService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListServiceServer is the server API for ListService service.
// All implementations must embed UnimplementedListServiceServer
// for forward compatibility
type ListServiceServer interface {
	Create(context.Context, *List) (*ListId, error)
	Read(context.Context, *ListId) (*List, error)
	Update(context.Context, *List) (*ListId, error)
	Delete(context.Context, *ListId) (*ListId, error)
	mustEmbedUnimplementedListServiceServer()
}

// UnimplementedListServiceServer must be embedded to have forward compatible implementations.
type UnimplementedListServiceServer struct {
}

func (UnimplementedListServiceServer) Create(context.Context, *List) (*ListId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedListServiceServer) Read(context.Context, *ListId) (*List, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedListServiceServer) Update(context.Context, *List) (*ListId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedListServiceServer) Delete(context.Context, *ListId) (*ListId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedListServiceServer) mustEmbedUnimplementedListServiceServer() {}

// UnsafeListServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ListServiceServer will
// result in compilation errors.
type UnsafeListServiceServer interface {
	mustEmbedUnimplementedListServiceServer()
}

func RegisterListServiceServer(s grpc.ServiceRegistrar, srv ListServiceServer) {
	s.RegisterService(&ListService_ServiceDesc, srv)
}

func _ListService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(List)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ListService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListServiceServer).Create(ctx, req.(*List))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ListService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListServiceServer).Read(ctx, req.(*ListId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(List)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ListService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListServiceServer).Update(ctx, req.(*List))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ListService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListServiceServer).Delete(ctx, req.(*ListId))
	}
	return interceptor(ctx, in, info, handler)
}

// ListService_ServiceDesc is the grpc.ServiceDesc for ListService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ListService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ListService",
	HandlerType: (*ListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ListService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ListService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ListService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ListService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
