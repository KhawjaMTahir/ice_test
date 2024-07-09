package config

import (
	"os"
)

type Config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func LoadConfig() Config {
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
