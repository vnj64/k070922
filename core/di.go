package core

import (
	v1 "project/api/v1"
	"project/domain"
	"project/domain/cases"
	"project/pkg/middleware"
)

type Di struct {
	Ctx         domain.Context
	UserHandler *v1.UserHandler
}

func NewDi() *Di {
	ctx := InitCtx()
	cfg := ctx.Infra().Config()

	mwr := middleware.NewMiddleware(cfg.JwtPublicPemPath())

	var (
		// Domain User
		userUseCase = cases.NewUserUseCase(ctx)
		userHandler = v1.NewUserHandler(*userUseCase, mwr)

		// Domain Refresh Token
	)

	return &Di{
		Ctx:         ctx,
		UserHandler: userHandler,
	}
}
