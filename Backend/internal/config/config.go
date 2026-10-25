package config

import (
	"os"
)

type Config struct {
	Port   string
	DBHost string
	DBUser string
	DBPass string
	DBName string
	DBPort string
}

func LoadConfig() *Config {
	return &Config{
		Port:   getEnv("PORT", "8080"),
		DBHost: getEnv("DB_HOST", "postgres"),
		DBUser: getEnv("DB_USER", "postgres"),
		DBPass: getEnv("DB_PASSWORD", "secret"),
		DBName: getEnv("DB_NAME", "mydb"),
		DBPort: getEnv("DB_PORT", "5432"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
