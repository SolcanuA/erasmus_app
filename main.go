package main

import (
	"log"

	"github.com/darkykek/config"

	"github.com/darkykek/db"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/darkykek/routes"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		panic("An error occurred while trying to read \".env\" file")
	}

	db.Connect()

	app := fiber.New()
	app.Use(cors.New())

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(config.GetKey("PORT")))
}
