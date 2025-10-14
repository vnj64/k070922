package v1

import (
	"github.com/gofiber/fiber/v2"
	"project/domain/cases"
	props2 "project/domain/props"
	"project/pkg/errs"
	"project/pkg/middleware"
	"project/shared/dto"
)

type UserHandler struct {
	useCase cases.UserUseCase
	mwr     *middleware.Middleware
}

func NewUserHandler(useCase cases.UserUseCase, mwr *middleware.Middleware) *UserHandler {
	return &UserHandler{
		useCase: useCase,
		mwr:     mwr,
	}
}

// CreateUserHandler godoc
// @Summary Создание пользователя
// @Description Создание пользователя
// @Tags User
// @Accept json
// @Produce json
// @Param request body dto.CreateUserDto true "Данные для создания пользователя"
// @Success 200 {object} dto.CreateUserResponseDto
// @Failure 400 {object} errs.Error "Invalid request"
// @Failure 500 {object} errs.Error "Internal server error"
// @Router /api/v1/user [post]
func (h *UserHandler) CreateUserHandler(ctx *fiber.Ctx) error {
	var args dto.CreateUserDto
	if err := ctx.BodyParser(&args); err != nil {
		return errs.SendError(ctx, errs.ErrBadRequest)
	}

	props := props2.DtoToProps(args)
	resp, err := h.useCase.CreateUser(props)
	if err != nil {
		return errs.SendError(ctx, err)
	}

	return errs.SendSuccess(ctx, int(errs.Created), resp)

}
