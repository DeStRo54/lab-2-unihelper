package user

import (
	"context"

	"homewormanager/internal/entity"
)

type UserRepo interface {
	Create(ctx context.Context, user entity.User) error
	GetByEmail(ctx context.Context, email string) (entity.User, error)
	GetById(ctx context.Context, id entity.UserID) (entity.User, error)
	GetFullById(ctx context.Context, id entity.UserID) (entity.UserFullInfo, error)
	UpdateRole(ctx context.Context, userID entity.UserID, role entity.Role) error
	GetAllFull(ctx context.Context) ([]entity.UserFullInfo, error)
}

type UserRedisRepo interface {
	CreateCreds(ctx context.Context, creds entity.UserCreds) (sessionKey string, err error)
	GetCreds(ctx context.Context, sessionKey string) (entity.UserCreds, error)
	RefreshCreds(ctx context.Context, sessionKey string) error
	DeleteCreds(ctx context.Context, sessionKey string) error
	CreateUserRole(ctx context.Context, id entity.UserID, role entity.Role) error
}

type GroupService interface {
	GetByID(ctx context.Context, groupID entity.GroupID) (entity.Group, error)
}
