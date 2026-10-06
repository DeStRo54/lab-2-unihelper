package user

import (
	"context"
	"encoding/json"
	"fmt"
	"homewormanager/internal/entity"
)

func (r *Repo) CreateUserRole(ctx context.Context, id entity.UserID, role entity.Role) error {

	roleBytes, err := json.Marshal(role)
	if err != nil {
		return err
	}

	if err := r.client.Set(ctx, fmt.Sprintf("user_id:%d", id), roleBytes, r.config.AuthTTL).Err(); err != nil {
		return err
	}

	return nil
}
