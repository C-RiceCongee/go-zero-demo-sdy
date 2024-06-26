// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: user/userrpc/userrpc.proto

package userrpc

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
	Userrpc_GetUserInfo_FullMethodName = "/userrpc.Userrpc/GetUserInfo"
)

// UserrpcClient is the client API for Userrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserrpcClient interface {
	GetUserInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type userrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewUserrpcClient(cc grpc.ClientConnInterface) UserrpcClient {
	return &userrpcClient{cc}
}

func (c *userrpcClient) GetUserInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Userrpc_GetUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserrpcServer is the server API for Userrpc service.
// All implementations must embed UnimplementedUserrpcServer
// for forward compatibility
type UserrpcServer interface {
	GetUserInfo(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedUserrpcServer()
}

// UnimplementedUserrpcServer must be embedded to have forward compatible implementations.
type UnimplementedUserrpcServer struct {
}

func (UnimplementedUserrpcServer) GetUserInfo(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUserrpcServer) mustEmbedUnimplementedUserrpcServer() {}

// UnsafeUserrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserrpcServer will
// result in compilation errors.
type UnsafeUserrpcServer interface {
	mustEmbedUnimplementedUserrpcServer()
}

func RegisterUserrpcServer(s grpc.ServiceRegistrar, srv UserrpcServer) {
	s.RegisterService(&Userrpc_ServiceDesc, srv)
}

func _Userrpc_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserrpcServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Userrpc_GetUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserrpcServer).GetUserInfo(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Userrpc_ServiceDesc is the grpc.ServiceDesc for Userrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Userrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "userrpc.Userrpc",
	HandlerType: (*UserrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfo",
			Handler:    _Userrpc_GetUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/userrpc/userrpc.proto",
}
