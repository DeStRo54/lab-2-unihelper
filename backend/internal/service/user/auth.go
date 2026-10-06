package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"homewormanager/internal/entity"
	"homewormanager/internal/errs"
)

type AuthUser struct {
	Email    string
	Password string
}

func (s *Service) Auth(ctx context.Context, req AuthUser) (sessionKey string, err error) {
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", errs.ErrInvalidPassword
	}

	sessionKey, err = s.userRedisRepo.CreateCreds(ctx, entity.UserCreds{
		UserID:  user.UserID,
		Role:    user.Role,
		GroupID: user.GroupID,
	})

	if err != nil {
		return "", err
	}
	return
}
