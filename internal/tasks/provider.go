package tasks

import "github.com/gofiber/fiber/v2"

func BootModule(app *fiber.App) {
	RegisterRoutes(app)
}
