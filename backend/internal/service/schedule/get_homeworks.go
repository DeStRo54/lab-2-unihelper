package schedule

import (
	"context"
	"time"

	"homewormanager/internal/entity"
	"homewormanager/internal/errs"
)

type GetHomework struct {
	UserID    entity.UserID
	GroupID   entity.GroupID
	FromTime  time.Time
	DaysCount int
}

func (s *Service) GetHomeworks(ctx context.Context, req GetHomework) ([]entity.HomeworkDay, error) {
	toTime := req.FromTime.Add(time.Duration(req.DaysCount*24) * time.Hour)

	homeworks, err := s.homeworkService.GetByGroupAndTime(ctx, req.UserID, req.GroupID, req.FromTime, toTime)
	if err != nil {
		if err != errs.HomeworksNotFound {
			return nil, err
		}
	}

	days := make([]entity.HomeworkDay, req.DaysCount)

	for i := 0; i < req.DaysCount; i++ {
		days[i] = entity.HomeworkDay{
			Date:     req.FromTime.Add(time.Duration(i) * 24 * time.Hour).Local(),
			Homework: []entity.Homework{},
		}
	}

	files, err := s.homeworkFilesService.GetByGroupID(ctx, req.GroupID)
	filesMap := make(map[entity.HomeworkID][]entity.HomeworkFile)

	for i := 0; i < len(files); i++ {
		filesMap[files[i].HomeworkID] = append(filesMap[files[i].HomeworkID], files[i])
	}

	homeworkMap := make(map[time.Time][]entity.Homework)
	for i := range homeworks {
		date := time.Date(
			homeworks[i].DueDate.Year(), homeworks[i].DueDate.Month(), homeworks[i].DueDate.Day(),
			0, 0, 0, 0,
			time.Local)
		homeworks[i].Files = filesMap[homeworks[i].HomeworkID]
		homeworkMap[date] = append(homeworkMap[date], homeworks[i])
	}

	for i, day := range days {
		dayHomeworks := homeworkMap[day.Date]
		for _, homework := range dayHomeworks {
			day.Homework = append(day.Homework, homework)
		}
		days[i] = day
	}
	return days, nil
}
