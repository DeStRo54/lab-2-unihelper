package homework_files

import (
	"context"
	"homewormanager/internal/entity"
)

func (r *Repo) Create(ctx context.Context, HomeworkFile entity.HomeworkFile) (entity.FileID, error) {
	q := "INSERT INTO homeworksfiles (homework_id, group_id, file_name, file_url, key) values ($1, $2, $3, $4, $5) RETURNING file_id"

	t := r.manager.GetTxOrDefault(ctx)

	var id entity.FileID

	err := t.QueryRowContext(ctx, q, HomeworkFile.HomeworkID, HomeworkFile.GroupID, HomeworkFile.FileName, HomeworkFile.FileURL, HomeworkFile.Key).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}
