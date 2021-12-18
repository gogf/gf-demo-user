package service

import (
	"context"

	"github.com/gogf/gf-demos/v2/internal/consts"
	"github.com/gogf/gf-demos/v2/internal/model/entity"
)

var Session = serviceSession{}

type serviceSession struct{}

// SetUser sets user into the session.
func (s *serviceSession) SetUser(ctx context.Context, user *entity.User) error {
	return Context.Get(ctx).Session.Set(consts.UserSessionKey, user)
}

// GetUser retrieves and returns the user from session.
// It returns nil if the user did not sign in.
func (s *serviceSession) GetUser(ctx context.Context) *entity.User {
	customCtx := Context.Get(ctx)
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
func (s *serviceSession) RemoveUser(ctx context.Context) error {
	customCtx := Context.Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(consts.UserSessionKey)
	}
	return nil
}
