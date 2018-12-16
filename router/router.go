package router

import (
    "gitee.com/johng/gf-demos/app/controller/chat"
    "gitee.com/johng/gf-demos/app/controller/user"
    "gitee.com/johng/gf/g"
)

// 统一路由注册.
func init() {
	s := g.Server()

	// 某些浏览器直接请求favicon.ico文件，特别是产生404时
	s.SetRewrite("/favicon.ico", "/resource/image/favicon.ico")

    // 用户模块 路由注册 - 使用执行对象注册方式
    s.BindObject("/user", new(ctl_user.Controller))

	// 聊天模块 路由注册 - 使用控制器注册方式
	s.BindController("/chat", new(ctl_chat.Controller))
}
