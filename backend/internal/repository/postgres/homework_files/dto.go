package homework_files

import (
	"homewormanager/internal/entity"
	"time"
)

type homeworkFile struct {
	FileID     entity.FileID     `db:"file_id"`
	HomeworkID entity.HomeworkID `db:"homework_id"`
	GroupID    entity.GroupID    `db:"group_id"`
	FileName   string            `db:"file_name"`
	FileURL    string            `db:"file_url"`
	Key        string            `db:"key"`
	CreatedAt  time.Time         `db:"uploaded_at"`
}

func (h homeworkFile) toHomeworkFile() entity.HomeworkFile {
	return entity.HomeworkFile{
		FileID:     h.FileID,
		HomeworkID: h.HomeworkID,
		GroupID:    h.GroupID,
		FileName:   h.FileName,
		FileURL:    h.FileURL,
		Key:        h.Key,
		CreatedAt:  h.CreatedAt,
	}
}
func toHomeworkFiles(f []homeworkFile) []entity.HomeworkFile {
	fs := make([]entity.HomeworkFile, len(f))
	for i := range f {
		fs[i] = f[i].toHomeworkFile()
	}
	return fs
}
