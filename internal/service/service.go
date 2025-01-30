package service

import (
	"mindscribe-be/internal/repository"
	baseservice "mindscribe-be/internal/service/base-service"
	userservice "mindscribe-be/internal/service/user-service"
	"mindscribe-be/pkg/config"
	"time"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Service struct {
	User *userservice.UserService
}

func NewService(db *sqlx.DB, log *zap.Logger, cfg *config.Config, repo *repository.Repository) *Service {
	start := time.Now()
	log.Info("Service: Starting...")
	base := baseservice.NewBaseService(db, log, cfg, repo)
	user := userservice.NewUserService(base)
	duration := time.Since(start)
	log.Info("Service: Completed", zap.Duration("duration", duration))
	return &Service{
		User: user,
	}
}
