package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf-demo-user/v2/internal/model/entity"
)

type ProfileReq struct {
	g.Meta `path:"/user/profile" method:"get" tags:"UserService" summary:"Get the profile of current user"`
}
type ProfileRes struct {
	*entity.User
}

type SignUpReq struct {
	g.Meta    `path:"/user/sign-up" method:"post" tags:"UserService" summary:"Sign up a new user account"`
	Passport  string `v:"required|length:6,16"`
	Password  string `v:"required|length:6,16"`
	Password2 string `v:"required|length:6,16|same:Password"`
	Nickname  string
}
type SignUpRes struct{}

type SignInReq struct {
	g.Meta   `path:"/user/sign-in" method:"post" tags:"UserService" summary:"Sign in with exist account"`
	Passport string `v:"required"`
	Password string `v:"required"`
}
type SignInRes struct{}

type CheckPassportReq struct {
	g.Meta   `path:"/user/check-passport" method:"post" tags:"UserService" summary:"Check passport available"`
	Passport string `v:"required"`
}
type CheckPassportRes struct{}

type CheckNickNameReq struct {
	g.Meta   `path:"/user/check-nick-name" method:"post" tags:"UserService" summary:"Check nickname available"`
	Nickname string `v:"required"`
}
type CheckNickNameRes struct{}

type IsSignedInReq struct {
	g.Meta `path:"/user/is-signed-in" method:"post" tags:"UserService" summary:"Check current user is already signed-in"`
}
type IsSignedInRes struct {
	OK bool `dc:"True if current user is signed in; or else false"`
}

type SignOutReq struct {
	g.Meta `path:"/user/sign-out" method:"post" tags:"UserService" summary:"Sign out current user"`
}
type SignOutRes struct{}
