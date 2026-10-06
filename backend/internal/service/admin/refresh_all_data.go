package admin

import (
	"context"

	"homewormanager/internal/utils/classParse"
)

// TODO: добавить отчищение статуса выполненности
func (s *Service) RefreshAllData(ctx context.Context) error {

	groups, err := s.groupService.GetAllGroups(ctx)
	if err != nil {
		return err
	}

	err = s.manager.Do(ctx, func(ctx context.Context) error {

		err = s.subjectNoteService.ClearAllNotes(ctx)
		if err != nil {
			return err
		}

		err = s.homeworkStatusService.ClearAllHomeworkStatuses(ctx)
		if err != nil {
			return err
		}

		err = s.homeworkService.ClearAllHomeworks(ctx)
		if err != nil {
			return err
		}

		err = s.classService.ClearAllClasses(ctx)
		if err != nil {
			return err
		}

		err = s.subjectService.ClearAllSubjects(ctx)
		if err != nil {
			return err
		}

		for _, group := range groups {

			classes, subjects, err := classParse.IcalParse(group.IcalLink)
			if err != nil {
				return err
			}

			err = s.subjectService.UpdGroupSubjects(ctx, group, subjects)
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
