package class

import (
	"context"
	"time"

	"homewormanager/internal/entity"
)

func (s *Service) GetByGroupAndTime(ctx context.Context, groupID entity.GroupID, fromTime time.Time, toTime time.Time) ([]entity.Class, error) {

	classes, err := s.classRepo.GetByGroupAndTime(ctx, groupID, fromTime, toTime)
	if err != nil {
		return nil, err
	}
	return classes, nil
}
