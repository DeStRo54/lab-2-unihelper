package admin

import (
	"github.com/gofiber/fiber/v2"
	"homewormanager/internal/entity"
	adminService "homewormanager/internal/service/admin"
	"strconv"
)

type UpdateRoleReq struct {
	UserID entity.UserID
	Role   entity.Role `json:"role"`
}

func (h *Handler) UpdateRole() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var (
			req UpdateRoleReq
			err error
		)

		id, err := strconv.Atoi(c.Params("userID"))
		if err != nil {
			return fiber.ErrBadRequest
		}

		req.UserID = entity.UserID(id)

		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		if req.Role == 0 && req.UserID == 0 {
			return fiber.ErrBadRequest
		}

		err = h.adminService.UpdateRole(c.Context(), adminService.UpdateUserRole{
			UserID: req.UserID,
			Role:   req.Role,
		})
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"data": "User role updated successfully",
		})
	}
}
