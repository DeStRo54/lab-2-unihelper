package admin

import (
	"context"
	"homewormanager/internal/entity"
)

func (s *Service) GetAllUsers(ctx context.Context) ([]entity.UserFullInfo, error) {
	users, err := s.userService.GetAllUsersFull(ctx)
	if err != nil {
		return []entity.UserFullInfo{}, err
	}
	return users, nil
}
