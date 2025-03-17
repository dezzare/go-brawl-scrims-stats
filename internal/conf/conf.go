package conf

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
