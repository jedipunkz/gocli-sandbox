package main

import (
	"github.com/mitchellh/cli"
	"github.com/spf13/viper"
	"github.com/fatih/color"
	"log"
	"os"
)

type Foo struct{}

func (f *Foo) Help() string {
	return "cli foo"
}

func (f *Foo) Synopsis() string {
	return "Print \"Foo\""
}

func (f *Foo) Run(args []string) int {
	log.Println("Foo!")
	return 0
}

type Config struct{}

func (c *Config) Help() string {
	return "read config"
}

func (c *Config) Synopsis() string {
	return "Load Config File and Output"
}

func (c *Config) Run(args []string) int {
	viper.SetConfigName(".gocli-sandbox")
	viper.AddConfigPath("$HOME/tmp")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Fatal error config file %s\n", err)
	}

	author := viper.GetString("author")
	title := viper.GetString("title")

	red := color.New(color.FgRed).SprintFunc()

	log.Println(red("Author: " + author))
	log.Println(red("Title: " + title))
	return 0
}

func main() {
	c := cli.NewCLI("cli", "0.0.1")

	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"foo": func() (cli.Command, error) {
			return &Foo{}, nil
		},
		"config": func() (cli.Command, error) {
			return &Config{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
