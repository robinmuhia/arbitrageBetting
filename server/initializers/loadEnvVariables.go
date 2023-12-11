package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// Loads environment variables to be used by the application.
func LoadEnvironmentVariables() {
	err := godotenv.Load()

	if err != nil{
		log.Fatal("Error loading .env file")
	}
}