package routes

import (
	"github.com/darkykek/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/login", handlers.Login)
	auth.Post("/register", handlers.Register)
}
