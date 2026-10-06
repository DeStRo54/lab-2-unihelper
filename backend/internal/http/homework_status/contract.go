package homework_status

import (
	"context"

	"homewormanager/internal/service/homework_status"
)

type HomeworkStatusService interface {
	SetHomeworkStatus(ctx context.Context, req homework_status.SetHomeworkStatus) error
}
