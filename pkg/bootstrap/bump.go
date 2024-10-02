package bootstrap

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"os"
	"strconv"
	"strings"
	"time"
)

type bumpStruct struct {
	repo *git.Repository
}

func newBump() *bumpStruct {
	fmt.Println("Opening repository")
	repo, err := git.PlainOpen(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return &bumpStruct{
		repo: repo,
	}
}

func Bump(patch, minor, major bool) {
	bump := newBump()
	_ = bump.validateWorkTree()
	newTag := bump.getNewTag(patch, minor, major)
	bump.createTag(newTag)
	bump.pushTag(newTag)
}

func bumpError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func (b *bumpStruct) validateWorkTree() *git.Worktree {
	w, err := b.repo.Worktree()
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
	err = b.repo.Fetch(&git.FetchOptions{
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

func (b *bumpStruct) getNewTag(patch, minor, major bool) string {
	tags, err := b.repo.Tags()
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

func (b *bumpStruct) createTag(newTag string) {
	headRef, err := b.repo.Head()
	if err != nil {
		bumpError(err)
	}

	cfg, err := b.repo.ConfigScoped(config.SystemScope)
	if err != nil {
		bumpError(err)
	}
	if cfg.User.Name == "" || cfg.User.Email == "" {
		globalCfg, err := b.repo.ConfigScoped(config.GlobalScope)
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
	userAccept("create tag " + newTag)
	tagRef, err := b.repo.CreateTag(newTag, headRef.Hash(), &git.CreateTagOptions{
		Tagger:  tagger,
		Message: fmt.Sprintf("Version %s release", newTag),
	})
	if err != nil {
		bumpError(err)
	}
	fmt.Println("tag created: ", tagRef.Strings())
}

func (b *bumpStruct) pushTag(newTag string) {
	pushOpt := &git.PushOptions{
		RemoteName: "origin",
		RefSpecs: []config.RefSpec{
			config.RefSpec("refs/tags/" + newTag + ":refs/tags/" + newTag),
		},
	}
	userAccept("push tag " + newTag)
	err := b.repo.Push(pushOpt)
	if err != nil {
		bumpError(err)
	}
	fmt.Println("push succeeded")
}

func userAccept(text string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(fmt.Sprintf("Are you sure you want to %s? (yes/no): ", text))
	if scanner.Scan() {
		input := strings.TrimSpace(strings.ToLower(scanner.Text()))
		if input == "yes" || input == "y" {
			fmt.Println("Proceeding with the operation...")
		} else {
			fmt.Println("Operation aborted by the user.")
			os.Exit(0)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
}
