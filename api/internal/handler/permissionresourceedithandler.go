package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tapi/internal/logic"
	"tapi/internal/svc"
	"tapi/internal/types"
)

func PermissionResourceEditHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PermissionResourceEditRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewPermissionResourceEditLogic(r.Context(), svcCtx)
		resp, err := l.PermissionResourceEdit(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
