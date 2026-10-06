package moderator

import (
	"context"

	"homewormanager/internal/entity"
)

func (s *Service) DeleteHomework(ctx context.Context, id entity.HomeworkID) error {
	files, err := s.homeworkFileService.GetByHomeworkID(ctx, id)
	if err != nil {
		return err
	}
	if len(files) != 0 {
		for _, file := range files {
			err = s.homeworkFileService.DeleteFile(ctx, file.FileID)
			if err != nil {
				return err
			}
		}
	}
	err = s.homeworkService.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
