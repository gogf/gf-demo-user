package middleware

import (
	"net/http"

	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

type (
	sMiddleware struct{}
)

func init() {
	service.RegisterMiddleware(New())
}

func New() service.IMiddleware {
	return &sMiddleware{}
}

// Ctx injects custom business context variable into context of current request.
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	customCtx := &model.Context{
		Session: r.Session,
	}
	service.BizCtx().Init(r, customCtx)
	if user := service.Session().GetUser(r.Context()); user != nil {
		customCtx.User = &model.ContextUser{
			Id:       user.Id,
			Passport: user.Passport,
			Nickname: user.Nickname,
		}
	}
	// Continue execution of next middleware.
	r.Middleware.Next()
}

// Auth validates the request to allow only signed-in users visit.
func (s *sMiddleware) Auth(r *ghttp.Request) {
	if service.User().IsSignedIn(r.Context()) {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatus(http.StatusForbidden)
	}
}

// CORS allows Cross-origin resource sharing.
func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
