package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvMongoUri() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGO_URI")
}

func getEnvVariable(name string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(name)
}
