package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVarialbe() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
