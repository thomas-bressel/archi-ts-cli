package prompts

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

// PromptProjectName to ask the name of the project
func PromptProjectName() (string, error) {
	prompt := promptui.Prompt{
		Label:   "Project Name",
		Default: "archi-project",
		Validate: func(input string) error {
			if len(input) < 1 {
				return fmt.Errorf("project name cannot be empty")
			}
			return nil
		},
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return result, nil
}

// PromptArchitecture ask the architecture type
func PromptArchitecture() (string, error) {
	prompt := promptui.Select{
		Label: "Select an architecture",
		Items: []string{
			"Layered Architecture",
			"Clean Architecture",
			"Hexagonal Architecture",
		},
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}:",
			Active:   color.New(color.FgCyan).Sprint("▸ {{ . | cyan }}"),
			Inactive: "  {{ . | faint }}",
			Selected: color.New(color.FgGreen).Sprint("✔ {{ . | green }}"),
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

// PromptORM ask Whick ORM to use
func PromptOrm() (string, error) {
	prompt := promptui.Select{
		Label: "Select an ORM",
		Items: []string{
			"I don't want to use any ORM",
			"TypeORM",
		},
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}:",
			Active:   color.New(color.FgCyan).Sprint("▸ {{ . | cyan }}"),
			Inactive: "  {{ . | faint }}",
			Selected: color.New(color.FgGreen).Sprint("✔ {{ . | green }}"),
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

// PromptExpress asks if the user wants to use ExpressJS library
func PromptExpress() (bool, error) {
	prompt := promptui.Select{
		Label: "Do you want to use ExpressJS library?",
		Items: []string{
			"No, I don't need Express",
			"Yes, install Express",
		},
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}:",
			Active:   color.New(color.FgYellow).Sprint("▸ {{ . | yellow }}"),
			Inactive: "  {{ . | faint }}",
			Selected: color.New(color.FgGreen).Sprint("✔ {{ . | green }}"),
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return false, err
	}

	// Return true if user wants Express, false otherwise
	return result == "Yes, install Express", nil
}
