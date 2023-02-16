package svc

import (
	"tapi/bkmodel/dao/query"
	"tapi/internal/config"
	"tapi/internal/middleware"

	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config

	LoginMiddle rest.Middleware

	BkModel *query.Query

	Redis *redis.Client

	AsynqClient *asynq.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, _ := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{})
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &ServiceContext{
		Config:      c,
		LoginMiddle: middleware.NewLoginMiddleMiddleware().Handle,
		BkModel:     query.Use(db),
		Redis:       rdb,
		AsynqClient: newAsynqClient(c),
	}
}
