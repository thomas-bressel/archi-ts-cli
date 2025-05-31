package cmd

import (
	"archi-ts-cli/internal/prompts"
	"archi-ts-cli/internal/templates/project"
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

	/////////////////////////////////////////////////
	// Start collecting the result of each prompts //
	/////////////////////////////////////////////////

	// Prompt 1 - Prompt to ask the user for the project name
	projectName, err := prompts.PromptProjectName()
	if err != nil {
		return fmt.Errorf("error with this project name: %w", err)
	}

	// Prompt 2 - Language Type (JS or TS)
	language, err := prompts.PromptLanguage()
	if err != nil {
		return fmt.Errorf("error when choosing a language: %w", err)
	}

	///////////////////////////////////////////////
	// End collecting the result of each prompts //
	///////////////////////////////////////////////

	// Step 3 - Create the project directory
	projectPath := filepath.Join(".", projectName)
	if err := os.MkdirAll(projectPath, 0755); err != nil {
		return fmt.Errorf("error when trying to create the folder: %w", err)
	}

	// Step 4 - Select the project directory
	if err := os.Chdir(projectPath); err != nil {
		return fmt.Errorf("error during the change of directory: %w", err)
	}

	color.New(color.FgYellow).Printf("📁 Project created at: %s\n", projectPath)
	fmt.Println()

	// Générer la structure du projet
	config := project.ProjectConfig{
		Name:     projectName,
		Language: language,
	}
	if err := project.GenerateProject(config); err != nil {
		return fmt.Errorf("error during the generation of the prject: %w", err)
	}

	return nil
}
