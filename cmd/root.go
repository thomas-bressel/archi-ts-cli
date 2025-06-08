package cmd

/*
Step 1 - Initialize the configuration for the CLI.
*/
func init() {
	// Add subcommands
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(versionCmd)
}

/*
Step 2 - Initialize the root command and add subcommands.
*/
func Execute() error {
	return rootCmd.Execute()
}
