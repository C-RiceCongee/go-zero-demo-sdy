package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-0-sd/common/response"
	"go-0-sd/order/api/internal/logic"
	"go-0-sd/order/api/internal/svc"
	"go-0-sd/order/api/internal/types"
	"net/http"
)

func ApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewApiLogic(r.Context(), svcCtx)
		resp, err := l.Api(&req)
		response.Response(w, resp, err)

	}
}
