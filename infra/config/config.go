package config

import (
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	JWTSecret   string
	Port        string
}

func LoadConfig() *Config {
	godotenv.Load()

	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", "root:password@tcp(localhost:3306)/news_db?charset=utf8mb4&parseTime=True&loc=Local"),
		JWTSecret:   getEnv("JWT_SECRET", "your-secret-key"),
		Port:        getEnv("PORT", ":8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}