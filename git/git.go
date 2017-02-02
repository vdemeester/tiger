package git

import (
	"os"
	"os/exec"
)

type Remote struct {
	Name string
	URI  string
}

func RunCommand(args []string, cmdModifier ...func(cmd *exec.Cmd)) error {
	cmd := exec.Command("git", args...)

	for _, modifier := range cmdModifier {
		modifier(cmd)
	}

	if cmd.Stdout == nil {
		cmd.Stdout = os.Stdout
	}
	if cmd.Stderr == nil {
		cmd.Stderr = os.Stderr
	}
	if cmd.Stdin == nil {
		cmd.Stdin = os.Stdin
	}

	return cmd.Run()
}
