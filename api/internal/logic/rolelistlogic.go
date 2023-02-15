package logic

import (
	"context"

	"tapi/internal/svc"
	"tapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type RoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleListLogic) RoleList(req *types.RoleListRequest) (resp *types.RoleListResponse, err error) {
	// 数据表
	role := l.svcCtx.BkModel.Role
	// 查询
	list, err := role.WithContext(l.ctx).Where(role.Status.Eq(1)).Order(role.ID).Find()
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return &types.RoleListResponse{
				Code: 500,
				Msg:  "查询失败",
			}, nil
		}
	}

	var data []types.Role

	if list != nil {
		for _, item := range list {
			d := types.Role{
				Id:    item.ID,
				Nmae:  item.Name,
				Type:  int64(item.Type),
				Ctime: int64(item.Ctime),
				Utime: int64(item.Utime),
			}
			data = append(data, d)
		}
	}

	return &types.RoleListResponse{
		Code: 200,
		Msg:  "成功",
		Data: data,
	}, nil
}
