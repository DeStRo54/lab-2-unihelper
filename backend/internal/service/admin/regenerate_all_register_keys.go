package admin

import "context"

func (s *Service) RegenerateAllRegisterKeys(ctx context.Context) error {
	err := s.groupService.RegenerateRegisterKeys(ctx)
	if err != nil {
		return err
	}
	return nil
}
