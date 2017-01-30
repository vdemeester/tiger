package main

import (
	"flag"
	"fmt"
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
	fmt.Println(args)
	fs := flag.NewFlagSet(args[0], flag.ExitOnError)
	flUpstream := fs.String("upstream", "upstream", "upstream remote name")
	fs.Parse(args[1:])
	fmt.Println(*flUpstream)
	fmt.Println(fs.Args())
	return nil
}
