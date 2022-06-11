package services

import (
	"github.com/gofiber/fiber/v2"
)


func Root(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map {
		"status":	true,
		"message":	"Here go all possible endpoints",
		"data":		nil,
	})
}