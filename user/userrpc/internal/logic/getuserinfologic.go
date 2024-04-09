package logic

import (
	"context"

	"go-0-sd/user/userrpc/internal/svc"
	"go-0-sd/user/userrpc/userrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *userrpc.Request) (*userrpc.Response, error) {
	//假装 经过一系列 缓存查询 数据查询
	return &userrpc.Response{
		Id:     "1123",
		Name:   "user from user rpc",
		Gender: "男",
	}, nil
}
