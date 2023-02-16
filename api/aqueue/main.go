package main

import (
	"context"
	"fmt"
	"os"
	"tapi/aqueue/queue"
	"tapi/internal/config"
	"tapi/internal/svc"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

func main() {
	var c config.Config
	conf.MustLoad("../etc/backend.yaml", &c)
	svcCtx := svc.NewServiceContext(c)
	ctx := context.Background()
	job := queue.NewQueue(ctx, svcCtx)
	mux := job.Register()
	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: svcCtx.Config.Redis.Host},
		asynq.Config{
			IsFailure: func(err error) bool {
				fmt.Printf("asynq server exec task IsFailure ======== >>>>>>>>>>>  err : %+v \n", err)
				return true
			},
			Concurrency: 20, //max concurrent process job task num
		},
	)

	if err := server.Run(mux); err != nil {
		logx.WithContext(ctx).Errorf("!!!CronJobErr!!! run err:%+v", err)
		os.Exit(1)
	}
}
