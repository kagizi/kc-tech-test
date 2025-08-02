package entity

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Wallet struct {
	ID        uuid.UUID       `json:"id" db:"id"`
	UserID    uuid.UUID       `json:"user_id" db:"user_id"`
	Balance   decimal.Decimal `json:"balance" db:"balance"`
	CreatedAt time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt time.Time       `json:"updated_at" db:"updated_at"`
}

func NewWallet(userID uuid.UUID) (*Wallet, error) {
	id, err := uuid.NewV7()

	if err != nil {
		return nil, fmt.Errorf("failed to generate UUID: %w", err)
	}

	now := time.Now()

	wallet := &Wallet{
		ID:        id,
		UserID:    userID,
		Balance:   decimal.Zero,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := wallet.validate(); err != nil {
		return nil, err
	}

	return wallet, nil
}

func (w *Wallet) validate() error {
	if w.Balance.IsNegative() {
		return errors.New("balance cannot be negative")
	}

	return nil
}

func (w *Wallet) Withdraw(amount decimal.Decimal) error {
	if amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("withdrawal amount must be greater than zero")
	}

	if w.Balance.LessThan(amount) {
		return errors.New("insufficient funds")
	}

	w.Balance = w.Balance.Sub(amount)
	w.UpdatedAt = time.Now()

	return w.validate()
}

func (w *Wallet) Deposit(amount decimal.Decimal) error {
	if amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("deposit amount must be greater than zero")
	}

	w.Balance = w.Balance.Add(amount)
	w.UpdatedAt = time.Now()

	return w.validate()
}
