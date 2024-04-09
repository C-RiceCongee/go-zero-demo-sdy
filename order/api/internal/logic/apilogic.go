package logic

import (
	"context"
	"go-0-sd/user/userrpc/userrpc"

	"go-0-sd/order/api/internal/svc"
	"go-0-sd/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApiLogic {
	return &ApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApiLogic) Api(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	r, err := l.svcCtx.UserRpcClient.GetUserInfo(l.ctx, &userrpc.Request{Id: "123"})
	if err != nil {
		return nil, err
	}
	return &types.Response{Id: r.GetId(), Name: r.Name, Gender: r.Gender}, nil
}
