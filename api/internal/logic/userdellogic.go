package logic

import (
	"context"
	"time"

	"tapi/bkmodel/dao/model"
	"tapi/internal/svc"
	"tapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDelLogic {
	return &UserDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDelLogic) UserDel(req *types.UserDelRequest) (resp *types.UserDelResponse, err error) {

	u := l.svcCtx.BkModel.User

	// 删除用户
	user := model.User{
		Status: 2,
		Utime:  int32(time.Now().Unix()),
	}

	_, err = u.WithContext(l.ctx).Where(u.ID.Eq(req.Id)).Updates(user)

	if err != nil {
		return &types.UserDelResponse{
			Code: 500,
			Msg:  err.Error(),
		}, nil
	}

	return &types.UserDelResponse{
		Code: 200,
		Msg:  "删除成功",
	}, nil
}
