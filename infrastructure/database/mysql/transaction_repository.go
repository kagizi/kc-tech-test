package mysql

import (
	"context"
	"database/sql"
	"github.com/kagizi/kc-tech-test/domain/entity"
	"github.com/kagizi/kc-tech-test/domain/repository"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) repository.TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) CreateWithTx(ctx context.Context, tx *sql.Tx, transaction *entity.Transaction) error {
	query := `
        INSERT INTO transactions (id, wallet_id, amount, type, created_at)
        VALUES (UUID_TO_BIN(?), UUID_TO_BIN(?), ?, ?, ?)
    `
	_, err := tx.ExecContext(ctx, query,
		transaction.ID.String(),
		transaction.WalletID.String(),
		transaction.Amount,
		transaction.Type,
		transaction.CreatedAt,
	)
	return err
}
