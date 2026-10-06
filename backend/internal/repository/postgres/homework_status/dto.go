package homework_status

import "homewormanager/internal/entity"

type homeworkStatus struct {
	ID         entity.StatusID   `db:"id"`
	UserID     entity.UserID     `db:"user_id"`
	HomeworkID entity.HomeworkID `db:"homework_id"`
}

func (s *homeworkStatus) toHomeworkStatus() entity.HomeworkStatus {
	return entity.HomeworkStatus{
		ID:         s.ID,
		UserID:     s.UserID,
		HomeworkID: s.HomeworkID,
		Status:     true,
	}
}
