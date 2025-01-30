package basehandler

import (
	"mindscribe-be/pkg/config"
	"mindscribe-be/pkg/response"

	"go.uber.org/zap"
)

type BaseHandler struct {
	Logger *zap.Logger
	Res    *response.Response
	Config *config.Config
}

func NewBaseHandler(log *zap.Logger, cfg *config.Config) *BaseHandler {
	return &BaseHandler{
		Logger: log,
		Res:    response.New(),
		Config: cfg,
	}
}
