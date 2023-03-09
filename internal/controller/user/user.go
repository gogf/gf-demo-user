package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf-demo-user/v2/api/user/v1"
	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// SignUp is the API for user sign up.
func (c *Controller) SignUp(ctx context.Context, req *v1.SignUpReq) (res *v1.SignUpRes, err error) {
	err = service.User().Create(ctx, model.UserCreateInput{
		Passport: req.Passport,
		Password: req.Password,
		Nickname: req.Nickname,
	})
	return
}

// SignIn is the API for user sign in.
func (c *Controller) SignIn(ctx context.Context, req *v1.SignInReq) (res *v1.SignInRes, err error) {
	err = service.User().SignIn(ctx, model.UserSignInInput{
		Passport: req.Passport,
		Password: req.Password,
	})
	return
}

// IsSignedIn checks and returns whether the user is signed in.
func (c *Controller) IsSignedIn(ctx context.Context, req *v1.IsSignedInReq) (res *v1.IsSignedInRes, err error) {
	res = &v1.IsSignedInRes{
		OK: service.User().IsSignedIn(ctx),
	}
	return
}

// SignOut is the API for user sign out.
func (c *Controller) SignOut(ctx context.Context, req *v1.SignOutReq) (res *v1.SignOutRes, err error) {
	err = service.User().SignOut(ctx)
	return
}

// CheckPassport checks and returns whether the user passport is available.
func (c *Controller) CheckPassport(ctx context.Context, req *v1.CheckPassportReq) (res *v1.CheckPassportRes, err error) {
	available, err := service.User().IsPassportAvailable(ctx, req.Passport)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf(`Passport "%s" is already token by others`, req.Passport)
	}
	return
}

// CheckNickName checks and returns whether the user nickname is available.
func (c *Controller) CheckNickName(ctx context.Context, req *v1.CheckNickNameReq) (res *v1.CheckNickNameRes, err error) {
	available, err := service.User().IsNicknameAvailable(ctx, req.Nickname)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf(`Nickname "%s" is already token by others`, req.Nickname)
	}
	return
}

// Profile returns the user profile.
func (c *Controller) Profile(ctx context.Context, req *v1.ProfileReq) (res *v1.ProfileRes, err error) {
	res = &v1.ProfileRes{
		User: service.User().GetProfile(ctx),
	}
	return
}
