package moderator

import (
	"github.com/gofiber/fiber/v2"
	"homewormanager/internal/entity"
	"strconv"
)

type DeleteNoteReq struct {
	NoteID entity.NoteID
}

func (h *Handler) DeleteNote() fiber.Handler {
	return func(c *fiber.Ctx) error {

		_, ok := c.Locals(entity.Claims).(entity.UserCreds)
		if !ok {
			return fiber.ErrUnauthorized
		}

		var (
			req DeleteNoteReq
			err error
		)

		id, err := strconv.Atoi(c.Params("noteID"))
		if err != nil {
			return fiber.ErrBadRequest
		}
		req.NoteID = entity.NoteID(id)

		if req.NoteID == 0 {
			return fiber.ErrBadRequest
		}

		err = h.moderatorService.DeleteNote(c.Context(), req.NoteID)

		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"data": "Note successfully deleted",
		})

	}
}
