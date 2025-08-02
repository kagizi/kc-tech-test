package dto

import "github.com/shopspring/decimal"

type DepositRequest struct {
	WalletID string          `json:"wallet_id"`
	Amount   decimal.Decimal `json:"amount"`
}

type WithdrawRequest struct {
	WalletID string          `json:"wallet_id"`
	Amount   decimal.Decimal `json:"amount"`
}

type DepositResponse struct {
	Message    string          `json:"message"`
	NewBalance decimal.Decimal `json:"balance,omitempty"`
}

type WithdrawResponse struct {
	Message    string          `json:"message"`
	NewBalance decimal.Decimal `json:"balance,omitempty"`
}

type BalanceResponse struct {
	WalletID string          `json:"wallet_id"`
	Balance  decimal.Decimal `json:"balance"`
}
