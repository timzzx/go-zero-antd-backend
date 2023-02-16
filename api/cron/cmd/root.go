package cmd

import (
	"fmt"
	"tapi/internal/config"
	"tapi/internal/svc"

	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/conf"
)

var svcCtx *svc.ServiceContext

func Execute() {
	c := cron.New(cron.WithSeconds())

	ScheduleRun(c)

	fmt.Println("定时任务启动...")
	go c.Start()
	defer c.Stop()
	select {}
}

func init() {
	var c config.Config
	conf.MustLoad("../etc/backend.yaml", &c)
	svcCtx = svc.NewServiceContext(c)
}
