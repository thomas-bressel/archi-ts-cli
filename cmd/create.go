package cmd

import (
	"archi-ts-cli/internal/prompts"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project architecture",
	Long:  `Create a new backend project with predefined architecture and structure.`,
	RunE:  runCreate,
}

func runCreate(cmd *cobra.Command, args []string) error {

	// Prompt to ask the user for the project name
	projectName, err := prompts.PromptProjectName()
	if err != nil {
		return fmt.Errorf("error with this project name: %w", err)
	}
	// Create the project directory
	projectPath := filepath.Join(".", projectName)
	if err := os.MkdirAll(projectPath, 0755); err != nil {
		return fmt.Errorf("error when trying to create the folder: %w", err)
	}

	return nil
}
