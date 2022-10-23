package db

import (
	"context"
	"log"
	"time"

	"github.com/darkykek/config"
	"github.com/darkykek/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MInstance struct {
	Client            *mongo.Client
	Database          *mongo.Database
	UsersCollection   *mongo.Collection
	CoursesCollection *mongo.Collection
	VideosCollection  *mongo.Collection
}

var Mg MInstance

func Connect() error {

	client, err := mongo.NewClient(options.Client().ApplyURI(config.GetKey("DB_URI")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)

	db := client.Database(config.GetKey("DATABASE_NAME"))

	users_collection := db.Collection("users")
	courses_collection := db.Collection("courses")
	videos_collection := db.Collection("videos")

	// dummy data
	user_create_res, err := users_collection.InsertOne(ctx, models.User{
		FirstName: "Kek",
		LastName:  "Lol",
		Username:  "darky.exe",
		Email:     "kek@lol.com",
		Password:  "123kek123lol",
	})

	if err != nil {
		log.Fatal("Failed to create user")
	}

	log.Printf("Address in memory -> %p", &user_create_res)

	Mg = MInstance{
		Client:            client,
		Database:          db,
		UsersCollection:   users_collection,
		CoursesCollection: courses_collection,
		VideosCollection:  videos_collection,
	}

	return nil
}
