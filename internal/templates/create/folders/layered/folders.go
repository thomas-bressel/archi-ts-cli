package layered

// getLayeredDirectories returns the directory structure for a refined Layered Architecture
func GetDirectories() []string {
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
