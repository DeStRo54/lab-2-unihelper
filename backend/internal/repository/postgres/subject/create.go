package subject

import (
	"context"

	"homewormanager/internal/entity"
)

func (r *Repo) Create(ctx context.Context, subject entity.Subject) error {

	q := "INSERT INTO subjects(group_id, subject_name) VALUES ($1, $2)"

	t := r.manager.GetTxOrDefault(ctx)

	_, err := t.ExecContext(ctx, q, subject.GroupID, subject.SubjectName)

	if err != nil {
		return err
	}
	return nil

}
