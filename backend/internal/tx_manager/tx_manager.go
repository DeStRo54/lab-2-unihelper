package tx_manager

import (
	"context"

	"github.com/jmoiron/sqlx"
)

const txkey = "manager.txkey"

type TxManager struct {
	db *sqlx.DB
}

func NewTxManager(db *sqlx.DB) *TxManager {
	return &TxManager{db: db}
}

func (m *TxManager) injectTx(ctx context.Context, tx *sqlx.Tx) context.Context {
	return context.WithValue(ctx, txkey, tx)
}

func (m *TxManager) extractTx(ctx context.Context) (*sqlx.Tx, bool) {
	tx, ok := ctx.Value(txkey).(*sqlx.Tx)
	return tx, ok
}

func (m *TxManager) GetTxOrDefault(ctx context.Context) DBOps {
	if tx, ok := m.extractTx(ctx); ok {
		return TxWrapper{Tx: tx}
	}
	return m.db
}

func (m *TxManager) Do(ctx context.Context, f func(ctx context.Context) error) error {
	tx, err := m.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		tx.Commit()
	}()
	if err := f(m.injectTx(ctx, tx)); err != nil {
		tx.Rollback()
		return err
	}
	return nil

}
