package prompts

import (
	"fmt"
	"strconv"

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

// PromptPort asks the number of listening port, 3000 is by default
func PromptPort() (int, error) {
	validate := func(input string) error {
		port, err := strconv.Atoi(input)
		if err != nil {
			return fmt.Errorf("invalid number: %v", err)
		}
		if port < 1024 || port > 65535 {
			return fmt.Errorf("port must be between 1024 and 65535")
		}

		// liste des ports interdits (exemple)
		forbidden := []int{3306, 5432, 6379, 27017}
		for _, p := range forbidden {
			if port == p {
				return fmt.Errorf("port %d is reserved and cannot be used", port)
			}
		}

		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Port Number",
		Default:  "3000",
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		return 0, err
	}

	port, err := strconv.Atoi(result)
	if err != nil {
		return 0, err
	}

	return port, nil
}
