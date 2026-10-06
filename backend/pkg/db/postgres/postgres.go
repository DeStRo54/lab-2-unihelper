package postgres

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
	"strings"
)

type PGConfig struct {
	Host           string `envconfig:"POSTGRES_HOST"`
	Port           string `envconfig:"POSTGRES_PORT"`
	User           string `envconfig:"POSTGRES_USER"`
	Password       string `envconfig:"POSTGRES_PASSWORD"`
	DBName         string `envconfig:"POSTGRES_DB_NAME"`
	SSLMode        string `envconfig:"POSTGRES_SSLMODE"`
	ConnectTimeout string `envconfig:"POSTGRES_CONNECT_TIMEOUT"`
}

func Connect(cfg *PGConfig) *sqlx.DB {
	password := cfg.Password
	if passwordFile := os.Getenv("POSTGRES_PASSWORD_FILE"); passwordFile != "" {
		secret, err := os.ReadFile(passwordFile)
		if err != nil {
			panic(fmt.Errorf("read database password secret: %w", err))
		}
		password = strings.TrimSpace(string(secret))
	}
	if password == "" {
		panic(errors.New("database password is not configured"))
	}
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=%s connect_timeout=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.DBName,
		cfg.SSLMode,
		cfg.ConnectTimeout,
	)

	if password != "" {
		dsn += fmt.Sprintf(" password=%s", password)
	}

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return db
}
