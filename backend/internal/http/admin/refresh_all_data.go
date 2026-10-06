package admin

import "github.com/gofiber/fiber/v2"

func (h *Handler) RefreshAllData() fiber.Handler {
	return func(c *fiber.Ctx) error {

		err := h.adminService.RefreshAllData(c.Context())
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"data": "Data refresh successfully",
		})
	}
}
