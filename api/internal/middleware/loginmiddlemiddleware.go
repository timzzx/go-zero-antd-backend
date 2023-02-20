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
		ro := varx.Ctx.Role
		rp := varx.Ctx.RolePermissionResource
		p := varx.Ctx.PermissionResource
		ctx := r.Context()
		v, ok := ctx.Value("uid").(json.Number)
		if !ok {
			httpx.OkJson(w, &types.CodeErrorResponse{
				Code: 401,
				Msg:  "ctx没有uid", //string(debug.Stack()),
			})
			return
		}
		uid, err := v.Int64()

		if err != nil {
			httpx.OkJson(w, &types.CodeErrorResponse{
				Code: 401,
				Msg:  err.Error(), //string(debug.Stack()),
			})
			return
		}
		var resource []struct {
			Name string
			Url  string
		}

		res, err := ro.WithContext(ctx).Where(u.ID.Eq(uid)).LeftJoin(u, u.RoleID.EqCol(ro.ID)).Where(u.Status.Eq(1)).Where(ro.Status.Eq(1)).Select(ro.Type).First()
		if err != nil {
			httpx.OkJson(w, &types.CodeErrorResponse{
				Code: 401,
				Msg:  err.Error(), //string(debug.Stack()),
			})
			return
		}

		if res.Type == 1 {

			u.WithContext(ctx).Where(u.ID.Eq(uid)).LeftJoin(rp, u.RoleID.EqCol(rp.RoleID)).LeftJoin(p, rp.Prid.EqCol(p.ID)).Where(u.Status.Eq(1)).Where(rp.Status.Eq(1)).Where(p.Status.Eq(1)).Where(p.URL.Eq(r.URL.Path)).Select(p.Name, p.URL).Scan(&resource)

			if resource == nil {
				httpx.OkJson(w, &types.CodeErrorResponse{
					Code: 401,
					Msg:  "没有权限", //string(debug.Stack()),
				})
				return
			}
		}

		next(w, r)
	}
}
