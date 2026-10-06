package schedule

import (
	"context"

	"homewormanager/internal/entity"
	scheduleService "homewormanager/internal/service/schedule"
)

type ScheduleService interface {
	GetAllByGroupAndTime(ctx context.Context, req scheduleService.GetSchedule) ([]entity.ScheduleDay, error)
	GetHomeworks(ctx context.Context, req scheduleService.GetHomework) ([]entity.HomeworkDay, error)
}
