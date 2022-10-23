package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/darkykek/routes"
)

func main() {

	log.Print("Initializing Environment variables...")
	if err := godotenv.Load(".env"); err != nil {
		panic("An error occurred while trying to read \".env\" file")
	}

	var mongouri string = os.Getenv("MONGO_CONNECTION_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongouri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// erasmus_database := client.Database("erasmus_database")

	// users_collection := erasmus_database.Collection("users")
	// courses_collection := erasmus_database.Collection("courses")
	// videos_collection := erasmus_database.Collection("videos")

	app := fiber.New()
	app.Use(cors.New())

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("PORT")))
}
