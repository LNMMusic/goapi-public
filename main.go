package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	
	"github.com/LNMMusic/goapi/config"
	"github.com/LNMMusic/goapi/handlers/db"
	"github.com/LNMMusic/goapi/models"
	"github.com/LNMMusic/goapi/routes"
)

func main() {
	// Env
	config.EnvLoad()

	// DB
	db.Psql.ConnectClient()
	db.Psql.Db.AutoMigrate(
		&models.User{},
	)
	db.Redis.ConnectClient()

	// App
	app := fiber.New()
	app.Use(
		// middlewares [can be applied to specific routes groups]
		cors.New(),
		logger.New(),
	)

	routes.SetUp(app)

	log.Fatal(app.Listen("127.0.0.1:8080"))
}