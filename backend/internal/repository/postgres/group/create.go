package group

import (
	"context"

	"homewormanager/internal/entity"
)

func (r *Repo) Create(ctx context.Context, group entity.Group) (entity.GroupID, error) {
	var id entity.GroupID

	q := "INSERT INTO groups(group_name, course, ical_link, register_key) VALUES ($1, $2, $3, $4) RETURNING group_id"

	t := r.manager.GetTxOrDefault(ctx)

	err := t.QueryRowContext(ctx, q, group.Name, group.Course, group.IcalLink, group.RegisterKey).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, err

}
