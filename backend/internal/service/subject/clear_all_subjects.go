package subject

import "context"

func (s *Service) ClearAllSubjects(ctx context.Context) error {
	err := s.subjectRepo.Clear(ctx)
	if err != nil {
		return err
	}
	return nil
}
