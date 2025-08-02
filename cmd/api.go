package cmd

import (
	"log"
	"net/http"

	"github.com/kagizi/kc-tech-test/domain/service"
	"github.com/kagizi/kc-tech-test/infrastructure"
	"github.com/kagizi/kc-tech-test/infrastructure/config"
	"github.com/kagizi/kc-tech-test/infrastructure/database/mysql"
)

func StartAPI() {
	cfg := config.LoadConfig()

	db, err := mysql.NewDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	userRepo := mysql.NewUserRepository(db)
	transactionRepo := mysql.NewTransactionRepository(db)
	walletRepo := mysql.NewWalletRepository(db, transactionRepo)

	userService := service.NewUserService(userRepo, walletRepo)
	walletService := service.NewWalletService(walletRepo)

	handler := infrastructure.NewWalletHandler(walletService, userService)

	http.HandleFunc("/user", handler.GetUserByPhone)
	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/deposit", handler.Deposit)
	http.HandleFunc("/withdraw", handler.Withdraw)
	http.HandleFunc("/balance", handler.GetBalance)

	log.Printf("Server starting on port %d", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
