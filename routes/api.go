package routes

import (
	"github.com/darkykek/db/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupAPIRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", handlers.HandleAPIDefault)
}
