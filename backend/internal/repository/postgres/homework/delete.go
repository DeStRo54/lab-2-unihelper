package homework

import (
	"context"

	"homewormanager/internal/entity"
)

func (r *Repo) Delete(ctx context.Context, id entity.HomeworkID) error {
	q := "DELETE FROM homeworks WHERE homework_id = $1"

	t := r.manager.GetTxOrDefault(ctx)

	_, err := t.ExecContext(ctx, q, id)

	if err != nil {
		return err
	}

	return nil
}
