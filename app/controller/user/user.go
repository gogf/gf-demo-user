package ctl_user

import (
    "github.com/gogf/gf-demos/app/library/response"
    "github.com/gogf/gf-demos/app/library/user"
    "github.com/gogf/gf/g/net/ghttp"
    "github.com/gogf/gf/g/util/gvalid"
)

// 用户API管理对象
type Controller struct { }

// 用户注册接口
func (c *Controller) SignUp(r *ghttp.Request) {
    if err := lib_user.SignUp(r.GetPostMap()); err != nil {
        lib_response.Json(r, 1, err.Error())
    } else {
        lib_response.Json(r, 0, "ok")
    }
}

// 用户登录接口
func (c *Controller) SignIn(r *ghttp.Request) {
    data  := r.GetPostMap()
    rules := map[string]string {
        "passport"  : "required",
        "password"  : "required",
    }
    msgs  := map[string]interface{} {
        "passport" : "账号不能为空",
        "password" : "密码不能为空",
    }
    if e := gvalid.CheckMap(data, rules, msgs); e != nil {
        lib_response.Json(r, 1, e.String())
    }
    if err := lib_user.SignIn(data["passport"], data["password"], r.Session); err != nil {
        lib_response.Json(r, 1, err.Error())
    } else {
        lib_response.Json(r, 0, "ok")
    }
}

// 判断用户是否已经登录
func (c *Controller) IsSignedIn(r *ghttp.Request) {
    if lib_user.IsSignedIn(r.Session) {
        lib_response.Json(r, 0, "ok")
    } else {
        lib_response.Json(r, 1, "")
    }
}

// 用户注销/退出接口
func (c *Controller) SignOut(r *ghttp.Request) {
    lib_user.SignOut(r.Session)
    lib_response.Json(r, 0, "ok")
}

// 检测用户账号接口(唯一性校验)
func (c *Controller) CheckPassport(r *ghttp.Request) {
    passport := r.Get("passport")
    if e := gvalid.Check(passport, "required", "请输入账号"); e != nil {
        lib_response.Json(r, 1, e.String())
    }
    if lib_user.CheckPassport(passport) {
        lib_response.Json(r, 0, "ok")
    }
    lib_response.Json(r, 1, "账号已经存在")
}

// 检测用户昵称接口(唯一性校验)
func (c *Controller) CheckNickName(r *ghttp.Request) {
    nickname := r.Get("nickname")
    if e := gvalid.Check(nickname, "required", "请输入昵称"); e != nil {
        lib_response.Json(r, 1, e.String())
    }
    if lib_user.CheckNickName(r.Get("nickname")) {
        lib_response.Json(r, 0, "ok")
    }
    lib_response.Json(r, 1, "昵称已经存在")
}

