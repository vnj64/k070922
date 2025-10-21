package core

import (
	v1 "project/api/v1"
	"project/domain"
	"project/domain/cases/auth"
	"project/domain/cases/user"
	"project/pkg/middleware"
)

type Di struct {
	Ctx         domain.Context
	UserHandler *v1.UserHandler
	AuthHandler *v1.AuthHandler
}

func NewDi() *Di {
	ctx := InitCtx()
	cfg := ctx.Infra().Config()

	mwr := middleware.NewMiddleware(cfg.JwtPublicPemPath())

	var (
		// Domain User
		userUseCase = user.NewUserUseCase(ctx)
		userHandler = v1.NewUserHandler(*userUseCase, mwr)

		// Domain Refresh Token
		// Domain Auth
		authUseCase = auth.NewAuthUseCase(ctx)
		authHandler = v1.NewAuthHandler(authUseCase)
	)

	return &Di{
		Ctx:         ctx,
		UserHandler: userHandler,
		AuthHandler: authHandler,
	}
}
