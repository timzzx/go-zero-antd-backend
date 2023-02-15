package middleware

import (
	"encoding/json"
	"net/http"
	"tapi/common/varx"
	"tapi/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type LoginMiddleMiddleware struct {
}

func NewLoginMiddleMiddleware() *LoginMiddleMiddleware {
	return &LoginMiddleMiddleware{}
}

func (m *LoginMiddleMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := varx.Ctx.User
		rp := varx.Ctx.RolePermissionResource
		p := varx.Ctx.PermissionResource
		ctx := r.Context()
		uid, _ := ctx.Value("uid").(json.Number).Int64()
		var resource []struct {
			Name string
			Url  string
		}
		u.WithContext(ctx).Where(u.ID.Eq(uid)).LeftJoin(rp, u.RoleID.EqCol(rp.RoleID)).LeftJoin(p, rp.Prid.EqCol(p.ID)).Where(u.Status.Eq(1)).Where(rp.Status.Eq(1)).Where(p.Status.Eq(1)).Where(p.URL.Eq(r.URL.Path)).Select(p.Name, p.URL).Scan(&resource)

		if resource == nil {
			httpx.OkJson(w, &types.CodeErrorResponse{
				Code: 500,
				Msg:  "没有权限", //string(debug.Stack()),
			})
			return
		}

		next(w, r)
	}
}
