package config

import (
	"archi-ts-cli/internal/models"
	"encoding/json"
	"os"
	"path/filepath"
)

/*
SaveProjectConfig stores the project config
*/
func SaveProjectConfig(config models.ProjectConfigFile) error {
	configPath := filepath.Join(".", ".archi", "config.json")

	// Create .archi folder if it doesn't exists
	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return err
	}

	// JSON encode
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	// Fill the file
	return os.WriteFile(configPath, data, 0644)
}

/*
LoadProjectConfig load the project config
*/
func LoadProjectConfig() (*models.ProjectConfigFile, error) {
	configPath := filepath.Join(".", ".archi", "config.json")

	// check if the file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, err
	}

	// Read the file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	// JSOn decode
	var config models.ProjectConfigFile
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// GetDefaultConfig return the default config of an archiTS project
func GetDefaultConfig() models.ArchiConfig {
	return models.ArchiConfig{
		DefaultArchitecture: "Entity Clean Architecture",
		DefaultLanguage:     "typescript",
		TemplatesPath:       "./templates",
	}
}
