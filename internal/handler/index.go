package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type IndexHandler struct {
	DB     *sqlx.DB
	logger *zap.Logger
}

func NewIndexHandler(db *sqlx.DB, log *zap.Logger) *IndexHandler {
	return &IndexHandler{
		DB:     db,
		logger: log,
	}
}

func (h *IndexHandler) HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "OK",
	})
}
