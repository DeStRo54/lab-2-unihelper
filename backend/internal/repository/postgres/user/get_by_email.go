package user

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"homewormanager/internal/entity"
	"homewormanager/internal/errs"
)

func (r *Repo) GetByEmail(ctx context.Context, email string) (entity.User, error) {

	q := "SELECT * FROM users WHERE email = $1"

	var user user

	t := r.manager.GetTxOrDefault(ctx)

	err := t.GetContext(ctx, &user, q, email)
	if errors.Is(err, sql.ErrNoRows) {
		return entity.User{}, errs.UserNotFound
	}
	if err != nil {
		return entity.User{}, err
	}
	return user.toUser(), nil
}
