package routes

import (
	"github.com/darkykek/db/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupPostRoutes(app *fiber.App) {
	course := app.Group("/users")
	course.Get("/", handlers.GetCourses)
	course.Get("/:post_id", handlers.GetCourse)
	course.Patch("/:post_id", handlers.UpdateCourse)
	course.Delete("/delete", handlers.DeleteCourse)
}
