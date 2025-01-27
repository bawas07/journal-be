package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"

	"mindscribe-be/pkg/config"
	"mindscribe-be/pkg/logger"
	"mindscribe-be/pkg/server"
)

func main() {
	// Load configuration
	cfg := config.Load()
	logger.Init(cfg)
	log := logger.Logger()
	log.Sugar().Info("Standing By...")

	// Initialize database connection
	db, err := sqlx.Connect("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.SSLMode,
	))
	if err != nil {
		log.Fatal("Failed to connect to database", zap.Error(err))
	}
	defer db.Close()

	// Initialize server with database connection
	srv := server.NewServer(cfg, db)

	log.Sugar().Infof("Server started on %s", cfg.Port)
	if err := srv.Listen(cfg.Port); err != nil {
		log.Fatal("Server error", zap.Error(err))
	}
}
