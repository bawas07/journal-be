package handler

import (
	basehandler "mindscribe-be/internal/handler/base-handler"
	indexhandler "mindscribe-be/internal/handler/index-handler"
	userhandler "mindscribe-be/internal/handler/user-handler"
	"mindscribe-be/pkg/config"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Handler struct {
	Index *indexhandler.IndexHandler
	User  *userhandler.UserHandler
}

func NewHandler(db *sqlx.DB, log *zap.Logger, cfg *config.Config) *Handler {
	log.Info("Handler: Starting...")
	base := basehandler.NewBaseHandler(log, cfg)
	log.Info("Handler: Setting Index Handler")
	indexHandler := indexhandler.NewIndexHandler(db, base)
	log.Info("Handler: Setting User Handler")
	userHandler := userhandler.NewUserHandler(db, base)
	log.Info("Handler: Completed")
	return &Handler{
		Index: indexHandler,
		User:  userHandler,
	}
}
