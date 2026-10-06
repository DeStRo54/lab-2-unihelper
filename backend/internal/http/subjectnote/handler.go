package subjectnote

import (
	"github.com/gofiber/fiber/v2"
	"homewormanager/internal/middleware"
)

type Handler struct {
	subjectNoteService SubjectNoteService
}

func NewSubjectNoteHandler(subjectNoteService SubjectNoteService) *Handler {
	return &Handler{
		subjectNoteService: subjectNoteService,
	}
}

func MapSubjectNoteRoutes(g fiber.Router, h *Handler, mw *middleware.MwManager) {
	g.Get("", mw.Auth(), h.GetSubjectNotes())
}
