// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: job.proto

package pb

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

// TqJobClient is the client API for TqJob service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TqJobClient interface {
	Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error)
	Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type tqJobClient struct {
	cc grpc.ClientConnInterface
}

func NewTqJobClient(cc grpc.ClientConnInterface) TqJobClient {
	return &tqJobClient{cc}
}

func (c *tqJobClient) Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error) {
	out := new(SubmitResponse)
	err := c.cc.Invoke(ctx, "/pb.TqJob/Submit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tqJobClient) Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error) {
	out := new(CancelResponse)
	err := c.cc.Invoke(ctx, "/pb.TqJob/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tqJobClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/pb.TqJob/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TqJobServer is the server API for TqJob service.
// All implementations must embed UnimplementedTqJobServer
// for forward compatibility
type TqJobServer interface {
	Submit(context.Context, *SubmitRequest) (*SubmitResponse, error)
	Cancel(context.Context, *CancelRequest) (*CancelResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	mustEmbedUnimplementedTqJobServer()
}

// UnimplementedTqJobServer must be embedded to have forward compatible implementations.
type UnimplementedTqJobServer struct {
}

func (UnimplementedTqJobServer) Submit(context.Context, *SubmitRequest) (*SubmitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}
func (UnimplementedTqJobServer) Cancel(context.Context, *CancelRequest) (*CancelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (UnimplementedTqJobServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTqJobServer) mustEmbedUnimplementedTqJobServer() {}

// UnsafeTqJobServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TqJobServer will
// result in compilation errors.
type UnsafeTqJobServer interface {
	mustEmbedUnimplementedTqJobServer()
}

func RegisterTqJobServer(s grpc.ServiceRegistrar, srv TqJobServer) {
	s.RegisterService(&TqJob_ServiceDesc, srv)
}

func _TqJob_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TqJobServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TqJob/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TqJobServer).Submit(ctx, req.(*SubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TqJob_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TqJobServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TqJob/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TqJobServer).Cancel(ctx, req.(*CancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TqJob_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TqJobServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TqJob/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TqJobServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TqJob_ServiceDesc is the grpc.ServiceDesc for TqJob service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TqJob_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TqJob",
	HandlerType: (*TqJobServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Submit",
			Handler:    _TqJob_Submit_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _TqJob_Cancel_Handler,
		},
		{
			MethodName: "List",
			Handler:    _TqJob_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job.proto",
}
