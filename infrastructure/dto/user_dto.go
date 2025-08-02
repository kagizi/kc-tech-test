package dto

import (
	"github.com/shopspring/decimal"
	"time"
)

type UserDetailResponse struct {
	UserID    string       `json:"user_id"`
	Name      string       `json:"name"`
	Phone     string       `json:"phone"`
	CreatedAt time.Time    `json:"created_at"`
	Wallet    WalletDetail `json:"wallet"`
}

type WalletDetail struct {
	WalletID  string          `json:"wallet_id"`
	Balance   decimal.Decimal `json:"balance"`
	CreatedAt time.Time       `json:"created_at"`
}

type RegisterRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type RegisterResponse struct {
	UserID   string          `json:"user_id"`
	WalletID string          `json:"wallet_id"`
	Balance  decimal.Decimal `json:"balance"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
