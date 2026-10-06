package class

import (
	"context"
	"strings"

	"homewormanager/internal/entity"
)

type classCount struct {
	entity.ClassCategory
	entity.SubjectID
}

func (s *Service) UpdGroupClasses(ctx context.Context, group entity.Group, classes []entity.Class) error {
	var err error

	classCounter := make(map[classCount]int64)

	for i := range classes {

		subject, err := s.subjectService.GetBySubjectNameAndGroup(ctx, strings.Join(strings.Split(classes[i].Summary, " ")[1:], " "), group.GroupID)
		if err != nil {
			return err
		}

		classC := classCount{
			ClassCategory: classes[i].Category,
			SubjectID:     subject.SubjectID,
		}

		classCounter[classC]++

		classes[i].SubjectID = subject.SubjectID
		classes[i].GroupID = group.GroupID
		classes[i].SemClassNumber = classCounter[classC]
	}

	err = s.classRepo.Create(ctx, classes)

	if err != nil {
		return err
	}

	return nil

}
