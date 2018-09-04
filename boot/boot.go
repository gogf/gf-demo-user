package boot

import (
    "gitee.com/johng/gf/g"
    "gitee.com/johng/gf/g/os/glog"
    "gitee.com/johng/gf/g/net/ghttp"
)

// 用于应用初始化。
func init() {
    v := g.View()
    c := g.Config()
    s := g.Server()

    // 配置对象及视图对象配置
    c.AddPath("config")
    v.AddPath("static/template")
    v.SetDelimiters("${", "}")

    // glog配置
    logpath := c.GetString("setting.logpath")
    glog.SetPath(logpath)
    glog.SetStdPrint(true)

    // Web Server配置
    s.SetDenyRoutes([]string{
        "/config/*",
    })
    s.SetLogPath(logpath)
    s.SetNameToUriType(ghttp.NAME_TO_URI_TYPE_ALLLOWER)
    s.SetErrorLogEnabled(true)
    s.SetAccessLogEnabled(true)
    s.SetPort(8199)
}

