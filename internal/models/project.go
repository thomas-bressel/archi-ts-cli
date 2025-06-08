package models

import "time"

// Type af architecture
type Architecture string

const (
	LayeredArchitecture   Architecture = "Layered Architecture"
	CleanArchitecture     Architecture = "Clean Architecture"
	HexagonalArchitecture Architecture = "Hexagonal Architecture"
)

// Type of Language
type Language string

const (
	TypeScript Language = "typescript"
	JavaScript Language = "javascript"
)

// ProjectConfig type to defnie the configuration of the project
type ProjectConfigBuilder struct {
	Name         string       `json:"name"`
	Architecture Architecture `json:"architecture"`
	Language     Language     `json:"language"`
	Express      bool         `json:"express"`
}

/*
ProjectConfig is about the project pre-config
*/
type ProjectConfigFile struct {
	Name         string       `json:"name"`
	Version      string       `json:"version"`
	Architecture Architecture `json:"architecture"`
	Language     Language     `json:"language"`
	Express      bool         `json:"express"`
	CreatedAt    time.Time    `json:"created_at"`
}

// GetFileExtension return the language
func (p ProjectConfigBuilder) GetFileExtension() string {
	if p.Language == TypeScript {
		return "ts"
	}
	return "js"
}

// GetMainFile return the index main file
func (p ProjectConfigBuilder) GetMainFile() string {
	if p.Language == TypeScript {
		return "dist/index.js"
	}
	return "src/index.js"
}

func (p ProjectConfigBuilder) ToConfigFile(version string) ProjectConfigFile {
	return ProjectConfigFile{
		Name:         p.Name,
		Version:      version,
		Architecture: p.Architecture,
		Language:     p.Language,
		Express:      p.Express,
		CreatedAt:    time.Now(),
	}
}
