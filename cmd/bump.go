package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"ocserv/pkg/bootstrap"
)

var (
	patch bool
	minor bool
	major bool
)

func init() {
	rootCmd.AddCommand(bumpCmd)
	bumpCmd.Flags().BoolVarP(&minor, "minor", "m", false, "Minor version")
	bumpCmd.Flags().BoolVarP(&major, "major", "j", false, "Major version")
}

var bumpCmd = &cobra.Command{
	Use:   "bump",
	Short: "Bump version",
	Long:  "Bump Version and push in git",
	Run: func(cmd *cobra.Command, args []string) {
		Bump()
	},
}

func Bump() {
	if !minor && !major {
		fmt.Println("Setting patch version")
		patch = true
	} else if minor {
		fmt.Println("Setting minor version")
	} else {
		fmt.Println("Setting major version")
	}
	bootstrap.Bump(patch, minor, major)
}
