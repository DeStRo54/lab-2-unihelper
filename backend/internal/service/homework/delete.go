package homework

import (
	"context"

	"homewormanager/internal/entity"
)

func (s *Service) Delete(ctx context.Context, homeworkID entity.HomeworkID) error {
	err := s.manager.Do(ctx, func(ctx context.Context) error {
		err := s.homeworkStatusService.DeleteHomeworkID(ctx, homeworkID)
		if err != nil {
			return err
		}
		err = s.homeworkRepo.Delete(ctx, homeworkID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
