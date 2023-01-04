package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Environment(keyword string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file")
	}

	return os.Getenv(keyword)
}
