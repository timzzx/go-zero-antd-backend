package logic

import (
	"context"
	"time"

	"tapi/bkmodel/dao/model"
	"tapi/internal/svc"
	"tapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserSetRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserSetRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSetRoleLogic {
	return &UserSetRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserSetRoleLogic) UserSetRole(req *types.UserSetRoleRequest) (resp *types.UserSetRoleResponse, err error) {
	u := l.svcCtx.BkModel.User
	_, err = u.WithContext(l.ctx).Where(u.ID.Eq(req.UserId)).Updates(model.User{
		RoleID: req.RoleId,
		Utime:  int32(time.Now().Unix()),
	})

	if err != nil {
		return &types.UserSetRoleResponse{
			Code: 500,
			Msg:  err.Error(),
		}, nil
	}

	return &types.UserSetRoleResponse{
		Code: 200,
		Msg:  "设置成功",
	}, nil
}
