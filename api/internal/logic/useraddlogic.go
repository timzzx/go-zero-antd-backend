package logic

import (
	"context"
	"time"

	"tapi/bkmodel/dao/model"
	"tapi/internal/svc"
	"tapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UserAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAddLogic) UserAdd(req *types.UserAddRequest) (resp *types.UserAddResponse, err error) {

	table := l.svcCtx.BkModel.User
	// 查询用户是否存在
	u, err := table.WithContext(l.ctx).Where(table.Name.Eq(req.Name)).First()
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return &types.UserAddResponse{
				Code: 500,
				Msg:  err.Error(),
			}, nil
		}

	}
	if u != nil && u.Name == req.Name {
		return &types.UserAddResponse{
			Code: 500,
			Msg:  "用户已存在",
		}, nil
	}

	// 新建用户
	currTime := time.Now().Unix()
	user := model.User{
		Name:     req.Name,
		Password: req.Password,
		Status:   1,
		Utime:    int32(currTime),
		Ctime:    int32(currTime),
	}

	err = table.WithContext(l.ctx).Create(&user)
	if err != nil {
		return &types.UserAddResponse{
			Code: 500,
			Msg:  err.Error(),
		}, nil
	}
	return &types.UserAddResponse{
		Code: 200,
		Msg:  "成功",
	}, nil
}
