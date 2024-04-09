package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-0-sd/order/api/internal/config"
	"go-0-sd/user/userrpc/userrpcclient"
)

type ServiceContext struct {
	Config        config.Config
	UserRpcClient userrpcclient.Userrpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserRpcClient: userrpcclient.NewUserrpc(zrpc.MustNewClient(c.UserRpc)),
	}
}
