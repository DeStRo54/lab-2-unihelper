package user

import (
	"context"
	"encoding/json"

	"homewormanager/internal/entity"
)

func (r *Repo) GetCreds(ctx context.Context, sessionKey string) (entity.UserCreds, error) {
	var creds credsData

	credsBytes, err := r.client.Get(ctx, sessionKey).Result()
	if err != nil {
		return entity.UserCreds{}, err
	}

	if err = json.Unmarshal([]byte(credsBytes), &creds); err != nil {
		return entity.UserCreds{}, err
	}

	return creds.toCreds(), nil

}
