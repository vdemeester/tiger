package main

import (
	"flag"
	"fmt"
)

type rebaseCmd struct {
}

func init() {
	commands["rb"] = &rebaseCmd{}
}

func (c *rebaseCmd) Execute(args []string) error {
	fs := flag.NewFlagSet(args[0], flag.ExitOnError)
	flUpstream := fs.String("upstream", "upstream", "upstream remote name")
	fs.Parse(args[1:])
	fmt.Println(*flUpstream)
	fmt.Println(fs.Args())
	return nil
}
