package config

import (
	"log"

	"github.com/joho/godotenv"
)

type config struct {
	Port        string
	JWTScret    string
	DatabaseURL string
}

func LoadConfig() *config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("no se encontro .env")
	}
	return nil
}
