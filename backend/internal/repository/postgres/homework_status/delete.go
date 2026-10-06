package homework_status

import (
	"context"

	"homewormanager/internal/entity"
)

func (r *Repo) Delete(ctx context.Context, id entity.StatusID) error {
	q := "DELETE FROM homeworkstatuses WHERE id = $1"

	t := r.manager.GetTxOrDefault(ctx)

	_, err := t.ExecContext(ctx, q, id)

	if err != nil {
		return err
	}

	return nil
}
