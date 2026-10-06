package user

import (
	"context"
)

func (r *Repo) DeleteCreds(ctx context.Context, sessionKey string) error {
	err := r.client.Del(ctx, sessionKey).Err()
	if err != nil {
		return err
	}
	return nil
}
