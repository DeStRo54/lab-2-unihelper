package group

import (
	"context"
	"github.com/google/uuid"
	"homewormanager/internal/entity"
)

func (s *Service) RegenerateGroupRegisterKey(ctx context.Context, groupID entity.GroupID) (string, error) {
	group := entity.Group{
		GroupID:     groupID,
		RegisterKey: uuid.NewString()[:8],
	}
	err := s.groupRepo.ChangeRegisterKey(ctx, group)
	if err != nil {
		return "", err
	}
	return group.RegisterKey, nil
}
