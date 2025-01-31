package handler

import (
	"mindscribe-be/pkg/response"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type IndexHandler struct {
	DB   *sqlx.DB
	base *BaseHandler
}

func newIndexHandler(base *BaseHandler) *IndexHandler {
	return &IndexHandler{
		base: base,
	}
}

func (h *IndexHandler) HealthCheck(c *fiber.Ctx) error {
	data := fiber.Map{
		"status":    "OK",
		"timestamp": time.Now(),
	}
	return h.base.Res.OkWithMessage(c, response.GeneralSuccess, "System is healthy", data)
}
