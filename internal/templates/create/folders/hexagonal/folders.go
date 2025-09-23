package hexagonal

/*
 * getHexagonalDirectories returns the directory structure for Hexagonal Architecture
 */
func GetDirectories() []string {
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
