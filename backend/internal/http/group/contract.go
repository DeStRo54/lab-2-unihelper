package group

import (
	"context"

	"homewormanager/internal/entity"
)

type GroupService interface {
	GetAllGroups(ctx context.Context) ([]entity.Group, error)
}
