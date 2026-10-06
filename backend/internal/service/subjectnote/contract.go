package subjectnote

import (
	"context"
	"homewormanager/internal/entity"
)

type SubjectNoteRepo interface {
	Create(ctx context.Context, note entity.SubjectNote) (entity.NoteID, error)
	Clear(ctx context.Context) error
	Delete(ctx context.Context, id entity.NoteID) error
	Update(ctx context.Context, id entity.NoteID, noteText string) error
	GetByGroup(ctx context.Context, groupID entity.GroupID) ([]entity.SubjectNote, error)
}

type SubjectService interface {
	GetByGroupId(ctx context.Context, groupId entity.GroupID) ([]entity.Subject, error)
}
