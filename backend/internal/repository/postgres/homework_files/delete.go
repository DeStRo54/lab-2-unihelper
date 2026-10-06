package homework_files

import (
	"context"
	"homewormanager/internal/entity"
)

func (r *Repo) Delete(ctx context.Context, fileID entity.FileID) error {
	q := "DELETE FROM homeworksfiles WHERE file_id = $1"

	t := r.manager.GetTxOrDefault(ctx)

	_, err := t.ExecContext(ctx, q, fileID)
	if err != nil {
		return err
	}
	return nil
}
