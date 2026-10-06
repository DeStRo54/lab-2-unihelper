package admin

import (
	"context"

	"homewormanager/internal/entity"
)

type UserService interface {
	UpdateRole(ctx context.Context, userID entity.UserID, role entity.Role) error
	GetAllUsersFull(ctx context.Context) ([]entity.UserFullInfo, error)
}

type GroupService interface {
	Create(ctx context.Context, group entity.Group) (entity.GroupID, error)
	GetAllGroups(ctx context.Context) ([]entity.Group, error)
	RegenerateRegisterKeys(ctx context.Context) error
}

type ClassService interface {
	UpdGroupClasses(ctx context.Context, group entity.Group, classes []entity.Class) error
	ClearAllClasses(ctx context.Context) error
}

type SubjectService interface {
	UpdGroupSubjects(ctx context.Context, group entity.Group, subjects []string) error
	ClearAllSubjects(ctx context.Context) error
}

type HomeworkService interface {
	ClearAllHomeworks(ctx context.Context) error
}

type SubjectNoteService interface {
	ClearAllNotes(ctx context.Context) error
}

type HomeworkStatusService interface {
	ClearAllHomeworkStatuses(ctx context.Context) error
}
