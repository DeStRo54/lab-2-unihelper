package moderator

import (
	"encoding/json"
	"homewormanager/internal/errs"
	"homewormanager/internal/utils"
	"time"

	"github.com/gofiber/fiber/v2"

	"homewormanager/internal/entity"
	moderatorService "homewormanager/internal/service/moderator"
)

type AddHomeworkToClassReq struct {
	ClassSemNumber int64            `json:"classSemNumber"`
	SubjectID      entity.SubjectID `json:"subjectId"`
	Category       string           `json:"category"`
	HomeworkText   string           `json:"homeworkText"`
	DueDate        time.Time        `json:"dueDate"`
}

func (h *Handler) AddHomeworkToClass() fiber.Handler {
	return func(c *fiber.Ctx) error {

		creds, ok := c.Locals(entity.Claims).(entity.UserCreds)
		if !ok {
			return fiber.ErrUnauthorized
		}

		var req AddHomeworkToClassReq

		form, err := c.MultipartForm()
		if err != nil {
			return fiber.ErrBadRequest
		}

		files := form.File["files"]
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

		data := form.Value["data"]
		if err = json.Unmarshal([]byte(data[0]), &req); err != nil {
			return fiber.ErrBadRequest
		}

		if req.ClassSemNumber == 0 || req.SubjectID == 0 || req.HomeworkText == "" || req.DueDate.IsZero() {
			return fiber.ErrBadRequest
		}

		category, ok := entity.CategoryToNumber[req.Category]
		if !ok {
			return fiber.ErrBadRequest
		}

		id, filesIdMap, filesURLMap, filesErrMap, err := h.moderatorService.AddHomework(c.Context(), moderatorService.AddHomework{
			ClassSemNumber: &req.ClassSemNumber,
			GroupID:        creds.GroupID,
			SubjectID:      req.SubjectID,
			Category:       &category,
			HomeworkText:   req.HomeworkText,
			FilesHeader:    files,
			DueDate:        req.DueDate,
		})
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"homework_id": id,
			"filesIdMap":  filesIdMap,
			"filesErrMap": filesErrMap,
			"filesURLMap": filesURLMap,
			"data":        "Homework successfully added",
		})
	}
}
