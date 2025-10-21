package auth

import (
	log2 "log"
	"project/domain"
	"project/domain/props"
	"project/pkg/errs"
	"project/pkg/jwt_helpers"
)

type AuthUseCase struct {
	ctx domain.Context
}

func NewAuthUseCase(ctx domain.Context) *AuthUseCase {
	return &AuthUseCase{
		ctx: ctx,
	}
}

func (uc *AuthUseCase) Login(args props.LoginUserReq) (resp props.LoginUserResp, err error) {
	log := uc.ctx.Infra().Config().AdminLogin()
	pass := uc.ctx.Infra().Config().AdminPassword()

	log2.Println(args.Username)
	log2.Println(args.Password)
	if args.Username == log && args.Password == pass {
		user, err := uc.ctx.Connection().UserRepository().GetByEmail(args.Username)
		if err != nil {
			return resp, errs.NewErrorWithDetails(errs.ErrInternalServerError, "database error")
		}
		if user == nil {
			return resp, errs.NewErrorWithDetails(errs.ErrNotFound, "user not found")
		}

		cfg := uc.ctx.Infra().Config()
		acToken, err := jwt_helpers.GenerateAuthToken(user.UUID.String(), user.Role, cfg)
		refToken, err := jwt_helpers.GenerateRefreshToken(user.UUID.String(), cfg)
		if err != nil {
			return resp, errs.NewErrorWithDetails(errs.ErrInternalServerError, err.Error())
		}

		resp.AccessToken = acToken
		resp.RefreshToken = refToken.Token

		return resp, nil
	}

	return resp, nil
}
