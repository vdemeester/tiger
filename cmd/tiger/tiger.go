package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/vdemeester/tiger/hooks"
)

var (
	commands = map[string]TigerCmd{}
)

type TigerCmd interface {
	Execute([]string) error
	Name() string
}

func main() {
	runHook("before-all")
	if len(os.Args) == 1 {
		delegateToGit(os.Args)
		return
	}
	if command, ok := commands[os.Args[1]]; ok {
		runHook(fmt.Sprintf("before-%s", command.Name()))
		err := command.Execute(os.Args[1:])
		if err != nil {
			os.Exit(1)
		}
		runHook(fmt.Sprintf("after-%s", command.Name()))
		runHook("after-all")
		return
	}
	delegateToGit(os.Args[1:])
	runHook("after-all")
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

func runHook(hook string) {
	if err := hooks.Run(hook); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
}
