package utils

import (
	"os/exec"

	"github.com/fatih/color"
)

// InstallDependencies install dependencies for TypeScript only
func InstallDependencies() error {

	// Install main dependencies
	color.New(color.FgBlue).Println("   Installing dependencies & dev dependencies ...")
	cmd := exec.Command("npm", "install")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// InstallPnpm install global pnpm
func InstallPnpm() error {
	color.New(color.FgBlue).Println("   Installing pnpm...")
	cmd := exec.Command("sh", "-c", "curl -fsSL https://get.pnpm.io/install.sh | sh")
	return cmd.Run()
}
