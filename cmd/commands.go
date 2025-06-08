package cmd

import (
	"archi-ts-cli/internal/utils"

	"github.com/spf13/cobra"
)

// All CLI commands are listed here
var (
	rootCmd = &cobra.Command{
		Use:   "archi",
		Short: "Archi TS CLI - Backend Architecture Generator",
		Long: `Archi TS is a CLI tool for generating backend architectures and components.
It helps you quickly scaffold projects with clean architecture patterns.`,
		Version: "1.7.0",
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Display detailed version information",
		Run: func(cmd *cobra.Command, args []string) {
			utils.DisplayBanner()
		},
	}

	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new project architecture",
		Long:  `Create a new backend project with predefined architecture and structure.`,
		RunE:  runCreate,
	}
)
