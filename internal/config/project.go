package config

import (
	"archi-ts-cli/internal/models"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	ConfigDir  = ".archi"
	ConfigFile = "config.json"
)

// SaveProjectConfig saves the project configuration into .archi/config.json
func SaveProjectConfig(config models.ProjectConfigFile) error {
	// Create the .archi directory if it doesn't exist
	configPath := filepath.Join(".", ConfigDir)
	if err := os.MkdirAll(configPath, 0755); err != nil {
		return fmt.Errorf("unable to create config directory: %w", err)
	}

	// Convert the configuration to JSON
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to marshal configuration: %w", err)
	}

	// Write the configuration file
	configFilePath := filepath.Join(configPath, ConfigFile)
	if err := os.WriteFile(configFilePath, data, 0644); err != nil {
		return fmt.Errorf("unable to write configuration file: %w", err)
	}

	return nil
}

// LoadProjectConfig loads the project configuration from .archi/config.json
func LoadProjectConfig() (*models.ProjectConfigFile, error) {
	configFilePath := filepath.Join(".", ConfigDir, ConfigFile)

	// Check if the file exists
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("configuration file not found")
	}

	// Read the file
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read configuration file: %w", err)
	}

	// Parse the JSON
	var config models.ProjectConfigFile
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("unable to parse configuration: %w", err)
	}

	return &config, nil
}

// ProjectExists checks if an Archi project exists in the current directory
func ProjectExists() bool {
	configFilePath := filepath.Join(".", ConfigDir, ConfigFile)
	_, err := os.Stat(configFilePath)
	return err == nil
}

// GetProjectArchitecture returns the architecture type of the project
func GetProjectArchitecture() (string, error) {
	config, err := LoadProjectConfig()
	if err != nil {
		return "", err
	}
	return string(config.Architecture), nil
}
