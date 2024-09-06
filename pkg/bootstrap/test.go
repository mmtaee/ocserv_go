package bootstrap

import (
	"log"
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
	log.Println("Configuring test environment and database ...")
	config.LoadTestEnv()
	config.Set()
	testutils.DropAndCreateDB("test")
	Migrate()
	log.Println("Running tests...")

	if verbose {
		command = append(command, "-v")
	}

	out, err = exec.Command("go", command...).CombinedOutput()
	if err != nil {
		log.Println("Error running tests:", err)
		os.Exit(1)
	}
	log.Println(string(out))
	os.Exit(0)
}
