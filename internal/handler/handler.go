package handler

import (
	userhandlers "mindscribe-be/internal/handler/user-handlers"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Handler struct {
	Index *IndexHandler
	User  *userhandlers.UserHandler
}

func NewHandler(db *sqlx.DB, log *zap.Logger) *Handler {
	log.Info("Handler: Starting...")
	log.Info("Handler: Setting Index Handler")
	indexHandler := NewIndexHandler(db, log)
	log.Info("Handler: Setting User Handler")
	userHandler := userhandlers.NewUserHandler(db, log)
	log.Info("Handler: Completed")
	return &Handler{
		Index: indexHandler,
		User:  userHandler,
	}
}
