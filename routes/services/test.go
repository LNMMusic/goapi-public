package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)


// SECURITY MANAGMENT TEST
func Public(c *fiber.Ctx) error {
	// auth [nil]
	// ...

	return c.Status(200).JSON(&NewResponse {
		Message:	"Succeed to Access to Public Service",
		Data:		nil,
	})
}


func PrivateFree(c *fiber.Ctx) error {
	// auth [should be set as middleware in next versions]
	token := c.Locals("user").(*jwt.Token)//;		claims := token.Claims.(jwt.MapClaims)

	// response
	return c.Status(200).JSON(&NewResponse {
		Message:	"Token Information",
		Data:		token,
	})
}

func PrivatePremium(c *fiber.Ctx) error {
	// auth [should be set as middleware in next versions]
	token := c.Locals("user").(*jwt.Token)//;		claims := token.Claims.(jwt.MapClaims)
	
	// response
	return c.Status(200).JSON(&NewResponse {
		Message:	"Token Information",
		Data:		token,
	})
}

func PrivateAdmin(c *fiber.Ctx) error {
	// auth [should be set as middleware in next versions]
	token := c.Locals("user").(*jwt.Token)//;		claims := token.Claims.(jwt.MapClaims)

	// response
	return c.Status(200).JSON(&NewResponse {
		Message:	"Token Information",
		Data:		token,
	})
}