package cmd

import (
	"archi-ts-cli/cmd/prompts"
	"archi-ts-cli/internal/models"

	project "archi-ts-cli/internal/templates/create"
	"archi-ts-cli/internal/utils"
	"fmt"

	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Create the project with user CLI choices
// Architecture is the most important thing to know.
// Always ask for the choice of a library or not : If there is no library, then it will impact the following prompt choice
// ORM choice only if a library has been previously choosen
func runCreate(cmd *cobra.Command, args []string) error {
	// Display a beautifull Banner
	utils.DisplayBanner()
	color.New(color.FgCyan, color.Bold).Println("🚀 ArchiTS CLI - Project Scaffolding")
	fmt.Println()

	// Step 1: Start collecting the result of each prompts

	// Prompt 1 - Project name
	projectName, err := prompts.PromptProjectName()
	if err != nil {
		return fmt.Errorf("error with this project name: %w", err)
	}

	// Prompt 2 - Architecture Type
	architecture, err := prompts.PromptArchitecture()
	if err != nil {
		return fmt.Errorf("error when choosing an architecture: %w", err)
	}

	// Prompt 2 - Express Library?
	var express bool
	var library string
	if architecture != string(models.HexagonalArchitecture) {
		express, err = prompts.PromptExpress()
		if err != nil {
			return fmt.Errorf("error when choosing express library: %w", err)
		}
	} else {
		// Prompt 2alt - Which Library?
		library, err = prompts.PromptExpressOnly()
		if library == "express" {
			express = true
		}

		if err != nil {
			return fmt.Errorf("error when choosing express library: %w", err)
		}
	}

	// Prompt 3 - ORM Type (only if a library was chosen)
	var orm string
	if express || library != "" {
		orm, err = prompts.PromptOrm(architecture)
		if err != nil {
			return fmt.Errorf("error when choosing an ORM: %w", err)
		}
	}

	// Prompt 4 - Port listen
	port, err := prompts.PromptPort()
	if err != nil {
		return fmt.Errorf("error when entering listen port: %w", err)
	}

	// Step 2: Create the project architecture

	// Create the project directory
	projectPath := filepath.Join(".", projectName)
	if err := os.MkdirAll(projectPath, 0755); err != nil {
		return fmt.Errorf("error when trying to create the folder: %w", err)
	}

	// Select the project directory
	if err := os.Chdir(projectPath); err != nil {
		return fmt.Errorf("error during the change of directory: %w", err)
	}

	color.New(color.FgYellow).Printf("📁 Project created at: %s\n", projectPath)
	color.New(color.FgBlue).Printf("Architecture: %s\n", architecture)
	if orm != "none" {
		color.New(color.FgCyan).Printf("ORM: %s\n", orm)
	}
	color.New(color.FgGreen).Printf("Language: TypeScript\n")
	expressStatus := "No"
	if express {
		expressStatus = "Yes"
	}
	color.New(color.FgHiMagenta).Printf("Express Library: %s\n", expressStatus)
	color.New(color.FgHiMagenta).Printf("Library Name: %s\n", library)
	fmt.Println()

	// Step 3: Generation and dependencies installation

	//Store the configuration in a struct and generate the project
	config := models.ProjectConfigBuilder{
		Name:         projectName,
		Architecture: models.Architecture(architecture),
		Orm:          models.Orm(orm),
		Port:         port,
		Express:      express,
		Library:      library,
	}

	if err := project.GenerateProject(config); err != nil {
		return fmt.Errorf("error during the generation of the project: %w", err)
	}

	// Dependencies installation
	color.New(color.FgBlue).Println("📦 Dependencies installation...")
	if err := utils.InstallDependencies(); err != nil {
		color.New(color.FgRed).Printf("⚠️  Error during dependencies installation: %v\n", err)
	}

	color.New(color.FgBlue).Println("📦 Installation de pnpm...")
	if err := utils.InstallPnpm(); err != nil {
		color.New(color.FgRed).Printf("⚠️  Error during pnpm installation: %v\n", err)
	}

	// Step 4: Success messages
	fmt.Println()
	color.New(color.FgGreen, color.Bold).Println("✅ Project structure created successfully!")
	fmt.Println()
	color.New(color.FgCyan).Printf("📂 Project name: %s\n", projectName)
	color.New(color.FgCyan).Printf("Architecture: %s\n", architecture)
	color.New(color.FgCyan).Printf("Language: TypeScript\n")

	// if an orm exist
	if orm == "" || orm == "none" {
		color.New(color.FgCyan).Printf("ORM: %s\n", orm)
	}
	color.New(color.FgCyan).Printf("ExpressJS: %s\n", expressStatus)
	color.New(color.FgCyan).Printf("Library Name: %s\n", library)

	fmt.Println()
	color.New(color.FgYellow).Println("To start:")
	fmt.Printf("  cd %s\n", projectName)
	fmt.Printf("  npm run dev (local development with nodemon)\n")
	fmt.Printf("  npm run build (before npm start)\n")
	fmt.Printf("  npm start (after npm build)\n")
	fmt.Println()

	if orm != "none" {
		fmt.Printf("Don't forget to fill environnement DB_USERNAME and DB_PASSWORD in the .env file, before running npm command.\n")
	}
	fmt.Println()
	return nil
}
