package admin

import "github.com/gofiber/fiber/v2"

func (h *Handler) RegenerateAllRegisterKeys() fiber.Handler {
	return func(c *fiber.Ctx) error {

		err := h.adminService.RegenerateAllRegisterKeys(c.Context())
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"data": "Keys updated successfully",
		})
	}
}
