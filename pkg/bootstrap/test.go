package bootstrap

import (
	"fmt"
	"ocserv/pkg/config"
	"ocserv/pkg/testutils"
	"os"
	"os/exec"
)

func Test(verbose bool) {
	var (
		out     []byte
		err     error
		command = []string{
			"test",
			"./...",
		}
	)
	fmt.Println("Configuring test environment and database ...")
	config.LoadTestEnv()
	config.Set()
	testutils.DropAndCreateDB("test")
	Migrate()
	fmt.Println("Running tests...")

	if verbose {
		command = append(command, "-v")
	}

	out, err = exec.Command("go", command...).CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		fmt.Println("Error running tests:", err)
		os.Exit(1)
	}
	fmt.Println(string(out))
	fmt.Println("Test Ok")
	os.Exit(0)
}
