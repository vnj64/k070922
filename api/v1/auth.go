package v1

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"project/domain/cases/auth"
	"project/domain/props"
	"project/pkg/errs"
)

type AuthHandler struct {
	useCase *auth.AuthUseCase
}

func NewAuthHandler(useCase *auth.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		useCase: useCase,
	}
}

func RegisterAuthRoutes(router fiber.Router, h *AuthHandler) {
	log.Println("routes init")
	router.Post("/login", h.LoginHandler)
}

// LoginHandler godoc
// @Summary Хэндлер для авторизации пользователя
// @Description Хэндлер для авторизации пользователя
// @Tags User
// @Accept json
// @Produce json
// @Param request body props.LoginUserReq true "Данные для авторизации пользователя"
// @Success 200 {object} props.LoginUserResp
// @Failure 400 {object} errs.Error "Invalid request"
// @Failure 500 {object} errs.Error "Internal server error"
// @Router /api/v1/auth/login [post]
func (h *AuthHandler) LoginHandler(ctx *fiber.Ctx) error {
	log.Println("login handler")
	var args props.LoginUserReq
	if err := ctx.BodyParser(&args); err != nil {
		return errs.SendError(ctx, errs.ErrBadRequest)
	}

	resp, err := h.useCase.Login(args)
	if err != nil {
		return errs.SendError(ctx, err)
	}

	return errs.SendSuccess(ctx, 200, resp)

}
