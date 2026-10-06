package moderator

import (
	"context"
	"homewormanager/internal/entity"
	moderatorService "homewormanager/internal/service/moderator"
)

type ModeratorService interface {
	AddHomework(ctx context.Context, req moderatorService.AddHomework) (entity.HomeworkID, map[string]entity.FileID, map[entity.FileID]string, map[string]error, error)
	DeleteHomework(ctx context.Context, id entity.HomeworkID) error
	UpdateHomework(ctx context.Context, homeworkId entity.HomeworkID, homeworkText string) error
	AddNote(ctx context.Context, req moderatorService.AddNote) (entity.NoteID, error)
	DeleteNote(ctx context.Context, id entity.NoteID) error
	UpdateNote(ctx context.Context, noteID entity.NoteID, noteText string) error
	RegenerateGroupRegisterKey(ctx context.Context, groupID entity.GroupID) (string, error)
	GetGroupRegisterKey(ctx context.Context, groupID entity.GroupID) (string, error)
	AddFileToHomework(ctx context.Context, req moderatorService.AddFileReq) (map[string]entity.FileID, map[entity.FileID]string, map[string]error, error)
	DeleteFileFromHomework(ctx context.Context, fileID entity.FileID) error
}
