package cases

import (
	"github.com/google/uuid"
	"project/domain"
	"project/domain/models"
	"project/domain/props"
	"project/pkg/errs"
	"project/pkg/timestamps"
	"project/pkg/validation"
	"time"
)

type UserUseCase struct {
	ctx domain.Context
}

func NewUserUseCase(ctx domain.Context) *UserUseCase {
	return &UserUseCase{
		ctx: ctx,
	}
}

func (uc *UserUseCase) CreateUser(args props.CreateUserReq) (resp props.CreateUserResp, err error) {
	if err := validation.CheckEmailValidation(args.Email); err != nil {
		return resp, errs.NewErrorWithDetails(errs.ErrUnprocessableEntity, err.Error())
	}

	stamp := timestamps.Timestamps{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user := &models.User{
		UUID:       uuid.New(),
		Timestamps: stamp,
	}
	if args.FirstName != nil {
		user.FirstName = args.FirstName
	}
	if args.SecondName != nil {
		user.SecondName = args.SecondName
	}

	if err := uc.ctx.Connection().UserRepository().Add(user); err != nil {
		return resp, errs.NewErrorWithDetails(errs.ErrInternalServerError, err.Error())
	}
	resp.User = user

	return resp, nil
}
