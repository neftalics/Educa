package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar el archivo .env, usando variables del sistema si existen")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
