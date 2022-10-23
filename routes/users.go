package routes

import (
	"github.com/darkykek/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	users := app.Group("/users")
	users.Get("/", handlers.GetUsers)
	users.Get("/:user_id", handlers.GetUser)
	users.Patch("/update/:user_id", handlers.UpdateUser)
	users.Delete("/delete/:user_id", handlers.DeleteUser)
}
