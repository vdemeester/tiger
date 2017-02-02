package main

import (
	"flag"
	"strings"

	"github.com/pkg/errors"
	"github.com/vdemeester/tiger/git"
)

type prCmd struct {
}

func init() {
	cmd := &prCmd{}
	commands[cmd.Name()] = cmd
}

func (c *prCmd) Name() string {
	return "pr"
}

func (c *prCmd) Execute(args []string) error {
	fs := flag.NewFlagSet(args[0], flag.ExitOnError)
	flUpstream := fs.String("upstream", "upstream", "upstream remote name with/without the remote (remote or remote:git-url)")
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
	return git.CheckoutPR(config)
}

func parseUpstream(remote string) (git.Remote, error) {
	remoteElts := strings.SplitN(remote, ":", 2)
	if len(remoteElts) == 0 {
		return git.Remote{}, errors.Errorf("invalid remote : %q", remote)
	}
	name := remoteElts[0]
	var uri string
	if len(remoteElts) == 2 {
		uri = remoteElts[1]
	}
	return git.Remote{
		Name: name,
		URI:  uri,
	}, nil
}
