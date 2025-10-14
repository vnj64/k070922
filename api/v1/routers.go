package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type Handlers struct {
	UserHandler *UserHandler
}

func SetupRoutes(app *fiber.App, h *Handlers) {
	api := app.Group("/api/v1")

	docs := api.Group("/docs")
	docs.Get("/swagger/*", swagger.HandlerDefault)
}
