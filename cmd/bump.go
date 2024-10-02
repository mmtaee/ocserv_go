package cmd

import (
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/spf13/cobra"
	"ocserv/pkg/bootstrap"
	"os"
	"strconv"
	"strings"
	"time"
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

type bump struct {
	r *git.Repository
}

var bumpCmd = &cobra.Command{
	Use:   "bump",
	Short: "Bump version",
	Long:  "Bump Version and push in git",
	Run: func(cmd *cobra.Command, args []string) {

		Bump()
		//Bump(&bump{
		//	r: r,
		//})
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

//func Bump(b *bump) {
//	_ = b.validateWorkTree()
//	newTag := b.getNewTag()
//	b.createTag(newTag)
//	b.pushTag(newTag)
//}

func bumpError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func (b *bump) validateWorkTree() *git.Worktree {
	w, err := b.r.Worktree()
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
	err = b.r.Fetch(&git.FetchOptions{
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
	return w
}

func (b *bump) getNewTag() string {
	tags, err := b.r.Tags()
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

	lastTag = strings.TrimPrefix(lastTag, "v")
	tagDelimitersStr := strings.Split(lastTag, ".")
	if len(tagDelimitersStr) != 3 {
		bumpError(errors.New("invalid tag format"))
	}
	tagDelimiters := make([]int, 0, 3)
	for _, s := range tagDelimitersStr {
		if num, err := strconv.Atoi(s); err == nil {
			tagDelimiters = append(tagDelimiters, num)
		}
	}

	if patch {
		tagDelimiters[2]++
	} else if minor {
		tagDelimiters[1]++
		tagDelimiters[2] = 0
	} else {
		tagDelimiters[0]++
		tagDelimiters[1] = 0
		tagDelimiters[2] = 0
	}
	return "v" + strings.Trim(strings.Replace(fmt.Sprint(tagDelimiters), " ", ".", -1), "[]")
}

func (b *bump) createTag(newTag string) {
	headRef, err := b.r.Head()
	if err != nil {
		bumpError(err)
	}

	cfg, err := b.r.ConfigScoped(config.SystemScope)
	if err != nil {
		bumpError(err)
	}
	if cfg.User.Name == "" || cfg.User.Email == "" {
		globalCfg, err := b.r.ConfigScoped(config.GlobalScope)
		if err != nil {
			bumpError(err)
		}
		if cfg.User.Name == "" {
			cfg.User.Name = globalCfg.User.Name
		}
		if cfg.User.Email == "" {
			cfg.User.Email = globalCfg.User.Email
		}
	}
	if cfg.User.Name == "" || cfg.User.Email == "" {
		bumpError(errors.New("could not find user name or email in Git config"))
	}
	tagger := &object.Signature{
		Name:  cfg.User.Name,
		Email: cfg.User.Email,
		When:  time.Now(),
	}

	tagRef, err := b.r.CreateTag(newTag, headRef.Hash(), &git.CreateTagOptions{
		Tagger:  tagger,
		Message: fmt.Sprintf("Version %s release", newTag),
	})
	if err != nil {
		bumpError(err)
	}
	fmt.Println("tag created: ", tagRef.Strings())
}

func (b *bump) pushTag(newTag string) {
	pushOpt := &git.PushOptions{
		RemoteName: "origin",
		RefSpecs: []config.RefSpec{
			config.RefSpec("refs/tags/" + newTag + ":refs/tags/" + newTag),
		},
	}
	err := b.r.Push(pushOpt)
	if err != nil {
		bumpError(err)
	}
	fmt.Println("push succeeded")
}
