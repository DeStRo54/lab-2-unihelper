package moderator

import (
	"context"
	"mime/multipart"
	"time"

	"homewormanager/internal/entity"
)

type AddHomework struct {
	ClassSemNumber *int64
	GroupID        entity.GroupID
	SubjectID      entity.SubjectID
	Category       *entity.ClassCategory
	HomeworkText   string
	FilesHeader    []*multipart.FileHeader
	DueDate        time.Time
}

// TODO: добавить проверку на то что такое занятие существует
func (s *Service) AddHomework(ctx context.Context, req AddHomework) (entity.HomeworkID, map[string]entity.FileID, map[entity.FileID]string, map[string]error, error) {
	homework := req.toHomework()
	id, err := s.homeworkService.Create(ctx, homework)
	if err != nil {
		return 0, nil, nil, nil, err
	}

	filesIdMap := make(map[string]entity.FileID)
	filesURLMap := make(map[entity.FileID]string)
	filesErrMap := make(map[string]error)

	//TODO: добавить параллельность
	for _, file := range req.FilesHeader {
		fileID, fileURL, err := s.homeworkFileService.AddFileToHomework(ctx, file, id, req.GroupID)
		if err != nil {
			filesErrMap[file.Filename] = err
			continue
		}
		filesIdMap[file.Filename] = fileID
		filesURLMap[fileID] = fileURL
	}

	return id, filesIdMap, filesURLMap, filesErrMap, nil
}

func (r *AddHomework) toHomework() entity.Homework {
	return entity.Homework{
		SemClassNumber: r.ClassSemNumber,
		GroupID:        r.GroupID,
		SubjectID:      r.SubjectID,
		Category:       r.Category,
		HomeworkText:   r.HomeworkText,
		DueDate:        r.DueDate,
	}
}
