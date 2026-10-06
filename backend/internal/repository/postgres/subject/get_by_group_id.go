package subject

import (
	"context"

	"homewormanager/internal/entity"
	"homewormanager/internal/errs"
)

func (r *Repo) GetByGroupId(ctx context.Context, groupId entity.GroupID) ([]entity.Subject, error) {

	q := "SELECT * FROM subjects WHERE group_id = $1"

	var subjects []subject

	t := r.manager.GetTxOrDefault(ctx)

	err := t.SelectContext(ctx, &subjects, q, groupId)
	if len(subjects) == 0 {
		return []entity.Subject{}, errs.SubjectsNotFound
	}
	if err != nil {
		return []entity.Subject{}, err
	}

	return toSubjects(subjects), nil

}
