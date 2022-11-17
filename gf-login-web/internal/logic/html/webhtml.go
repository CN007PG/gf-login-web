package html

import (
	"gf-login-web/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	sWebHtml struct{}
)

// 此方法 为依赖注入时框架执行的匿名包初始化方法init
// 此方法 要在terminal执行gf gen service 命令导出service下的html.go后再加入
func init() {
	service.RegisterWebHtml(New())
}

// 此方法在init调用的 同上要在terminal执行gf gen service 命令导出service下的html.go后再加入
func New() *sWebHtml {
	return &sWebHtml{}
}

// 直接返回主页html
func (s *sWebHtml) ReturnIndexHtml(r *ghttp.Request, username string) {
	// 返回登录成功页面
	r.Response.WriteTpl("index.html", g.Map{
		"tips":           "欢迎您！",
		"username":       username,
		"button_display": "inline-block",
	})
	r.Exit()
}
