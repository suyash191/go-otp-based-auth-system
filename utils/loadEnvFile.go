package utils

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnvVariables loads environment variables from .env file
func LoadEnvVariables() {
	// read .env file in the root directory
	err := godotenv.Load("config/database.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
