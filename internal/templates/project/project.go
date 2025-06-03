package project

import (
	"archi-ts-cli/internal/files"
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
GenerateProject to generate the project structure and files based on the configuration
*/
func GenerateProject(cfg ProjectConfig) error {
	// Create the project directory
	if err := createDirectoryStructure(cfg.Architecture); err != nil {
		return fmt.Errorf("error during the folders creation: %w", err)
	}

	// Generate the base files
	if err := generateBaseFiles(cfg); err != nil {
		return fmt.Errorf("error during the files creation: %w", err)
	}

	// .env
	if err := utils.WriteFile(".env", files.GetEnvTemplate()); err != nil {
		return err
	}

	return nil
}

/*
* createDirectoryStructure creates the directory structure based on the chosen architecture
 */
func createDirectoryStructure(architecture string) error {
	var directories []string

	switch architecture {
	case "Layered Architecture":
		directories = getLayeredDirectories()
	case "Clean Architecture":
		directories = getCleanDirectories()
	case "Hexagonal Architecture":
		directories = getHexagonalDirectories()
	default:
		directories = getLayeredDirectories() // Default
	}

	for _, dir := range directories {
		if err := utils.CreateDirectory(dir); err != nil {
			return err
		}
	}
	return nil
}

// generateBaseFiles create the base files for the project
func generateBaseFiles(cfg ProjectConfig) error {

	// Package.json
	packageContent, err := files.GeneratePackageJson(cfg.Name, cfg.Language)
	if err != nil {
		return err
	}
	if err := utils.WriteFile("package.json", packageContent); err != nil {
		return err
	}

	// tsconfig.json (if the option typescript is selected)
	if cfg.Language == "typescript" {
		if err := utils.WriteFile("tsconfig.json", files.GetTsconfigTemplate()); err != nil {
			return err
		}
	}

	return nil
}
