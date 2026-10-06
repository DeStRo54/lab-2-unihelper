package user

import (
	"context"
	"database/sql"
	"errors"

	"homewormanager/internal/entity"
	"homewormanager/internal/errs"
)

func (r *Repo) GetById(ctx context.Context, id entity.UserID) (entity.User, error) {

	q := "SELECT * FROM users WHERE user_id = $1"

	var user user

	t := r.manager.GetTxOrDefault(ctx)

	err := t.GetContext(ctx, &user, q, id)
	if errors.Is(err, sql.ErrNoRows) {
		return entity.User{}, errs.UserNotFound
	}
	if err != nil {
		return entity.User{}, err
	}
	return user.toUser(), nil
}
