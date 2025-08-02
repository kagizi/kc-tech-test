package mysql

import (
	"context"
	"database/sql"
)

type TxDB struct {
	*sql.DB
}

func NewTxDB(db *sql.DB) *TxDB {
	return &TxDB{DB: db}
}

func (db *TxDB) ExecTX(ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
