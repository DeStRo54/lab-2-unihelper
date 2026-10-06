package class

import (
	"context"
)

func (s *Service) ClearAllClasses(ctx context.Context) error {
	err := s.classRepo.Clear(ctx)
	if err != nil {
		return err
	}
	return nil
}
