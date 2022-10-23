package routes

import (
	"github.com/darkykek/handlers"
	"github.com/darkykek/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupPostRoutes(app *fiber.App) {
	course := app.Group("/courses")
	course.Get("/", handlers.GetCourses)
	course.Get("/:course_id", middleware.TokenProtection(), handlers.GetCourse)
	course.Patch("/:course_id", middleware.TokenProtection(), handlers.UpdateCourse)
	course.Delete("/delete", middleware.TokenProtection(), handlers.DeleteCourse)
}
