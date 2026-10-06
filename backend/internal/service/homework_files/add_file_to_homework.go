package homework_files

import (
	"context"
	"fmt"
	"homewormanager/internal/entity"
	"mime/multipart"
)

func (s *Service) AddFileToHomework(ctx context.Context, fileHeader *multipart.FileHeader, homeworkID entity.HomeworkID, groupID entity.GroupID) (entity.FileID, string, error) {
	key := fmt.Sprintf("homeworks/%d/%s", homeworkID, fileHeader.Filename)
	fileName := fileHeader.Filename
	fileUrl := "https://global.s3.cloud.ru/unihelper/" + key
	homeworkFile := entity.HomeworkFile{
		HomeworkID: homeworkID,
		FileName:   fileName,
		GroupID:    groupID,
		FileURL:    fileUrl,
		Key:        key,
	}

	var id entity.FileID

	err := s.manager.Do(ctx, func(ctx context.Context) error {
		var err error
		id, err = s.homeworkFilesRepo.Create(ctx, homeworkFile)
		if err != nil {
			return err
		}
		err = s.homeworkFilesS3.Upload(ctx, fileHeader, key)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return 0, "", err
	}

	return id, fileUrl, err
}
