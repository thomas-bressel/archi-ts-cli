package project

import "fmt"

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
		"src/data/database/migrations",
		"src/data/database/seeds",

		// Common utils
		"src/common",
		"src/common/utils",
		"src/common/config",
		"src/common/constants",
		"src/common/errors",
		"src/common/logging",

		// Base directories
		"scripts",
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
		"tests/integration",
		"tests/integration/api",
		"tests/integration/database",
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
		"src/infrastructure/database/migrations",
		"src/infrastructure/cache",
		"src/infrastructure/email",
		"src/infrastructure/mappers",

		// Shared
		"src/shared",

		// Base directories
		"scripts",
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
	fmt.Println("Génération d'une architecture Hexagonale")
	return []string{
		"src",

		// Domain (hexagone center)
		"src/domain",
		"src/domain/entities",
		"src/domain/value-objects",
		"src/domain/events",
		"src/domain/exceptions",
		"src/domain/services",

		// Application (Use Cases)
		"src/application",
		"src/application/use-cases", // serevices
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
		"src/adapters/secondary/persistence/models",
		"src/adapters/secondary/persistence/orm",
		"src/adapters/secondary/persistence/repositories",
		"src/adapters/secondary/email",
		"src/adapters/secondary/cache",
		"src/adapters/secondary/storage",

		// Configuration
		"src/config",
		"src/config/dependencies", // dependencies injection

		// External directories
		"scripts",
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

		// Tests - Structure alignée avec src/
		"tests",
		"tests/unit",
		"tests/unit/domain",
		"tests/unit/domain/entities",
		"tests/unit/domain/value-objects",
		"tests/unit/domain/services",
		"tests/unit/application",
		"tests/unit/application/use-cases",
		"tests/unit/application/ports",
		"tests/unit/application/ports/in",
		"tests/unit/application/ports/out",

		"tests/integration",
		"tests/integration/adapters",
		"tests/integration/adapters/primary",
		"tests/integration/adapters/secondary",
		"tests/integration/application", // tests d'intégration use-cases

		"tests/functional", // tests bout-en-bout des fonctionnalités
		"tests/e2e",        // tests end-to-end complets

		".github",
		".github/workflows",
	}
}
