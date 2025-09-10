package cmd

import (
	"archi-ts-cli/internal/config"
	"archi-ts-cli/internal/templates/generate"
	"archi-ts-cli/internal/utils"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// EntityPaths contains all paths for generated files INCLUDING TESTS
type EntityPaths struct {
	Entity         string
	Controller     string
	Service        string
	Repository     string
	Route          string
	ControllerTest string
	ServiceTest    string
	RepositoryTest string
}

func runGenerateEntity(cmd *cobra.Command, args []string) error {
	entityName := args[0]
	caser := cases.Title(language.English)
	entityName = caser.String(strings.ToLower(entityName))

	projectConfig, err := config.LoadProjectConfig()
	if err != nil {
		return fmt.Errorf("impossible to load project configuration: %w", err)
	}

	entityConfig := generate.EntityConfig{
		Name:    entityName,
		Variant: "typescript",
	}

	ext := "ts"
	paths := getEntityPaths(string(projectConfig.Architecture), entityName, ext, string(projectConfig.Orm))
	filesToGenerate := []struct {
		path     string
		template string
		name     string
	}{
		{paths.Entity, generate.GetEntityTemplate(entityConfig), "Entity"},
		{paths.Controller, generate.GetControllerTemplate(entityConfig), "Controller"},
		{paths.Service, generate.GetServiceTemplate(entityConfig), "Service"},
		{paths.Repository, generate.GetRepositoryTemplate(entityConfig), "Repository"},
		{paths.Route, generate.GetRouteTemplate(entityConfig), "Route"},
		{paths.ControllerTest, generate.GetControllerTestTemplate(entityConfig), "Controller Test"},
		{paths.ServiceTest, generate.GetServiceTestTemplate(entityConfig), "Service Test"},
		{paths.RepositoryTest, generate.GetRepositoryTestTemplate(entityConfig), "Repository Test"},
	}

	color.New(color.FgCyan, color.Bold).Println("📁 Generating source files...")
	for i, file := range filesToGenerate {
		if i == 5 {
			fmt.Println()
			color.New(color.FgCyan, color.Bold).Println("🧪 Generating test files...")
		}

		dir := filepath.Dir(file.path)
		if err := utils.CreateDirectory(dir); err != nil {
			return fmt.Errorf("error while creating the directory %s: %w", dir, err)
		}

		if err := utils.WriteFile(file.path, file.template); err != nil {
			return fmt.Errorf("error while creating the file %s: %w", file.path, err)
		}
		color.New(color.FgGreen).Printf("  ✅ %s created: %s\n", file.name, file.path)
	}

	fmt.Println()
	color.New(color.FgGreen, color.Bold).Printf("✨ '%s' entity has been generated with tests!\n", entityName)
	color.New(color.FgYellow).Println("💡 Don't forget to:")
	fmt.Println("  - Run tests: npm test")
	fmt.Println("  - Update test files with your specific business logic")

	return nil
}

func getEntityPaths(architecture string, entityName string, ext string, orm string) EntityPaths {
	lowerName := strings.ToLower(entityName)

	switch architecture {
	case "Clean Architecture":
		return EntityPaths{
			Entity:         fmt.Sprintf("src/domain/entities/%s.entity.%s", lowerName, ext),
			Controller:     fmt.Sprintf("src/presentation/controllers/%s.controller.%s", lowerName, ext),
			Service:        fmt.Sprintf("src/data/services/%s.service.%s", lowerName, ext),
			Repository:     fmt.Sprintf("src/data/repositories/%s.repository.%s", lowerName, ext),
			Route:          fmt.Sprintf("src/presentation/routes/%s.routes.%s", lowerName, ext),
			ControllerTest: fmt.Sprintf("tests/unit/controllers/%s.controller.test.%s", lowerName, ext),
			ServiceTest:    fmt.Sprintf("tests/unit/services/%s.service.test.%s", lowerName, ext),
			RepositoryTest: fmt.Sprintf("tests/unit/repositories/%s.repository.test.%s", lowerName, ext),
		}
	case "Hexagonal Architecture":
		repositoryPath := ""
		if orm == "TypeORM" {
			repositoryPath = fmt.Sprintf("src/adapters/typeorm/repositories/%s.repository.%s", lowerName, ext)
		} else {
			repositoryPath = fmt.Sprintf("src/adapters/repositories/%s.repository.%s", lowerName, ext)
		}

		return EntityPaths{
			Entity:         fmt.Sprintf("src/domain/entities/%s.entity.%s", lowerName, ext),
			Controller:     fmt.Sprintf("src/interfaces/controllers/%s.controller.%s", lowerName, ext),
			Service:        fmt.Sprintf("src/application/services/%s.service.%s", lowerName, ext),
			Repository:     repositoryPath,
			Route:          fmt.Sprintf("src/interfaces/routes/%s.routes.%s", lowerName, ext),
			ControllerTest: fmt.Sprintf("tests/unit/controllers/%s.controller.test.%s", lowerName, ext),
			ServiceTest:    fmt.Sprintf("tests/unit/services/%s.service.test.%s", lowerName, ext),
			RepositoryTest: fmt.Sprintf("tests/unit/repositories/%s.repository.test.%s", lowerName, ext),
		}
	default: // Layered Architecture
		return EntityPaths{
			Entity:         fmt.Sprintf("src/entities/%s.entity.%s", lowerName, ext),
			Controller:     fmt.Sprintf("src/controllers/%s.controller.%s", lowerName, ext),
			Service:        fmt.Sprintf("src/services/%s.service.%s", lowerName, ext),
			Repository:     fmt.Sprintf("src/repositories/%s.repository.%s", lowerName, ext),
			Route:          fmt.Sprintf("src/routes/%s.routes.%s", lowerName, ext),
			ControllerTest: fmt.Sprintf("tests/unit/controllers/%s.controller.test.%s", lowerName, ext),
			ServiceTest:    fmt.Sprintf("tests/unit/services/%s.service.test.%s", lowerName, ext),
			RepositoryTest: fmt.Sprintf("tests/unit/repositories/%s.repository.test.%s", lowerName, ext),
		}
	}
}
