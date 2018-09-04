package router

import (
	"gitee.com/johng/gf-cases/app/ctl/chat"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/net/ghttp"
)

// 统一路由注册.
func init() {
	s := g.Server()
	// 某些浏览器直接请求favicon.ico文件，特别是产生404时
	s.BindHandler("/favicon.ico", func(r *ghttp.Request) {
		r.Response.ServeFile("/static/resource/image/favicon.ico")
	})
	// 聊天室示例程序使用**控制器注册方式**
	s.BindController("/chat", new(ctlchat.Controller))
}
