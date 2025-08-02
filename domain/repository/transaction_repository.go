package repository

import (
	"context"
	"database/sql"
	"github.com/kagizi/kc-tech-test/domain/entity"
)

type TransactionRepository interface {
	CreateWithTx(ctx context.Context, tx *sql.Tx, transaction *entity.Transaction) error
}
