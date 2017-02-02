package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/vdemeester/tiger/git"
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
	runGit([]string{"status"}, func(cmd *exec.Cmd) {
		cmd.Stdout = ioutil.Discard
	})
	runHook("before-all")
	if len(os.Args) == 1 {
		runGit(os.Args)
		return
	}
	if command, ok := commands[os.Args[1]]; ok {
		runHook(fmt.Sprintf("before-%s", command.Name()))
		err := command.Execute(os.Args[1:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(1)
		}
		runHook(fmt.Sprintf("after-%s", command.Name()))
		runHook("after-all")
		return
	}
	runGit(os.Args[1:])
	runHook("after-all")
}

func runGit(args []string, cmdModifier ...func(*exec.Cmd)) {
	err := git.RunCommand(args, cmdModifier...)
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
