package logic

import (
	"context"

	"tapi/common/varx"
	"tapi/internal/svc"
	"tapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RouterListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRouterListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RouterListLogic {
	return &RouterListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RouterListLogic) RouterList(req *types.RouterListResquest) (resp *types.RouterListResponse, err error) {
	// todo: add your logic here and delete this line
	list := varx.RouterList
	if list == nil {
		return &types.RouterListResponse{
			Code: 500,
			Msg:  "获取失败",
		}, nil
	}

	var data []types.Router

	for _, item := range list {
		d := types.Router{
			Method: item.Method,
			Path:   item.Path,
		}
		if d.Path != "/api/login" {
			data = append(data, d)
		}
	}

	return &types.RouterListResponse{
		Code: 200,
		Msg:  "成功",
		Data: data,
	}, nil
}
