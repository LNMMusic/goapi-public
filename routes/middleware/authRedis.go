package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	
	"github.com/LNMMusic/goapi/handlers/auth"
)

// CONFIG [struct]
type ConfigAuthRedis struct {
	Filter			func(token_id string) bool
	ErrorHandler	func(c *fiber.Ctx) error
}

// MIDDLEWARE
func NewAuthRedis(config ConfigAuthRedis) fiber.Handler {
	// main cfg
	// var cfg = config

	// middleware
	return func(c *fiber.Ctx) error {
		token := c.Locals("user").(*jwt.Token); claims := token.Claims.(jwt.MapClaims)
		tokenID:= claims["id"].(string)
		userID := claims["user_id"].(string)

		// redis check
		if err := auth.JWTRedisHandlerValidate(userID, tokenID); err != nil {return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map {
			"Message":	"Failed to Validate Token [It was dropped manually in Redis]",
			"Data":		err,
		})}

		// skip
		return c.Next()
	}
}

// CUSTOM
func AuthRedis() fiber.Handler {
	var config = ConfigAuthRedis {}
	return NewAuthRedis(config)
}