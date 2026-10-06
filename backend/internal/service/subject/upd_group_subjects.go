package subject

import (
	"context"

	"homewormanager/internal/entity"
)

func (s *Service) UpdGroupSubjects(ctx context.Context, group entity.Group, subjects []string) error {

	for _, subject := range subjects {
		err := s.subjectRepo.Create(ctx, entity.Subject{
			GroupID:     group.GroupID,
			SubjectName: subject,
		})
		if err != nil {
			return err
		}
	}
	return nil

}
