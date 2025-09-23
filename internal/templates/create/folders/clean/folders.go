package clean

/*
getCleanDirectories returns the directory structure for Clean Architecture
*/
func GetDirectories() []string {
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
