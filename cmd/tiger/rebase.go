package main

import (
	"flag"
	"github.com/pkg/errors"
	"github.com/vdemeester/tiger/git"
)

type rebasePrCmd struct {
}

func init() {
	cmd := &rebasePrCmd{}
	commands[cmd.Name()] = cmd
}

func (c *rebasePrCmd) Name() string {
	return "rebase-pr"
}

func (c *rebasePrCmd) Execute(args []string) error {
	fs := flag.NewFlagSet(args[0], flag.ExitOnError)
	flUpstream := fs.String("upstream", "upstream", "upstream remote name")
	fs.Parse(args[1:])
	if len(fs.Args()) != 1 {
		return errors.Errorf("Wrong number of arguments, expected 1, got %d (%v)", len(fs.Args()), fs.Args())
	}
	upstream, err := parseUpstream(*flUpstream)
	if err != nil {
		return errors.Wrap(err, "error parsing upstream")
	}
	config := git.PRConfig{
		Number: fs.Arg(0),
		Remote: upstream,
	}
	return git.RebasePR(config)
}
