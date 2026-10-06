package moderator

import (
	"github.com/gofiber/fiber/v2"
	"homewormanager/internal/entity"
	"strconv"
)

func (h *Handler) RegenerateGroupRegisterKey() fiber.Handler {
	return func(c *fiber.Ctx) error {

		creds, ok := c.Locals(entity.Claims).(entity.UserCreds)
		if !ok {
			return fiber.ErrUnauthorized
		}

		var groupID entity.GroupID

		if creds.Role == entity.RoleGroupModerator {
			groupID = creds.GroupID
		} else {
			id, err := strconv.Atoi(c.Params("groupID"))
			if err != nil {
				return fiber.ErrBadRequest
			}
			if id == 0 {
				return fiber.ErrBadRequest
			}
			groupID = entity.GroupID(id)
		}

		registerKey, err := h.moderatorService.RegenerateGroupRegisterKey(c.Context(), groupID)

		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"register_key": registerKey,
			"data":         "Key successfully regenerated",
		})

	}
}
