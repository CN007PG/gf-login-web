# 一、示例简介：

 1. 示例通过ghttp.Server注册一个中间件来做session的认证(如果有相应的session表示刚才登录过，直接返回登录过的结果)，接着处理四个协议："登录界面 获取 api"，"登录 获取 api"，"登录 请求 api"，"登出 请求 api"，最后把相应的结果（其中包括html网页）返回
 2. 通过cmd.go，consts.go，login_web.go，webhtml.go，midddleware.go，html.go，config.yaml，web_html_util.go，main.go（有部分文件是框架生成的）等文件了解GoFrame框架的规范和如何使用
