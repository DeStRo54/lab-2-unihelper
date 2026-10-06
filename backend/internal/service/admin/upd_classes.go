package admin

import (
	"context"

	"homewormanager/internal/utils/classParse"
)

func (s *Service) UpdateClasses(ctx context.Context) error {

	groups, err := s.groupService.GetAllGroups(ctx)
	if err != nil {
		return err
	}

	err = s.manager.Do(ctx, func(ctx context.Context) error {
		err = s.classService.ClearAllClasses(ctx)
		if err != nil {
			return err
		}
		for _, group := range groups {

			classes, _, err := classParse.IcalParse(group.IcalLink)
			if err != nil {
				return err
			}

			err = s.classService.UpdGroupClasses(ctx, group, classes)

			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
