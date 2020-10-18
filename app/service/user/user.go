package user

import (
	"errors"
	"fmt"
	"github.com/gogf/gf-demos/app/model/user"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

const (
	SessionMark = "user_info"
)

// 注册输入参数
type SignUpParam struct {
	Passport  string `v:"required|length:6,16#账号不能为空|账号长度应当在:min到:max之间"`
	Password  string `v:"required|length:6,16#请输入确认密码|密码长度应当在:min到:max之间"`
	Password2 string `v:"required|length:6,16|same:Password#密码不能为空|密码长度应当在:min到:max之间|两次密码输入不相等"`
	Nickname  string
}

// 用户注册
func SignUp(param *SignUpParam) error {
	// 昵称为非必需参数，默认使用账号名称
	if param.Nickname == "" {
		param.Nickname = param.Passport
	}
	// 账号唯一性数据检查
	if !CheckPassport(param.Passport) {
		return errors.New(fmt.Sprintf("账号 %s 已经存在", param.Passport))
	}
	// 昵称唯一性数据检查
	if !CheckNickName(param.Nickname) {
		return errors.New(fmt.Sprintf("昵称 %s 已经存在", param.Nickname))
	}
	// 将输入参数赋值到数据库实体对象上
	var entity *user.Entity
	if err := gconv.Struct(param, &entity); err != nil {
		return err
	}
	if _, err := user.Model.Save(entity); err != nil {
		return err
	}
	return nil
}

// 判断用户是否已经登录
func IsSignedIn(session *ghttp.Session) bool {
	return session.Contains(SessionMark)
}

// 用户登录，成功返回用户信息，否则返回nil; passport应当会md5值字符串
func SignIn(passport, password string, session *ghttp.Session) error {
	one, err := user.Model.FindOne("passport=? and password=?", passport, password)
	if err != nil {
		return err
	}
	if one == nil {
		return errors.New("账号或密码错误")
	}
	return session.Set(SessionMark, one)
}

// 用户注销
func SignOut(session *ghttp.Session) error {
	return session.Remove(SessionMark)
}

// 检查账号是否符合规范(目前仅检查唯一性),存在返回false,否则true
func CheckPassport(passport string) bool {
	if i, err := user.Model.FindCount("passport", passport); err != nil {
		return false
	} else {
		return i == 0
	}
}

// 检查昵称是否符合规范(目前仅检查唯一性),存在返回false,否则true
func CheckNickName(nickname string) bool {
	if i, err := user.Model.FindCount("nickname", nickname); err != nil {
		return false
	} else {
		return i == 0
	}
}

// 获得用户信息详情
func GetProfile(session *ghttp.Session) (u *user.Entity) {
	_ = session.GetStruct(SessionMark, &u)
	return
}
