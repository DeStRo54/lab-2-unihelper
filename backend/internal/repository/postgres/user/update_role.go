package user

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"homewormanager/internal/entity"
	"homewormanager/internal/errs"
)

func (r *Repo) UpdateRole(ctx context.Context, userID entity.UserID, role entity.Role) error {
	q := "UPDATE users SET role = $1 WHERE user_id = $2"

	t := r.manager.GetTxOrDefault(ctx)
	_, err := t.ExecContext(ctx, q, role, userID)
	if errors.Is(err, sql.ErrNoRows) {
		return errs.UserNotFound
	}

	return nil
}
