package homework

import (
	"context"

	"homewormanager/internal/entity"
)

func (s *Service) Update(ctx context.Context, homeworkId entity.HomeworkID, homeworkText string) error {
	err := s.homeworkRepo.Update(ctx, homeworkId, homeworkText)
	if err != nil {
		return err
	}
	return nil
}
