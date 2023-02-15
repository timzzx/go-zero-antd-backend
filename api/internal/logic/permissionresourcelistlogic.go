package logic

import (
	"context"

	"tapi/internal/svc"
	"tapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PermissionResourceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPermissionResourceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionResourceListLogic {
	return &PermissionResourceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PermissionResourceListLogic) PermissionResourceList(req *types.PermissionResourceListRequest) (resp *types.PermissionResourceListResponse, err error) {
	p := l.svcCtx.BkModel.PermissionResource
	pr, err := p.WithContext(l.ctx).Where(p.Status.Eq(1)).Find()

	if err != nil {
		return &types.PermissionResourceListResponse{
			Code: 500,
			Msg:  err.Error(),
		}, nil
	}
	var data []types.PermissionResource
	if pr != nil {
		for _, item := range pr {
			d := types.PermissionResource{
				Name:  item.Name,
				Url:   item.URL,
				Ctime: int64(item.Ctime),
			}
			data = append(data, d)
		}
	}

	return &types.PermissionResourceListResponse{
		Code: 200,
		Msg:  "成功",
		Data: data,
	}, nil
}
