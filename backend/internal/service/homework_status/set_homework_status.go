package homework_status

import (
	"context"
	"errors"

	"homewormanager/internal/entity"
	"homewormanager/internal/errs"
)

type SetHomeworkStatus struct {
	UserID     entity.UserID
	HomeworkId entity.HomeworkID
	Status     bool
}

func (s *Service) SetHomeworkStatus(ctx context.Context, req SetHomeworkStatus) error {
	reqStatus := req.toHomeworkStatus()

	prevStatus, err := s.homeworkStatusRepo.GetByUserAndHomework(ctx, reqStatus.UserID, reqStatus.HomeworkID)
	if errors.Is(err, errs.HomeworkStatusNotFound) {
		if reqStatus.Status {
			err := s.homeworkStatusRepo.Create(ctx, reqStatus)
			if err != nil {
				return err
			}
			return nil
		}
		return nil
	}
	if err != nil {
		return err
	}
	if !reqStatus.Status {
		err := s.homeworkStatusRepo.Delete(ctx, prevStatus.ID)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (s *SetHomeworkStatus) toHomeworkStatus() entity.HomeworkStatus {
	return entity.HomeworkStatus{
		UserID:     s.UserID,
		HomeworkID: s.HomeworkId,
		Status:     s.Status,
	}
}
