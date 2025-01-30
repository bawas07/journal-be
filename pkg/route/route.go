package route

import (
	"mindscribe-be/internal/handler"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// Router handles all route definitions
type Router struct {
	app     *fiber.App
	handler *handler.Handler
	log     *zap.Logger
}

// New creates a new Router instance
func NewRouter(app *fiber.App, h *handler.Handler, logger *zap.Logger) *Router {
	// Add static here for now
	return &Router{
		app:     app,
		handler: h,
		log:     logger,
	}
}

// Setup registers all application routes
func (r *Router) Setup() {
	// Log all routes being registered
	start := time.Now()

	r.log.Info("Routes: Starting...")

	// Add HTTP logger middleware

	// API v1 group
	v1 := r.app.Group("/api/v1")

	// Health check route
	v1.Get("/health", r.handler.Index.HealthCheck)

	duration := time.Since(start)
	r.log.Info("Routes: Completed", zap.Duration("duration", duration))
}

// RegisterRoutes is a helper function to setup routes on an app instance
func RegisterRoutes(app *fiber.App, h *handler.Handler, log *zap.Logger) error {
	router := NewRouter(app, h, log)
	router.Setup()
	return nil
}
