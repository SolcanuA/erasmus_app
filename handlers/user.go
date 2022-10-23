package handlers

import "github.com/gofiber/fiber/v2"

func GetUsers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"Get": "Users",
	})
}

func GetUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"Get": "User",
	})
}

func UpdateUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"Update": "User",
	})
}

func DeleteUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"Delete": "User",
	})
}
