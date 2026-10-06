package homework_files

import (
	"context"
	"database/sql"
	"errors"
	"homewormanager/internal/entity"
	"homewormanager/internal/errs"
)

func (r *Repo) GetByGroupID(ctx context.Context, id entity.GroupID) ([]entity.HomeworkFile, error) {
	q := "SELECT * FROM homeworksfiles WHERE group_id = $1"

	t := r.manager.GetTxOrDefault(ctx)

	var files []homeworkFile
	err := t.SelectContext(ctx, &files, q, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []entity.HomeworkFile{}, errs.FileNotFound
		}
		return []entity.HomeworkFile{}, err
	}

	return toHomeworkFiles(files), err
}
