package service

import (
	"mindscribe-be/internal/repository"
	"mindscribe-be/pkg/config"
	"time"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Service struct {
	User *UserService
	Post *PostService
}

func NewService(db *sqlx.DB, log *zap.Logger, cfg *config.Config, repo *repository.Repository) *Service {
	start := time.Now()
	log.Info("Service: Starting...")
	base := newBaseService(db, log, cfg, repo)
	user := newUserService(base)
	post := newPostService(base)
	duration := time.Since(start)
	log.Sugar().Infof("Service: Completed in %s", duration)
	return &Service{
		User: user,
		Post: post,
	}
}
