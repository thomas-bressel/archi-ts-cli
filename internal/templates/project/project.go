package project

import (
	"archi-ts-cli/internal/config"
	"archi-ts-cli/internal/models"

	"archi-ts-cli/internal/files"
	"archi-ts-cli/internal/utils"
	"fmt"
	"time"
)

/*
GenerateProject to generate the project structure and files based on the configuration
*/
func GenerateProject(cfg models.ProjectConfigBuilder) error {
	// Create the project directory
	if err := createDirectoryStructure(cfg.Architecture); err != nil {
		return fmt.Errorf("error during the folders creation: %w", err)
	}

	// Generate the base files
	if err := generateBaseFiles(cfg); err != nil {
		return fmt.Errorf("error during the files creation: %w", err)
	}

	// Generate the config file of the project
	projectConfigFile := models.ProjectConfigFile{
		Name:         cfg.Name,
		Version:      "1.0.0",
		Architecture: cfg.Architecture,
		Language:     cfg.Language,
		Express:      cfg.Express,
		CreatedAt:    time.Now(),
	}

	if err := config.SaveProjectConfig(projectConfigFile); err != nil {
		return fmt.Errorf("erreur lors de la sauvegarde de la configuration: %w", err)
	}

	return nil
}

/*
* createDirectoryStructure creates the directory structure based on the chosen architecture
 */
func createDirectoryStructure(architecture models.Architecture) error {
	var directories []string

	switch architecture {
	case models.LayeredArchitecture:
		directories = getLayeredDirectories()
	case models.CleanArchitecture:
		directories = getCleanDirectories()
	case models.HexagonalArchitecture:
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
func generateBaseFiles(cfg models.ProjectConfigBuilder) error {

	// Package.json
	packageContent, err := files.GeneratePackageJson(cfg.Name, string(cfg.Language), cfg.Express)
	if err != nil {
		return err
	}
	if err := utils.WriteFile("package.json", packageContent); err != nil {
		return err
	}

	// .env
	if err := utils.WriteFile(".env", files.GetEnvTemplate()); err != nil {
		return err
	}

	// tsconfig.json (if the option typescript is selected)
	if cfg.Language == models.TypeScript {
		if err := utils.WriteFile("tsconfig.json", files.GetTsconfigTemplate()); err != nil {
			return err
		}
	}

	// .gitignore
	if err := utils.WriteFile(".gitignore", files.GetGitignoreTemplate()); err != nil {
		return err
	}

	// jest.config.ts or js
	configFile := "jest.config." + cfg.GetFileExtension()
	if err := utils.WriteFile(configFile, files.GetJestConfigTemplate(cfg.GetFileExtension())); err != nil {
		return err
	}

	// src/index
	indexFile := "src/index." + cfg.GetFileExtension()
	if err := utils.WriteFile(indexFile, files.GetIndexTemplate(string(cfg.Language), cfg.Express)); err != nil {
		return err
	}

	return nil
}
