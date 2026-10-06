package homework_files

import (
	"context"
	"homewormanager/internal/entity"
	"mime/multipart"
)

type HomeworkFilesS3 interface {
	Upload(ctx context.Context, file *multipart.FileHeader, key string) error
	Delete(ctx context.Context, key string) error
}
type HomeworkFilesRepo interface {
	Create(ctx context.Context, HomeworkFile entity.HomeworkFile) (entity.FileID, error)
	GetByFileID(ctx context.Context, fileID entity.FileID) (entity.HomeworkFile, error)
	Delete(ctx context.Context, fileID entity.FileID) error
	GetByGroupID(ctx context.Context, id entity.GroupID) ([]entity.HomeworkFile, error)
	GetByHomeworkID(ctx context.Context, id entity.HomeworkID) ([]entity.HomeworkFile, error)
}
