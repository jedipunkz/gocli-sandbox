package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName(".gocli-sandbox")
	viper.AddConfigPath("$HOME/tmp")

	viper.SetDefault("author", "foo")
	viper.SetDefault("title", "foo")

	rootCmd.AddCommand(configCmd)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Read Config file",
	Long:  `Read config file named $HOME/.gocli-sandbox.yaml and Output values`,
	Run: func(cmd *cobra.Command, args []string) {
		cyan := color.New(color.FgCyan).SprintFunc()
		red := color.New(color.FgRed).SprintFunc()

		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("Fatal errror config file %s \n", err))
		}

		author := viper.GetString("author")
		title := viper.GetString("title")

		fmt.Println(cyan("Author: " + author))
		fmt.Println(red("Title: " + title))
	},
}
