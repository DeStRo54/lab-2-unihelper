package moderator

import (
	"github.com/gofiber/fiber/v2"
	"homewormanager/internal/entity"
)

func (h *Handler) GetGroupRegisterKey() fiber.Handler {
	return func(c *fiber.Ctx) error {

		creds, ok := c.Locals(entity.Claims).(entity.UserCreds)
		if !ok {
			return fiber.ErrUnauthorized
		}

		registerKey, err := h.moderatorService.GetGroupRegisterKey(c.Context(), creds.GroupID)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"register_key": registerKey,
		})
	}
}
