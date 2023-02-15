package logic

import (
	"context"
	"time"

	"tapi/bkmodel/dao/model"
	"tapi/internal/svc"
	"tapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PermissionResourceDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPermissionResourceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionResourceDeleteLogic {
	return &PermissionResourceDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PermissionResourceDeleteLogic) PermissionResourceDelete(req *types.PermissionResourceDeleteRequest) (resp *types.PermissionResourceDeleteResponse, err error) {
	p := l.svcCtx.BkModel.PermissionResource

	// 标记删除
	_, err = p.WithContext(l.ctx).Where(p.ID.Eq(req.Id)).Updates(&model.PermissionResource{
		Status: 2,
		Utime:  int32(time.Now().Unix()),
	})

	if err != nil {
		return &types.PermissionResourceDeleteResponse{
			Code: 500,
			Msg:  err.Error(),
		}, nil
	}

	return &types.PermissionResourceDeleteResponse{
		Code: 200,
		Msg:  "成功",
	}, nil
}
