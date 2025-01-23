package server

import (
	"journaling-be/internal/handler"
	"journaling-be/pkg/config"
	"journaling-be/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Server struct {
	app    *fiber.App
	logger *zap.Logger
}

func NewServer(cfg *config.Config) *Server {
	app := fiber.New()

	// Middleware
	app.Use(middleware.HttpLogger(cfg.Logger))

	// Routes
	api := app.Group("/api/v1")
	handler.RegisterRoutes(api)

	return &Server{
		app:    app,
		logger: cfg.Logger,
	}
}

func (s *Server) Listen(addr string) error {
	s.logger.Info("Starting server", zap.String("address", addr))
	return s.app.Listen(addr)
}
