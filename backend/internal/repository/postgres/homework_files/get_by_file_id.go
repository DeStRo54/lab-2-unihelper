package homework_files

import (
	"context"
	"database/sql"
	"errors"
	"homewormanager/internal/entity"
	"homewormanager/internal/errs"
)

func (r *Repo) GetByFileID(ctx context.Context, fileID entity.FileID) (entity.HomeworkFile, error) {
	q := "SELECT * FROM homeworksfiles WHERE file_id = $1"

	t := r.manager.GetTxOrDefault(ctx)

	var file homeworkFile
	err := t.GetContext(ctx, &file, q, fileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.HomeworkFile{}, errs.FileNotFound
		}
		return entity.HomeworkFile{}, err
	}

	return file.toHomeworkFile(), err
}
