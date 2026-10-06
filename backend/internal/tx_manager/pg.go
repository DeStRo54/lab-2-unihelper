package tx_manager

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type DBOps interface {
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type TxWrapper struct {
	Tx *sqlx.Tx
}

func (t TxWrapper) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return t.Tx.SelectContext(ctx, dest, query, args...)
}

func (t TxWrapper) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return t.Tx.GetContext(ctx, dest, query, args...)
}

func (t TxWrapper) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return t.Tx.ExecContext(ctx, query, args...)
}

func (t TxWrapper) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return t.Tx.QueryRowContext(ctx, query, args...)
}
