package service

import (
	"context"
	"fmt"
	"github.com/kagizi/kc-tech-test/domain/entity"
	"github.com/kagizi/kc-tech-test/domain/repository"
	"github.com/shopspring/decimal"
	"math/rand"
)

type UserService struct {
	userRepo   repository.UserRepository
	walletRepo repository.WalletRepository
}

func NewUserService(userRepo repository.UserRepository, walletRepo repository.WalletRepository) *UserService {
	return &UserService{
		userRepo:   userRepo,
		walletRepo: walletRepo,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, name, phone string) (*entity.User, *entity.Wallet, error) {
	existingUser, err := s.userRepo.GetByPhone(ctx, phone)
	if err == nil && existingUser != nil {
		return nil, nil, fmt.Errorf("phone number %s is already registered", phone)
	}

	user, err := entity.NewUser(name, phone)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create user: %w", err)
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, nil, fmt.Errorf("failed to save user: %w", err)
	}

	wallet, err := entity.NewWallet(user.ID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create wallet: %w", err)
	}

	randomBalance := decimal.NewFromFloat(float64(rand.Intn(100000) + 10000))
	wallet.Balance = randomBalance

	if err := s.walletRepo.Create(ctx, wallet); err != nil {
		return nil, nil, fmt.Errorf("failed to save wallet: %w", err)
	}

	return user, wallet, nil
}

func (s *UserService) GetUserByPhone(ctx context.Context, phone string) (*entity.User, *entity.Wallet, error) {
	user, err := s.userRepo.GetByPhone(ctx, phone)
	if err != nil {
		return nil, nil, err
	}

	wallet, err := s.walletRepo.GetByUserID(ctx, user.ID)
	if err != nil {
		return nil, nil, fmt.Errorf("wallet not found for user: %w", err)
	}

	return user, wallet, nil
}
