package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := os.Getenv("ENV")

	if env != "development" && env != "dev" {
		err = godotenv.Overload(".env.prod")
	}

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Get(key string) string {
	return os.Getenv(key)
}
