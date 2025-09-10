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
		"database",
		"database/backups",
		"database/schema",
		"storage",
		"storage/uploads",
		"storage/avatars",
		"storage/documents",
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
