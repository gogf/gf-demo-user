package router

import (
	"gitee.com/johng/gf-demos/app/controller/chat"
	"gitee.com/johng/gf/g"
)

// 统一路由注册.
func init() {
	s := g.Server()
	// 某些浏览器直接请求favicon.ico文件，特别是产生404时
	s.SetRewrite("/favicon.ico", "/resource/image/favicon.ico")
	// 聊天室示例程序使用**控制器注册方式**
	s.BindController("/chat", new(ctl_chat.Controller))
}
