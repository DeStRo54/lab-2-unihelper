package subject

import "homewormanager/internal/entity"

type subject struct {
	SubjectId   entity.SubjectID `db:"subject_id"`
	GroupId     entity.GroupID   `db:"group_id"`
	SubjectName string           `db:"subject_name"`
}

func (s subject) toSubject() entity.Subject {
	return entity.Subject{
		SubjectID:   s.SubjectId,
		GroupID:     s.GroupId,
		SubjectName: s.SubjectName,
	}
}

func toSubjects(subjects []subject) []entity.Subject {
	res := make([]entity.Subject, len(subjects))
	for i, g := range subjects {
		res[i] = g.toSubject()
	}
	return res
}
