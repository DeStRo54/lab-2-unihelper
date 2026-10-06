package homework

import (
	"context"
	"time"

	"homewormanager/internal/entity"
)

func (s *Service) GetByGroupAndTime(ctx context.Context, userID entity.UserID, groupID entity.GroupID, fromTime time.Time, toTime time.Time) ([]entity.Homework, error) {
	homeworks, err := s.homeworkRepo.GetByGroupAndTime(ctx, userID, groupID, fromTime, toTime)
	if err != nil {
		return nil, err
	}
	return homeworks, nil
}
