package middleware

import (
	"context"
	"errors"
	"homewormanager/internal/entity"
	"homewormanager/internal/errs"
)

func (mw *MwManager) lookupRole(ctx context.Context, UserID entity.UserID) (entity.Role, error) {
	role, err := mw.userRedisRepo.GetUserRole(ctx, UserID)
	if err != nil {
		if !errors.Is(err, errs.NoUserRole) {
			return 0, err
		}
	}
	user, err := mw.userRepo.GetById(ctx, UserID)
	if err != nil {
		return 0, err
	}
	role = user.Role

	err = mw.userRedisRepo.CreateUserRole(ctx, UserID, role)
	if err != nil {
		return 0, err
	}

	return role, nil

}
