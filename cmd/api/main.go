package main

import (
	"log"

	"journaling-be/pkg/config"
	"journaling-be/pkg/server"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize server
	srv := server.NewServer(cfg)

	// Start server
	log.Printf("Server started on %s", cfg.Port)
	if err := srv.Listen(cfg.Port); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
