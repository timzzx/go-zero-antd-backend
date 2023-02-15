package logic

import (
	"context"
	"time"

	"tapi/bkmodel/dao/model"
	"tapi/internal/svc"
	"tapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PermissionResourceEditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPermissionResourceEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionResourceEditLogic {
	return &PermissionResourceEditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PermissionResourceEditLogic) PermissionResourceEdit(req *types.PermissionResourceEditRequest) (resp *types.PermissionResourceEditResponse, err error) {
	p := l.svcCtx.BkModel.PermissionResource

	// 只做新增，不做修改
	// 查询数据是否存在
	rp, _ := p.WithContext(l.ctx).Where(p.URL.Eq(req.Url)).Where(p.Status.Eq(1)).First()
	if rp != nil {
		return &types.PermissionResourceEditResponse{
			Code: 500,
			Msg:  "资源已存在",
		}, nil
	}

	// 新建
	err = p.WithContext(l.ctx).Create(&model.PermissionResource{
		Name:  req.Name,
		URL:   req.Url,
		Ctime: int32(time.Now().Unix()),
		Utime: int32(time.Now().Unix()),
	})

	if err != nil {
		return &types.PermissionResourceEditResponse{
			Code: 500,
			Msg:  "新建失败",
		}, nil
	}

	return &types.PermissionResourceEditResponse{
		Code: 200,
		Msg:  "成功",
	}, nil
}
