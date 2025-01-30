package basehandler

import (
	"mindscribe-be/internal/service"
	"mindscribe-be/pkg/config"
	"mindscribe-be/pkg/response"

	"go.uber.org/zap"
)

type BaseHandler struct {
	Logger  *zap.Logger
	Res     *response.Response
	Config  *config.Config
	Service *service.Service
}

func NewBaseHandler(log *zap.Logger, cfg *config.Config, service *service.Service) *BaseHandler {
	return &BaseHandler{
		Logger:  log,
		Res:     response.New(),
		Config:  cfg,
		Service: service,
	}
}
