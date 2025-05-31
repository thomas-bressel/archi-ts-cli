package cmd

import (
	"archi-ts-cli/internal/prompts"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project architecture",
	Long:  `Create a new backend project with predefined architecture and structure.`,
	RunE:  runCreate,
}

func runCreate(cmd *cobra.Command, args []string) error {

	// Banner
	color.New(color.FgCyan, color.Bold).Println("🚀 ArchiTS CLI - Project Scaffolding")
	fmt.Println()

	// Step 1 - Prompt to ask the user for the project name
	projectName, err := prompts.PromptProjectName()
	if err != nil {
		return fmt.Errorf("error with this project name: %w", err)
	}

	// Step 2 - Language Type (JS or TS)
	language, err := prompts.PromptLanguage()
	if err != nil {
		return fmt.Errorf("error when choosing a language: %w", err)
	}
	fmt.Println("Selected language : ", language) // Debugging purpose

	// Ending Step - Create the project directory
	projectPath := filepath.Join(".", projectName)
	if err := os.MkdirAll(projectPath, 0755); err != nil {
		return fmt.Errorf("error when trying to create the folder: %w", err)
	}

	return nil
}
