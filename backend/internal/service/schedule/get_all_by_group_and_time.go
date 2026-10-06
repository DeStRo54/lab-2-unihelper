package schedule

import (
	"context"
	"time"

	"homewormanager/internal/entity"
	"homewormanager/internal/errs"
	"homewormanager/internal/utils"
)

type GetSchedule struct {
	UserID    entity.UserID
	GroupID   entity.GroupID
	FromTime  time.Time
	DaysCount int
}

func (s *Service) GetAllByGroupAndTime(ctx context.Context, req GetSchedule) ([]entity.ScheduleDay, error) {
	toTime := req.FromTime.Add(time.Duration(req.DaysCount*24) * time.Hour)

	classes, err := s.classService.GetByGroupAndTime(ctx, req.GroupID, req.FromTime, toTime)
	if err != nil {
		if err != errs.ClassesNotFound {
			return nil, err
		}
	}

	homeworks, err := s.homeworkService.GetByGroupAndTime(ctx, req.UserID, req.GroupID, req.FromTime, toTime)
	if err != nil {
		if err != errs.HomeworksNotFound {
			return nil, err
		}
	}

	days := make([]entity.ScheduleDay, req.DaysCount)
	files, err := s.homeworkFilesService.GetByGroupID(ctx, req.GroupID)
	filesMap := make(map[entity.HomeworkID][]entity.HomeworkFile)

	for i := 0; i < len(files); i++ {
		filesMap[files[i].HomeworkID] = append(filesMap[files[i].HomeworkID], files[i])
	}

	for i := 0; i < req.DaysCount; i++ {
		days[i] = entity.ScheduleDay{
			Date:                req.FromTime.Add(time.Duration(i) * 24 * time.Hour).Local(),
			OutputClass:         []entity.OutputClass{},
			IndependentHomework: []entity.Homework{},
		}
	}

	classMap := make(map[time.Time][]entity.Class)
	for _, class := range classes {
		date := time.Date(
			class.StartTime.Year(), class.StartTime.Month(), class.StartTime.Day(),
			0, 0, 0, 0,
			time.Local)
		classMap[date] = append(classMap[date], class)
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
		dayClasses := classMap[day.Date]
		dayHomeworks := homeworkMap[day.Date]

		for _, class := range dayClasses {
			outputClass := entity.OutputClass{
				Class:    class,
				Homework: []entity.Homework{},
			}

			for _, homework := range dayHomeworks {
				if homework.SubjectID == class.SubjectID && utils.DeRef[int64](homework.SemClassNumber) == class.SemClassNumber && utils.DeRef[entity.ClassCategory](homework.Category) == class.Category {
					outputClass.Homework = append(outputClass.Homework, homework)
				}
			}

			day.OutputClass = append(day.OutputClass, outputClass)
		}

		for _, homework := range dayHomeworks {
			if utils.DeRef[int64](homework.SemClassNumber) == 0 {
				day.IndependentHomework = append(day.IndependentHomework, homework)
			}
		}

		days[i] = day

	}
	return days, nil
}
