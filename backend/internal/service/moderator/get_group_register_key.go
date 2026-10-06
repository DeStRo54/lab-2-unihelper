package moderator

import (
	"context"
	"homewormanager/internal/entity"
)

func (s *Service) GetGroupRegisterKey(ctx context.Context, groupID entity.GroupID) (string, error) {
	group, err := s.groupService.GetByID(ctx, groupID)
	if err != nil {
		return "", err
	}
	return group.RegisterKey, nil
}
