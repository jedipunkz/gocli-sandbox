package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/mitchellh/cli"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
)

type Exec struct{}

func (f *Exec) Help() string {
	return "exec command"
}

func (f *Exec) Synopsis() string {
	return "Execution command"
}

func (f *Exec) Run(args []string) int {
	out, err := exec.Command("ls", "/tmp").Output()
	if err != nil {
		log.Println("Fatal error.")
		return 1
	}

	cyan := color.New(color.FgCyan).SprintFunc()

	fmt.Printf(cyan(string(out)))
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
		"exec": func() (cli.Command, error) {
			return &Exec{}, nil
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
