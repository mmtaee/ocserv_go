package cmd

import (
	"github.com/spf13/cobra"
	"ocserv/pkg/bootstrap"
)

var verbose bool

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Testing Project",
	Long:  "Testing Project",
	Run: func(cmd *cobra.Command, args []string) {
		Test()
	},
}

func Test() {
	bootstrap.Test(verbose)
}
