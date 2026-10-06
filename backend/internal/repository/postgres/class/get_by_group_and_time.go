package class

import (
	"context"
	"time"

	"homewormanager/internal/entity"
	"homewormanager/internal/errs"
)

func (r *Repo) GetByGroupAndTime(ctx context.Context, groupID entity.GroupID, fromTime time.Time, toTime time.Time) ([]entity.Class, error) {
	q := "SELECT * FROM classes WHERE group_id = $1 AND start_time >= $2 AND end_time < $3"

	var classes []class

	t := r.manager.GetTxOrDefault(ctx)

	err := t.SelectContext(ctx, &classes, q, groupID, fromTime, toTime)
	if len(classes) == 0 {
		return []entity.Class{}, errs.ClassesNotFound
	}
	if err != nil {
		return []entity.Class{}, err
	}

	return toClasses(classes), nil
}

func toClasses(classes []class) []entity.Class {
	res := make([]entity.Class, len(classes))
	for i, g := range classes {
		res[i] = g.toClass()
	}
	return res
}
