package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Config struct {
	Port    string
	Logger  *zap.Logger
	BaseURL string
	Env     string
}

func Load() *Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Initialize logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	return &Config{
		Port:    ":" + getEnv("PORT", "8080"),
		Logger:  logger,
		BaseURL: getEnv("BASE_URL", "http://localhost:3000"),
		Env:     getEnv("ENV", "development"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return intValue
}
