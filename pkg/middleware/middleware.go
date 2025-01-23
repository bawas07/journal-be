package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func HttpLogger(logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger.Info("Incoming request",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
		)

		err := c.Next()

		logger.Info("Request completed",
			zap.Int("status", c.Response().StatusCode()),
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
		)

		return err
	}
}
