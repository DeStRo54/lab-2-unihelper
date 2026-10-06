package subject

import (
	"context"

	"homewormanager/internal/entity"
)

func (s *Service) GetBySubjectNameAndGroup(ctx context.Context, subjectName string, groupId entity.GroupID) (entity.Subject, error) {
	subject, err := s.subjectRepo.GetByNameAndGroup(ctx, subjectName, groupId)
	if err != nil {
		return subject, err
	}
	return subject, nil
}
