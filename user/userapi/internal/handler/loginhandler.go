package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-0-sd/common/response"
	"go-0-sd/user/userapi/internal/logic"
	"go-0-sd/user/userapi/internal/svc"
	"go-0-sd/user/userapi/internal/types"
	"net/http"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.Response(w, resp, err)

	}
}
