package internal

import "github.com/gofiber/fiber/v2"

func NewJSONResponse(c *fiber.Ctx, data interface{}, status int) error {
	c.Status(status)

	if data == nil {
		return nil
	}

	return c.JSON(&fiber.Map{
		"data": data,
	})
}
