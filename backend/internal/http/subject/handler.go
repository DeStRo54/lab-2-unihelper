package subject

import (
	"github.com/gofiber/fiber/v2"
	"homewormanager/internal/middleware"
)

type Handler struct {
	subjectService SubjectService
}

func NewSubjectHandler(subjectService SubjectService) *Handler {
	return &Handler{
		subjectService: subjectService,
	}
}

func MapSubjectRoutes(g fiber.Router, h *Handler, mw *middleware.MwManager) {
	g.Get("", mw.Auth(), h.GetSubjects())
}
