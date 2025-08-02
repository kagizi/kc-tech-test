package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/kagizi/kc-tech-test/domain/entity"
	"github.com/shopspring/decimal"
)

type WalletRepository interface {
	Create(ctx context.Context, wallet *entity.Wallet) error
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Wallet, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (*entity.Wallet, error)
	WithdrawWithTransaction(ctx context.Context, walletID uuid.UUID, amount decimal.Decimal) error
	DepositWithTransaction(ctx context.Context, walletID uuid.UUID, amount decimal.Decimal) error
}
