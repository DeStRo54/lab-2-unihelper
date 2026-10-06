package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"homewormanager/internal/entity"
	"homewormanager/internal/errs"
)

func (r *Repo) GetUserRole(ctx context.Context, id entity.UserID) (entity.Role, error) {

	var role entity.Role

	bytes, err := r.client.Get(ctx, fmt.Sprintf("user_id:%d", id)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, errs.NoUserRole
		}
		return 0, err
	}

	if err = json.Unmarshal([]byte(bytes), &role); err != nil {
		return 0, err
	}

	return role, nil

}
