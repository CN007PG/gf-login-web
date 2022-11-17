package utility

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

// 取session 的用户名
func GetSessionUsername(r *ghttp.Request, sessionKey string, usernameKey string) string {
	// 返回相应的 session key 的 username key 的用户名
	re_val, _ := r.Session.Get(sessionKey)
	return gconv.String(re_val.Map()[usernameKey])
}
