package server

import (
	"mindscribe-be/internal/handler"
	"mindscribe-be/internal/repository"
	"mindscribe-be/internal/service"
	"mindscribe-be/pkg/config"
	"mindscribe-be/pkg/logger"
	"mindscribe-be/pkg/middleware"
	"mindscribe-be/pkg/route"
	"mindscribe-be/pkg/validation"

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
	log := logger.Logger()
	logger := middleware.HttpLogger(log)
	// Middleware
	app.Use(logger)

	// Handlers
	log.Info("====================================")
	r := repository.NewRepo(log, cfg)
	log.Info("====================================")
	s := service.NewService(db, log, cfg, r)
	log.Info("====================================")
	v := validation.NewValidate(db, log)
	h := handler.NewHandler(log, cfg, s, v)
	log.Info("====================================")
	// Routes
	route.RegisterRoutes(app, h, log)

	return &Server{
		app:    app,
		logger: log,
		db:     db,
	}
}

func (s *Server) Listen(addr string) error {
	s.logger.Info("====================================")
	s.logger.Info("Complete!")
	s.logger.Info("====================================")
	return s.app.Listen(addr)
}
