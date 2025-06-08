package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

/*
ProjectConfig is about the project config
*/
type ProjectConfig struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	Architecture string `json:"architecture"`
	Language     string `json:"language"`
	Express      bool   `json:"express"`
	CreatedAt    string `json:"created_at"`
}

/*
ArchiConfig is about the whole CLI config
*/
type ArchiConfig struct {
	DefaultArchitecture string `json:"default_architecture"`
	DefaultLanguage     string `json:"default_language"`
	TemplatesPath       string `json:"templates_path"`
}

/*
SaveProjectConfig stores the project config
*/
func SaveProjectConfig(config ProjectConfig) error {
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
func LoadProjectConfig() (*ProjectConfig, error) {
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
	var config ProjectConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// GetDefaultConfig return the default config of an archiTS project
func GetDefaultConfig() ArchiConfig {
	return ArchiConfig{
		DefaultArchitecture: "Entity Clean Architecture",
		DefaultLanguage:     "typescript",
		TemplatesPath:       "./templates",
	}
}
