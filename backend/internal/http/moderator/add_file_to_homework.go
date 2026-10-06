package moderator

import (
	"github.com/gofiber/fiber/v2"
	"homewormanager/internal/entity"
	"homewormanager/internal/errs"
	moderatorService "homewormanager/internal/service/moderator"
	"homewormanager/internal/utils"
	"strconv"
)

func (h *Handler) AddFileToHomework() fiber.Handler {
	return func(c *fiber.Ctx) error {
		creds, ok := c.Locals(entity.Claims).(entity.UserCreds)
		if !ok {
			return fiber.ErrUnauthorized
		}

		homeworkID, err := strconv.Atoi(c.Params("homeworkID"))
		if err != nil {
			return fiber.ErrBadRequest
		}

		form, err := c.MultipartForm()

		files := form.File["files"]
		if len(files) == 0 || homeworkID == 0 {
			return fiber.ErrBadRequest
		}

		if len(files) > 10 {
			return errs.TooManyFiles
		}

		for _, file := range files {
			if err = utils.CheckFile(file); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
					"file":  file.Filename,
				})
			}
		}

		filesIdMap, filesURLMap, filesErrMap, err := h.moderatorService.AddFileToHomework(c.Context(), moderatorService.AddFileReq{
			FilesHeader: files,
			HomeworkID:  entity.HomeworkID(homeworkID),
			GroupID:     creds.GroupID,
		})
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"filesIdMap":  filesIdMap,
			"filesErrMap": filesErrMap,
			"filesURLMap": filesURLMap,
			"data":        "Files successfully added",
		})

	}
}
