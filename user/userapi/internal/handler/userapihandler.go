package handler

import (
	"go-0-sd/common/response"
	"go-0-sd/user/userapi/internal/logic"
	"go-0-sd/user/userapi/internal/svc"
	"go-0-sd/user/userapi/internal/types"
	"net/http"
)

func UserapiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserapiLogic(r.Context(), svcCtx)
		resp, err := l.Userapi(&req)
		response.Response(w, resp, err)

	}
}
