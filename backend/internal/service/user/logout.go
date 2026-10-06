package user

import "context"

func (s *Service) Logout(ctx context.Context, sessionKey string) error {
	err := s.userRedisRepo.DeleteCreds(ctx, sessionKey)
	if err != nil {
		return err
	}
	return nil
}
