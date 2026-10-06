package homework_status

import (
	"github.com/gofiber/fiber/v2"
	"homewormanager/internal/entity"
	"homewormanager/internal/service/homework_status"
	"strconv"
)

type SetHomeworkStatusReq struct {
	HomeworkId entity.HomeworkID
	Status     bool `json:"status"`
}

func (h *Handler) SetHomeworkStatus() fiber.Handler {
	return func(c *fiber.Ctx) error {

		creds, ok := c.Locals(entity.Claims).(entity.UserCreds)
		if !ok {
			return fiber.ErrUnauthorized
		}

		var (
			req SetHomeworkStatusReq
			err error
		)

		id, err := strconv.Atoi(c.Params("homeworkID"))
		if err != nil {
			return fiber.ErrBadRequest
		}
		req.HomeworkId = entity.HomeworkID(id)

		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		if req.HomeworkId == 0 {
			return fiber.ErrBadRequest
		}

		err = h.homeworkStatusService.SetHomeworkStatus(c.Context(), homework_status.SetHomeworkStatus{
			UserID:     creds.UserID,
			HomeworkId: req.HomeworkId,
			Status:     req.Status,
		})
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"data": "Status successfully updated",
		})

	}
}
