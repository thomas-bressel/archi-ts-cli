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

// EntityPaths contient tous les chemins des fichiers à générer
type EntityPaths struct {
	Entity     string
	Controller string
	Service    string
	Repository string
	Route      string
}

// runGenerateEntity is the function that will be called when the user run the command 'archi generate entity [name]'
func runGenerateEntity(cmd *cobra.Command, args []string) error {

	// Step 1 : Prepare and configuration
	entityName := args[0]

	// Uppcase the fist letter of entity name
	caser := cases.Title(language.English)
	entityName = caser.String(strings.ToLower(entityName))

	// Loading the project conf and checking if we are in a valid project
	projectConfig, err := config.LoadProjectConfig()
	if err != nil {
		return fmt.Errorf("impossible to load project configuration : %w\n must be sure to have created a project with 'archi create'", err)
	}

	// Creating the entity configuration
	entityConfig := generate.EntityConfig{
		Name:    entityName,
		Variant: string(projectConfig.Language),
	}

	// Managing file extension according to the language
	ext := "js"
	if projectConfig.Language == "typescript" {
		ext = "ts"
	}

	// Step 2 : Define the architecture paths
	paths := getEntityPaths(string(projectConfig.Architecture), entityName, ext)

	// Generating the files
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
	}

	// Step 3 : Create the files

	for _, file := range filesToGenerate {

		// Create the directory if not exists
		dir := filepath.Dir(file.path)
		if err := utils.CreateDirectory(dir); err != nil {
			return fmt.Errorf("error while creating the directory %s: %w", dir, err)
		}

		// Write the file
		if err := utils.WriteFile(file.path, file.template); err != nil {
			return fmt.Errorf("error while creating the file  %s: %w", file.path, err)
		}
		color.New(color.FgGreen).Printf("  ✅ %s created: %s\n", file.name, file.path)
	}

	// Step 4 : Success message
	color.New(color.FgGreen, color.Bold).Printf("✨ '%s' entity has been generated !\n", entityName)

	return nil
}

// getEntityPaths return the paths for each entity file based on the architecture
func getEntityPaths(architecture string, entityName string, ext string) EntityPaths {
	lowerName := strings.ToLower(entityName)

	switch architecture {
	case "Clean Architecture":
		return EntityPaths{
			Entity:     fmt.Sprintf("src/domain/entities/%s.entity.%s", lowerName, ext),
			Controller: fmt.Sprintf("src/presentation/controllers/%s.controller.%s", lowerName, ext),
			Service:    fmt.Sprintf("src/data/services/%s.service.%s", lowerName, ext),
			Repository: fmt.Sprintf("src/data/repositories/%s.repository.%s", lowerName, ext),
			Route:      fmt.Sprintf("src/presentation/routes/%s.routes.%s", lowerName, ext),
		}
	case "Hexagonal Architecture":
		return EntityPaths{
			Entity:     fmt.Sprintf("src/core/domain/entities/%s.entity.%s", lowerName, ext),
			Controller: fmt.Sprintf("src/adapters/inbound/http/controllers/%s.controller.%s", lowerName, ext),
			Service:    fmt.Sprintf("src/core/application/use-cases/%s.use-case.%s", lowerName, ext),
			Repository: fmt.Sprintf("src/adapters/outbound/repositories/%s.repository.%s", lowerName, ext),
			Route:      fmt.Sprintf("src/adapters/inbound/http/routes/%s.routes.%s", lowerName, ext),
		}
	default: // Layered Architecture
		return EntityPaths{
			Entity:     fmt.Sprintf("src/entities/%s.entity.%s", lowerName, ext),
			Controller: fmt.Sprintf("src/controllers/%s.controller.%s", lowerName, ext),
			Service:    fmt.Sprintf("src/services/%s.service.%s", lowerName, ext),
			Repository: fmt.Sprintf("src/repositories/%s.repository.%s", lowerName, ext),
			Route:      fmt.Sprintf("src/routes/%s.routes.%s", lowerName, ext),
		}
	}
}
