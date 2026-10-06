package user

import (
	"github.com/go-redis/redis/v8"

	"homewormanager/internal/config"
)

type Repo struct {
	client *redis.Client
	config *config.Config
}

func NewUserRepo(client *redis.Client, config *config.Config) *Repo {
	return &Repo{
		client: client,
		config: config,
	}
}
