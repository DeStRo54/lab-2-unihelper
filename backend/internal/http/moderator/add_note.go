package moderator

import (
	"github.com/gofiber/fiber/v2"
	"homewormanager/internal/entity"
	moderatorService "homewormanager/internal/service/moderator"
	"strconv"
)

type AddNoteReq struct {
	SubjectID entity.SubjectID
	NoteText  string `json:"noteText"`
}

func (h *Handler) AddNote() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var (
			req AddNoteReq
			err error
		)

		creds, ok := c.Locals(entity.Claims).(entity.UserCreds)
		if !ok {
			return fiber.ErrUnauthorized
		}

		subjectID, err := strconv.Atoi(c.Params("subjectID"))
		if err != nil {
			return fiber.ErrBadRequest
		}

		req.SubjectID = entity.SubjectID(subjectID)

		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		if req.NoteText == "" && req.SubjectID == 0 {
			return fiber.ErrBadRequest
		}

		noteID, err := h.moderatorService.AddNote(c.Context(), moderatorService.AddNote{
			SubjectID: req.SubjectID,
			GroupID:   creds.GroupID,
			NoteText:  req.NoteText,
		})

		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"note_id": noteID,
			"data":    "Note successfully added",
		})
	}
}
