package config

import (
	"errors"
	"github.com/kelseyhightower/envconfig"
	"homewormanager/internal/client/http/s3"
	"homewormanager/pkg/metrics"
	"time"

	"homewormanager/pkg/db/postgres"
	"homewormanager/pkg/db/redis"
)

type Config struct {
	Host   string `envconfig:"HOST"`
	Port   string `envconfig:"PORT"`
	Domain string `envconfig:"DOMAIN"`

	AuthTTL time.Duration `envconfig:"AUTH_TTL"`

	EncodingMode string `envconfig:"ENCODING_MODE"`

	FrontendHost string `envconfig:"FRONTEND_HOST"`
	FrontendPort string `envconfig:"FRONTEND_PORT"`

	metrics.MetricsConfig
	postgres.PGConfig
	redis.RedisConfig
	s3.S3Config
}

func NewCfg() *Config {
	var cfg Config

	local := time.FixedZone("MSK", 3*60*60)
	time.Local = local

	if err := envconfig.Process("", &cfg); err != nil {
		panic(err)
	}
	if cfg.Host == "" {
		panic(errors.New("cfg is required"))
	}

	return &cfg
}
