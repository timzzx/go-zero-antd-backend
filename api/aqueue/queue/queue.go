package queue

import (
	"context"
	"tapi/aqueue/handler"
	"tapi/aqueue/jobtype"
	"tapi/internal/svc"

	"github.com/hibiken/asynq"
)

type Queue struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueue(ctx context.Context, svcCtx *svc.ServiceContext) *Queue {
	return &Queue{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// register job
func (l *Queue) Register() *asynq.ServeMux {

	mux := asynq.NewServeMux()

	// job
	mux.Handle(jobtype.DesUserList, handler.NewUserListHandler(l.svcCtx))

	return mux
}
