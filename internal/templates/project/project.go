package project

import (
	"archi-ts-cli/internal/utils"
	"fmt"
)

// ProjectConfig type to defnie the configuration of the project
type ProjectConfig struct {
	Name         string
	Architecture string
	Language     string
}

/*
* GenerateProject to generate the project structure and files based on the configuration
 */
func GenerateProject(cfg ProjectConfig) error {
	// Create the project directory
	if err := createDirectoryStructure(cfg.Architecture); err != nil {
		return fmt.Errorf("erreur lors de la création des dossiers: %w", err)
	}
	return nil
}

/*
* createDirectoryStructure creates the directory structure based on the chosen architecture
 */
func createDirectoryStructure(architecture string) error {
	var directories []string

	switch architecture {
	case "Layered Architecture":
		directories = getLayeredDirectories()
	case "Clean Architecture":
		directories = getCleanDirectories()
	case "Hexagonal Architecture":
		directories = getHexagonalDirectories()
	default:
		directories = getLayeredDirectories() // Default
	}

	for _, dir := range directories {
		if err := utils.CreateDirectory(dir); err != nil {
			return err
		}
	}
	return nil
}

/*
* getLayeredDirectories returns the directory structure for Layered Architecture
 */
func getLayeredDirectories() []string {
	return []string{
		"src",
		"src/controllers",
		"src/controllers/base",
		"src/services",
		"src/services/base",
		"src/repositories",
		"src/repositories/base",
		"src/models",
		"src/models/base",
		"src/database",
		"src/database/connections",
		"src/database/migrations",
		"src/database/seeds",
		"src/middleware",
		"src/utils",
		"src/config",
		"src/routes",
		"src/routes/api",
		"database",
		"uploads",
		"logs",
	}
}

/*
* getCleanDirectories returns the directory structure for Clean Architecture
 */
func getCleanDirectories() []string {
	return []string{
		"src",
		// Domain Layer (Enterprise Business Rules)
		"src/domain",
		"src/domain/entities",
		"src/domain/entities/base",
		"src/domain/value-objects",
		"src/domain/repositories",
		"src/domain/services",

		// Application Layer (Application Business Rules)
		"src/application",
		"src/application/use-cases",
		"src/application/use-cases/user",
		"src/application/use-cases/role",
		"src/application/use-cases/permission",
		"src/application/dto",
		"src/application/dto/request",
		"src/application/dto/response",
		"src/application/ports",
		"src/application/ports/inbound",
		"src/application/ports/outbound",

		// Infrastructure Layer (Frameworks & Drivers)
		"src/infrastructure",
		"src/infrastructure/database",
		"src/infrastructure/database/repositories",
		"src/infrastructure/database/entities",
		"src/infrastructure/database/migrations",
		"src/infrastructure/database/seeds",
		"src/infrastructure/external-services",
		"src/infrastructure/web",
		"src/infrastructure/web/controllers",
		"src/infrastructure/web/middleware",
		"src/infrastructure/web/routes",
		"src/infrastructure/web/serializers",
		"src/infrastructure/config",

		// Main Layer (Composition Root)
		"src/main",
		"src/main/factories",

		// Base directories
		"database",
		"uploads",
		"logs",
	}
}

/*
* getHexagonalDirectories returns the directory structure for Hexagonal Architecture
 */
func getHexagonalDirectories() []string {
	return []string{
		"src",
		// Core (Application Hexagon)
		"src/core",
		"src/core/domain",
		"src/core/domain/models",
		"src/core/domain/value-objects",
		"src/core/domain/services",
		"src/core/domain/exceptions",
		"src/core/application",
		"src/core/application/services",
		"src/core/application/commands",
		"src/core/application/queries",
		"src/core/application/handlers",
		"src/core/application/handlers/commands",
		"src/core/application/handlers/queries",

		// Ports (Interfaces)
		"src/ports",
		"src/ports/inbound",
		"src/ports/outbound",
		"src/ports/outbound/repositories",
		"src/ports/outbound/external-services",
		"src/ports/outbound/infrastructure",

		// Adapters (Implementations)
		"src/adapters",
		"src/adapters/inbound",
		"src/adapters/inbound/http",
		"src/adapters/inbound/http/controllers",
		"src/adapters/inbound/http/middleware",
		"src/adapters/inbound/http/routes",
		"src/adapters/inbound/http/serializers",
		"src/adapters/inbound/cli",
		"src/adapters/inbound/cli/commands",
		"src/adapters/inbound/message-queue",
		"src/adapters/inbound/message-queue/consumers",

		// Outbound Adapters
		"src/adapters/outbound",
		"src/adapters/outbound/repositories",
		"src/adapters/outbound/repositories/database",
		"src/adapters/outbound/repositories/database/mysql",
		"src/adapters/outbound/repositories/database/postgresql",
		"src/adapters/outbound/repositories/database/entities",
		"src/adapters/outbound/repositories/in-memory",
		"src/adapters/outbound/repositories/file-system",
		"src/adapters/outbound/external-services",
		"src/adapters/outbound/external-services/email",
		"src/adapters/outbound/external-services/hash",
		"src/adapters/outbound/external-services/events",
		"src/adapters/outbound/external-services/logging",
		"src/adapters/outbound/infrastructure",
		"src/adapters/outbound/infrastructure/database",
		"src/adapters/outbound/infrastructure/message-queue",
		"src/adapters/outbound/infrastructure/file-storage",

		// Configuration
		"src/config",
		"src/config/dependencies",
		"src/config/database",
		"src/config/database/migrations",
		"src/config/environment",

		// Main (Bootstrap)
		"src/main",

		// Base directories
		"database",
		"uploads",
		"logs",
	}
}
