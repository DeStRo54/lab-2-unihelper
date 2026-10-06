package homework

import "context"

func (s *Service) ClearAllHomeworks(ctx context.Context) error {
	err := s.homeworkRepo.Clear(ctx)
	if err != nil {
		return err
	}
	return nil
}
