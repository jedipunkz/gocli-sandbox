package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

type FooCommand struct{}

func (c *FooCommand) Run(args []string) int {
	fmt.Printf("Foo sub-command execute\n")
	return 0
}
func (c *FooCommand) Synopsis() string {
	return "Foo Sub-Command."
}
func (c *FooCommand) Help() string {
	return "Usage: example foo"
}

func main() {
	c := cli.NewCLI("example", "0.0.1")

	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"foo": func() (cli.Command, error) {
			return &FooCommand{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(exitStatus)
}
