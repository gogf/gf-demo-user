package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf-demos/app/dao"
	"github.com/gogf/gf-demos/app/model"
)

// 中间件管理服务
var User = userService{}

type userService struct{}

// 用户注册
func (s *userService) SignUp(r *model.UserServiceSignUpReq) error {
	// 昵称为非必需参数，默认使用账号名称
	if r.Nickname == "" {
		r.Nickname = r.Passport
	}
	// 账号唯一性数据检查
	if !s.CheckPassport(r.Passport) {
		return errors.New(fmt.Sprintf("账号 %s 已经存在", r.Passport))
	}
	// 昵称唯一性数据检查
	if !s.CheckNickName(r.Nickname) {
		return errors.New(fmt.Sprintf("昵称 %s 已经存在", r.Nickname))
	}
	if _, err := dao.User.Save(r); err != nil {
		return err
	}
	return nil
}

// 判断用户是否已经登录
func (s *userService) IsSignedIn(ctx context.Context) bool {
	if v := Context.Get(ctx); v != nil && v.User != nil {
		return true
	}
	return false
}

// 用户登录，成功返回用户信息，否则返回nil; passport应当会md5值字符串
func (s *userService) SignIn(ctx context.Context, passport, password string) error {
	var user *model.User
	err := dao.User.Where("passport=? and password=?", passport, password).Scan(&user)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("账号或密码错误")
	}
	if err := Session.SetUser(ctx, user); err != nil {
		return err
	}
	Context.SetUser(ctx, &model.ContextUser{
		Id:       user.Id,
		Passport: user.Passport,
		Nickname: user.Nickname,
	})
	return nil
}

// 用户注销
func (s *userService) SignOut(ctx context.Context) error {
	return Session.RemoveUser(ctx)
}

// 检查账号是否符合规范(目前仅检查唯一性),存在返回false,否则true
func (s *userService) CheckPassport(passport string) bool {
	if i, err := dao.User.FindCount("passport", passport); err != nil {
		return false
	} else {
		return i == 0
	}
}

// 检查昵称是否符合规范(目前仅检查唯一性),存在返回false,否则true
func (s *userService) CheckNickName(nickname string) bool {
	if i, err := dao.User.FindCount("nickname", nickname); err != nil {
		return false
	} else {
		return i == 0
	}
}

// 获得用户信息详情
func (s *userService) GetProfile(ctx context.Context) *model.User {
	return Session.GetUser(ctx)
}
