package middleware

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"

	"homewormanager/internal/entity"
)

func (mw *MwManager) Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessionKey := c.Cookies(entity.SessionKey)

		if sessionKey == "" {
			return fiber.ErrUnauthorized
		}

		userCreds, err := mw.userRedisRepo.GetCreds(c.Context(), sessionKey)
		if err != nil {
			if errors.Is(err, redis.Nil) {
				return fiber.ErrUnauthorized
			}
			return err
		}
		userCreds.Role, err = mw.lookupRole(c.Context(), userCreds.UserID)
		if err != nil {
			return err
		}
		c.Locals(entity.Claims, userCreds)

		return c.Next()

	}
}
