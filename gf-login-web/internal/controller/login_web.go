package controller

import (
	"context"
	"gf-login-web/internal/consts"
	"gf-login-web/internal/service"
	"gf-login-web/utility"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"gf-login-web/api/v1"
)

var (
	LoginWeb = cLoginWeb{}
)

type cLoginWeb struct{}

// 登录界面 获取 api 监听回调
func (c *cLoginWeb) LoginGetUI(ctx context.Context, req *v1.LoginUIGetReq) (res *v1.LoginUIGetRes, err error) {
	g.RequestFromCtx(ctx).Response.WriteTpl("login.html", g.Map{
		"title": "登录页面",
	})
	return
}

// 登录 获取 api 监听回调
func (c *cLoginWeb) LoginGet(ctx context.Context, req *v1.LoginGetReq) (res *v1.LoginGetRes, err error) {
	username := utility.GetSessionUsername(g.RequestFromCtx(ctx), consts.SessionLoginWeb, "username")
	// 返回登录成功页面
	service.WebHtml().ReturnIndexHtml(g.RequestFromCtx(ctx), gconv.String(username))
	return
}

// 登录 请求 api 监听回调
func (c *cLoginWeb) LoginPost(ctx context.Context, req *v1.LoginPostReq) (res *v1.LoginPostRes, err error) {
	username := g.RequestFromCtx(ctx).Get("username")
	password := g.RequestFromCtx(ctx).Get("password")

	// 暂时写死 后面再用数据库
	dbUsername := "admin"
	dbPassword := "123456"
	if username.String() == dbUsername && password.String() == dbPassword {
		// 添加session
		g.RequestFromCtx(ctx).Session.Set(consts.SessionLoginWeb, g.Map{
			"username": dbUsername,
		})
		// 返回登录成功页面
		service.WebHtml().ReturnIndexHtml(g.RequestFromCtx(ctx), gconv.String(username))
	}
	// 返回登录失败
	g.RequestFromCtx(ctx).Response.WriteJson(g.Map{
		"code": -1,
		"msg":  "登录失败",
	})
	return
}

// 登出 请求 api 监听回调
func (c *cLoginWeb) LogoutPost(ctx context.Context, req *v1.LogoutPostReq) (res *v1.LogoutPostRes, err error) {
	// 删除session
	g.RequestFromCtx(ctx).Session.Remove(consts.SessionLoginWeb)
	// 返回登出页面
	g.RequestFromCtx(ctx).Response.WriteTpl("index.html", g.Map{
		"tips":           "再见！",
		"username":       "",
		"button_display": "none",
	})
	return
}
