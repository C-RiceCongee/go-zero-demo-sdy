package handler

import (
	"go-0-sd/common/response"
	"go-0-sd/user/userapi/internal/logic"
	"go-0-sd/user/userapi/internal/svc"
	"net/http"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		response.Response(w, resp, err)

	}
}
