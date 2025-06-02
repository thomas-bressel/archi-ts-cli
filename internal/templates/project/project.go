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

		// Controllers Layer
		"src/controllers",
		"src/controllers/base",

		// Services Layer
		"src/services",
		"src/services/base",

		// Repositories Layer
		"src/repositories",
		"src/repositories/base",

		// Models Layer
		"src/models",
		"src/models/entities",
		"src/models/dtos",
		"src/models/requests",
		"src/models/responses",
		"src/models/database",

		// Data Access Layer
		"src/database",
		"src/database/connections",
		"src/database/migrations",
		"src/database/seeds",
		"src/database/queries",

		// Infrastructure Layer
		"src/infrastructure",
		"src/infrastructure/cache",
		"src/infrastructure/email",
		"src/infrastructure/modules",

		// Middleware Layer
		"src/middleware",
		"src/middleware/auth",
		"src/middleware/validation",
		"src/middleware/security",

		// Routes Layer
		"src/routes",
		"src/routes/api",
		"src/routes/web",

		// Utilities & Config
		"src/utils",
		"src/config",
		"src/constants",
		"src/mappers",

		// External directories
		"database",
		"database/backups",
		"database/schema",
		"storage",
		"storage/uploads",
		"storage/uploads/avatars",
		"storage/uploads/documents",
		"storage/logs",
		"tests",
		"tests/unit",
		"tests/integration",
		"tests/functional",
		"tests/e2e",
		".github",
		".github/workflows",
	}
}

/*
getCleanDirectories returns the directory structure for Clean Architecture
*/
func getCleanDirectories() []string {
	return []string{
		"src",
		// Domain
		"src/domain",
		"src/domain/entities",

		// Data
		"src/data",
		"src/data/repositories",
		"src/data/data-sources",
		"src/data/data-sources/local",
		"src/data/services",
		"src/data/dtos",
		"src/data/models",
		"src/data/mappers",

		// Presentation
		"src/presentation",
		"src/presentation/controllers",
		"src/presentation/routes",
		"src/presentation/middlewares",
		"src/presentation/models",

		// Main Layer (Composition Root)
		"src/infrastructure",
		"src/infrastructure/database",
		"src/infrastructure/cache",
		"src/infrastructure/email",
		"src/infrastructure/server",

		// Shared
		"src/shared",
		"src/shared/utils",
		"src/shared/constants",

		// Base directories
		"database",
		"database/backups",
		"database/schema",
		"storage",
		"storage/uploads",
		"storage/avatars",
		"storage/documents",
		"__tests__",
		"__tests__/unit",
		"__tests__/functional",
		"logs",
		".github/",
		".github/workflows",
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
		"src/core/domain/entities",
		"src/core/domain/value-objects",
		"src/core/domain/services",
		"src/core/domain/exceptions",
		"src/core/application",
		"src/core/application/use-cases",
		"src/core/application/commands",
		"src/core/application/queries",
		"src/core/application/handlers",
		"src/core/application/handlers/commands",
		"src/core/application/handlers/queries",
		"src/core/application/dtos",

		// Ports (Interfaces)
		"src/ports",
		"src/ports/inbound",
		"src/ports/inbound/http",
		"src/ports/inbound/cli",
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
		"src/adapters/inbound/http/models",
		"src/adapters/inbound/cli",
		"src/adapters/inbound/cli/commands",

		// Outbound Adapters
		"src/adapters/outbound",
		"src/adapters/outbound/repositories",
		"src/adapters/outbound/repositories/mysql",
		"src/adapters/outbound/repositories/redis",
		"src/adapters/outbound/repositories/entities",
		"src/adapters/outbound/repositories/queries",
		"src/adapters/outbound/repositories/mappers",
		"src/adapters/outbound/external-services",
		"src/adapters/outbound/external-services/email",
		"src/adapters/outbound/external-services/auth",
		"src/adapters/outbound/external-services/storage",
		"src/adapters/outbound/infrastructure",
		"src/adapters/outbound/infrastructure/database",
		"src/adapters/outbound/infrastructure/cache",
		"src/adapters/outbound/infrastructure/server",

		// Shared/Common
		"src/shared",
		"src/shared/utils",
		"src/shared/constants",
		"src/shared/exceptions",
		"src/shared/types",

		// Configuration & DI
		"src/config",
		"src/config/dependencies",
		"src/config/database",
		"src/config/environment",

		// Main (Bootstrap)
		"src/main",

		// External directories
		"database",
		"database/backups",
		"database/schema",
		"database/migrations",
		"database/seeds",
		"storage",
		"storage/uploads",
		"storage/uploads/avatars",
		"storage/uploads/documents",
		"storage/logs",
		"tests",
		"tests/unit",
		"tests/integration",
		"tests/functional",
		"tests/e2e",
		".github",
		".github/workflows",
	}

}
