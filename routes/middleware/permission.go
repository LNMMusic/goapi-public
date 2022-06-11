package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// CONFIG [struct]
type ConfigPermission struct {
	AuthPremium			bool
	AuthAdmin			bool
}

// MIDDLEWARE [fiber.Handler = func(c *fiber.Ctx) => error]
func NewPermission(config ConfigPermission) fiber.Handler {
	// main config [w/ default config parsed with custom member's values]
	var cfg = config
	
	// middleware [endpoint] [to skip c.Next()]
	return func(c *fiber.Ctx) error {
		token := c.Locals("user").(*jwt.Token); claims := token.Claims.(jwt.MapClaims)
		isPremium := claims["user_isPremium"].(bool)
		isAdmin	  := claims["user_isAdmin"].(bool)

		// not skip
		if cfg.AuthPremium && !isPremium {return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map {
			"Message":	"Unauthorized Token! User has not Premium Permissions",
			"Data":		nil,
		})}
		if cfg.AuthAdmin && !isAdmin {return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map {
			"Message":	"Unauthorized Token! User has not Admin Permissions",
			"Data":		nil,
		})}

		// skip
		return c.Next()
	}
}

// CUSTOM
func PermissionPremium() fiber.Handler {
	var config = ConfigPermission {
		AuthPremium:	true,
		AuthAdmin:		true,
	}
	return NewPermission(config)
}
func PermissionAdmin() fiber.Handler {
	var config = ConfigPermission {
		AuthPremium:	false,
		AuthAdmin:		true,
	}
	return NewPermission(config)
}