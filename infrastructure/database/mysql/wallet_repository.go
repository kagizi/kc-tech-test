package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kagizi/kc-tech-test/domain/entity"
	"github.com/kagizi/kc-tech-test/domain/repository"
	"github.com/shopspring/decimal"
)

type WalletRepository struct {
	db              *sql.DB
	txDB            *TxDB
	transactionRepo repository.TransactionRepository
}

func NewWalletRepository(db *sql.DB, transactionRepo repository.TransactionRepository) repository.WalletRepository {
	return &WalletRepository{
		db:              db,
		txDB:            NewTxDB(db),
		transactionRepo: transactionRepo,
	}
}

func (r *WalletRepository) Create(ctx context.Context, wallet *entity.Wallet) error {
	query := `
        INSERT INTO wallets (id, user_id, balance, created_at, updated_at)
        VALUES (UUID_TO_BIN(?), UUID_TO_BIN(?), ?, ?, ?)
    `
	_, err := r.db.ExecContext(ctx, query,
		wallet.ID.String(),
		wallet.UserID.String(),
		wallet.Balance,
		wallet.CreatedAt,
		wallet.UpdatedAt,
	)
	return err
}

func (r *WalletRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.Wallet, error) {
	query := `
        SELECT 
            BIN_TO_UUID(id), 
            BIN_TO_UUID(user_id), 
            balance, 
            created_at, 
            updated_at
        FROM wallets
        WHERE id = UUID_TO_BIN(?)
    `
	row := r.db.QueryRowContext(ctx, query, id.String())

	var wallet entity.Wallet
	var idStr, userIDStr string
	err := row.Scan(
		&idStr,
		&userIDStr,
		&wallet.Balance,
		&wallet.CreatedAt,
		&wallet.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("wallet not found")
		}
		return nil, err
	}

	wallet.ID, err = uuid.Parse(idStr)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID format: %w", err)
	}

	wallet.UserID, err = uuid.Parse(userIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID format: %w", err)
	}

	return &wallet, nil
}

func (r *WalletRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*entity.Wallet, error) {
	query := `
        SELECT 
            BIN_TO_UUID(id), 
            BIN_TO_UUID(user_id), 
            balance, 
            created_at, 
            updated_at
        FROM wallets
        WHERE user_id = UUID_TO_BIN(?)
    `
	row := r.db.QueryRowContext(ctx, query, userID.String())

	var wallet entity.Wallet
	var idStr, userIDStr string
	err := row.Scan(
		&idStr,
		&userIDStr,
		&wallet.Balance,
		&wallet.CreatedAt,
		&wallet.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("wallet not found")
		}
		return nil, err
	}

	wallet.ID, err = uuid.Parse(idStr)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID format: %w", err)
	}

	wallet.UserID, err = uuid.Parse(userIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID format: %w", err)
	}

	return &wallet, nil
}

func (r *WalletRepository) WithdrawWithTransaction(ctx context.Context, walletID uuid.UUID, amount decimal.Decimal) error {
	return r.txDB.ExecTX(ctx, func(tx *sql.Tx) error {
		wallet, err := r.getWalletByIDWithTx(ctx, tx, walletID)
		if err != nil {
			return fmt.Errorf("failed to get wallet: %w", err)
		}

		if err := wallet.Withdraw(amount); err != nil {
			return err
		}

		if err := r.updateWalletWithTx(ctx, tx, wallet); err != nil {
			return fmt.Errorf("failed to update wallet: %w", err)
		}

		transaction, err := entity.NewTransaction(walletID, amount, entity.TransactionTypeWithdrawal)
		if err != nil {
			return fmt.Errorf("failed to create transaction: %w", err)
		}

		if err := r.transactionRepo.CreateWithTx(ctx, tx, transaction); err != nil {
			return fmt.Errorf("failed to create transaction: %w", err)
		}

		return nil
	})
}

func (r *WalletRepository) DepositWithTransaction(ctx context.Context, walletID uuid.UUID, amount decimal.Decimal) error {
	return r.txDB.ExecTX(ctx, func(tx *sql.Tx) error {
		wallet, err := r.getWalletByIDWithTx(ctx, tx, walletID)
		if err != nil {
			return fmt.Errorf("failed to get wallet: %w", err)
		}

		if err := wallet.Deposit(amount); err != nil {
			return err
		}

		if err := r.updateWalletWithTx(ctx, tx, wallet); err != nil {
			return fmt.Errorf("failed to update wallet: %w", err)
		}

		transaction, err := entity.NewTransaction(walletID, amount, entity.TransactionTypeDeposit)
		if err != nil {
			return fmt.Errorf("failed to create transaction: %w", err)
		}

		if err := r.transactionRepo.CreateWithTx(ctx, tx, transaction); err != nil {
			return fmt.Errorf("failed to create transaction: %w", err)
		}

		return nil
	})
}

func (r *WalletRepository) getWalletByIDWithTx(ctx context.Context, tx *sql.Tx, id uuid.UUID) (*entity.Wallet, error) {
	query := `
        SELECT 
            BIN_TO_UUID(id), 
            BIN_TO_UUID(user_id), 
            balance, 
            created_at, 
            updated_at
        FROM wallets
        WHERE id = UUID_TO_BIN(?)
    `
	row := tx.QueryRowContext(ctx, query, id.String())

	var wallet entity.Wallet
	var idStr, userIDStr string
	err := row.Scan(
		&idStr,
		&userIDStr,
		&wallet.Balance,
		&wallet.CreatedAt,
		&wallet.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("wallet not found")
		}
		return nil, err
	}

	wallet.ID, err = uuid.Parse(idStr)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID format: %w", err)
	}

	wallet.UserID, err = uuid.Parse(userIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID format: %w", err)
	}

	return &wallet, nil
}

func (r *WalletRepository) updateWalletWithTx(ctx context.Context, tx *sql.Tx, wallet *entity.Wallet) error {
	query := `
        UPDATE wallets
        SET balance = ?, updated_at = ?
        WHERE id = UUID_TO_BIN(?)
    `
	result, err := tx.ExecContext(ctx, query,
		wallet.Balance,
		time.Now(),
		wallet.ID.String(),
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("wallet not found")
	}

	return nil
}
