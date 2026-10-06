package user

import (
	"context"
)

func (r *Repo) RefreshCreds(ctx context.Context, sessionKey string) error {
	if err := r.client.Expire(ctx, sessionKey, r.config.AuthTTL).Err(); err != nil {
		return err
	}
	return nil
}
