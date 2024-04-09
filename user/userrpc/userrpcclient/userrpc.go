// Code generated by goctl. DO NOT EDIT.
// Source: userrpc.proto

package userrpcclient

import (
	"context"

	"go-0-sd/user/userrpc/userrpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = userrpc.Request
	Response = userrpc.Response

	Userrpc interface {
		GetUserInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultUserrpc struct {
		cli zrpc.Client
	}
)

func NewUserrpc(cli zrpc.Client) Userrpc {
	return &defaultUserrpc{
		cli: cli,
	}
}

func (m *defaultUserrpc) GetUserInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := userrpc.NewUserrpcClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}