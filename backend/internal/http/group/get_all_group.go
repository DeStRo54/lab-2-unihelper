package group

import (
	"github.com/gofiber/fiber/v2"

	"homewormanager/internal/entity"
)

type GroupResp struct {
	GroupID entity.GroupID `json:"group_id"`
	Name    string         `json:"name"`
	Course  int8           `json:"course"`
}

func (h *Handler) GetAllGroup() fiber.Handler {
	return func(c *fiber.Ctx) error {

		g, err := h.groupService.GetAllGroups(c.Context())
		if err != nil {
			return err
		}
		groups := toGroupResp(g)
		return c.JSON(groups)
	}
}

func toGroupResp(g []entity.Group) []GroupResp {
	res := make([]GroupResp, len(g))
	for i := range g {
		res[i] = GroupResp{
			GroupID: g[i].GroupID,
			Name:    g[i].Name,
			Course:  g[i].Course,
		}
	}
	return res
}
