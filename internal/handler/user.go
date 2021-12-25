package handler

import (
	"context"

	"github.com/gogf/gf-demos/v2/apiv1"
	"github.com/gogf/gf-demos/v2/internal/model"
	"github.com/gogf/gf-demos/v2/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
)

var User = hUser{}

type hUser struct{}

func (h *hUser) SignUp(ctx context.Context, req *apiv1.UserSignUpReq) (res *apiv1.UserSignUpRes, err error) {
	err = service.User().Create(ctx, model.UserCreateInput{
		Passport: req.Passport,
		Password: req.Password,
		Nickname: req.Nickname,
	})
	return
}
func (h *hUser) SignIn(ctx context.Context, req *apiv1.UserSignInReq) (res *apiv1.UserSignInRes, err error) {
	err = service.User().SignIn(ctx, model.UserSignInInput{
		Passport: req.Passport,
		Password: req.Password,
	})
	return
}

func (h *hUser) IsSignedIn(ctx context.Context, req *apiv1.UserIsSignedInReq) (res *apiv1.UserIsSignedInRes, err error) {
	res = &apiv1.UserIsSignedInRes{
		OK: service.User().IsSignedIn(ctx),
	}
	return
}

func (h *hUser) SignOut(ctx context.Context, req *apiv1.UserSignOutReq) (res *apiv1.UserSignOutRes, err error) {
	err = service.User().SignOut(ctx)
	return
}

func (h *hUser) CheckPassport(ctx context.Context, req *apiv1.UserCheckPassportReq) (res *apiv1.UserCheckPassportRes, err error) {
	available, err := service.User().IsPassportAvailable(ctx, req.Passport)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf(`Passport "%s" is already token by others`, req.Passport)
	}
	return
}

func (h *hUser) CheckNickName(ctx context.Context, req *apiv1.UserCheckNickNameReq) (res *apiv1.UserCheckNickNameRes, err error) {
	available, err := service.User().IsNicknameAvailable(ctx, req.Nickname)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf(`Nickname "%s" is already token by others`, req.Nickname)
	}
	return
}

func (h *hUser) Profile(ctx context.Context, req *apiv1.UserProfileReq) (res *apiv1.UserProfileRes, err error) {
	res = &apiv1.UserProfileRes{
		User: service.User().GetProfile(ctx),
	}
	return
}
