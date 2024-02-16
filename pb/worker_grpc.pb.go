// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: worker.proto

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

const (
	TqWorker_Register_FullMethodName   = "/pb.TqWorker/Register"
	TqWorker_Deregister_FullMethodName = "/pb.TqWorker/Deregister"
	TqWorker_Status_FullMethodName     = "/pb.TqWorker/Status"
)

// TqWorkerClient is the client API for TqWorker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TqWorkerClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Deregister(ctx context.Context, in *DeregisterRequest, opts ...grpc.CallOption) (*DeregisterResponse, error)
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type tqWorkerClient struct {
	cc grpc.ClientConnInterface
}

func NewTqWorkerClient(cc grpc.ClientConnInterface) TqWorkerClient {
	return &tqWorkerClient{cc}
}

func (c *tqWorkerClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, TqWorker_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tqWorkerClient) Deregister(ctx context.Context, in *DeregisterRequest, opts ...grpc.CallOption) (*DeregisterResponse, error) {
	out := new(DeregisterResponse)
	err := c.cc.Invoke(ctx, TqWorker_Deregister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tqWorkerClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, TqWorker_Status_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TqWorkerServer is the server API for TqWorker service.
// All implementations must embed UnimplementedTqWorkerServer
// for forward compatibility
type TqWorkerServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Deregister(context.Context, *DeregisterRequest) (*DeregisterResponse, error)
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	mustEmbedUnimplementedTqWorkerServer()
}

// UnimplementedTqWorkerServer must be embedded to have forward compatible implementations.
type UnimplementedTqWorkerServer struct {
}

func (UnimplementedTqWorkerServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedTqWorkerServer) Deregister(context.Context, *DeregisterRequest) (*DeregisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deregister not implemented")
}
func (UnimplementedTqWorkerServer) Status(context.Context, *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedTqWorkerServer) mustEmbedUnimplementedTqWorkerServer() {}

// UnsafeTqWorkerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TqWorkerServer will
// result in compilation errors.
type UnsafeTqWorkerServer interface {
	mustEmbedUnimplementedTqWorkerServer()
}

func RegisterTqWorkerServer(s grpc.ServiceRegistrar, srv TqWorkerServer) {
	s.RegisterService(&TqWorker_ServiceDesc, srv)
}

func _TqWorker_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TqWorkerServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TqWorker_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TqWorkerServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TqWorker_Deregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeregisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TqWorkerServer).Deregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TqWorker_Deregister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TqWorkerServer).Deregister(ctx, req.(*DeregisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TqWorker_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TqWorkerServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TqWorker_Status_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TqWorkerServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TqWorker_ServiceDesc is the grpc.ServiceDesc for TqWorker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TqWorker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TqWorker",
	HandlerType: (*TqWorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _TqWorker_Register_Handler,
		},
		{
			MethodName: "Deregister",
			Handler:    _TqWorker_Deregister_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _TqWorker_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "worker.proto",
}
