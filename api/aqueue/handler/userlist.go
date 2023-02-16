package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"tapi/aqueue/jobtype"
	"tapi/internal/svc"

	"github.com/hibiken/asynq"
)

type UserListHandler struct {
	svcCtx *svc.ServiceContext
}

func NewUserListHandler(svcCtx *svc.ServiceContext) *UserListHandler {
	return &UserListHandler{
		svcCtx: svcCtx,
	}
}

func (l *UserListHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	// 获取参数
	var p jobtype.PayloadUserList
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return errors.New("参数错误")
	}

	u := l.svcCtx.BkModel.User
	d, err := u.WithContext(context.Background()).Where(u.ID.Eq(p.Id)).Debug().First()
	if err != nil {
		return err
	}
	fmt.Println(d)
	return nil
}
