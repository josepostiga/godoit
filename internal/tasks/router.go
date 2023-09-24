package tasks

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Group("/tasks").
		Get("/", index).
		Post("/", store).
		Patch("/:id", update).
		Get("/:id", show).
		Delete("/:id", delete)
}
