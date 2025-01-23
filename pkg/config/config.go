package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type Config struct {
	Port    string
	Logger  *zap.Logger
	BaseURL string
	Env     string
	DB      DBConfig
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

	db := DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", ""),
		Name:     getEnv("DB_NAME", "journaling"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	return &Config{
		Port:    ":" + getEnv("PORT", "8080"),
		Logger:  logger,
		BaseURL: getEnv("BASE_URL", "http://localhost:3000"),
		Env:     getEnv("ENV", "development"),
		DB:      db,
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
