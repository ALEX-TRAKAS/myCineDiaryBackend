package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		log.Printf("DATABASE_URL: %q\n", os.Getenv("DATABASE_URL"))
		return strings.TrimSpace(value)
	}

	return defaultValue
}
