package health_check

import (
	"github.com/gofiber/fiber/v2"
	responses "github.com/josepostiga/godoit/internal"
	"net/http"
)

func show(c *fiber.Ctx) error {
	return responses.New(c, nil, http.StatusOK)
}
