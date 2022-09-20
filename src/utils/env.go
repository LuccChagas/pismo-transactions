package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("/Users/luccas/go/projects/pismo-transactions/.env") //Workdir
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}
