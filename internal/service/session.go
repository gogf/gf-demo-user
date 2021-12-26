package service

import (
	"context"

	"github.com/gogf/gf-demos/v2/internal/consts"
	"github.com/gogf/gf-demos/v2/internal/model/entity"
)

type (
	// sSession is service struct of module Session.
	sSession struct{}
)

var (
	// insSession is the instance of service Session.
	insSession = sSession{}
)

// Session returns the interface of Session service.
func Session() *sSession {
	return &insSession
}

// SetUser sets user into the session.
func (s *sSession) SetUser(ctx context.Context, user *entity.User) error {
	return Context().Get(ctx).Session.Set(consts.UserSessionKey, user)
}

// GetUser retrieves and returns the user from session.
// It returns nil if the user did not sign in.
func (s *sSession) GetUser(ctx context.Context) *entity.User {
	customCtx := Context().Get(ctx)
	if customCtx != nil {
		if v := customCtx.Session.MustGet(consts.UserSessionKey); !v.IsNil() {
			var user *entity.User
			_ = v.Struct(&user)
			return user
		}
	}
	return nil
}

// RemoveUser removes user rom session.
func (s *sSession) RemoveUser(ctx context.Context) error {
	customCtx := Context().Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(consts.UserSessionKey)
	}
	return nil
}
