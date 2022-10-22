package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() {

	//running locally; no .env required yet
	mongouri := "mongodb://localhost:27017" //uri subject to change in case of migration

	client, err := mongo.NewClient(options.Client().ApplyURI(mongouri))

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

}
