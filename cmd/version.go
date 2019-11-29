package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:	"version",
	Short:	"Print the version number of gocli-sandbox",
	Long:	`All Software has versions, This is GoCLI-sandbox`,
	Run:	func(cmd *cobra.Command, args []string) {
		fmt.Println("GoCLI-sandbox CLI Study Sample v0.1 -- HEAD")
	},
}
