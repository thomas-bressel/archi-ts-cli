package project

// getLayeredDirectories returns the directory structure for a refined Layered Architecture
func getLayeredDirectories() []string {
	return []string{
		"src",
		// Presentation
		"src/presentation",
		"src/presentation/controllers",
		"src/presentation/routes",
		"src/presentation/middlewares",
		"src/presentation/middlewares/auth",
		"src/presentation/middlewares/validation",
		"src/presentation/middlewares/security",

		// Business
		"src/business",
		"src/business/services",
		"src/business/models",
		"src/business/interfaces",

		// Data
		"src/data",
		"src/data/repositories",
		"src/data/models",
		"src/data/database",
		"src/data/database/connection",
		"src/data/database/migration",
		"src/data/database/seeds",

		// Common utils
		"src/common",
		"src/common/utils",
		"src/common/config",
		"src/common/constants",
		"src/common/errors",
		"src/common/logging",

		// Base directories
		"storage",
		"storage/uploads",
		"storage/avatars",
		"storage/documents",
		"storage/database",
		"storage/database/backups",
		"storage/database/schema",
		"tests",
		"tests/unit",
		"tests/unit/presentation",
		"tests/unit/business",
		"tests/unit/data",
		"tests/functional",
		"tests/intergation",
		"tests/intergation/api",
		"tests/intergation/database",
		"tests/e2e",
		"tests/e2e/api",
		"logs",
		".github/",
		".github/workflows",
	}
}

/*
getCleanDirectories returns the directory structure for Clean Architecture
*/
func getCleanDirectories() []string {
	return []string{
		"src",

		// 1. Domain (business entities)
		"src/domain",
		"src/domain/entities",
		"src/domain/value-objects",
		"src/domain/errors",

		// 2. Application/Use Cases (business logic)
		"src/application",
		"src/application/use-cases",
		"src/application/interfaces", // repositories interfaces
		"src/application/dtos",

		// 4. Presentation (Web)
		"src/presentation",
		"src/presentation/controllers",
		"src/presentation/routes",
		"src/presentation/middlewares",
		"src/presentation/validators",

		// 3. Infrastructure (Implementations)
		"src/infrastructure",
		"src/infrastructure/repositories", // Real Implementations
		"src/infrastructure/database",
		"src/infrastructure/database/config",
		"src/infrastructure/cache",
		"src/infrastructure/email",
		"src/infrastructure/mappers",

		// Shared
		"src/shared",

		// Base directories
		"storage",
		"storage/uploads",
		"storage/avatars",
		"storage/documents",
		"storage/database",
		"storage/database/backups",
		"storage/database/schema",
		"tests",

		// 1. Unit test by layers
		"tests/unit",
		"tests/unit/domain", // Entities, Value Objects
		"tests/unit/domain/entities",
		"tests/unit/domain/value-objects",

		"tests/unit/application", // Use Cases (IMPORTANT!)
		"tests/unit/application/use-cases",

		"tests/unit/infrastructure", // Repositories, exernal services
		"tests/unit/infrastructure/repositories",
		"tests/unit/infrastructure/services",

		"tests/unit/presentation", // Controllers, validators
		"tests/unit/presentation/controllers",
		"tests/unit/presentation/middlewares",

		// 2. Integration tests
		"tests/integration",
		"tests/integration/repositories", // real BDD
		"tests/integration/api",          // Routes and controllers
		"tests/integration/use-cases",    // Use cases with true dependancies
		"tests/integration/external",     // other services (email, etc.)

		// 3. E2E Tests
		"tests/e2e",
		"tests/e2e/scenarios", // Full scenari
		"tests/e2e/api",       // full API

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

		// Domain (hexagone center)
		"src/domain",
		"src/domain/entities",
		"src/domain/value-objects",
		"src/domain/events",
		"src/domain/exceptions",

		// Application (Use Cases)
		"src/application",
		"src/application/use-cases", // serevicfes
		"src/application/ports",     // interfaces / contracts
		"src/application/ports/in",  // entering ports (commands)
		"src/application/ports/out", // exit ports  (persistence)
		"src/application/dtos",

		// Primary adapters (left side - Enter)
		"src/adapters/primary",
		"src/adapters/primary/http", // API REST
		"src/adapters/primary/http/controllers",
		"src/adapters/primary/http/routes",
		"src/adapters/primary/http/middlewares",
		"src/adapters/primary/cli",  // Interface CLI
		"src/adapters/primary/grpc", // If gRPC

		// Secondary adapters (right side - Exit)
		"src/adapters/secondary",
		"src/adapters/secondary/persistence",
		"src/adapters/secondary/persistence/orm",
		"src/adapters/secondary/persistence/orm/entities",
		"src/adapters/secondary/persistence/orm/repositories",
		"src/adapters/secondary/email",
		"src/adapters/secondary/cache",
		"src/adapters/secondary/storage",

		// Configuration
		"src/config",
		"src/config/dependencies", // dependencies injection

		// External directories

		"storage",
		"storage/database",
		"storage/database/backups",
		"storage/database/schema",
		"storage/database/migrations",
		"storage/database/seeds",
		"storage/uploads",
		"storage/uploads/avatars",
		"storage/uploads/documents",
		"storage/logs",
		"tests/unit/core",
		"tests/unit/core/entities",
		"tests/unit/core/value-objects",
		"tests/unit/core/domain-services",
		"tests/unit/core/specifications",
		"tests/unit/ports",
		"tests/unit/ports/primary",   // Ports primaires (driving)
		"tests/unit/ports/secondary", // Ports secondaires (driven)
		"tests/integration/adapters",
		"tests/integration/adapters/primary",   // Adapters primaires (API, CLI, Web)
		"tests/integration/adapters/secondary", // Adapters secondaires (BDD, Services)
		"tests/integration/infrastructure",
		"tests/integration/infrastructure/persistence", // Real database
		"tests/integration/infrastructure/messaging",   // Queues, Events
		"tests/integration/infrastructure/external",
		"tests/integration/infrastructure/cache", // Cache (Redis, etc.)
		"tests/integration/application",
		"tests/integration/application/use-cases", // Use cases
		"tests/integration/application/workflows",

		"tests",
		"tests/unit",
		"tests/integration",
		"tests/functional",
		"tests/e2e",
		".github",
		".github/workflows",
	}
}
