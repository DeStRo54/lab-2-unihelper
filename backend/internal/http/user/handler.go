package user

import (
	"github.com/gofiber/fiber/v2"

	"homewormanager/internal/config"
	"homewormanager/internal/middleware"
)

type Handler struct {
	config      *config.Config
	userService UserService
}

func NewUserHandler(config *config.Config, service UserService) *Handler {
	return &Handler{
		config:      config,
		userService: service,
	}
}

func MapUserRoutes(g fiber.Router, h *Handler, mw *middleware.MwManager) {
	g.Post("/register", h.Register())
	g.Post("/auth", h.Auth())
	g.Get("/get", mw.Auth(), h.Get())
	g.Get("/refresh", h.Refresh())
	g.Delete("/logout", mw.Auth(), h.Logout())
}
