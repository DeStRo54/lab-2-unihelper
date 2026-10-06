package group

import (
	"context"
	"homewormanager/internal/entity"
)

func (r *Repo) ChangeRegisterKey(ctx context.Context, group entity.Group) error {

	q := "UPDATE groups SET register_key=$1 WHERE group_id=$2"

	t := r.manager.GetTxOrDefault(ctx)

	_, err := t.ExecContext(ctx, q, group.RegisterKey, group.GroupID)

	if err != nil {
		return err
	}

	return err

}
