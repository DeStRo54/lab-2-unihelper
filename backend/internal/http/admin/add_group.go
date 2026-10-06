package admin

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	adminService "homewormanager/internal/service/admin"
)

type AddGroupReq struct {
	Name     string `json:"name"`
	Course   int8   `json:"course"`
	IcalLink string `json:"icalLink"`
}

func (h *Handler) AddGroup() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req AddGroupReq

		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		//TODO добавить проверку на формат строки группы "XXXX-YY-YY"
		if req.Name == "" || req.Course == 0 || req.IcalLink == "" || req.Name != strings.ToUpper(req.Name) {
			return fiber.ErrBadRequest
		}

		err := h.adminService.AddGroup(c.Context(), req.toAdd())
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"data": "Group successfully added",
		})
	}
}
func (r *AddGroupReq) toAdd() adminService.AddGroup {
	return adminService.AddGroup{
		Name:     r.Name,
		Course:   r.Course,
		IcalLink: r.IcalLink,
	}
}
