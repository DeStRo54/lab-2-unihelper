package schedule

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"homewormanager/internal/entity"
	scheduleService "homewormanager/internal/service/schedule"
)

type GetReq struct {
	FromTime  time.Time `query:"from_time"`
	DaysCount int       `query:"days_count"`
}

func (h *Handler) GetSchedule() fiber.Handler {
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

		days, err := h.scheduleService.GetAllByGroupAndTime(c.Context(), scheduleService.GetSchedule{
			UserID:    creds.UserID,
			GroupID:   creds.GroupID,
			FromTime:  req.FromTime.Local(),
			DaysCount: req.DaysCount,
		})
		if err != nil {
			return err
		}

		daysMap := make(map[string]scheduleDay)
		for _, d := range days {
			key := d.Date.Format("2006-01-02")
			daysMap[key] = scheduleDay{
				OutputClass:         toOutputClass(d.OutputClass),
				IndependentHomework: toHomework(d.IndependentHomework),
			}
		}
		return c.JSON(daysMap)
	}
}
