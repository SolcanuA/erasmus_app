package handlers

import "github.com/gofiber/fiber/v2"

func HandleAPIDefault(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"API_STATUS": "running"})
}
