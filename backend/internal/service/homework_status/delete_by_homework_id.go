package homework_status

import (
	"context"
	"errors"
	"homewormanager/internal/entity"
	"homewormanager/internal/errs"
)

func (s *Service) DeleteHomeworkID(ctx context.Context, homeworkID entity.HomeworkID) error {
	err := s.homeworkStatusRepo.DeleteByHomeworkID(ctx, homeworkID)
	if errors.Is(err, errs.HomeworkStatusNotFound) {
		return nil
	}
	if err != nil {
		return err
	}

	return nil
}
