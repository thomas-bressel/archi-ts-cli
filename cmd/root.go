package cmd

// Step 1 - Initialize the configuration for the CLI.
func init() {
	// Add subcommands to archi command
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(generateCmd)

	// Add subcommand to generate command
	generateCmd.AddCommand(entityCmd)
}

// Step 2 - Initialize the root command and add subcommands.
func Execute() error {
	return rootCmd.Execute()
}
