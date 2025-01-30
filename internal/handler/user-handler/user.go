package userhandler

import (
	"database/sql"
	"errors"

	basehandler "mindscribe-be/internal/handler/base-handler"
	"mindscribe-be/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type UserHandler struct {
	DB   *sqlx.DB
	base *basehandler.BaseHandler
}

func NewUserHandler(db *sqlx.DB, base *basehandler.BaseHandler) *UserHandler {
	return &UserHandler{
		DB:   db,
		base: base,
	}
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	// Validate input
	if req.Email == "" || req.Username == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "All fields are required",
		})
	}

	err, user := h.base.Service.User.CreateUser(c.Context(), req.Email, req.Username, req.Password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return h.base.Res.FailWithMessage(c, response.GeneralBadRequest, "Failed to create user", fiber.Map{})
		}
		return h.base.Res.FailWithMessage(c, response.GeneralServerError, "Failed to create user", fiber.Map{})
	}

	return h.base.Res.Ok(c, response.GeneralAccepted, fiber.Map{
		"id":         user.ID,
		"email":      user.Email,
		"username":   user.Username,
		"created_at": user.CreatedAt,
	})
}
