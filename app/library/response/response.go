package lib_response

import (
    "github.com/gogf/gf/g"
    "github.com/gogf/gf/g/net/ghttp"
)

// 标准返回结果数据结构封装。
// 返回固定数据结构的JSON:
// err:  错误码(0:成功, 1:失败, >1:错误码);
// msg:  请求结果信息;
// data: 请求结果,根据不同接口返回结果的数据结构不同;
func Json(r *ghttp.Request, err int, msg string, data...interface{}) {
    responseData := interface{}(nil)
    if len(data) > 0 {
        responseData = data[0]
    }
    r.Response.WriteJson(g.Map{
        "err"  : err,
        "msg"  : msg,
        "data" : responseData,
    })
    r.ExitAll()
}
