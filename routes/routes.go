package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	SetupAPIRoutes(app)
	SetupAuthRoutes(app)
	SetupUserRoutes(app)
	SetupPostRoutes(app)
}
