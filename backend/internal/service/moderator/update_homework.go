package moderator

import (
	"context"

	"homewormanager/internal/entity"
)

func (s *Service) UpdateHomework(ctx context.Context, homeworkId entity.HomeworkID, homeworkText string) error {
	err := s.homeworkService.Update(ctx, homeworkId, homeworkText)
	if err != nil {
		return err
	}
	return nil
}
