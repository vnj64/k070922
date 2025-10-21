package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type Handlers struct {
	UserHandler *UserHandler
	AuthHandler *AuthHandler
}

func SetupRoutes(app *fiber.App, h *Handlers) {
	api := app.Group("/api/v1")

	docs := app.Group("/docs")
	docs.Get("/swagger/*", swagger.HandlerDefault)

	RegisterAuthRoutes(api.Group("/auth"), h.AuthHandler)
}
