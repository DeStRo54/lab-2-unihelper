package homework_status

import (
	"context"

	"homewormanager/internal/entity"
)

func (r *Repo) Create(ctx context.Context, HomeworkStatus entity.HomeworkStatus) error {
	q := "INSERT INTO homeworkstatuses (user_id, homework_id) VALUES ($1, $2);"

	t := r.manager.GetTxOrDefault(ctx)

	_, err := t.ExecContext(ctx, q, HomeworkStatus.UserID, HomeworkStatus.HomeworkID)
	if err != nil {
		return err
	}

	return nil
}
