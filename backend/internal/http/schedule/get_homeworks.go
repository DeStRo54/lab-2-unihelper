package schedule

import (
	"github.com/gofiber/fiber/v2"
	"homewormanager/internal/entity"
	scheduleService "homewormanager/internal/service/schedule"
)

func (h *Handler) GetHomeworks() fiber.Handler {
	return func(c *fiber.Ctx) error {
		creds, ok := c.Locals(entity.Claims).(entity.UserCreds)
		if !ok {
			return fiber.ErrUnauthorized
		}

		var req GetReq
		if err := c.QueryParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		if req.FromTime.IsZero() || req.DaysCount == 0 {
			return fiber.ErrBadRequest
		}

		days, err := h.scheduleService.GetHomeworks(c.Context(), scheduleService.GetHomework{
			UserID:    creds.UserID,
			GroupID:   creds.GroupID,
			FromTime:  req.FromTime,
			DaysCount: req.DaysCount,
		})
		if err != nil {
			return err
		}

		daysMap := make(map[string]homeworkDay)
		for _, d := range days {
			key := d.Date.Format("2006-01-02")
			if len(d.Homework) == 0 {
				continue
			}
			daysMap[key] = homeworkDay{
				Homework: toHomework(d.Homework),
			}
		}
		return c.JSON(daysMap)
	}
}
