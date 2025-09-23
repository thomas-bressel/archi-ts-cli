package files

import (
	"archi-ts-cli/internal/models"
	"encoding/json"
	"fmt"
)

// CreatePackageJson generate the package.json file content for Layered and Clean architectures
func CreatePackageJson(cfg models.ProjectConfigBuilder) (string, error) {

	// Header generation
	pkg := models.PackageJSON{
		Name:        cfg.Name,
		Version:     "1.0.0",
		Description: getDescription(cfg),
		Main:        getMain(cfg.Architecture),
		Scripts:     getScripts(cfg.Architecture, cfg.Orm),
		Keywords:    getKeywords(cfg.Architecture, cfg.Express, cfg.Orm),
		Author:      "",
		License:     "MIT",
	}

	// Global dependencies generation
	pkg.Dependencies = getDependencies(cfg.Orm, cfg.Express)

	// Global development dependencies
	// Later it will be replaced by No library, Express of Fatify
	pkg.DevDependencies = getDevDependencies(cfg.Express)

	jsonData, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

// Return a basic description with details
func getDescription(cfg models.ProjectConfigBuilder) string {
	library := "no library"
	if cfg.Express {
		library = "Express"
	}
	orm := "no ORM"
	if cfg.Orm == "typeorm" {
		orm = "Type ORM"
	}

	architecture := "None"
	switch cfg.Architecture {
	case models.CleanArchitecture:
		architecture = "clean"
	case models.LayeredArchitecture:
		architecture = "layered"
	case models.HexagonalArchitecture:
		architecture = "hexagonal"
	default:
		architecture = "None"
	}

	description := fmt.Sprintf(
		"Project '%s' uses a %s architecture, %s and %s library.",
		cfg.Name,
		architecture,
		orm,
		library,
	)

	return description
}

// Return the production entry file path
func getMain(architecture models.Architecture) string {
	var main string
	switch architecture {
	case models.LayeredArchitecture:
		main = "dist/index.js"
	case models.CleanArchitecture:
		main = "dist/index.js"
	case models.HexagonalArchitecture:
		main = "index.js"
	default:
		main = "dist/index.js"
	}
	return main
}

// getScripts returns npm scripts for TypeScript (depending architectures type)
func getScripts(architecture models.Architecture, orm models.Orm) map[string]string {
	var scripts map[string]string

	switch architecture {
	case models.HexagonalArchitecture:
		scripts = map[string]string{
			"build":         "tsc",
			"dev":           "nodemon --exec ts-node -r tsconfig-paths/register src/index.ts",
			"lint":          "eslint src/**/*.ts",
			"lint:fix":      "eslint src/**/*.ts --fix",
			"start":         "node dist/src/index.js",
			"test":          "jest",
			"test:coverage": "jest --coverage",
			"test:watch":    "jest --watch",
			"watch":         "tsc --watch",
		}
	default:
		scripts = map[string]string{
			"build":         "tsc",
			"dev":           "nodemon --watch 'src/**/*.ts' --exec 'ts-node -r tsconfig-paths/register' src/index.ts",
			"lint":          "eslint src/**/*.ts",
			"lint:fix":      "eslint src/**/*.ts --fix",
			"start":         "node dist/src/index.js",
			"test":          "jest",
			"test:coverage": "jest --coverage",
			"test:watch":    "jest --watch",
			"watch":         "tsc --watch",
		}
	}

	// Add TypeORM migration scripts if needed
	if orm == "typeorm" {
		typeOrmScripts := getTypeOrmMigrationScripts(architecture)
		for key, value := range typeOrmScripts {
			scripts[key] = value
		}
	}

	return scripts
}

// getTypeOrmMigrationScripts returns TypeORM specific migration scripts
func getTypeOrmMigrationScripts(architecture models.Architecture) map[string]string {
	migrationPath := ""
	datasourcePath := ""

	switch architecture {
	case models.LayeredArchitecture:
		migrationPath = "src/data/database"
		datasourcePath = "connection"
	case models.CleanArchitecture:
		migrationPath = "src/infrastructure/database"
		datasourcePath = "config"
	case models.HexagonalArchitecture:
		migrationPath = "src/adapters/secondary/persistence"
		datasourcePath = "orm"
	}

	return map[string]string{
		"migration:generate": "TS_NODE_PROJECT=./tsconfig.json ts-node -r tsconfig-paths/register scripts/generate-migration.ts",
		"migration:run":      fmt.Sprintf(`TS_NODE_PROJECT=./tsconfig.json typeorm-ts-node-commonjs migration:run -d ./%s/%s/data-source.ts`, migrationPath, datasourcePath),
		"migration:revert":   fmt.Sprintf(`TS_NODE_PROJECT=./tsconfig.json typeorm-ts-node-commonjs migration:run -d ./%s/%s/data-source.ts`, migrationPath, datasourcePath),
		"migration:show":     fmt.Sprintf(`TS_NODE_PROJECT=./tsconfig.json typeorm-ts-node-commonjs migration:run -d ./%s/%s/data-source.ts`, migrationPath, datasourcePath),
		"m:gen":              "npm run migration:generate",
		"m:run":              "npm run migration:run",
		"m:revert":           "npm run migration:revert",
		"m:show":             "npm run migration:show",
	}
}

// getKeywords returns keywords based on Express usage (standard architectures)
func getKeywords(architecture models.Architecture, express bool, orm models.Orm) []string {

	// Default keywords list
	baseKeywords := []string{"api", "backend", "nodejs", "typescript"}

	// Express.js
	if express {
		baseKeywords = append(baseKeywords, "express")
	}

	// If hexagonal architecture
	if architecture == models.HexagonalArchitecture {
		baseKeywords = append(baseKeywords, "hexagonal", "ports-adapters")
	}

	// If typeORM
	if orm == "typeorm" {
		baseKeywords = append(baseKeywords, "typeorm")
	}
	return baseKeywords
}

// getDependencies returns runtime dependencies (standard architectures)
func getDependencies(orm models.Orm, express bool) map[string]string {
	deps := make(map[string]string)

	// Dependencies enable on all architecture
	deps["class-validator"] = "^0.14.1"
	deps["class-transformer"] = "^0.5.1"
	deps["dotenv"] = "^16.6.1"
	deps["cors"] = "^2.8.5"
	deps["tsconfig-paths"] = "^4.2.0"

	// TypeORM needs
	if orm == "typeorm" {
		deps["reflect-metadata"] = "^0.2.2"
		deps["typeorm"] = "^0.3.26"
		deps["mysql2"] = "^3.14.4"
		deps["sqlite3"] = "^5.1.7"
		deps["inversify"] = "^7.10.0"
	}

	// Express.js
	if express {
		deps["express"] = "^4.18.2"
	}

	return deps
}

// getDevDependencies returns development dependencies for TypeScript (standard architectures)
func getDevDependencies(express bool) map[string]string {

	// Dependencies enable on all architecture
	devDeps := map[string]string{
		"nodemon":                          "^3.0.2",
		"typescript":                       "^5.3.3",
		"eslint":                           "^8.56.0",
		"jest":                             "^29.7.0",
		"ts-jest":                          "^29.1.1",
		"ts-node":                          "^10.9.2",
		"ts-node-dev":                      "^2.0.0",
		"tsconfig-paths":                   "^4.2.0",
		"supertest":                        "^6.3.3",
		"@types/validator":                 "^13.12.2",
		"@types/cors":                      "^2.8.19",
		"@types/node":                      "^20.10.5",
		"@typescript-eslint/eslint-plugin": "^6.15.0",
		"@typescript-eslint/parser":        "^6.15.0",
		"@types/jest":                      "^29.5.14",
		"@types/supertest":                 "^2.0.12",
	}

	// Express.js
	if express {
		devDeps["@types/express"] = "^4.17.21"
	}

	return devDeps
}
