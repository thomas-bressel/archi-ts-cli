package cmd

import (
	"archi-ts-cli/internal/config"
	"archi-ts-cli/internal/models"
	"archi-ts-cli/internal/templates/generate"
	"archi-ts-cli/internal/templates/generate/tests"

	"archi-ts-cli/internal/utils"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// runGenerateEntity generate different files corresponding to layers and entity
func runGenerateEntity(cmd *cobra.Command, args []string) error {

	// Get the argument and format to english type (the name of the entity)
	entityName := args[0]
	caser := cases.Title(language.English)
	entityName = caser.String(strings.ToLower(entityName))

	// Read information in .archi/config.json file
	projectConfig, err := config.LoadProjectConfig()
	if err != nil {
		return fmt.Errorf("impossible to load project configuration: %w", err)
	}

	// Get config of the entity
	entityConfig := models.EntityConfig{
		Name:    entityName,
		Express: projectConfig.Express,
		Orm:     projectConfig.Orm,
	}

	// Define suffixe and paths of each files
	ext := "ts"
	paths := getEntityPaths(string(projectConfig.Architecture), entityName, ext, string(projectConfig.Orm))

	createListFilesToGenerate(string(projectConfig.Architecture), entityConfig, paths)

	return nil
}

// getEntityPaths return paths where each files should be created
// [X] Layered Architecture
// [ ] Clean Architecture
// [ ] Hexagonal Architecture
func getEntityPaths(architecture string, entityName string, ext string, orm string) models.EntityPaths {
	lowerName := strings.ToLower(entityName)

	switch architecture {
	case "Clean Architecture":
		return models.EntityPaths{
			Entity:             fmt.Sprintf("src/domain/entities/%s.entity.%s", lowerName, ext),
			Controller:         fmt.Sprintf("src/presentation/controllers/%s.controller.%s", lowerName, ext),
			Service:            fmt.Sprintf("src/data/services/%s.service.%s", lowerName, ext),
			Repository:         fmt.Sprintf("src/data/repositories/%s.repository.%s", lowerName, ext),
			Route:              fmt.Sprintf("src/presentation/routes/%s.routes.%s", lowerName, ext),
			ControllerUnitTest: fmt.Sprintf("tests/unit/controllers/%s.controller.test.%s", lowerName, ext),
			ServiceUnitTest:    fmt.Sprintf("tests/unit/services/%s.service.test.%s", lowerName, ext),
			RepositoryUnitTest: fmt.Sprintf("tests/unit/repositories/%s.repository.test.%s", lowerName, ext),
		}
	case "Hexagonal Architecture":
		repositoryPath := ""
		if orm == "TypeORM" {
			repositoryPath = fmt.Sprintf("src/adapters/typeorm/repositories/%s.repository.%s", lowerName, ext)
		} else {
			repositoryPath = fmt.Sprintf("src/adapters/repositories/%s.repository.%s", lowerName, ext)
		}

		return models.EntityPaths{
			Entity:             fmt.Sprintf("src/domain/entities/%s.entity.%s", lowerName, ext),
			Controller:         fmt.Sprintf("src/interfaces/controllers/%s.controller.%s", lowerName, ext),
			Service:            fmt.Sprintf("src/application/services/%s.service.%s", lowerName, ext),
			Repository:         repositoryPath,
			Route:              fmt.Sprintf("src/interfaces/routes/%s.routes.%s", lowerName, ext),
			ControllerUnitTest: fmt.Sprintf("tests/unit/controllers/%s.controller.test.%s", lowerName, ext),
			ServiceUnitTest:    fmt.Sprintf("tests/unit/services/%s.service.test.%s", lowerName, ext),
			RepositoryUnitTest: fmt.Sprintf("tests/unit/repositories/%s.repository.test.%s", lowerName, ext),
		}
	default: // Layered Architecture
		return models.EntityPaths{
			Route:      fmt.Sprintf("src/presentation/routes/%s.routes.%s", lowerName, ext),
			Controller: fmt.Sprintf("src/presentation/controllers/%s.controller.%s", lowerName, ext),

			Model:   fmt.Sprintf("src/business/models/%s.model.%s", lowerName, ext),
			Service: fmt.Sprintf("src/business/services/%s.service.%s", lowerName, ext),

			Repository: fmt.Sprintf("src/data/repositories/%s.repository.%s", lowerName, ext),
			RawModel:   fmt.Sprintf("src/data/models/%s.model.%s", lowerName, ext),

			ControllerUnitTest: fmt.Sprintf("tests/unit/presentation/%s.controller.test.%s", lowerName, ext),
			ServiceUnitTest:    fmt.Sprintf("tests/unit/business/%s.service.test.%s", lowerName, ext),
			RepositoryUnitTest: fmt.Sprintf("tests/unit/data/%s.repository.test.%s", lowerName, ext),

			ApiIntegrationTest: fmt.Sprintf("tests/integration/api/%s.integration.test.%s", lowerName, ext),
		}
	}
}

// createListFilesToGenerate return an array with all files to create
func createListFilesToGenerate(architecture string, entityConfig models.EntityConfig, paths models.EntityPaths) {

	switch architecture {
	case string(models.CleanArchitecture):
	case string(models.HexagonalArchitecture):
	default:

		// List all file to create
		filesToGenerate := []struct {
			path     string
			template string
			name     string
		}{
			{paths.Route, generate.GetRouteTemplate(entityConfig), "Route"},
			{paths.Controller, generate.GetControllerTemplate(entityConfig), "Controller"},
			{paths.Model, generate.GetModelTemplate(entityConfig), "Model"},
			{paths.Service, generate.GetServiceTemplate(entityConfig), "Service"},
			{paths.Repository, generate.GetRepositoryTemplate(entityConfig), "Repository"},
			{paths.RawModel, generate.GetRawModelTemplate(entityConfig), "Raw Model"},
			{paths.ControllerUnitTest, tests.GetControllerUnitTestTemplate(entityConfig), "Controller Unit Test"},
			{paths.ServiceUnitTest, tests.GetServiceUnitTestTemplate(entityConfig), "Service Unit Test"},
			{paths.RepositoryUnitTest, tests.GetRepositoryUnitTestTemplate(entityConfig), "Repository Unit Test"},
			{paths.ApiIntegrationTest, tests.GetApiIntegrationTestTemplate(entityConfig), "Api Integration. Test"},
		}

		GenerateAllFiles(filesToGenerate)
	}

}

// Generate all files in right folders
// [X] Layered - route.go
// [X] Layered - controller.go
// [X] Layered - service.go
// [X] Layered - repository.go
// [X] Layered - model.go
// [X] Layered - rawmodel.go
// [X] Layered - unittests.go
// [X] Layered - integrationtests.go
func GenerateAllFiles(filesToGenerate []struct {
	path     string
	template string
	name     string
}) error {
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
	return nil
}
