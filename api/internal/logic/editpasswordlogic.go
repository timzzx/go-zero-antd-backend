package logic

import (
	"context"
	"encoding/json"
	"time"

	"tapi/bkmodel/dao/model"
	"tapi/common/md5x"
	"tapi/common/varx"
	"tapi/internal/svc"
	"tapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditPasswordLogic {
	return &EditPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditPasswordLogic) EditPassword(req *types.EditPasswordRequest) (resp *types.EditPasswordResponse, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	// user表
	u := l.svcCtx.BkModel.User
	// 密码加密
	password := md5x.Md5(req.Password + varx.PasswordSalt)

	// 更新用户密码
	_, err = u.WithContext(l.ctx).Where(u.ID.Eq(uid)).Updates(model.User{
		Password: password,
		Utime:    int32(time.Now().Unix()),
	})

	if err != nil {
		return &types.EditPasswordResponse{
			Code: 500,
			Msg:  "更新失败",
		}, nil
	}
	return &types.EditPasswordResponse{
		Code: 200,
		Msg:  "更新成功",
	}, nil
}
