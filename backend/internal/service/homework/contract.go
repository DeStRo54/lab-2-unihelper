package homework

import (
	"context"
	"homewormanager/internal/entity"
	"time"
)

//go:generate mockgen -source $GOFILE -destination contract_mocks_test.go -package $GOPACKAGE

type HomeworkRepo interface {
	Clear(ctx context.Context) error
	Create(ctx context.Context, homework entity.Homework) (entity.HomeworkID, error)
	Delete(ctx context.Context, id entity.HomeworkID) error
	Update(ctx context.Context, id entity.HomeworkID, homeworkText string) error
	GetByGroupAndTime(ctx context.Context, userID entity.UserID, groupID entity.GroupID, fromTime time.Time, toTime time.Time) ([]entity.Homework, error)
}

type HomeworkStatusService interface {
	DeleteHomeworkID(ctx context.Context, homeworkID entity.HomeworkID) error
}
