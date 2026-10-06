package subjectnote

import (
	"context"
	"homewormanager/internal/entity"
)

func (s *Service) AddNote(ctx context.Context, note entity.SubjectNote) (entity.NoteID, error) {
	id, err := s.subjectNoteRepo.Create(ctx, note)
	if err != nil {
		return 0, err
	}
	return id, nil
}
