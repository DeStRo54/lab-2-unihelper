package homework_files

import (
	"context"
	"homewormanager/internal/entity"
)

func (s *Service) DeleteFile(ctx context.Context, fileID entity.FileID) error {
	homeworkFile, err := s.homeworkFilesRepo.GetByFileID(ctx, fileID)
	if err != nil {
		return err
	}

	err = s.manager.Do(ctx, func(ctx context.Context) error {
		err = s.homeworkFilesRepo.Delete(ctx, homeworkFile.FileID)
		if err != nil {
			return err
		}

		err = s.homeworkFilesS3.Delete(ctx, homeworkFile.Key)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
