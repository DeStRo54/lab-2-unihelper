package group

import (
	"context"

	"homewormanager/internal/entity"
)

//go:generate mockgen -source $GOFILE -destination contract_mocks_test.go -package $GOPACKAGE

type GroupRepo interface {
	Create(ctx context.Context, group entity.Group) (entity.GroupID, error)
	GetByName(ctx context.Context, name string) (entity.Group, error)
	GetAllGroups(ctx context.Context) ([]entity.Group, error)
	ChangeRegisterKey(ctx context.Context, group entity.Group) error
	GetByID(ctx context.Context, groupID entity.GroupID) (entity.Group, error)
}
