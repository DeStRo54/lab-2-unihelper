package group

import (
	"context"
)

func (s *Service) RegenerateRegisterKeys(ctx context.Context) error {
	//TODO: добавить тест на функцию
	groups, err := s.groupRepo.GetAllGroups(ctx)
	if err != nil {
		return err
	}
	for i := range groups {
		_, err = s.RegenerateGroupRegisterKey(ctx, groups[i].GroupID)
		if err != nil {
			return err
		}
	}
	return nil
}
