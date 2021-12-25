package service

import (
	"context"

	"github.com/gogf/gf-demos/v2/internal/consts"
	"github.com/gogf/gf-demos/v2/internal/model"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	// SContext is service struct of module Context.
	SContext struct{}
)

var (
	// insContext is the instance of service Context.
	insContext = SContext{}
)

// Context returns the interface of Context service.
func Context() SContext {
	return insContext
}

// Init initializes and injects custom business context object into request context.
func (s SContext) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.ContextKey, customCtx)
}

// Get retrieves and returns the user object from context.
// It returns nil if nothing found in given context.
func (s SContext) Get(ctx context.Context) *model.Context {
	value := ctx.Value(consts.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// SetUser injects business user object into context.
func (s SContext) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	s.Get(ctx).User = ctxUser
}
