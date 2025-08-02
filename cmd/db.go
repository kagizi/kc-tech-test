package cmd

import (
	"log"

	"github.com/kagizi/kc-tech-test/infrastructure/config"
	"github.com/kagizi/kc-tech-test/infrastructure/database/mysql"
)

func RunMigrations() {
	cfg := config.LoadConfig()

	db, err := mysql.NewDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	if err := mysql.RunMigrationsUp(db, cfg); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	log.Println("Migrations up completed successfully")
}

func RunMigrationsDown() {
	cfg := config.LoadConfig()

	db, err := mysql.NewDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	if err := mysql.RunMigrationsDown(db, cfg); err != nil {
		log.Fatal("Failed to run migrations down:", err)
	}

	log.Println("Migrations down completed successfully")
}
