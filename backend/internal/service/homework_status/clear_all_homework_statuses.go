package homework_status

import "context"

func (s *Service) ClearAllHomeworkStatuses(ctx context.Context) error {
	err := s.homeworkStatusRepo.Clear(ctx)
	if err != nil {
		return err
	}
	return nil
}
