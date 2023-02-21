package logic

import (
	"context"

	"tapi/common/varx"
	"tapi/internal/svc"
	"tapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
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

	// 查出表里所有的
	p := l.svcCtx.BkModel.PermissionResource
	plist, err := p.WithContext(l.ctx).Where(p.Status.Eq(1)).Find()
	if err != nil && err != gorm.ErrRecordNotFound {
		return &types.RouterListResponse{
			Code: 500,
			Msg:  err.Error(),
		}, nil
	}

	var data []types.Router

	for _, item := range list {
		d := types.Router{
			Method: item.Method,
			Path:   item.Path,
		}
		if d.Path != "/api/login" {
			var flag bool
			for _, pdata := range plist {
				if d.Path == pdata.URL {
					flag = true
				}
			}
			if flag != true {
				data = append(data, d)
			}
		}
	}

	return &types.RouterListResponse{
		Code: 200,
		Msg:  "成功",
		Data: data,
	}, nil
}
