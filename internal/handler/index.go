package handler

import (
	"mindscribe-be/pkg/response"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type IndexHandler struct {
	DB     *sqlx.DB
	logger *zap.Logger
	res    *response.Response
}

func NewIndexHandler(db *sqlx.DB, log *zap.Logger) *IndexHandler {
	return &IndexHandler{
		DB:     db,
		logger: log,
		res:    response.New(),
	}
}

func (h *IndexHandler) HealthCheck(c *fiber.Ctx) error {
	data := fiber.Map{
		"status":    "OK",
		"timestamp": time.Now(),
	}
	return h.res.OkWithMessage(c, response.GeneralSuccess, "System is healthy", data)
}
