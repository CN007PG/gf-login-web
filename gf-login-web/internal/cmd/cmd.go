package cmd

import (
	"context"
	"gf-login-web/internal/controller"
	"gf-login-web/internal/middleware"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			//设置组
			s.Group("/", func(group *ghttp.RouterGroup) {
				//设置中间件
				group.Middleware(middleware.Middleware.MiddlewareAuth)
				//绑定LoginWeb
				group.Bind(
					controller.LoginWeb,
				)
			})
			//////////////
			s.Run()
			return nil
		},
	}
)
