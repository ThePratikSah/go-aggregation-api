package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetCollectionName() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	collection := os.Getenv("COLLECTION")

	return collection
}
