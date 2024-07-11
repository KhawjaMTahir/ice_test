package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func LoadConfig() Config {
	rootDir := filepath.Join("..", "..")
	err := os.Chdir(rootDir)
	if err != nil {
		log.Fatalf("Error changing working directory: %v", err)
	}

	// Load .env file
	err = godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	return Config{
		DBUsername: getEnv("DB_USERNAME", "root"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBName:     getEnv("DB_NAME", "ice"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
