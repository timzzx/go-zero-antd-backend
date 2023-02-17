package logic

import (
	"context"
	"time"

	"tapi/common/jwtx"
	"tapi/common/md5x"
	"tapi/common/varx"
	"tapi/internal/svc"
	"tapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// user表
	table := l.svcCtx.BkModel.User
	// 查询用户
	user, err := table.WithContext(l.ctx).Where(table.Name.Eq(req.Name)).Debug().First()
	if err != nil {
		return &types.LoginResponse{
			Code: 500,
			Msg:  err.Error(),
		}, nil
	}

	// 密码加密验证
	password := md5x.Md5(req.Password + varx.PasswordSalt)

	// 判断密码是否正确
	if user.Password != password {
		return &types.LoginResponse{
			Code: 500,
			Msg:  "密码错误",
		}, nil
	}

	// 获取accessToken
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire

	accessToken, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, user.ID)
	if err != nil {
		return &types.LoginResponse{
			Code: 500,
			Msg:  err.Error(),
		}, nil
	}

	return &types.LoginResponse{
		Code:  200,
		Token: accessToken,
		Msg:   "成功",
	}, nil
}
