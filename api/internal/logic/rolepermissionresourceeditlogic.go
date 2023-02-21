package logic

import (
	"context"
	"strconv"
	"strings"
	"time"

	"tapi/bkmodel/dao/model"
	"tapi/internal/svc"
	"tapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RolePermissionResourceEditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRolePermissionResourceEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RolePermissionResourceEditLogic {
	return &RolePermissionResourceEditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RolePermissionResourceEditLogic) RolePermissionResourceEdit(req *types.RolePermissionResourceEditRequest) (resp *types.RolePermissionResourceEditResponse, err error) {
	if req.Data == "" {
		return &types.RolePermissionResourceEditResponse{
			Code: 200,
			Msg:  "成功",
		}, nil
	}
	r := l.svcCtx.BkModel.RolePermissionResource
	// 先删除
	_, err = r.WithContext(l.ctx).Where(r.RoleID.Eq(req.RoleId)).Delete()
	if err != nil {
		return &types.RolePermissionResourceEditResponse{
			Code: 500,
			Msg:  err.Error(),
		}, nil
	}

	var data []*model.RolePermissionResource

	ids := strings.Split(req.Data, ",")

	for _, item := range ids {
		i, _ := strconv.Atoi(item)
		if i != -1 {
			p := &model.RolePermissionResource{
				RoleID: req.RoleId,
				Prid:   int64(i),
				Ctime:  int32(time.Now().Unix()),
				Utime:  int32(time.Now().Unix()),
			}
			data = append(data, p)
		}
	}

	// 分配权限
	if data != nil {
		r.WithContext(l.ctx).Create(data...)
	}
	return &types.RolePermissionResourceEditResponse{
		Code: 200,
		Msg:  "成功",
	}, nil
}
