package health_check

import (
	"github.com/gofiber/fiber/v2"
	"github.com/josepostiga/godoit/internal"
	"net/http"
)

func show(c *fiber.Ctx) error {
	return internal.NewJSONResponse(c, nil, http.StatusOK)
}
