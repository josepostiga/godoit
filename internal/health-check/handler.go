package health_check

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func show(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).
		JSON(&fiber.Map{
			"status": "OK",
		})
}
