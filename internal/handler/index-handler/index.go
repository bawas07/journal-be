package indexhandler

import (
	basehandler "mindscribe-be/internal/handler/base-handler"
	"mindscribe-be/pkg/response"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type IndexHandler struct {
	DB   *sqlx.DB
	base *basehandler.BaseHandler
}

func NewIndexHandler(db *sqlx.DB, base *basehandler.BaseHandler) *IndexHandler {
	return &IndexHandler{
		DB:   db,
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
