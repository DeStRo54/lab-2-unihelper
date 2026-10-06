package homework_files

import (
	"context"
	"homewormanager/internal/entity"
)

func (s *Service) GetByHomeworkID(ctx context.Context, id entity.HomeworkID) ([]entity.HomeworkFile, error) {
	files, err := s.homeworkFilesRepo.GetByHomeworkID(ctx, id)
	if err != nil {
		return nil, err
	}
	return files, nil
}
