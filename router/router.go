package router

import (
	"github.com/gogf/gf-demos/app/api/chat"
	"github.com/gogf/gf-demos/app/api/user"
	"github.com/gogf/gf/frame/g"
)

// 统一路由注册.
func init() {
	s := g.Server()

	// 某些浏览器直接请求favicon.ico文件，特别是产生404时
	s.SetRewrite("/favicon.ico", "/resource/image/favicon.ico")

	// 用户模块 路由注册 - 使用执行对象注册方式
	s.BindObject("/user", new(user.Controller))

	// 聊天模块 路由注册 - 使用控制器注册方式
	s.BindController("/chat", new(chat.Controller))
}
