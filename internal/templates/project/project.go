package project

import (
	"archi-ts-cli/internal/utils"
	"fmt"
)

// ProjectConfig type to defnie the configuration of the project
type ProjectConfig struct {
	Name         string
	Architecture string
	Language     string
}

/*
* GenerateProject to generate the project structure and files based on the configuration
 */
func GenerateProject(cfg ProjectConfig) error {
	// Create the project directory
	if err := createDirectoryStructure(); err != nil {
		return fmt.Errorf("erreur lors de la création des dossiers: %w", err)
	}
	return nil
}

/*
* createDirectoryStructure creates the basic directory structure for the project
 */
func createDirectoryStructure() error {
	directories := []string{
		"src",
		"src/Databases",
		"src/Middlewares",
		"src/Models",
		"src/Modules",
		"database",
		"uploads",
	}

	for _, dir := range directories {
		if err := utils.CreateDirectory(dir); err != nil {
			return err
		}
	}
	return nil
}
