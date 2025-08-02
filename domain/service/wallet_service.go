package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/kagizi/kc-tech-test/domain/repository"
	"github.com/shopspring/decimal"
)

type WalletService struct {
	walletRepo repository.WalletRepository
}

func NewWalletService(
	walletRepo repository.WalletRepository,
) *WalletService {
	return &WalletService{
		walletRepo: walletRepo,
	}
}

func (s *WalletService) Withdraw(ctx context.Context, walletID uuid.UUID, amount decimal.Decimal) error {
	return s.walletRepo.WithdrawWithTransaction(ctx, walletID, amount)
}

func (s *WalletService) GetBalance(ctx context.Context, walletID uuid.UUID) (decimal.Decimal, error) {
	wallet, err := s.walletRepo.GetByID(ctx, walletID)
	if err != nil {
		return decimal.Zero, fmt.Errorf("failed to get wallet: %w", err)
	}

	return wallet.Balance, nil
}

func (s *WalletService) Deposit(ctx context.Context, walletID uuid.UUID, amount decimal.Decimal) error {
	return s.walletRepo.DepositWithTransaction(ctx, walletID, amount)
}
