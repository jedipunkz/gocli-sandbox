package main

import (
	"github.com/mitchellh/cli"
	"log"
	"os"
)

type Foo struct{}

func (f *Foo)  Help() string {
	return "cli foo"
}

func (f *Foo) Synopsis() string {
	return "Print \"Foo\""
}

func (f *Foo) Run(args []string) int {
	log.Println("Foo!")
	return 0
}

func main() {
	c := cli.NewCLI("cli", "0.0.1")

	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"foo": func() (cli.Command, error) {
			return &Foo{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}

