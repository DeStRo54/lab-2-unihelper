package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Host string `envconfig:"REDIS_HOST"`
	Port string `envconfig:"REDIS_PORT"`
}

func Connect(cfg *RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Host + ":" + cfg.Port,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	return client
}
