package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "kc-test-tech",
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start the API server",
	Run: func(cmd *cobra.Command, args []string) {
		StartAPI()
	},
}

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Database operations",
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		RunMigrations()
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Run database migrations down",
	Run: func(cmd *cobra.Command, args []string) {
		RunMigrationsDown()
	},
}

func init() {
	dbCmd.AddCommand(migrateCmd)
	dbCmd.AddCommand(migrateDownCmd)
	rootCmd.AddCommand(apiCmd)
	rootCmd.AddCommand(dbCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
