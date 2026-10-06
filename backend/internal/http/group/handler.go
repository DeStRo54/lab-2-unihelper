package group

import (
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	groupService GroupService
}

func NewGroupHandler(groupService GroupService) *Handler {
	return &Handler{
		groupService: groupService,
	}
}

func MapGroupRoutes(g fiber.Router, h *Handler) {
	g.Get("/getAll", h.GetAllGroup())
}
