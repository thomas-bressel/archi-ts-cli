package models

import "time"

// Type of architecture
type Architecture string
type Orm string

const (
	LayeredArchitecture   Architecture = "Layered Architecture"
	CleanArchitecture     Architecture = "Clean Architecture"
	HexagonalArchitecture Architecture = "Hexagonal Architecture"
)
const (
	TypeOrm Orm = "TypeORM"
	Prisma  Orm = "Prisma"
)

// ProjectConfig type to define the configuration of the project
type ProjectConfigBuilder struct {
	Name         string       `json:"name"`
	Architecture Architecture `json:"architecture"`
	Orm          Orm          `json:"orm"`
	Port         int          `json:"port"`
	Express      bool         `json:"express"`
	Library      string       `json:"library"`
}

// ProjectConfig is about the project pre-config
type ProjectConfigFile struct {
	Name         string       `json:"name"`
	Version      string       `json:"version"`
	Architecture Architecture `json:"architecture"`
	Orm          Orm          `json:"orm"`
	Express      bool         `json:"express"`
	Library      string       `json:"library"`
	Port         int          `json:"port"`
	CreatedAt    time.Time    `json:"created_at"`
}

// PackageJSON represents the structure of package.json with ordered fields
type PackageJSON struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Description     string            `json:"description"`
	Main            string            `json:"main"`
	Scripts         map[string]string `json:"scripts"`
	Keywords        []string          `json:"keywords"`
	Author          string            `json:"author"`
	License         string            `json:"license"`
	Dependencies    map[string]string `json:"dependencies,omitempty"`
	DevDependencies map[string]string `json:"devDependencies"`
}

type LibraryOption struct {
	Display string
	ID      string
}

type ORMOption struct {
	Display string
	ID      string
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
		Orm:          p.Orm,
		Express:      p.Express,
		Library:      p.Library,
		Port:         p.Port,
		CreatedAt:    time.Now(),
	}
}
