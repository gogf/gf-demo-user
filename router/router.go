package router

import (
	"github.com/gogf/gf-demos/app/api/chat"
	"github.com/gogf/gf-demos/app/api/user"
	"github.com/gogf/gf-demos/app/service/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 统一路由注册.
func init() {
	s := g.Server()

	// 某些浏览器直接请求favicon.ico文件，特别是产生404时
	s.SetRewrite("/favicon.ico", "/resource/image/favicon.ico")

	s.Group("/", func(group *ghttp.RouterGroup) {
		ctlChat := new(chat.Controller)
		ctlUser := new(user.Controller)
		group.Middleware(middleware.CORS)
		//// 聊天模块 路由注册 - 使用控制器注册方式(未来不推荐)
		group.ALL("/chat", ctlChat)
		//// 用户模块 路由注册 - 使用执行对象注册方式
		group.ALL("/user", ctlUser)

		// 鉴权路由注册
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.Auth)
			group.ALL("/user/profile", ctlUser, "Profile")
		})
	})
}
