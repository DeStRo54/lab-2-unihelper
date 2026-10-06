package group

import (
	"context"
	"errors"
	"github.com/google/uuid"

	"homewormanager/internal/entity"
	"homewormanager/internal/errs"
)

func (s *Service) Create(ctx context.Context, group entity.Group) (entity.GroupID, error) {

	_, err := s.groupRepo.GetByName(ctx, group.Name)

	if err == nil {
		return 0, errs.GroupExists
	}
	if !(errors.Is(err, errs.GroupNotFound)) {
		return 0, err
	}

	group.RegisterKey = uuid.NewString()[:8] //TODO:вынести это в env

	id, err := s.groupRepo.Create(ctx, group)
	if err != nil {
		return 0, err
	}

	return id, nil

}
