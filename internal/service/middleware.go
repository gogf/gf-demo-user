package service

import (
	"net/http"

	"github.com/gogf/gf-demos/v2/internal/model"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	// sMiddleware is service struct of module Middleware.
	sMiddleware struct{}
)

var (
	// insMiddleware is the instance of service Middleware.
	insMiddleware = sMiddleware{}
)

// Middleware returns the interface of Middleware service.
func Middleware() *sMiddleware {
	return &insMiddleware
}

// Ctx injects custom business context variable into context of current request.
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	customCtx := &model.Context{
		Session: r.Session,
	}
	Context().Init(r, customCtx)
	if user := Session().GetUser(r.Context()); user != nil {
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
	if User().IsSignedIn(r.Context()) {
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
