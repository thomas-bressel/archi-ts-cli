package project

import (
	"archi-ts-cli/internal/config"
	"archi-ts-cli/internal/models"
	"archi-ts-cli/internal/templates/create/files"
	"archi-ts-cli/internal/utils"
	"fmt"
	"path/filepath"
	"time"
)

// GenerateProject to generate the project structure and files based on the configuration
func GenerateProject(cfg models.ProjectConfigBuilder) error {

	// Create the project directory structure based on architecture and ORM
	if err := createDirectoryStructure(cfg.Architecture); err != nil {
		return fmt.Errorf("error during the folders creation: %w", err)
	}

	// Generate the base files
	if err := generateBaseFiles(cfg); err != nil {
		return fmt.Errorf("error during the files creation: %w", err)
	}

	// Generate ORM specific files if needed
	if cfg.Orm == "typeorm" {
		if err := generateORMFiles(cfg); err != nil {
			return fmt.Errorf("error during ORM files creation: %w", err)
		}
	}

	// Generate the config file of the project
	projectConfigFile := models.ProjectConfigFile{
		Name:         cfg.Name,
		Version:      "1.0.0",
		Architecture: cfg.Architecture,
		Orm:          cfg.Orm,
		Express:      cfg.Express,
		Library:      cfg.Library,
		Port:         cfg.Port,
		CreatedAt:    time.Now(),
	}

	if err := config.SaveProjectConfig(projectConfigFile); err != nil {
		return fmt.Errorf("erreur lors de la sauvegarde de la configuration: %w", err)
	}

	return nil
}

// createDirectoryStructure creates the directory structure based on the chosen architecture and ORM
// [X] Layered Architecture
// [X] Clean Architecture
// [X] Hexagonal Architecture
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
// [X] Layered Architecture
// [X] Clean Architecture
// [X] Hexagonal Architecture
func generateBaseFiles(cfg models.ProjectConfigBuilder) error {
	// Package.json - Dissociate architecture and ORM
	var packageContent string
	var err error

	// Generate package.json based on architecture and ORM
	packageContent, err = files.GeneratePackageJson(cfg)

	if err != nil {
		return err
	}

	// Write package.json
	if err := utils.WriteFile("package.json", packageContent); err != nil {
		return err
	}

	// .env files
	envFiles := []string{".env", ".env.exemple", ".env.docker", ".env.production", ".env.staging"}
	for _, file := range envFiles {
		if err := utils.WriteFile(file, files.GetEnvTemplate(cfg.Port)); err != nil {
			return err
		}
	}

	// [X] tsconfig.json
	if err := utils.WriteFile("tsconfig.json", files.GetTsconfigTemplate(cfg.Architecture)); err != nil {
		return err
	}

	// [X] .gitignore
	if err := utils.WriteFile(".gitignore", files.GetGitignoreTemplate()); err != nil {
		return err
	}

	// [ ] jest.config.ts
	if err := utils.WriteFile("jest.config.ts", files.GetJestConfigTemplate(cfg.Architecture, cfg.Orm)); err != nil {
		return err
	}

	// src/index.ts - Use the new template with ORM support
	// [X] Node.js template for Layered Architecture
	// [X] Express.js template for Layered Architecture
	// [X] Express.js & Type ORM template for Layered Architecture

	// [X] Node.js template for Clean Architecture
	// [X] Express.js template for Clean Architecture
	// [X] Express.js & Type ORM template for Clean Architecture

	// [X] Node.js template for Hexagonal Architecture
	// [X] Express.js template for Hexagonal Architecture
	// [X] Express.js & Type ORM template for Hexagonal Architecture

	if err := utils.WriteFile("src/index.ts", files.GetIndexTemplate(cfg.Architecture, cfg.Express, cfg.Orm)); err != nil {
		return err
	}

	return nil
}

// generateORMFiles generates ORM specific files if the architecture is Hexagonal and ORM is TypeORM
// [X] Layered Architecture
// [X] Clean Architecture
// [X] Hexagonal Architecture
func generateORMFiles(cfg models.ProjectConfigBuilder) error {
	ormPath := ""
	scriptPath := ""

	switch cfg.Architecture {
	case models.LayeredArchitecture:
		ormPath = filepath.Join("src", "data", "database", "connection")
		scriptPath = filepath.Join("scripts")
	case models.CleanArchitecture:
		ormPath = filepath.Join("src", "infrastructure", "database", "config")
		scriptPath = filepath.Join("scripts")
	case models.HexagonalArchitecture:
		ormPath = filepath.Join("src", "adapters", "secondary", "persistence", "orm")
		scriptPath = filepath.Join("scripts")
	default:
		ormPath = ""
		scriptPath = ""
	}

	// Generate data-source.ts
	dataSourcePath := filepath.Join(ormPath, "data-source.ts")
	if err := utils.WriteFile(dataSourcePath, files.GetTypeORMDataSourceTemplate(string(cfg.Architecture))); err != nil {
		return fmt.Errorf("error creating data-source.ts: %w", err)
	}

	// Generate create-database.ts
	createDbPath := filepath.Join(ormPath, "create-database.ts")
	if err := utils.WriteFile(createDbPath, files.GetTypeORMCreateDatabaseTemplate()); err != nil {
		return fmt.Errorf("error creating create-database.ts: %w", err)
	}

	// Generate generate-migration.ts
	createScriptsPath := filepath.Join(scriptPath, "generate-migration.ts")
	if err := utils.WriteFile(createScriptsPath, files.GetTypeORMHelperScriptTemplate(string(cfg.Architecture))); err != nil {
		return fmt.Errorf("error creating generate-migration.ts: %w", err)
	}

	return nil
}
