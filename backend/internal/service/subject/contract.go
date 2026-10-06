package subject

import (
	"context"

	"homewormanager/internal/entity"
)

type SubjectRepo interface {
	GetByGroupId(ctx context.Context, groupId entity.GroupID) ([]entity.Subject, error)
	GetByNameAndGroup(ctx context.Context, name string, groupID entity.GroupID) (entity.Subject, error)
	Create(ctx context.Context, subject entity.Subject) error
	Clear(ctx context.Context) error
}
