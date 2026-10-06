package user

import (
	"context"

	"homewormanager/internal/entity"
	userService "homewormanager/internal/service/user"
)

type UserService interface {
	Create(ctx context.Context, req userService.CreateUser) error
	Auth(ctx context.Context, req userService.AuthUser) (sessionKey string, err error)
	GetUserFull(ctx context.Context, id entity.UserID) (entity.UserFullInfo, error)
	RefreshCookie(ctx context.Context, sessionKey string) error
	Logout(ctx context.Context, sessionKey string) error
}
