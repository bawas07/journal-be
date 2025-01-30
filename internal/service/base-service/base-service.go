package baseservice

import (
	"mindscribe-be/pkg/config"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type BaseService struct {
	Logger *zap.Logger
	DB     *sqlx.DB
	Config *config.Config
}

func NewBaseService(db *sqlx.DB, log *zap.Logger, cfg *config.Config) *BaseService {
	return &BaseService{
		Logger: log,
		Config: cfg,
		DB:     db,
	}
}
