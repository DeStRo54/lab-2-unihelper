package subjectnote

import (
	"context"
	"homewormanager/internal/entity"
)

func (s *Service) Update(ctx context.Context, noteID entity.NoteID, noteText string) error {
	err := s.subjectNoteRepo.Update(ctx, noteID, noteText)
	if err != nil {
		return err
	}
	return nil
}
