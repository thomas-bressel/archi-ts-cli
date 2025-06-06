package utils

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/fatih/color"
)

// DisplayBanner for a beautiful name banner !
func DisplayBanner() {

	banner := `
    ___              __    _ _______ _____     _____ _      _____ 
   /   |  __________/ /_  (_)__   __/ ___/    / ____| |    |_   _|
  / /| | / ___/ ___/ __ \/ /  | |   \__\     | |    | |      | |  
 / ___ |/ /  / /__/ / / / /   | |  ___/ /    | |____| |___  _| |_  
/_/  |_/_/   \___/_/ /_/_/    |_| /____/      \_____|_____||_____|  
`

	// Afficher le banner en couleur
	color.New(color.FgCyan, color.Bold).Print(banner)

	// Information
	fmt.Println()
	color.New(color.FgWhite, color.Bold).Println("ArchiTS CLI: 1.7.0")
	color.New(color.FgGreen).Printf("Node: %s\n", getNodeVersion())
	color.New(color.FgBlue).Printf("Go: %s\n", runtime.Version())
	color.New(color.FgYellow).Printf("OS: %s\n", runtime.GOOS)
	color.New(color.FgWhite, color.Bold).Println("Author: Thomas Bressel")
	color.New(color.FgMagenta).Println("Architecture Generator for Backend Projects")
	fmt.Println()
}

// ✅ Ajout des fonctions manquantes
func getNodeVersion() string {
	cmd := exec.Command("node", "--version")
	output, err := cmd.Output()
	if err != nil {
		return "Not installed"
	}
	return strings.TrimSpace(string(output))
}

func DisplayVersionTable() {
	color.New(color.FgCyan, color.Bold).Println("🚀 ArchiTS CLI - Project Architecture Generator")
	fmt.Println()

	// Table header
	color.New(color.FgWhite, color.Bold).Println("Environment Information:")
	color.New(color.FgWhite).Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	// ✅ Ajout des infos système
	printInfoLine("CLI Version", "1.7.0", color.FgGreen)
	printInfoLine("Go Version", runtime.Version(), color.FgBlue)
	printInfoLine("Node Version", getNodeVersion(), color.FgYellow)
	printInfoLine("OS", runtime.GOOS, color.FgMagenta)

	fmt.Println()
	color.New(color.FgWhite, color.Bold).Println("Available Commands:")
	color.New(color.FgWhite).Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	printCommandLine("create", "Generate a new project architecture")
	printCommandLine("version", "Display version information")
	printCommandLine("help", "Show help information")

	fmt.Println()
}

func printInfoLine(label, value string, valueColor color.Attribute) {
	color.New(color.FgWhite).Printf("%-15s", label+":")
	color.New(valueColor, color.Bold).Println(value)
}

func printCommandLine(command, description string) {
	color.New(color.FgCyan).Printf("  %-10s", command)
	color.New(color.FgWhite).Println(description)
}
