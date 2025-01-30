package service

import (
	baseservice "mindscribe-be/internal/service/base-service"
	userservice "mindscribe-be/internal/service/user-service"
	"mindscribe-be/pkg/config"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Service struct {
	User *userservice.UserService
}

func NewHandler(db *sqlx.DB, log *zap.Logger, cfg *config.Config) *Service {
	log.Info("Service: Starting...")
	base := baseservice.NewBaseService(db, log, cfg)
	user := userservice.NewUserService(base)
	log.Info("Service: Completed")
	return &Service{
		User: user,
	}
}
