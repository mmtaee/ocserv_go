package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "help",
	Short: "Help Command",
	Long:  "Display Help Command",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, err2 := fmt.Fprintln(os.Stderr, err)
		if err2 != nil {
			os.Exit(1)
		}
	}
}
