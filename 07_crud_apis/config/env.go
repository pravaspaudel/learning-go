package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	PORT   string
	DB_URL string
}

var App *EnvConfig

func ConfigureEnv() {

	err := godotenv.Load()

	if err != nil {
		log.Println("no .env file")
	}

	App = &EnvConfig{
		PORT:   getEnvWithOption("PORT", "8080"),
		DB_URL: getEnv("DB_URL"),
	}
}

func getEnv(key string) string {
	val := os.Getenv(key)

	if val == "" {
		log.Fatalf("environment variable %s is not provided", key)
	}

	return val
}

func getEnvWithOption(key, secondary string) string {
	val := os.Getenv(key)
	if val == "" {
		return secondary
	}
	return val
}
