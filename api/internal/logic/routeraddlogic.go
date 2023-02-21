package logic

import (
	"context"
	"time"

	"tapi/bkmodel/dao/model"
	"tapi/internal/svc"
	"tapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RouterAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRouterAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RouterAddLogic {
	return &RouterAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RouterAddLogic) RouterAdd(req *types.RouterAddRequest) (resp *types.RouterAddResponse, err error) {
	r := l.svcCtx.BkModel.PermissionResource

	data := model.PermissionResource{
		Name:  req.Name,
		URL:   req.Path,
		Ctime: int32(time.Now().Unix()),
		Utime: int32(time.Now().Unix()),
	}

	// 新增
	err = r.WithContext(l.ctx).Create(&data)
	if err != nil {
		return &types.RouterAddResponse{
			Code: 500,
			Msg:  err.Error(),
		}, nil
	}

	return &types.RouterAddResponse{
		Code: 200,
		Msg:  "添加成功",
	}, nil
}
