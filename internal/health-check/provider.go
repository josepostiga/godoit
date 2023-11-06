package health_check

import "github.com/gofiber/fiber/v2"

func BootModule(app *fiber.App) {
	RegisterRoutes(app)
}
