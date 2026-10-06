package admin

import (
	"context"
	"homewormanager/internal/entity"
)

type UpdateUserRole struct {
	UserID entity.UserID
	Role   entity.Role
}

func (s *Service) UpdateRole(ctx context.Context, req UpdateUserRole) error {
	err := s.userService.UpdateRole(ctx, req.UserID, req.Role)
	if err != nil {
		return err
	}
	return nil
}
