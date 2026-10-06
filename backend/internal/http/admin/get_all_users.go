package admin

import (
	"github.com/gofiber/fiber/v2"
	"homewormanager/internal/entity"
)

type UserResp struct {
	UserID    entity.UserID `json:"user_id"`
	Name      string        `json:"name"`
	Surname   *string       `json:"surname"`
	Email     string        `json:"email"`
	Role      entity.Role   `json:"role"`
	GroupName string        `json:"group_name"`
}

func (h *Handler) GetAllUsers() fiber.Handler {
	return func(c *fiber.Ctx) error {

		u, err := h.adminService.GetAllUsers(c.Context())
		if err != nil {
			return err
		}

		return c.JSON(toUserResp(u))
	}
}

func toUserResp(u []entity.UserFullInfo) []UserResp {
	res := make([]UserResp, len(u))
	for i := range u {
		res[i] = UserResp{
			UserID:    u[i].UserID,
			Name:      u[i].Name,
			Surname:   u[i].Surname,
			Email:     u[i].Email,
			Role:      u[i].Role,
			GroupName: u[i].GroupName,
		}
	}
	return res
}
