package handler

import (
	"mindscribe-be/internal/service"
	"mindscribe-be/pkg/config"
	"mindscribe-be/pkg/validation"
	"time"

	"go.uber.org/zap"
)

type Handler struct {
	Index *IndexHandler
	User  *UserHandler
}

func NewHandler(log *zap.Logger, cfg *config.Config, service *service.Service, validate *validation.Validate) *Handler {
	start := time.Now()

	log.Info("Handler: Starting...")
	base := newBaseHandler(log, cfg, service, validate)
	log.Info("Handler: Setting Index Handler")
	indexHandler := newIndexHandler(base)
	log.Info("Handler: Setting User Handler")
	userHandler := newUserHandler(base)
	duration := time.Since(start)
	log.Sugar().Infof("Handler: Completed in %s", duration)
	return &Handler{
		Index: indexHandler,
		User:  userHandler,
	}
}
