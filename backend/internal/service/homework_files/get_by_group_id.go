package homework_files

import (
	"context"
	"homewormanager/internal/entity"
)

func (s *Service) GetByGroupID(ctx context.Context, groupID entity.GroupID) ([]entity.HomeworkFile, error) {
	files, err := s.homeworkFilesRepo.GetByGroupID(ctx, groupID)
	if err != nil {
		return nil, err
	}
	return files, nil
}
