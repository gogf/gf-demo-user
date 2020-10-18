package user

import (
	"github.com/gogf/gf-demos/app/service/user"
	"github.com/gogf/gf-demos/library/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// 用户API管理对象
type C struct{}

// 注册请求参数，用于前后端交互参数格式约定
type SignUpRequest struct {
	Passport  string `v:"required|length:6,16#账号不能为空|账号长度应当在:min到:max之间"`
	Password  string `v:"required|length:6,16#请输入确认密码|密码长度应当在:min到:max之间"`
	Password2 string `v:"required|length:6,16|same:Password#密码不能为空|密码长度应当在:min到:max之间|两次密码输入不相等"`
	Nickname  string
}

// @summary 用户注册接口
// @tags    用户服务
// @produce json
// @param   passport  formData string  true "用户账号名称"
// @param   password  formData string  true "用户密码"
// @param   password2 formData string  true "确认密码"
// @param   nickname  formData string false "用户昵称"
// @router  /user/signup [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (c *C) SignUp(r *ghttp.Request) {
	var (
		data        *SignUpRequest
		signUpParam *user.SignUpParam
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(data, &signUpParam); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := user.SignUp(signUpParam); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok")
	}
}

// 登录请求参数，用于前后端交互参数格式约定
type SignInRequest struct {
	Passport string `v:"required#账号不能为空"`
	Password string `v:"required#密码不能为空"`
}

// @summary 用户登录接口
// @tags    用户服务
// @produce json
// @param   passport formData string true "用户账号"
// @param   password formData string true "用户密码"
// @router  /user/signin [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (c *C) SignIn(r *ghttp.Request) {
	var (
		data *SignInRequest
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := user.SignIn(data.Passport, data.Password, r.Session); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok")
	}
}

// @summary 判断用户是否已经登录
// @tags    用户服务
// @produce json
// @router  /user/is-signed-in [GET]
// @success 200 {object} response.JsonResponse "执行结果:`true/false`"
func (c *C) IsSignedIn(r *ghttp.Request) {
	response.JsonExit(r, 0, "", user.IsSignedIn(r.Session))
}

// @summary 用户注销/退出接口
// @tags    用户服务
// @produce json
// @router  /user/sign-out [GET]
// @success 200 {object} response.JsonResponse "执行结果, 1: 未登录"
func (c *C) SignOut(r *ghttp.Request) {
	if err := user.SignOut(r.Session); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "ok")
}

// 账号唯一性检测请求参数，用于前后端交互参数格式约定
type CheckPassportRequest struct {
	Passport string `v:"required#账号不能为空"`
}

// @summary 检测用户账号接口(唯一性校验)
// @tags    用户服务
// @produce json
// @param   passport query string true "用户账号"
// @router  /user/check-passport [GET]
// @success 200 {object} response.JsonResponse "执行结果:`true/false`"
func (c *C) CheckPassport(r *ghttp.Request) {
	var (
		data *CheckPassportRequest
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if data.Passport != "" && !user.CheckPassport(data.Passport) {
		response.JsonExit(r, 0, "账号已经存在", false)
	}
	response.JsonExit(r, 0, "", true)
}

// 账号唯一性检测请求参数，用于前后端交互参数格式约定
type CheckNickNameRequest struct {
	Nickname string `v:"required#昵称不能为空"`
}

// @summary 检测用户昵称接口(唯一性校验)
// @tags    用户服务
// @produce json
// @param   nickname query string true "用户昵称"
// @router  /user/checkpassport [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (c *C) CheckNickName(r *ghttp.Request) {
	var (
		data *CheckNickNameRequest
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if data.Nickname != "" && !user.CheckNickName(data.Nickname) {
		response.JsonExit(r, 0, "昵称已经存在", false)
	}
	response.JsonExit(r, 0, "ok", true)
}

// @summary 获取用户详情信息
// @tags    用户服务
// @produce json
// @router  /user/profile [GET]
// @success 200 {object} user.Entity "用户信息"
func (c *C) Profile(r *ghttp.Request) {
	response.JsonExit(r, 0, "", user.GetProfile(r.Session))
}
