package moderator

import (
	"context"
	"homewormanager/internal/entity"
)

func (s *Service) DeleteFileFromHomework(ctx context.Context, fileID entity.FileID) error {
	err := s.homeworkFileService.DeleteFile(ctx, fileID)
	if err != nil {
		return err
	}
	return nil
}
