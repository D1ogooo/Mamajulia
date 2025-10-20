package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env não encontrado, usando variáveis do sistema")
	}
}

func Get(key string) string {
	return os.Getenv(key)
}
