package tasks

import (
	"github.com/gofiber/fiber/v2"
	"github.com/josepostiga/godoit/internal/authentication/middleware"
)

func RegisterRoutes(app *fiber.App) {
	app.Group("/tasks", middleware.Authenticated).
		Get("/", index).
		Post("/", store).
		Patch("/:id", update).
		Get("/:id", show).
		Delete("/:id", delete).
		Post("/:id/status", status)
}
