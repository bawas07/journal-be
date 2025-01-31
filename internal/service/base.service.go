package service

import (
	"mindscribe-be/internal/repository"
	"mindscribe-be/pkg/config"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type BaseService struct {
	Logger *zap.Logger
	DB     *sqlx.DB
	Config *config.Config
	Repo   *repository.Repository
}

func newBaseService(db *sqlx.DB, log *zap.Logger, cfg *config.Config, repo *repository.Repository) *BaseService {
	return &BaseService{
		Logger: log,
		Config: cfg,
		DB:     db,
		Repo:   repo,
	}
}
