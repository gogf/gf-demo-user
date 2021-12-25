package service

import (
	"net/http"

	"github.com/gogf/gf-demos/v2/internal/model"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	// SMiddleware is service struct of module Middleware.
	SMiddleware struct{}
)

var (
	// insMiddleware is the instance of service Middleware.
	insMiddleware = SMiddleware{}
)

// Middleware returns the interface of Middleware service.
func Middleware() SMiddleware {
	return insMiddleware
}

// Ctx injects custom business context variable into context of current request.
func (s SMiddleware) Ctx(r *ghttp.Request) {
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
func (s SMiddleware) Auth(r *ghttp.Request) {
	if User().IsSignedIn(r.Context()) {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatus(http.StatusForbidden)
	}
}

// CORS allows Cross-origin resource sharing.
func (s SMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
