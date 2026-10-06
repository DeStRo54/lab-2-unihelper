package user

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"homewormanager/internal/entity"
	userService "homewormanager/internal/service/user"
)

type AuthReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req AuthReq
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}
		if req.Email == "" || req.Password == "" {
			return fiber.ErrBadRequest
		}

		sessionKey, err := h.userService.Auth(c.Context(), userService.AuthUser{
			Email:    req.Email,
			Password: req.Password,
		})
		if err != nil {
			return err
		}

		expiresTime := time.Now().Add(h.config.AuthTTL)

		c.Cookie(&fiber.Cookie{
			Name:    entity.SessionKey,
			Value:   sessionKey,
			Path:    "/",
			Domain:  h.config.Domain,
			Expires: expiresTime,
		})
		c.Cookie(&fiber.Cookie{
			Name:    "session_expires",
			Value:   expiresTime.String(),
			Path:    "/",
			Domain:  h.config.Domain,
			Expires: expiresTime,
		})

		return c.JSON(fiber.Map{
			"data": "User successfully authed",
		})
	}
}
