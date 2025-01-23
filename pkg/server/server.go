package server

import (
	"journaling-be/internal/handler"
	"journaling-be/pkg/config"
	"journaling-be/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Server struct {
	app    *fiber.App
	logger *zap.Logger
	db     *sqlx.DB
}

func NewServer(cfg *config.Config, db *sqlx.DB) *Server {
	app := fiber.New()

	// Middleware
	app.Use(middleware.HttpLogger(cfg.Logger))

	// Routes
	api := app.Group("/api/v1")
	handler.RegisterRoutes(api, db)

	return &Server{
		app:    app,
		logger: cfg.Logger,
		db:     db,
	}
}

func (s *Server) Listen(addr string) error {
	s.logger.Info("Starting server", zap.String("address", addr))
	return s.app.Listen(addr)
}
