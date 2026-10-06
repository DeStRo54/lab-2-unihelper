package homework

import (
	"context"

	"homewormanager/internal/entity"
)

func (s *Service) Create(ctx context.Context, homework entity.Homework) (entity.HomeworkID, error) {

	id, err := s.homeworkRepo.Create(ctx, homework)
	if err != nil {
		return 0, err
	}
	return id, nil
}
