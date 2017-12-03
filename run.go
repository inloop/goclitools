package goclitools

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli"
)

// RunInDir ...
func RunInDir(cmd, dir string) ([]byte, error) {
	command := exec.Command("sh", "-c", "set -o pipefail && "+strings.Replace(cmd, "'", "\\'", -1))
	command.Dir = dir
	return command.Output()
}

// Run ...
func Run(cmd string) ([]byte, error) {
	return RunInDir(cmd, "")
}

// RunInteractiveInDir ...
func RunInteractiveInDir(cmd, dir string) error {
	if os.Getenv("DEBUG") != "" {
		log.Println(cmd)
	}
	command := exec.Command("sh", "-c", "set -o pipefail && "+strings.Replace(cmd, "'", "\\'", -1))
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	command.Dir = dir
	err := command.Run()
	if err != nil {
		return cli.NewExitError(err, 1)
	}
	return nil
}

// RunInteractive ...
func RunInteractive(cmd string) error {
	return RunInteractiveInDir(cmd, "")
}