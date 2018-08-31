package router

import (
    "gitee.com/johng/gf/g"
    "gitee.com/johng/gf/g/net/ghttp"
    "gitee.com/johng/gf-cases/app/ctl/chat"
)

// 统一路由注册.
func init() {
    s := g.Server()
    // 某些浏览器直接请求favicon.ico文件，特别是产生404时
    s.BindHandler("/favicon.ico", func(r *ghttp.Request) {
        r.Response.ServeFile("/static/resource/image/favicon.ico")
    })
    s.BindController("/chat", new(ctl_chat.Controller))
}
