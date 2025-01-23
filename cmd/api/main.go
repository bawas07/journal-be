package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"journaling-be/pkg/config"
	"journaling-be/pkg/server"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database connection
	db, err := sqlx.Connect("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.SSLMode,
	))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize server with database connection
	srv := server.NewServer(cfg, db)

	// Start server
	log.Printf("Server started on %s", cfg.Port)
	if err := srv.Listen(cfg.Port); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
