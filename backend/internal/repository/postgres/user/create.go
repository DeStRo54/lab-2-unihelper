package user

import (
	"context"

	"homewormanager/internal/entity"
)

func (r *Repo) Create(ctx context.Context, user entity.User) error {

	t := r.manager.GetTxOrDefault(ctx)
	q := "INSERT INTO users(name, surname, email, password, group_id) VALUES ($1, $2, $3, $4, $5)"

	_, err := t.ExecContext(ctx, q, user.Name, user.Surname, user.Email, user.Password, user.GroupID)
	if err != nil {
		return err
	}

	return nil

}
