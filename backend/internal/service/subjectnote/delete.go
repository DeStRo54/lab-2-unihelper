package subjectnote

import (
	"context"
	"homewormanager/internal/entity"
)

func (s *Service) Delete(ctx context.Context, noteID entity.NoteID) error {
	err := s.subjectNoteRepo.Delete(ctx, noteID)
	if err != nil {
		return err
	}
	return nil
}
