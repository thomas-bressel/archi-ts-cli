package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "archi",
	Short: "Archi TS CLI - Backend Architecture Generator by Thomas Bressel",
	Long: `Archi TS is a CLI tool for generating backend architectures and components.
It helps you quickly scaffold projects with clean architecture patterns.`,
	Version: "1.0.0",
}

func Execute() error {
	return rootCmd.Execute()
}
