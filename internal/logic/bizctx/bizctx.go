package bizctx

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/gogf/gf-demo-user/v2/internal/consts"
	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

type (
	sBizCtx struct{}
)

func init() {
	service.RegisterBizCtx(New())
}

func New() service.IBizCtx {
	return &sBizCtx{}
}

// Init initializes and injects custom business context object into request context.
func (s *sBizCtx) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.ContextKey, customCtx)
}

// Get retrieves and returns the user object from context.
// It returns nil if nothing found in given context.
func (s *sBizCtx) Get(ctx context.Context) *model.Context {
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
func (s *sBizCtx) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	s.Get(ctx).User = ctxUser
}
