package cmd

import (
	"github.com/spf13/cobra"
	"ocserv/internal/database"
)

func init() {
	rootCmd.AddCommand(makeMigrationsCmd)
}

var makeMigrationsCmd = &cobra.Command{
	Use:   "make-migrations",
	Short: "Make migration files",
	Long:  "Make migration files and check database diff",
	Run: func(cmd *cobra.Command, args []string) {
		makeMigrations()
	},
}

func makeMigrations() {
	database.MakeMigrations()
}
