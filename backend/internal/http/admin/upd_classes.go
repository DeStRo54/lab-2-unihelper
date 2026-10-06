package admin

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) UpdateClasses() fiber.Handler {
	return func(c *fiber.Ctx) error {

		err := h.adminService.UpdateClasses(c.Context())
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"data": "Class updated successfully",
		})
	}
}
