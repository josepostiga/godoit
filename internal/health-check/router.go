package health_check

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Group("/health-check").
		Get("/", show)
}
