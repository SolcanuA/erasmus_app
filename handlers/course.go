package handlers

import "github.com/gofiber/fiber/v2"

func GetCourses(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"Get": "Course",
	})
}

func GetCourse(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"Get": "Course",
	})
}

func UpdateCourse(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"Update": "Course",
	})
}

func DeleteCourse(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"Delete": "Course",
	})
}
