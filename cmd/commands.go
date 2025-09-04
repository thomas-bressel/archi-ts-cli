package cmd

import (
	"archi-ts-cli/internal/utils"

	"github.com/spf13/cobra"
)

// All CLI commands are listed here
var (

	// rootCmd is the main command of Archi TS CLI
	rootCmd = &cobra.Command{
		Use:   "archi",
		Short: "Archi TS CLI - Backend Architecture Generator",
		Long: `Archi TS is a CLI tool for generating backend architectures and components.
It helps you quickly scaffold projects with clean architecture patterns.`,
		Version: "1.7.0",
	}

	// versionCmd displays detailed version information
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Display detailed version information",
		Run: func(cmd *cobra.Command, args []string) {
			utils.DisplayBanner()
		},
	}

	// createCmd is the Cobra command to create a new backend project architecture.
	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new project architecture",
		Long:  `Create a new backend project with predefined architecture and structure.`,
		RunE:  runCreate,
	}

	// generateCmd is the Cobra command to generate specific components within an existing project.
	generateCmd = &cobra.Command{
		Use:     "generate",
		Aliases: []string{"g"},
		Short:   "Generate specific components",
		Long:    `Generate specific components like controllers, services, or repositories within an existing project.`,
	}

	// entityCmd is the Cobra command to generate a complete entity with all its associated files.
	entityCmd = &cobra.Command{
		Use:     "entity [name]",
		Aliases: []string{"e"},
		Short:   "Generate a new entity",
		Long:    `Generate a complete entity with all associated files (Entity, Route, Controller, Service, Repository, Query, DTOs).`,
		Args:    cobra.ExactArgs(1),
		RunE:    runGenerateEntity,
	}
)
