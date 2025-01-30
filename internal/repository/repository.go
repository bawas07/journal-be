package repository

import (
	baserepository "mindscribe-be/internal/repository/base-repository"
	userrepository "mindscribe-be/internal/repository/user-repository"
	"mindscribe-be/pkg/config"
	"time"

	"go.uber.org/zap"
)

type Repository struct {
	UserRepo *userrepository.UserRepo
}

func NewRepo(log *zap.Logger, cfg *config.Config) *Repository {
	start := time.Now()
	log.Info("Repository: Starting...")
	base := baserepository.New(cfg)
	user := userrepository.NewUserRepo(base)
	duration := time.Since(start)

	log.Info("Repository: Completed", zap.Duration("duration", duration))
	return &Repository{
		UserRepo: user,
	}
}
