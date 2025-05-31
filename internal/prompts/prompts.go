package prompts

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

/*
*	PromptProjectName to ask the name of the project.
 */
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

/*
* PromptLanguage ask the language type :  JavaScript or TypeScript
 */
func PromptLanguage() (string, error) {
	prompt := promptui.Select{
		Label: "Select a language",
		Items: []string{
			"TypeScript",
			"JavaScript (please don't, if you're a true developer)",
		},
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}:",
			Active:   color.New(color.FgYellow).Sprint("▸ {{ . | yellow }}"),
			Inactive: "  {{ . | faint }}",
			Selected: color.New(color.FgGreen).Sprint("✓ {{ . | green }}"),
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	// lowercase the result
	if result == "TypeScript" {
		return "typescript", nil
	}
	return "javascript", nil
}
