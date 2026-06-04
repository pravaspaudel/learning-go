package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfigs struct {
	PORT              string
	DB_CONNECTION_URL string
	JWT_SECRET        string
}

var AppConfig EnvConfigs

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("failed to load .env file")
	}

	AppConfig.PORT = getEnv("PORT", "8080")
	AppConfig.DB_CONNECTION_URL = getSureEnv("DB_CONNECTION_URL")
	AppConfig.JWT_SECRET = getSureEnv("JWT_SECRET")
}

func getEnv(key string, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

func getSureEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("%s is missing in .env", key)
	}
	return val
}
