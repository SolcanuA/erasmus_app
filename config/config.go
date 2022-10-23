package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetKey(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env")
	}

	return os.Getenv(key)
}
