package cmd

import (
	"github.com/spf13/cobra"
	"ocserv/pkg/bootstrap"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate models",
	Long:  "Migrate models in to database",
	Run: func(cmd *cobra.Command, args []string) {
		Migrate()
	},
}

func Migrate() {
	bootstrap.Migrate()
}
