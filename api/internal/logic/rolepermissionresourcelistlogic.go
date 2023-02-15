package logic

import (
	"context"

	"tapi/internal/svc"
	"tapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RolePermissionResourceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRolePermissionResourceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RolePermissionResourceListLogic {
	return &RolePermissionResourceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RolePermissionResourceListLogic) RolePermissionResourceList(req *types.RolePermissionResourceListRequest) (resp *types.RolePermissionResourceListResponse, err error) {
	// join查询
	r := l.svcCtx.BkModel.RolePermissionResource
	p := l.svcCtx.BkModel.PermissionResource

	var data []types.RolePermissionResource

	err = r.WithContext(l.ctx).Select(p.ID.As("Id"), r.Ctime, p.Name, p.URL).LeftJoin(p, r.Prid.EqCol(p.ID)).Scan(&data)

	if err != nil {
		return &types.RolePermissionResourceListResponse{
			Code: 200,
			Msg:  err.Error(),
		}, nil
	}

	return &types.RolePermissionResourceListResponse{
		Code: 200,
		Msg:  "成功",
		Data: data,
	}, nil
}
