package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	
	"github.com/LNMMusic/goapi/config"
)

// Handlers
func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"message": "Missing or malformed JWT",
			"data": nil,
		})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"status": "error",
		"message": "Invalid or expired JWT",
		"data": nil,
	})
}

// MIDDLEWARE [Requires Auth Header on Endpoint]
func Auth() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(config.EnvGet("SECRET_KEY")),
		ErrorHandler: jwtError,
		SuccessHandler: nil,	// in case success calls c.Next() to skip to other fiber handler [other middleware or service]
	})
}