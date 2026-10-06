package user

import (
	"github.com/gofiber/fiber/v2"

	"homewormanager/internal/entity"
	userService "homewormanager/internal/service/user"
)

type RegisterReq struct {
	Name        string         `json:"name"`
	Surname     *string        `json:"surname"`
	Email       string         `json:"email"`
	Password    string         `json:"password"`
	GroupID     entity.GroupID `json:"groupId"`
	RegisterKey string         `json:"registerKey"`
}

func (h *Handler) Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req RegisterReq

		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		//TODO добавить проверку на валидность email
		if req.Name == "" || req.Email == "" || req.Password == "" || req.GroupID == 0 || req.RegisterKey == "" {
			return fiber.ErrBadRequest
		}

		err := h.userService.Create(c.Context(), req.toCreate())
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"data": "Successfully registered",
		})
	}
}

func (r *RegisterReq) toCreate() userService.CreateUser {
	return userService.CreateUser{
		Name:        r.Name,
		Surname:     r.Surname,
		Email:       r.Email,
		Password:    r.Password,
		GroupID:     r.GroupID,
		RegisterKey: r.RegisterKey,
	}
}
