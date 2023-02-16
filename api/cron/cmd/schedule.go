package cmd

import (
	"fmt"
	"tapi/cron/cronx"

	"github.com/robfig/cron/v3"
)

func ScheduleRun(c *cron.Cron) {
	c.AddFunc(cronx.EveryMinute(), func() {
		fmt.Println("定时任务")
	})
	// 每分钟定时查询用户信息
	c.AddFunc(cronx.Every5s(), userlist)

	// 计划任务执行写在这里
}
