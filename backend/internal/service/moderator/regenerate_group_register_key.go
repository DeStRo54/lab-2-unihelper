package moderator

import (
	"context"
	"homewormanager/internal/entity"
)

func (s *Service) RegenerateGroupRegisterKey(ctx context.Context, groupID entity.GroupID) (string, error) {
	key, err := s.groupService.RegenerateGroupRegisterKey(ctx, groupID)
	if err != nil {
		return "", err
	}
	return key, nil
}
