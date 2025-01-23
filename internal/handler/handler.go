package handler

import (
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	// Add service dependencies here
}

func NewHandler() *Handler {
	return &Handler{
		// Initialize services here
	}
}

func RegisterRoutes(router fiber.Router) {
	h := NewHandler()
	
	router.Get("/health", h.HealthCheck)
}

func (h *Handler) HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "OK",
	})
}