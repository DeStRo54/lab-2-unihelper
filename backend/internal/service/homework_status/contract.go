package homework_status

import (
	"context"

	"homewormanager/internal/entity"
)

type HomeworkStatusRepo interface {
	GetByUserAndHomework(ctx context.Context, userID entity.UserID, homeworkID entity.HomeworkID) (entity.HomeworkStatus, error)
	Delete(ctx context.Context, id entity.StatusID) error
	Create(ctx context.Context, HomeworkStatus entity.HomeworkStatus) error
	DeleteByHomeworkID(ctx context.Context, id entity.HomeworkID) error
	Clear(ctx context.Context) error
}
