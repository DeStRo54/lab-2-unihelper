package moderator

import (
	"github.com/gofiber/fiber/v2"
	"homewormanager/internal/entity"
	"strconv"
)

type UpdNoteReq struct {
	NoteID   entity.NoteID
	NoteText string `json:"noteText"`
}

func (h *Handler) UpdateNote() fiber.Handler {
	return func(c *fiber.Ctx) error {

		_, ok := c.Locals(entity.Claims).(entity.UserCreds)
		if !ok {
			return fiber.ErrUnauthorized
		}

		var req UpdNoteReq

		id, err := strconv.Atoi(c.Params("noteID"))
		if err != nil {
			return fiber.ErrBadRequest
		}
		req.NoteID = entity.NoteID(id)

		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		if req.NoteID == 0 || req.NoteText == "" {
			return fiber.ErrBadRequest
		}

		err = h.moderatorService.UpdateNote(c.Context(), req.NoteID, req.NoteText)

		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"data": "Note successfully updated",
		})

	}
}
