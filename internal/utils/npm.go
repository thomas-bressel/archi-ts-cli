package utils

import (
	"os/exec"

	"github.com/fatih/color"
)

// InstallDependencies install dependencies depending of the language
func InstallDependencies(language string, express bool) error {
	var dependencies []string
	var devDependencies []string

	// Common dependency
	dependencies = []string{
		"cors",
		"dotenv",
	}

	if language == "typescript" {
		// TypeScript dependancies
		devDependencies = []string{
			"typescript",
			"@types/node",
			"@types/cors",
			"ts-node",
			"nodemon",
			"@typescript-eslint/eslint-plugin",
			"@typescript-eslint/parser",
		}
	} else {
		// JavaScript dependencies
		devDependencies = []string{
			"nodemon",
			"eslint",
		}
	}

	// Add Express if selected in prompts
	if express {
		devDependencies = append(devDependencies, "@types/express")
		dependencies = append(dependencies, "express")
	}

	// Install main dependencies
	if len(dependencies) > 0 {
		color.New(color.FgBlue).Println("   Installing dependencies...")
		args := append([]string{"install"}, dependencies...)
		cmd := exec.Command("npm", args...)
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	// Install dev dependzencies
	if len(devDependencies) > 0 {
		color.New(color.FgBlue).Println("   Installing dev dependencies...")
		args := append([]string{"install", "--save-dev"}, devDependencies...)
		cmd := exec.Command("npm", args...)
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}

// InstallPnpm install global pnpm
func InstallPnpm() error {
	color.New(color.FgBlue).Println("   Installing pnpm...")
	cmd := exec.Command("sh", "-c", "curl -fsSL https://get.pnpm.io/install.sh | sh")
	return cmd.Run()
}
