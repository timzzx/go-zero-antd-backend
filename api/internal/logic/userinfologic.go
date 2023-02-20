package logic

import (
	"context"
	"encoding/json"

	"tapi/aqueue/jobtype"
	"tapi/internal/svc"
	"tapi/internal/types"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	// 获取token中的uid，具体自行查看go-zero的文档和源码，access的验证框架已经实现，我们只需要配置Auth的对应参数
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return &types.UserInfoResponse{
			Code: 500,
			Msg:  err.Error(),
		}, nil
	}
	table := l.svcCtx.BkModel.User
	user, err := table.WithContext(l.ctx).Where(table.ID.Eq(uid)).First()

	if err != nil {
		return &types.UserInfoResponse{
			Code: 500,
			Msg:  err.Error(),
		}, nil
	}

	// 测试一下写入job
	payload, err := json.Marshal(jobtype.PayloadUserList{Id: 2})
	if err != nil {
		return &types.UserInfoResponse{
			Code: 500,
			Msg:  err.Error(),
		}, nil
	} else {
		_, err = l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.DesUserList, payload))
		if err != nil {
			return &types.UserInfoResponse{
				Code: 500,
				Msg:  err.Error(),
			}, nil
		}
	}

	return &types.UserInfoResponse{
		Code: 200,
		Msg:  "成功",
		Data: types.UserInfo{
			Id:    user.ID,
			Name:  user.Name,
			Ctime: int64(user.Ctime),
			Utime: int64(user.Utime),
		},
	}, nil
}
