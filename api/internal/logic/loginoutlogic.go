package logic

import (
	"context"

	"tapi/internal/svc"
	"tapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginOutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginOutLogic {
	return &LoginOutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginOutLogic) LoginOut(req *types.LoginOutRequest) (resp *types.LoginOutResponse, err error) {
	// todo: add your logic here and delete this line
	// 退出 可以做一些清除处理

	return &types.LoginOutResponse{
		Code: 200,
		Msg:  "退出成功",
	}, nil
}
