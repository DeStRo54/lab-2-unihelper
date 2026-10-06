package subject

import (
	"context"
	"database/sql"
	"errors"

	"homewormanager/internal/entity"
	"homewormanager/internal/errs"
)

func (r *Repo) GetByNameAndGroup(ctx context.Context, name string, groupID entity.GroupID) (entity.Subject, error) {

	q := "SELECT * FROM subjects WHERE subject_name = $1 and group_id = $2"

	var subject subject

	t := r.manager.GetTxOrDefault(ctx)

	err := t.GetContext(ctx, &subject, q, name, groupID)
	if errors.Is(err, sql.ErrNoRows) {
		return entity.Subject{}, errs.SubjectsNotFound
	}
	if err != nil {
		return entity.Subject{}, err
	}

	return subject.toSubject(), nil

}
