package handler

import (
	"database/sql"
	"errors"

	"mindscribe-be/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	base *BaseHandler
}

func newUserHandler(base *BaseHandler) *UserHandler {
	return &UserHandler{
		base: base,
	}
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email,isUnique=users;email"`
	Username string `json:"username" validate:"required,min=2"`
	Password string `json:"password" validate:"required,min=6"`
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	if err := h.base.Validate.ValidateStruct(req); err != nil {
		return h.base.Res.ValidationError(c, err)
	}

	// // Validate input
	// if req.Email == "" || req.Username == "" || req.Password == "" {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": "All fields are required",
	// 	})
	// }

	err, user := h.base.Service.User.Create(c.Context(), req.Email, req.Username, req.Password)

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
