package user

import (
	"context"

	"homewormanager/internal/entity"
)

func (s *Service) GetUser(ctx context.Context, id entity.UserID) (entity.User, error) {
	user, err := s.userRepo.GetById(ctx, id)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
