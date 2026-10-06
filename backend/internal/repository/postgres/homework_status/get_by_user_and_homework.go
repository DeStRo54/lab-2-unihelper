package homework_status

import (
	"context"
	"database/sql"
	"errors"

	"homewormanager/internal/entity"
	"homewormanager/internal/errs"
)

func (r *Repo) GetByUserAndHomework(ctx context.Context, userID entity.UserID, homeworkID entity.HomeworkID) (entity.HomeworkStatus, error) {
	q := "SELECT * FROM homeworkstatuses WHERE user_id = $1 AND homework_id = $2"

	t := r.manager.GetTxOrDefault(ctx)

	var homeworkStatus homeworkStatus

	err := t.GetContext(ctx, &homeworkStatus, q, userID, homeworkID)
	if errors.Is(err, sql.ErrNoRows) {
		return entity.HomeworkStatus{}, errs.HomeworkStatusNotFound
	}
	if err != nil {
		return entity.HomeworkStatus{}, err
	}

	return homeworkStatus.toHomeworkStatus(), nil
}
