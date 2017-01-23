package main

import (
	"os"
	"os/exec"
)

var (
	commands = map[string]TigerCmd{}
)

type TigerCmd interface {
	Execute([]string) error
}

func main() {
	if len(os.Args) == 1 {
		delegateToGit(os.Args)
		return
	}
	if command, ok := commands[os.Args[1]]; ok {
		err := command.Execute(os.Args[1:])
		if err != nil {
			os.Exit(1)
		}
		return
	}
	delegateToGit(os.Args[1:])
}
func delegateToGit(args []string) {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}
}
