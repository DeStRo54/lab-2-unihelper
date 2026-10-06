package subjectnote

import "context"

func (s *Service) ClearAllNotes(ctx context.Context) error {
	err := s.subjectNoteRepo.Clear(ctx)
	if err != nil {
		return err
	}
	return nil
}
