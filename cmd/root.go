package cmd

var rootCmd = &cobra.Command{
  Use:   "GoCLI-sandbox",
  Short: "GO CLI Study",
  Long: `GoCLI : CLI Sandbox for studying golang`,
  Run: func(cmd *cobra.Command, args []string) {
    // Do Stuff Here
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
