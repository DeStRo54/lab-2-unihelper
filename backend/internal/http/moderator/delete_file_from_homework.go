package moderator

import (
	"github.com/gofiber/fiber/v2"
	"homewormanager/internal/entity"
	"strconv"
)

func (h *Handler) DeleteFileFromHomework() fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, ok := c.Locals(entity.Claims).(entity.UserCreds)
		if !ok {
			return fiber.ErrUnauthorized
		}
		fileID, err := strconv.Atoi(c.Params("fileID"))
		if err != nil {
			return fiber.ErrBadRequest
		}

		err = h.moderatorService.DeleteFileFromHomework(c.Context(), entity.FileID(fileID))
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"data": "File successfully deleted",
		})
	}
}
