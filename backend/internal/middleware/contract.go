package middleware

import (
	"context"

	"homewormanager/internal/entity"
)

type UserRedisRepo interface {
	GetCreds(ctx context.Context, sessionKey string) (entity.UserCreds, error)
	CreateUserRole(ctx context.Context, id entity.UserID, role entity.Role) error
	GetUserRole(ctx context.Context, id entity.UserID) (entity.Role, error)
}
type UserRepo interface {
	GetById(ctx context.Context, id entity.UserID) (entity.User, error)
}
