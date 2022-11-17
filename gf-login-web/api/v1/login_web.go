package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginUIGetReq struct {
	g.Meta `path:"/" tags:"LoginWeb" method:"get" summary:"登录界面 获取 api"`
}

type LoginUIGetRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type LoginGetReq struct {
	g.Meta `path:"/login" tags:"LoginWeb" method:"get" summary:"登录 获取 api"`
}

type LoginGetRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type LoginPostReq struct {
	g.Meta `path:"/login" tags:"LoginWeb" method:"post" summary:"登录 请求 api"`
}

type LoginPostRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type LogoutPostReq struct {
	g.Meta `path:"/logout" tags:"LoginWeb" method:"post" summary:"登出 请求 api"`
}

type LogoutPostRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
