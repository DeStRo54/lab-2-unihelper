package group

import (
	"context"
	"homewormanager/internal/entity"
)

func (s *Service) GetByID(ctx context.Context, groupID entity.GroupID) (entity.Group, error) {
	groups, err := s.groupRepo.GetByID(ctx, groupID)
	if err != nil {
		return entity.Group{}, err
	}
	return groups, nil
}
