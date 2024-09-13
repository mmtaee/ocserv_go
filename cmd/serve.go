package cmd

import (
	"github.com/spf13/cobra"
	"ocserv/pkg/bootstrap"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve Api",
	Long:  "Serve Web Api Service",
	Run: func(cmd *cobra.Command, args []string) {
		Serve()
	},
}

func Serve() {
	bootstrap.Serve()
}
