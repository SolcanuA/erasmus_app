package main

import (
	"encoding/json"
	"fmt"
	"log"

	database "github.com/darkykek/database"

	"github.com/gofiber/fiber/v2"
)

func main() {

	fmt.Println("Initializing database...")
	database.Init()

	fmt.Println("Creating server instance...")
	app := fiber.New(fiber.Config{
		AppName:       "Erasmus Server",
		CaseSensitive: true,
		BodyLimit:     100 * 1024,
		Concurrency:   1024 * 512,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	})

	//wildcard path;
	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendString("init")
	})

	log.Fatal(app.Listen(":8080"))
}
