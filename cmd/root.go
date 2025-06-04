package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "archi",
	Short: "Archi TS CLI - Backend Architecture Generator by Thomas Bressel",
	Long: `Archi TS is a CLI tool for generating backend architectures and components.
It helps you quickly scaffold projects with clean architecture patterns.`,
	Version: "1.3.0",
}

/*
Step 1 - Initialize the configuration for the CLI.
*/
func init() {
	// Add subcommands
	rootCmd.AddCommand(createCmd)
}

/*
Step 2 - Initialize the root command and add subcommands.
*/
func Execute() error {
	return rootCmd.Execute()
}
