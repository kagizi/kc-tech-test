package entity

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type TransactionType string

const (
	TransactionTypeWithdrawal TransactionType = "WITHDRAWAL"
	TransactionTypeDeposit    TransactionType = "DEPOSIT"
)

type Transaction struct {
	ID        uuid.UUID       `json:"id" db:"id"`
	WalletID  uuid.UUID       `json:"wallet_id" db:"wallet_id"`
	Amount    decimal.Decimal `json:"amount" db:"amount"`
	Type      TransactionType `json:"type" db:"type"`
	CreatedAt time.Time       `json:"created_at" db:"created_at"`
}

func NewTransaction(walletID uuid.UUID, amount decimal.Decimal, txType TransactionType) (*Transaction, error) {
	id, err := uuid.NewV7()

	if err != nil {
		return nil, fmt.Errorf("failed to generate UUID: %w", err)
	}

	tx := &Transaction{
		ID:        id,
		WalletID:  walletID,
		Amount:    amount,
		Type:      txType,
		CreatedAt: time.Now(),
	}

	if err := tx.validate(); err != nil {
		return nil, err
	}

	return tx, nil
}

func (t *Transaction) validate() error {
	if t.Amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("amount must be greater than zero")
	}

	if t.Type != TransactionTypeWithdrawal && t.Type != TransactionTypeDeposit {
		return errors.New("invalid transaction type")
	}

	return nil
}
