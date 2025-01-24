package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"journaling-be/pkg/config"
)

func main() {
	cmd := flag.String("cmd", "", "Migration command (up/down/create)")
	steps := flag.Int("steps", 0, "Number of migration steps (optional)")
	name := flag.String("name", "", "Migration name (required for create)")
	flag.Parse()

	if *cmd == "" {
		fmt.Println("Usage: migrate -cmd [up|down|create] [-steps N] [-name migration_name]")
		os.Exit(1)
	}

	if *cmd == "create" {
		if *name == "" {
			fmt.Println("Migration name is required for create command")
			os.Exit(1)
		}
		if err := createMigration(*name); err != nil {
			log.Fatalf("Failed to create migration: %v", err)
		}
		return
	}

	cfg := config.Load()

	// Get absolute path to migrations directory
	migrationsPath, err := filepath.Abs("migrations")
	if err != nil {
		log.Fatalf("Failed to get migrations path: %v", err)
	}

	m, err := migrate.New(
		fmt.Sprintf("file://%s", migrationsPath),
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name, cfg.DB.SSLMode),
	)
	if err != nil {
		log.Fatalf("Failed to initialize migrations: %v", err)
	}
	defer m.Close()

	var migrationErr error
	switch *cmd {
	case "up":
		if *steps > 0 {
			migrationErr = m.Steps(*steps)
		} else {
			migrationErr = m.Up()
		}
	case "down":
		if *steps > 0 {
			migrationErr = m.Steps(-*steps)
		} else {
			migrationErr = m.Down()
		}
	default:
		log.Fatalf("Invalid command: %s", *cmd)
	}

	if migrationErr != nil && migrationErr != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", migrationErr)
	}

	log.Println("Migration completed successfully")
}

func createMigration(name string) error {
	timestamp := time.Now().Format("20060102150405")
	baseName := fmt.Sprintf("migrations/%s_%s", timestamp, name)

	files := []string{
		fmt.Sprintf("%s.up.sql", baseName),
		fmt.Sprintf("%s.down.sql", baseName),
	}

	for _, file := range files {
		f, err := os.Create(file)
		if err != nil {
			return fmt.Errorf("failed to create migration file: %w", err)
		}
		defer f.Close()

		// Add basic SQL comment to each file
		_, err = f.WriteString(fmt.Sprintf("-- Migration: %s\n", name))
		if err != nil {
			return fmt.Errorf("failed to write to migration file: %w", err)
		}
	}

	log.Printf("Created migration files: %s.up.sql, %s.down.sql", baseName, baseName)
	return nil
}
