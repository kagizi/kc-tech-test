package infrastructure

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/kagizi/kc-tech-test/domain/service"
	"github.com/kagizi/kc-tech-test/infrastructure/dto"
	"net/http"
)

type WalletHandler struct {
	walletService *service.WalletService
	userService   *service.UserService
}

func NewWalletHandler(walletService *service.WalletService, userService *service.UserService) *WalletHandler {
	return &WalletHandler{
		walletService: walletService,
		userService:   userService,
	}
}

func (h *WalletHandler) GetUserByPhone(w http.ResponseWriter, r *http.Request) {
	phone := r.URL.Query().Get("phone")
	if phone == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorResponse{Error: "phone required"})
		return
	}

	user, wallet, err := h.userService.GetUserByPhone(r.Context(), phone)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(dto.ErrorResponse{Error: err.Error()})
		return
	}

	resp := dto.UserDetailResponse{
		UserID:    user.ID.String(),
		Name:      user.Name,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		Wallet: dto.WalletDetail{
			WalletID:  wallet.ID.String(),
			Balance:   wallet.Balance,
			CreatedAt: wallet.CreatedAt,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *WalletHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorResponse{Error: "invalid request"})
		return
	}

	user, wallet, err := h.userService.RegisterUser(r.Context(), req.Name, req.Phone)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.ErrorResponse{Error: err.Error()})
		return
	}

	resp := dto.RegisterResponse{
		UserID:   user.ID.String(),
		WalletID: wallet.ID.String(),
		Balance:  wallet.Balance,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *WalletHandler) Deposit(w http.ResponseWriter, r *http.Request) {
	var req dto.DepositRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorResponse{Error: "invalid request"})
		return
	}

	walletID, err := uuid.Parse(req.WalletID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorResponse{Error: "invalid wallet_id"})
		return
	}

	if err := h.walletService.Deposit(r.Context(), walletID, req.Amount); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorResponse{Error: err.Error()})
		return
	}

	newBalance, _ := h.walletService.GetBalance(r.Context(), walletID)

	resp := dto.DepositResponse{
		Message:    "deposit successful",
		NewBalance: newBalance,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *WalletHandler) Withdraw(w http.ResponseWriter, r *http.Request) {
	var req dto.WithdrawRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorResponse{Error: "invalid request"})
		return
	}

	walletID, err := uuid.Parse(req.WalletID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorResponse{Error: "invalid wallet_id"})
		return
	}

	if err := h.walletService.Withdraw(r.Context(), walletID, req.Amount); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorResponse{Error: err.Error()})
		return
	}

	newBalance, _ := h.walletService.GetBalance(r.Context(), walletID)
	resp := dto.WithdrawResponse{
		Message:    "withdrawal successful",
		NewBalance: newBalance,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *WalletHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	walletID := r.URL.Query().Get("wallet_id")
	if walletID == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorResponse{Error: "wallet_id required"})
		return
	}

	id, err := uuid.Parse(walletID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorResponse{Error: "invalid wallet_id"})
		return
	}

	balance, err := h.walletService.GetBalance(r.Context(), id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(dto.ErrorResponse{Error: err.Error()})
		return
	}

	resp := dto.BalanceResponse{
		WalletID: walletID,
		Balance:  balance,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
