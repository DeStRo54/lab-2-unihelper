package middleware

import (
	"github.com/gofiber/fiber/v2"

	"homewormanager/internal/entity"
)

func (mw *MwManager) AllowedRoles(allowedRoles []entity.Role) fiber.Handler {
	return func(c *fiber.Ctx) error {
		cl, ok := c.Locals(entity.Claims).(entity.UserCreds)
		if !ok {
			return fiber.ErrUnauthorized
		}

		flag := false
		for _, role := range allowedRoles {
			if cl.Role == role {
				flag = true
				break
			}
		}

		if !flag {
			return fiber.ErrForbidden
		}

		return c.Next()
	}
}
