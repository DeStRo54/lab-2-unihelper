package admin

import (
	"context"
	"homewormanager/internal/entity"
)

func (s *Service) GetAllGroups(ctx context.Context) ([]entity.Group, error) {
	users, err := s.groupService.GetAllGroups(ctx)
	if err != nil {
		return []entity.Group{}, err
	}
	return users, nil
}
