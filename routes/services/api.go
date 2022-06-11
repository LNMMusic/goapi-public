package services

import (
	"github.com/gofiber/fiber/v2"
)


func Api(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map {
		"status":	true,
		"message":	"Welcome to the root of the API",
		"data":		nil,
	})
}