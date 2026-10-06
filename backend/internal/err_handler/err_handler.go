package err_handler

import (
	"github.com/gofiber/fiber/v2"
	"homewormanager/internal/errs"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := errs.GetStatusFromError(err)
	ctx.Status(code)
	return ctx.JSON(fiber.Map{
		"error": err.Error(),
	})
}
