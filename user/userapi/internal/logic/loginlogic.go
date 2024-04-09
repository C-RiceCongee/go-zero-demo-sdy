package logic

import (
	"context"

	"go-0-sd/common/jwt"
	"go-0-sd/user/userapi/internal/svc"
	"go-0-sd/user/userapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp string, err error) {
	// todo: add your logic here and delete this line
	token, err := jwt.GenToken(&jwt.JwtPayLoad{
		UserID:   123123,
		Username: req.UserName,
		Role:     1,
	}, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		return "", err
	}
	return token, nil
}
