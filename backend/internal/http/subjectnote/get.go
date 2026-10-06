package subjectnote

import (
	"github.com/gofiber/fiber/v2"
	"homewormanager/internal/entity"
)

func (h *Handler) GetSubjectNotes() fiber.Handler {
	return func(c *fiber.Ctx) error {
		creds, ok := c.Locals(entity.Claims).(entity.UserCreds)
		if !ok {
			return fiber.ErrUnauthorized
		}

		subjectWithNotes, err := h.subjectNoteService.GetByGroup(c.Context(), creds.GroupID)
		if err != nil {
			return err
		}

		resp := toSubjectNotes(subjectWithNotes)

		return c.JSON(resp)
	}
}
