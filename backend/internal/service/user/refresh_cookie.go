package user

import (
	"context"
)

func (s *Service) RefreshCookie(ctx context.Context, sessionKey string) error {
	_, err := s.userRedisRepo.GetCreds(ctx, sessionKey)
	if err != nil {
		return err
	}

	err = s.userRedisRepo.RefreshCreds(ctx, sessionKey)
	if err != nil {
		return err
	}
	return nil
}
