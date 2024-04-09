package logic

import (
	"context"
	"encoding/json"

	"go-0-sd/user/userapi/internal/svc"
	"go-0-sd/user/userapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	userId := l.ctx.Value("user_id").(json.Number)
	username := l.ctx.Value("username").(string)
	uid, _ := userId.Int64()
	return &types.UserInfoResponse{
		Id:       uint(uid),
		UserName: username,
	}, nil
}
