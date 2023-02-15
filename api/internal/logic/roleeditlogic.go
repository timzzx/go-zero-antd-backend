package logic

import (
	"context"
	"time"

	"tapi/bkmodel/dao/model"
	"tapi/internal/svc"
	"tapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleEditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleEditLogic {
	return &RoleEditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleEditLogic) RoleEdit(req *types.RoleEditRequest) (resp *types.RoleEditResponse, err error) {
	// 数据表
	role := l.svcCtx.BkModel.Role.WithContext(l.ctx)
	if req.Id == 0 {
		// 查询看是否存在
		rs, _ := role.Where(l.svcCtx.BkModel.Role.Name.Eq(req.Name)).Where(l.svcCtx.BkModel.Role.Status.Eq(1)).First()
		if rs != nil {
			return &types.RoleEditResponse{
				Code: 500,
				Msg:  "角色已存在，请重新创建",
			}, nil
		}
		// 新增
		err = role.Create(&model.Role{
			Name:  req.Name,
			Type:  int32(req.Type),
			Ctime: int32(time.Now().Unix()),
			Utime: int32(time.Now().Unix()),
		})
		if err != nil {
			return &types.RoleEditResponse{
				Code: 500,
				Msg:  err.Error(),
			}, nil
		}
	} else {
		// 更新
		r, err := role.Where(l.svcCtx.BkModel.Role.ID.Eq(req.Id)).Updates(model.Role{
			Name:  req.Name,
			Utime: int32(time.Now().Unix()),
		})
		if err != nil {
			return &types.RoleEditResponse{
				Code: 500,
				Msg:  err.Error(),
			}, nil
		}
		if r.Error != nil {
			return &types.RoleEditResponse{
				Code: 500,
				Msg:  r.Error.Error(),
			}, nil
		}
	}

	return &types.RoleEditResponse{
		Code: 200,
		Msg:  "成功",
	}, nil
}
