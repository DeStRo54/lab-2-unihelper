package user

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

func (h Handler) Get() fiber.Handler {
	return func(c *fiber.Ctx) error {
		creds, ok := c.Locals(entity.Claims).(entity.UserCreds)

		if !ok {
			return fiber.ErrUnauthorized
		}

		user, err := h.userService.GetUserFull(c.Context(), creds.UserID)
		if err != nil {
			return err
		}

		return c.JSON(toUserResp(user))
	}
}

func toUserResp(user entity.UserFullInfo) UserResp {
	return UserResp{
		UserID:    user.UserID,
		Name:      user.Name,
		Surname:   user.Surname,
		Email:     user.Email,
		Role:      user.Role,
		GroupName: user.GroupName,
	}
}
