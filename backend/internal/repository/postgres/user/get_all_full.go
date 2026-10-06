package user

import (
	"context"
	"database/sql"
	"errors"
	"homewormanager/internal/entity"
	"homewormanager/internal/errs"
)

func (r *Repo) GetAllFull(ctx context.Context) ([]entity.UserFullInfo, error) {

	q := "SELECT user_id, name, surname, email, role, group_name FROM users JOIN groups ON users.group_id = groups.group_id "

	var users []userFull

	t := r.manager.GetTxOrDefault(ctx)

	err := t.SelectContext(ctx, &users, q)
	if errors.Is(err, sql.ErrNoRows) {
		return []entity.UserFullInfo{}, errs.UserNotFound
	}
	if err != nil {
		return []entity.UserFullInfo{}, err
	}
	return toUsersFull(users), nil
}
