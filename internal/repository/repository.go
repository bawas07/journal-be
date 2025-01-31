package repository

import (
	"mindscribe-be/pkg/config"
	"time"

	"go.uber.org/zap"
)

type Repository struct {
	UserRepo *UserRepo
}

func NewRepo(log *zap.Logger, cfg *config.Config) *Repository {
	start := time.Now()
	log.Info("Repository: Starting...")
	base := newBaseRepo(cfg)
	user := NewUserRepo(base)
	duration := time.Since(start)

	log.Sugar().Infof("Repository: Completed in %s", duration)
	return &Repository{
		UserRepo: user,
	}
}
