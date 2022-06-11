package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/LNMMusic/goapi/routes/services"
	"github.com/LNMMusic/goapi/routes/middleware"
)

// Routes Groups
func SetUp(app *fiber.App) {
	// Root
	app.Get("/", services.Root)

	// Api
	api := app.Group("/api")
	api.Get("/", services.Api)

	// Models [user]
	models := app.Group("/models")
	models.Get("/user", services.GetUsers)
	models.Get("/user/:id", services.GetUser)
	models.Post("/user/create", services.CreateUser)
	models.Put("/user/update/:id", services.UpdateUser)
	models.Delete("/user/delete/:id", services.DeleteUser)

	// Auth
	auth := app.Group("/auth")
	auth.Post("/sign-up", services.SignUp)
	auth.Post("/sign-in", services.SignIn)
	auth.Post("/sign-out",
				middleware.Auth(), middleware.AuthRedis(),
				services.SignOut)

	auth.Post("/token",
				middleware.Auth(), middleware.AuthRedis(),
				services.Token)
	auth.Post("/sessions", services.Sessions)

	// Test
	test := app.Group("/test")
	test.Get("/public", services.Public)
	test.Post("/private-free",
				middleware.Auth(), middleware.AuthRedis(),
				services.PrivateFree)
	test.Post("/private-premium",
				middleware.Auth(), middleware.AuthRedis(), middleware.PermissionPremium(),
				services.PrivatePremium)
	test.Post("/private-admin",
				middleware.Auth(), middleware.AuthRedis(), middleware.PermissionAdmin(),
				services.PrivateAdmin)
}