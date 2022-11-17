package middleware

import (
	"gf-login-web/internal/consts"
	"gf-login-web/internal/service"
	"gf-login-web/utility"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	Middleware = sMiddleware{}
)

type (
	sMiddleware struct{}
)

// 认证中间件
func (s *sMiddleware) MiddlewareAuth(r *ghttp.Request) {
	if r.Method == "GET" {
		//对GET模式下的“/"和”/login“申请作Session认证
		if r.Router.Uri == "/" {
			re, _ := r.Session.Contains(consts.SessionLoginWeb)
			if re {
				username := utility.GetSessionUsername(r, consts.SessionLoginWeb, "username")
				// 直接返回登录成功页面
				service.WebHtml().ReturnIndexHtml(r, gconv.String(username))
			} else {
				// 正常 登录
				r.Middleware.Next()
			}
		} else if r.Router.Uri == "/login" {
			re, _ := r.Session.Contains(consts.SessionLoginWeb)
			if re {
				r.Middleware.Next()
			} else {
				// 返回错误码
				r.Response.WriteJson(g.Map{
					"code": 403,
					"msg":  "您访问超时或已登出",
				})
			}
		} else {
			//get的其它申请直接下一步
			r.Middleware.Next()
		}
	} else {
		//post的申请直接下一步
		r.Middleware.Next()
	}

}
