package service

import (
	"context"
	"github.com/gogf/gf-demos/app/model"
)

// Session管理服务
var Session = new(serviceSession)

type serviceSession struct{}

const (
	// 用户信息存放在Session中的Key
	sessionKeyUser = "SessionKeyUser"
)

// 设置用户Session.
func (s *serviceSession) SetUser(ctx context.Context, user *model.User) error {
	return Context.Get(ctx).Session.Set(sessionKeyUser, user)
}

// 获取当前登录的用户信息对象，如果用户未登录返回nil。
func (s *serviceSession) GetUser(ctx context.Context) *model.User {
	customCtx := Context.Get(ctx)
	if customCtx != nil {
		if v := customCtx.Session.GetVar(sessionKeyUser); !v.IsNil() {
			var user *model.User
			_ = v.Struct(&user)
			return user
		}
	}
	return nil
}

// 删除用户Session。
func (s *serviceSession) RemoveUser(ctx context.Context) error {
	customCtx := Context.Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(sessionKeyUser)
	}
	return nil
}
