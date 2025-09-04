package utils

import (
	"os/exec"

	"github.com/fatih/color"
)

// InstallDependencies install dependencies for TypeScript only
func InstallDependencies() error {
	var dependencies []string
	var devDependencies []string

	// Common dependencies
	dependencies = []string{
		"cors",
		"dotenv",
	}

	// TypeScript dependencies
	devDependencies = []string{
		"typescript",
		"@types/node",
		"@types/cors",
		"ts-node",
		"nodemon",
		"@typescript-eslint/eslint-plugin",
		"@typescript-eslint/parser",
		"@types/express",
		"express",
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

	// Install dev dependencies
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
