package homework

import (
	"context"

	"homewormanager/internal/entity"
)

func (r *Repo) Update(ctx context.Context, id entity.HomeworkID, homeworkText string) error {
	q := "UPDATE homeworks SET homework_text = $1 WHERE homework_id = $2"

	t := r.manager.GetTxOrDefault(ctx)

	_, err := t.ExecContext(ctx, q, homeworkText, id)

	if err != nil {
		return err
	}

	return nil
}
