package handler

import (
	"mindscribe-be/internal/service"
	"mindscribe-be/pkg/config"
	"mindscribe-be/pkg/response"
	"mindscribe-be/pkg/validation"

	"go.uber.org/zap"
)

type BaseHandler struct {
	Logger   *zap.Logger
	Res      *response.Response
	Config   *config.Config
	Service  *service.Service
	Validate *validation.Validate
}

func newBaseHandler(log *zap.Logger, cfg *config.Config, service *service.Service, validate *validation.Validate) *BaseHandler {
	return &BaseHandler{
		Logger:   log,
		Res:      response.New(validate),
		Config:   cfg,
		Service:  service,
		Validate: validate,
	}
}
