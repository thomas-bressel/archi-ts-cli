package project

/*
 * getLayeredDirectories returns the directory structure for a refined Layered Architecture
 */
func getLayeredDirectories() []string {
	return []string{
		"src",

		// Config & Core
		"src/config",    // configuration (db, env, logger)
		"src/errors",    // custom errors & handlers
		"src/utils",     // utilities & helpers
		"src/constants", // constants
		"src/mappers",   // entity <-> dto mappers

		// Database Layer
		"src/database",
		"src/database/connections",
		"src/database/migrations",
		"src/database/seeds",
		"src/database/queries",

		// Entities (Domain Models)
		"src/entities",

		// DTOs (Data Transfer Objects)
		"src/dtos",
		"src/dtos/requests",
		"src/dtos/responses",

		// Repositories Layer
		"src/repositories",
		"src/repositories/base",

		// Services Layer
		"src/services",
		"src/services/base",

		// Controllers Layer
		"src/controllers",
		"src/controllers/base",

		// Middleware Layer
		"src/middleware",
		"src/middleware/auth",
		"src/middleware/validation",
		"src/middleware/security",

		// Routes Layer
		"src/routes",
		"src/routes/api",
		"src/routes/web",

		// Infrastructure (external integrations)
		"src/infrastructure",
		"src/infrastructure/cache",
		"src/infrastructure/email",
		"src/infrastructure/modules",

		// External directories
		"database",
		"database/backups",
		"database/schema",
		"storage",
		"storage/uploads",
		"storage/uploads/avatars",
		"storage/uploads/documents",
		"storage/logs",

		// Tests
		"tests",
		"tests/unit",
		"tests/integration",
		"tests/functional",
		"tests/e2e",

		// CI/CD
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
