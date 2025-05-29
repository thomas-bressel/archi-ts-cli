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

	// Prompt pour le nom du projet
	projectName, err := prompts.PromptProjectName()
	if err != nil {
		return fmt.Errorf("erreur lors de la saisie du nom de projet: %w", err)
	}
	// Créer le répertoire du projet
	projectPath := filepath.Join(".", projectName)
	if err := os.MkdirAll(projectPath, 0755); err != nil {
		return fmt.Errorf("erreur lors de la création du répertoire: %w", err)
	}

	return nil
}
