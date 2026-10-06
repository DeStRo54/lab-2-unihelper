package homework_status

import (
	"github.com/gofiber/fiber/v2"

	"homewormanager/internal/middleware"
)

type Handler struct {
	homeworkStatusService HomeworkStatusService
}

func NewHomeworkStatusHandler(homeworkStatusService HomeworkStatusService) *Handler {
	return &Handler{
		homeworkStatusService: homeworkStatusService,
	}
}

func MapHomeworkStatusRoutes(g fiber.Router, h *Handler, mw *middleware.MwManager) {
	g.Post("/status/:homeworkID", mw.Auth(), h.SetHomeworkStatus())
}
