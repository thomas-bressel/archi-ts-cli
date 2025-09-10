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
	if err := createDirectoryStructure(cfg.Architecture, cfg.Orm); err != nil {
		return fmt.Errorf("error during the folders creation: %w", err)
	}

	// Generate the base files
	if err := generateBaseFiles(cfg); err != nil {
		return fmt.Errorf("error during the files creation: %w", err)
	}

	// Generate ORM specific files if needed
	// [ ] Layered Architecture
	// [ ] Clean Architecture
	// [ ] Hexagonal Architecture
	if err := generateORMFiles(cfg); err != nil {
		return fmt.Errorf("error during ORM files creation: %w", err)
	}

	// Generate the config file of the project
	projectConfigFile := models.ProjectConfigFile{
		Name:         cfg.Name,
		Version:      "1.0.0",
		Architecture: cfg.Architecture,
		Orm:          cfg.Orm,
		Express:      cfg.Express,
		CreatedAt:    time.Now(),
	}

	if err := config.SaveProjectConfig(projectConfigFile); err != nil {
		return fmt.Errorf("erreur lors de la sauvegarde de la configuration: %w", err)
	}

	return nil
}

// createDirectoryStructure creates the directory structure based on the chosen architecture and ORM
// [X] Layered Architecture
// [ ] Clean Architecture
// [ ] Hexagonal Architecture
func createDirectoryStructure(architecture models.Architecture, orm models.Orm) error {
	var directories []string

	switch architecture {
	case models.LayeredArchitecture:
		directories = getLayeredDirectories()
	case models.CleanArchitecture:
		directories = getCleanDirectories()
	case models.HexagonalArchitecture:
		directories = getHexagonalDirectories()
		// Si TypeORM est choisi, ajouter les dossiers spécifiques
		if orm == models.TypeOrm {
			directories = append(directories, getHexagonalORMDirectories(orm)...)
		}
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
// [ ] Clean Architecture
// [ ] Hexagonal Architecture
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

	// .env - Add  TypeORM variables if needed
	envContent := files.GetEnvTemplate()
	if cfg.Orm == models.TypeOrm {
		envContent = files.GetEnvTemplateWithTypeORM()
	}
	if err := utils.WriteFile(".env", envContent); err != nil {
		return err
	}

	// [X] tsconfig.json
	if err := utils.WriteFile("tsconfig.json", files.GetTsconfigTemplate(cfg.Architecture, cfg.Orm)); err != nil {
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
	// [ ] Express.js template for Layered Architecture
	// [ ] Express.js & Type ORM template for Layered Architecture

	// [ ] Node.js template for Clean Architecture
	// [ ] Express.js template for Clean Architecture
	// [ ] Express.js & Type ORM template for Clean Architecture

	// [ ] Node.js template for Hexagonal Architecture
	// [ ] Express.js template for Hexagonal Architecture
	// [ ] Express.js & Type ORM template for Hexagonal Architecture

	if err := utils.WriteFile("src/index.ts", files.GetIndexTemplate(cfg.Architecture, cfg.Express, cfg.Orm)); err != nil {
		return err
	}

	return nil
}

// generateORMFiles generates ORM specific files if the architecture is Hexagonal and ORM is TypeORM
func generateORMFiles(cfg models.ProjectConfigBuilder) error {
	// Only generate TypeORM files for Hexagonal architecture with TypeORM
	if cfg.Architecture != models.HexagonalArchitecture || cfg.Orm != models.TypeOrm {
		return nil
	}

	// Create adapters/typeorm directory (structure simplifiée)
	ormPath := filepath.Join("src", "adapters", "typeorm")
	if err := utils.CreateDirectory(ormPath); err != nil {
		return err
	}

	// Generate data-source.ts
	dataSourcePath := filepath.Join(ormPath, "data-source.ts")
	if err := utils.WriteFile(dataSourcePath, files.GetTypeORMDataSourceTemplate()); err != nil {
		return fmt.Errorf("error creating data-source.ts: %w", err)
	}

	// Generate create-database.ts
	createDbPath := filepath.Join(ormPath, "create-database.ts")
	if err := utils.WriteFile(createDbPath, files.GetTypeORMCreateDatabaseTemplate()); err != nil {
		return fmt.Errorf("error creating create-database.ts: %w", err)
	}

	return nil
}
