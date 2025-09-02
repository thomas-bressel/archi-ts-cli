package models

import "time"

// Type of architecture
type Architecture string

const (
	LayeredArchitecture   Architecture = "Layered Architecture"
	CleanArchitecture     Architecture = "Clean Architecture"
	HexagonalArchitecture Architecture = "Hexagonal Architecture"
)

// ProjectConfig type to define the configuration of the project
type ProjectConfigBuilder struct {
	Name         string       `json:"name"`
	Architecture Architecture `json:"architecture"`
	Express      bool         `json:"express"`
}

// ProjectConfig is about the project pre-config
type ProjectConfigFile struct {
	Name         string       `json:"name"`
	Version      string       `json:"version"`
	Architecture Architecture `json:"architecture"`
	Express      bool         `json:"express"`
	CreatedAt    time.Time    `json:"created_at"`
}

// GetFileExtension returns ts extension
func (p ProjectConfigBuilder) GetFileExtension() string {
	return "ts"
}

// GetMainFile returns the index main file
func (p ProjectConfigBuilder) GetMainFile() string {
	return "dist/index.js"
}

func (p ProjectConfigBuilder) ToConfigFile(version string) ProjectConfigFile {
	return ProjectConfigFile{
		Name:         p.Name,
		Version:      version,
		Architecture: p.Architecture,
		Express:      p.Express,
		CreatedAt:    time.Now(),
	}
}
