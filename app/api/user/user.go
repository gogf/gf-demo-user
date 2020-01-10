package user

import (
	"github.com/gogf/gf-demos/app/service/user"
	"github.com/gogf/gf-demos/library/response"
	"github.com/gogf/gf/net/ghttp"
)

// 用户API管理对象
type Controller struct{}

// 注册请求参数，用于前后端交互参数格式约定
type SignUpRequest struct {
	user.SignUpInput
}

// 用户注册接口
func (c *Controller) SignUp(r *ghttp.Request) {
	var data *SignUpRequest
	// 这里没有使用Parse而是仅用GetStruct获取对象，
	// 数据校验交给后续的service层统一处理
	if err := r.GetStruct(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := user.SignUp(&data.SignUpInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok")
	}
}

// 注册请求参数，用于前后端交互参数格式约定
type SignInRequest struct {
	Passport string `v:"required#账号不能为空"`
	Password string `v:"required#密码不能为空"`
}

// 用户登录接口
func (c *Controller) SignIn(r *ghttp.Request) {
	var data *SignInRequest
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := user.SignIn(data.Passport, data.Password, r.Session); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok")
	}
}

// 判断用户是否已经登录
func (c *Controller) IsSignedIn(r *ghttp.Request) {
	if user.IsSignedIn(r.Session) {
		response.JsonExit(r, 0, "ok")
	} else {
		response.JsonExit(r, 1, "")
	}
}

// 用户注销/退出接口
func (c *Controller) SignOut(r *ghttp.Request) {
	if err := user.SignOut(r.Session); err != nil {
		response.JsonExit(r, 1, "")
	}
	response.JsonExit(r, 0, "ok")
}

// 账号唯一性检测请求参数，用于前后端交互参数格式约定
type CheckPassportRequest struct {
	Passport string
}

// 检测用户账号接口(唯一性校验)
func (c *Controller) CheckPassport(r *ghttp.Request) {
	var data *CheckPassportRequest
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if data.Passport != "" && !user.CheckPassport(data.Passport) {
		response.JsonExit(r, 1, "账号已经存在")
	}
	response.JsonExit(r, 0, "ok")
}

// 账号唯一性检测请求参数，用于前后端交互参数格式约定
type CheckNickNameRequest struct {
	Nickname string
}

// 检测用户昵称接口(唯一性校验)
func (c *Controller) CheckNickName(r *ghttp.Request) {
	var data *CheckNickNameRequest
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if data.Nickname != "" && !user.CheckNickName(data.Nickname) {
		response.JsonExit(r, 1, "昵称已经存在")
	}
	response.JsonExit(r, 0, "ok")
}

// 获取用户详情
func (c *Controller) Profile(r *ghttp.Request) {
	response.JsonExit(r, 0, "", user.GetProfile(r.Session))
}
