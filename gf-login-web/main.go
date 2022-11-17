package main

import (
	//增加logic的引入 这样可以框架就会执行logic匿名包初始化方法
	_ "gf-login-web/internal/logic"

	_ "gf-login-web/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gf-login-web/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}

//readme: 这只是一个示例，不是真正的项目
