package group

import (
	"context"
	"database/sql"
	"errors"
	"homewormanager/internal/entity"
	"homewormanager/internal/errs"
)

func (r *Repo) GetByID(ctx context.Context, groupID entity.GroupID) (entity.Group, error) {

	q := "SELECT * FROM groups WHERE group_id = $1"

	var group group

	t := r.manager.GetTxOrDefault(ctx)

	err := t.GetContext(ctx, &group, q, groupID)
	if errors.Is(err, sql.ErrNoRows) {
		return entity.Group{}, errs.GroupNotFound
	}
	if err != nil {
		return entity.Group{}, err
	}
	return group.toGroup(), nil
}
