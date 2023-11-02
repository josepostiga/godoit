package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	responses "github.com/josepostiga/godoit/internal"
	"os"
)

func Authenticated(c *fiber.Ctx) error {
	if len(c.Get("Authorization")) < 7 {
		return responses.New(c, &fiber.Map{"error": "Invalid authorization token."}, fiber.StatusUnauthorized)
	}

	_, err := jwt.Parse(c.Get("Authorization")[7:], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return responses.New(c, &fiber.Map{"error": "Unauthorized"}, fiber.StatusUnauthorized)
	}

	return c.Next()
}
