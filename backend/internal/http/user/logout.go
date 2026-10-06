package user

import (
	"github.com/gofiber/fiber/v2"
	"homewormanager/internal/entity"
)

func (h *Handler) Logout() fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessionKey := c.Cookies(entity.SessionKey)

		if sessionKey == "" {
			return fiber.ErrUnauthorized
		}

		err := h.userService.Logout(c.Context(), sessionKey)
		if err != nil {
			return fiber.ErrUnauthorized
		}

		return c.JSON(fiber.Map{
			"data": "user logout success",
		})

	}
}
