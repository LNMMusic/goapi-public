package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"github.com/LNMMusic/goapi/models"
	"github.com/LNMMusic/goapi/handlers/db"
	"github.com/LNMMusic/goapi/handlers/auth"
)

// TOKEN MANAGEMENT
func SignUp(c *fiber.Ctx) error {
	// request
	req := &models.User{}
	if err := c.BodyParser(req); err != nil {return c.Status(422).JSON(&NewResponse {
		Message:	"Failed to parse request",
		Data:		err.Error(),
	})}

	// process
	if err := req.HashPassword(); err != nil {return c.Status(422).JSON(&NewResponse {
		Message:	"Failed to Hash Password!",
		Data:		err.Error(),
	})}
	if err := db.Psql.Db.Create(&req).Error; err != nil {return c.Status(422).JSON(&NewResponse {
		Message:	"Failed to Sign Up!",
		Data:		err,
	})}

	// response
	return c.Status(200).JSON(&NewResponse {
		Message:	"Succeed to Sign Up!",
		Data:		req.Response(),
	})
}

func SignIn(c *fiber.Ctx) error {
	// request
	req := &models.User{}
	if err := c.BodyParser(req); err != nil {return c.Status(422).JSON(&NewResponse {
		Message:	"Failed to parse request",
		Data:		err.Error(),
	})}

	// process [validation + token]
	var user = &models.User{}
	if err := db.Psql.Db.Where(models.User{Username: req.Username}).Find(user).Error; err != nil {return c.Status(fiber.StatusUnauthorized).JSON(&NewResponse {
		Message:	"Failed to Sign In! Invalid Username",
		Data:		err,
	})}
	if len(user.Username) == 0 {return c.Status(fiber.StatusUnauthorized).JSON(&NewResponse {
		Message:	"Failed to Sign In! Invalid Username",
		Data:		nil,
	})}
	if !(user.ValidPassword(req.Password)) {return c.Status(fiber.StatusUnauthorized).JSON(&NewResponse {
		Message:	"Failed to Sign In! Invalid Password",
		Data:		nil,
	})}

	token, err := auth.CreateJWTToken(user.Id, user.IsPremium, user.IsAdmin)
	if err != nil {return c.Status(500).JSON(&NewResponse {
		Message:	"Failed to Create JWT Token!",
		Data:		err.Error(),
	})}

	// response
	return c.Status(200).JSON(&NewResponse {
		Message:	"Succeed to Sign In!",
		Data:		token,
	})
}

func SignOut(c *fiber.Ctx) error {
	// auth
	token := c.Locals("user").(*jwt.Token); claims := token.Claims.(jwt.MapClaims)
	tokenID:= claims["id"].(string)
	userID := claims["user_id"].(string)

	// request

	// process
	if err := auth.JWTRedisHandlerDropSession(userID, tokenID); err != nil {return c.Status(500).JSON(&NewResponse {
		Message:	"Failed to Sign Out! [couldn't drop token in redis]",
		Data:		err,
	})}

	// response
	return c.Status(200).JSON(&NewResponse {
		Message:	"Succeed to Sign Out!",
		Data:		nil,
	})
}


// TOKEN METADATA
func Token(c *fiber.Ctx) error {
	// auth info
	token := c.Locals("user").(*jwt.Token)//; claims := token.Claims.(jwt.MapClaims)

	// request

	// response
	return c.Status(200).JSON(&NewResponse {
		Message:	"Token Metadata",
		Data:		token,
	})
}


// SESSIONS [By User]
func Sessions(c *fiber.Ctx) error {
	// request [auth manually by username and password]
	req := &models.User{}
	if err := c.BodyParser(req); err !=  nil {return c.Status(422).JSON(&NewResponse {
		Message:	"Failed to parse request",
		Data:		err.Error(),
	})}

	// process [validation + no token + sessions info]
	var user = &models.User{}
	if err := db.Psql.Db.Where(models.User{Username: req.Username}).Find(user).Error; err != nil {return c.Status(fiber.StatusUnauthorized).JSON(&NewResponse {
		Message:	"Failed to Sign In! Invalid Username",
		Data:		err,
	})}
	if len(user.Username) == 0 {return c.Status(fiber.StatusUnauthorized).JSON(&NewResponse {
		Message:	"Failed to Sign In! Invalid Username",
		Data:		nil,
	})}
	if !(user.ValidPassword(req.Password)) {return c.Status(fiber.StatusUnauthorized).JSON(&NewResponse {
		Message:	"Failed to Sign In! Invalid Password",
		Data:		nil,
	})}

	sessions, err := auth.JWTRedisHandlerGetSessions(user.Id.String()); if err != nil {return c.Status(422).JSON(&NewResponse {
		Message:	"Failed to Get Sessions!",
		Data:		err.Error(),
	})}

	// response
	return c.Status(200).JSON(&NewResponse {
		Message:	"Sessions Metadata",
		Data:		sessions,
	})
}