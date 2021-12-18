package apiv1

import (
	"github.com/gogf/gf-demos/v2/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type UserProfileReq struct {
	g.Meta `path:"/user/profile" method:"get" tags:"User" summary:"Get the profile of current user"`
}
type UserProfileRes struct {
	*entity.User
}

type UserSignUpReq struct {
	g.Meta    `path:"/user/sign-up" method:"post" tags:"User" summary:"Sign up a new user account"`
	Passport  string `v:"required|length:6,16"`
	Password  string `v:"required|length:6,16"`
	Password2 string `v:"required|length:6,16|same:Password"`
	Nickname  string
}
type UserSignUpRes struct{}

type UserSignInReq struct {
	g.Meta   `path:"/user/sign-in" method:"post" tags:"User" summary:"Sign in with exist account"`
	Passport string `v:"required"`
	Password string `v:"required"`
}
type UserSignInRes struct{}

type UserCheckPassportReq struct {
	g.Meta   `path:"/user/check-passport" method:"post" tags:"User" summary:"Check passport available"`
	Passport string `v:"required"`
}
type UserCheckPassportRes struct{}

type UserCheckNickNameReq struct {
	g.Meta   `path:"/user/check-passport" method:"post" tags:"User" summary:"Check nickname available"`
	Nickname string `v:"required"`
}
type UserCheckNickNameRes struct{}

type UserIsSignedInReq struct {
	g.Meta `path:"/user/is-signed-in" method:"post" tags:"User" summary:"Check current user is already signed-in"`
}
type UserIsSignedInRes struct {
	OK bool `dc:"True if current user is signed in; or else false"`
}

type UserSignOutReq struct {
	g.Meta `path:"/user/sign-out" method:"post" tags:"User" summary:"Sign out current user"`
}
type UserSignOutRes struct{}
