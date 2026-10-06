package homework_status

import (
	"context"
	"homewormanager/internal/entity"
)

func (r *Repo) DeleteByHomeworkID(ctx context.Context, id entity.HomeworkID) error {
	q := "DELETE FROM homeworkstatuses WHERE homework_id = $1"

	t := r.manager.GetTxOrDefault(ctx)

	_, err := t.ExecContext(ctx, q, id)

	if err != nil {
		return err
	}

	return nil
}
