package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:       "Erasmus Server",
		CaseSensitive: true,
		BodyLimit:     100 * 1024,
		Concurrency:   1024 * 512,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	})

	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendString("wildcard path")
	})

	app.Listen(":8080")
}
