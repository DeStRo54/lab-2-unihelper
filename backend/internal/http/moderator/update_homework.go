package moderator

import (
	"github.com/gofiber/fiber/v2"

	"homewormanager/internal/entity"
)

type UpdHomeworkReq struct {
	HomeworkId   entity.HomeworkID `json:"homeworkId"`
	HomeworkText string            `json:"homeworkText"`
}

func (h *Handler) UpdateHomework() fiber.Handler {
	return func(c *fiber.Ctx) error {

		_, ok := c.Locals(entity.Claims).(entity.UserCreds)
		if !ok {
			return fiber.ErrUnauthorized
		}

		var req UpdHomeworkReq

		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		if req.HomeworkId == entity.HomeworkID(0) || req.HomeworkText == "" {
			return fiber.ErrBadRequest
		}

		err := h.moderatorService.UpdateHomework(c.Context(), req.HomeworkId, req.HomeworkText)

		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"data": "Homework successfully updated",
		})

	}
}
