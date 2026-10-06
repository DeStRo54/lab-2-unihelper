package group

import (
	"context"

	"homewormanager/internal/entity"
	"homewormanager/internal/errs"
)

func (r *Repo) GetAllGroups(ctx context.Context) ([]entity.Group, error) {

	q := "SELECT * FROM groups"

	var groups []group

	t := r.manager.GetTxOrDefault(ctx)

	err := t.SelectContext(ctx, &groups, q)
	if len(groups) == 0 {
		return []entity.Group{}, errs.GroupNotFound
	}
	if err != nil {
		return []entity.Group{}, err
	}

	return toGroups(groups), nil

}

func toGroups(groups []group) []entity.Group {
	res := make([]entity.Group, len(groups))
	for i, g := range groups {
		res[i] = g.toGroup()
	}
	return res
}
