package main

import (
	"flag"
	"fmt"
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
	fmt.Println(*flUpstream)
	fmt.Println(fs.Args())
	return nil
}
