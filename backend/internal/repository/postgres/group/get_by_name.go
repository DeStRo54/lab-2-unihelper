package group

import (
	"context"
	"database/sql"
	"errors"

	"homewormanager/internal/entity"
	"homewormanager/internal/errs"
)

func (r *Repo) GetByName(ctx context.Context, name string) (entity.Group, error) {

	q := "SELECT * FROM groups WHERE groups.group_name = $1"

	var group group

	t := r.manager.GetTxOrDefault(ctx)

	err := t.GetContext(ctx, &group, q, name)
	if errors.Is(err, sql.ErrNoRows) {
		return entity.Group{}, errs.GroupNotFound
	}
	if err != nil {
		return entity.Group{}, err
	}
	return group.toGroup(), nil
}
