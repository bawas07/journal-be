package handler

import (
	basehandler "mindscribe-be/internal/handler/base-handler"
	indexhandler "mindscribe-be/internal/handler/index-handler"
	userhandler "mindscribe-be/internal/handler/user-handler"
	"mindscribe-be/internal/service"
	"mindscribe-be/pkg/config"
	"time"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Handler struct {
	Index *indexhandler.IndexHandler
	User  *userhandler.UserHandler
}

func NewHandler(db *sqlx.DB, log *zap.Logger, cfg *config.Config, service *service.Service) *Handler {
	start := time.Now()

	log.Info("Handler: Starting...")
	base := basehandler.NewBaseHandler(log, cfg, service)
	log.Info("Handler: Setting Index Handler")
	indexHandler := indexhandler.NewIndexHandler(db, base)
	log.Info("Handler: Setting User Handler")
	userHandler := userhandler.NewUserHandler(db, base)
	duration := time.Since(start)
	log.Info("Handler: Completed", zap.Duration("duration", duration))
	return &Handler{
		Index: indexHandler,
		User:  userHandler,
	}
}
