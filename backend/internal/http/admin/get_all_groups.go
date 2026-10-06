package admin

import (
	"github.com/gofiber/fiber/v2"
	"homewormanager/internal/entity"
)

type GroupResp struct {
	GroupID     entity.GroupID `json:"group_id"`
	Name        string         `json:"group_name"`
	Course      int8           `json:"course"`
	IcalLink    string         `json:"ical_link"`
	RegisterKey string         `json:"register_key"`
}

func (h *Handler) GetAllGroups() fiber.Handler {
	return func(c *fiber.Ctx) error {

		u, err := h.adminService.GetAllGroups(c.Context())
		if err != nil {
			return err
		}

		return c.JSON(toGroupResp(u))
	}
}

func toGroupResp(u []entity.Group) []GroupResp {
	res := make([]GroupResp, len(u))
	for i, g := range u {
		res[i] = GroupResp{
			GroupID:     g.GroupID,
			Name:        g.Name,
			Course:      g.Course,
			IcalLink:    g.IcalLink,
			RegisterKey: g.RegisterKey,
		}
	}
	return res
}
