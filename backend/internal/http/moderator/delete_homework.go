package moderator

import (
	"github.com/gofiber/fiber/v2"

	"homewormanager/internal/entity"
)

type DeleteHomeworkReq struct {
	HomeworkId entity.HomeworkID `json:"homeworkId"`
}

func (h *Handler) DeleteHomework() fiber.Handler {
	return func(c *fiber.Ctx) error {

		_, ok := c.Locals(entity.Claims).(entity.UserCreds)
		if !ok {
			return fiber.ErrUnauthorized
		}

		var req DeleteHomeworkReq

		if err := c.QueryParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		if req.HomeworkId == entity.HomeworkID(0) {
			return fiber.ErrBadRequest
		}

		err := h.moderatorService.DeleteHomework(c.Context(), req.HomeworkId)

		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"data": "Homework successfully deleted",
		})

	}
}
