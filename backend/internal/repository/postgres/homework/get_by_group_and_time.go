package homework

import (
	"context"
	"time"

	"homewormanager/internal/entity"
	"homewormanager/internal/errs"
)

func (r *Repo) GetByGroupAndTime(ctx context.Context, userID entity.UserID, groupID entity.GroupID, fromTime time.Time, toTime time.Time) ([]entity.Homework, error) {
	q := "WITH tmp AS(SELECT h.*," +
		"COALESCE(EXISTS (" +
		"SELECT 1 FROM homeworkstatuses AS hs " +
		"WHERE hs.homework_id = h.homework_id AND hs.user_id = $4" +
		"), TRUE) AS Status " +
		"FROM homeworks AS h " +
		"WHERE h.group_id = $1 " +
		"AND h.due_date >= $2 " +
		"AND h.due_date <= $3) " +
		"SELECT tmp.*, subject_name FROM tmp JOIN subjects ON tmp.subject_id = subjects.subject_id ORDER BY due_date"

	var homeworks []homework

	t := r.manager.GetTxOrDefault(ctx)

	err := t.SelectContext(ctx, &homeworks, q, groupID, fromTime, toTime, userID)
	if len(homeworks) == 0 {
		return []entity.Homework{}, errs.HomeworksNotFound
	}
	if err != nil {
		return []entity.Homework{}, err
	}

	return toHomeworks(homeworks), nil
}

func toHomeworks(homeworks []homework) []entity.Homework {
	res := make([]entity.Homework, len(homeworks))
	for i, g := range homeworks {
		res[i] = g.toHomework()
	}
	return res
}
