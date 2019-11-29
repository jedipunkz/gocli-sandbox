package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/fatih/color"
)

func init() {
	viper.SetDefault("TestVar", "foo")
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:	"version",
	Short:	"Print the version number of gocli-sandbox",
	Long:	`All Software has versions, This is GoCLI-sandbox`,
	Run:	func(cmd *cobra.Command, args []string) {
		cyan := color.New(color.FgCyan).SprintFunc()
		author, _ := cmd.Flags().GetString("author")
		v := viper.New()
		v.SetConfigName(".gocli-sandbox")
		v.AddConfigPath("$HOME")
		err = v.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal errror config file %s \n", err))
		}

		author = v.GetString("author")

		fmt.Println(cyan("GoCLI-sandbox CLI Study Sample v0.1 -- HEAD" + " " + author))
	},
}
