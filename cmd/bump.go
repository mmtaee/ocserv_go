package cmd

import (
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func init() {
	rootCmd.AddCommand(bumpCmd)
}

var bumpCmd = &cobra.Command{
	Use:   "bump",
	Short: "Bump version",
	Long:  "Bump Version and push in git",
	Run: func(cmd *cobra.Command, args []string) {
		Bump()
	},
}

func bumpError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func Bump() {
	fmt.Println("Opening repository")
	r, err := git.PlainOpen(".")
	if err != nil {
		bumpError(err)
	}

	w, err := r.Worktree()
	if err != nil {
		bumpError(err)
	}

	status, err := w.Status()
	if err != nil {
		bumpError(err)
	}
	if !status.IsClean() {
		statusResult := strings.Replace(status.String(), "\n", "\n\t", -1)
		bumpError(errors.New("Repository is not clean\n\t" + statusResult))
	}

	fmt.Println("Fetching tags")
	err = r.Fetch(&git.FetchOptions{
		RemoteName: "origin",
		Tags:       git.AllTags,
	})
	if err != nil && !strings.Contains(err.Error(), "up-to-date") {
		bumpError(err)
	}

	fmt.Println("Pulling branch")
	err = w.Pull(&git.PullOptions{
		RemoteName: "origin",
	})
	if err != nil && !strings.Contains(err.Error(), "up-to-date") {
		bumpError(err)
	}

	tags, err := r.Tags()
	if err != nil {
		bumpError(err)
	}

	var lastTag string
	err = tags.ForEach(func(ref *plumbing.Reference) error {
		lastTag = ref.Name().Short()
		return nil
	})
	if err != nil {
		bumpError(err)
	}
	if lastTag == "" {
		lastTag = "v0.0.0"
	}

	// check minor or other
	lastTag = strings.TrimPrefix(lastTag, "v")
	tagDelimiters := strings.Split(lastTag, ".")
	fmt.Println(tagDelimiters)
	//r.CreateTag()
}
