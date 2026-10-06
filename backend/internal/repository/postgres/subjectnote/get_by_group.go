package subjectnote

import (
	"context"
	"homewormanager/internal/entity"
)

func (r *Repo) GetByGroup(ctx context.Context, groupID entity.GroupID) ([]entity.SubjectNote, error) {
	q := "SELECT * FROM subjectnotes WHERE group_id = $1"

	var subjectNotes []subjectNote

	t := r.manager.GetTxOrDefault(ctx)

	err := t.SelectContext(ctx, &subjectNotes, q, groupID)
	if err != nil {
		return []entity.SubjectNote{}, err
	}

	return toSubjectNotes(subjectNotes), nil
}

func toSubjectNotes(notes []subjectNote) []entity.SubjectNote {
	res := make([]entity.SubjectNote, len(notes))
	for i, g := range notes {
		res[i] = g.toSubjectNote()
	}
	return res
}
