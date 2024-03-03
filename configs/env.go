package configs

import (
	"os"

	"log"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error .env")
	}
	mongoURI := os.Getenv("MONGO_URI")
	return mongoURI
}
