package user

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"

	"homewormanager/internal/entity"
)

func (r *Repo) CreateCreds(ctx context.Context, creds entity.UserCreds) (sessionKey string, err error) {
	sessionKey = uuid.NewString()

	credsBytes, err := json.Marshal(credsData{
		UserID:  creds.UserID,
		GroupID: creds.GroupID,
	})
	if err != nil {
		return "", err
	}

	if err = r.client.Set(ctx, sessionKey, credsBytes, r.config.AuthTTL).Err(); err != nil {
		return "", err
	}

	return
}
