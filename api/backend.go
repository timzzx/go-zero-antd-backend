package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"tapi/common/varx"
	"tapi/internal/config"
	"tapi/internal/handler"
	"tapi/internal/svc"
	"tapi/internal/types"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/backend.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 日志 （能用不够优雅）
	// logx.DisableStat()

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, interface{}) {
		fmt.Println(err.Error())
		return http.StatusOK, &types.CodeErrorResponse{
			Code: 500,
			Msg:  err.Error(),
		}
	})

	// 全局recover中间件
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if result := recover(); result != nil {
					log.Println(fmt.Sprintf("%v\n%s", result, debug.Stack()))
					httpx.OkJson(w, &types.CodeErrorResponse{
						Code: 500,
						Msg:  "服务器错误", //string(debug.Stack()),
					})
				}
			}()

			next.ServeHTTP(w, r)
		})
	})
	// 路由列表
	varx.RouterList = server.Routes()
	varx.Ctx = ctx.BkModel

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
