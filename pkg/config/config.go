package config

import "go.uber.org/zap"

type Config struct {
	ServerAddress string
	Logger        *zap.Logger
}

func Load() *Config {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	return &Config{
		ServerAddress: ":8080",
		Logger:        logger,
	}
}