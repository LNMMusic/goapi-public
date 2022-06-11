package services

import (
	"github.com/gofiber/fiber/v2"

	"github.com/LNMMusic/goapi/handlers/db"
	"github.com/LNMMusic/goapi/models"
)

// READ [GET]
func GetUsers(c *fiber.Ctx) error {
	// request

	// process
	users := []models.User{}
	db.Psql.Db.Find(&users)

	// response
	return c.Status(200).JSON(&users)
}
func GetUser(c *fiber.Ctx) error {
	// Request
	id := c.Params("id")

	// process
	user := &models.User{}
	db.Psql.Db.First(user, "id = ?", id)

	// Response
	return c.Status(200).JSON(user.Response())
}


// WRITE [POST]
func CreateUser(c *fiber.Ctx) error {
	// request
	req := &models.User{}
	if err := c.BodyParser(req); err != nil {return c.Status(422).JSON(&NewResponse {
		Message:	"Failed to parse request",
		Data:		err.Error(),
	})}

	// process
	if err := req.HashPassword(); err != nil {return c.Status(422).JSON(&NewResponse {
		Message:	"Failed to Create User! Hash Password Failed",
		Data:		err.Error(),
	})}
	if err := db.Psql.Db.Create(&req).Error; err != nil {return c.Status(422).JSON(&NewResponse {
		Message:	"Failed to Create User!",
		Data:		err,
	})}

	// response
	return c.Status(200).JSON(&NewResponse {
		Message:	"Succeed to Create User!",
		Data:		req.Response(),
	})
}

func UpdateUser(c *fiber.Ctx) error {
	// request
	req := &models.User{};		id := c.Params("id")
	if err := c.BodyParser(req); err != nil {return c.Status(422).JSON(&NewResponse {
		Message:	"Failed to parse request",
		Data:		err.Error(),
	})}

	// process
	user := &models.User{}
	db.Psql.Db.First(user, "id = ?", id)
	db.Psql.Db.Model(user).Updates(req)

	// response
	return c.Status(200).JSON(&NewResponse {
		Message:	"Succeed to Update User!",
		Data:		nil,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	// request
	id := c.Params("id")

	// process
	user := &models.User{}
	db.Psql.Db.First(user, "id = ?", id)
	db.Psql.Db.Delete(user)

	// response
	return c.Status(200).JSON(&NewResponse {
		Message:	"Succeed to Delete User!",
		Data:		nil,
	})
}