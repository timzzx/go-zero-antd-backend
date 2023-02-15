package logic

import (
	"context"
	"time"

	"tapi/bkmodel/dao/model"
	"tapi/internal/svc"
	"tapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDeleteLogic {
	return &RoleDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleDeleteLogic) RoleDelete(req *types.RoleDeleteRequest) (resp *types.RoleDeleteResponse, err error) {
	// 数据表
	role := l.svcCtx.BkModel.Role
	// 删除(标记删除)
	r, err := role.WithContext(l.ctx).Where(role.ID.Eq(req.Id)).Updates(model.Role{
		Status: 2,
		Utime:  int32(time.Now().Unix()),
	})
	if err != nil {
		return &types.RoleDeleteResponse{
			Code: 500,
			Msg:  err.Error(),
		}, nil
	}
	if r.Error != nil {
		return &types.RoleDeleteResponse{
			Code: 500,
			Msg:  r.Error.Error(),
		}, nil
	}

	return &types.RoleDeleteResponse{
		Code: 200,
		Msg:  "成功",
	}, nil
}
