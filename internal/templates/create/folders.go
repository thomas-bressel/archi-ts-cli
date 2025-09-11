package project

import "archi-ts-cli/internal/models"

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
		"tests/unit/presnetation",
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

		// Adapters (vide pour le moment, TypeORM sera ajouté dynamiquement si choisi)
		"src/adapters",

		// Application
		"src/application",
		"src/application/dto",
		"src/application/services",

		// Domain
		"src/domain",
		"src/domain/entities",
		"src/domain/ports",

		// Interfaces
		"src/interfaces",
		"src/interfaces/controllers",
		"src/interfaces/routes",

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

/*
 * getHexagonalORMDirectories returns ORM-specific directories for Hexagonal Architecture
 */
func getHexagonalORMDirectories(orm models.Orm) []string {
	switch orm {
	case models.TypeOrm:
		return []string{
			"src/adapters/typeorm",
			"src/adapters/typeorm/repositories",
		}
	default:
		return []string{}
	}
}
