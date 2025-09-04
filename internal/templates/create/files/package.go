package files

import (
	"archi-ts-cli/internal/models"
	"encoding/json"
)

// GeneratePackageJson generate the package.json file content based on the project name and express usage
func GeneratePackageJson(name string, express bool) (string, error) {
	pkg := models.PackageJSON{
		Name:        name,
		Version:     "1.0.0",
		Description: "Backend API generated with Archi CLI",
		Main:        "dist/index.js",
		Scripts:     getScripts(),
		Keywords:    getKeywords(express),
		Author:      "",
		License:     "MIT",
	}

	// Add dependencies if Express is needed
	if express {
		pkg.Dependencies = getDependencies(express)
	}

	// Always add devDependencies
	pkg.DevDependencies = getDevDependencies(express)

	jsonData, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

// getScripts returns npm scripts for TypeScript
func getScripts() map[string]string {
	return map[string]string{
		"start":         "node dist/src/index.js",
		"dev":           "nodemon src/index.ts",
		"build":         "tsc",
		"watch":         "tsc --watch",
		"lint":          "eslint src/**/*.ts",
		"lint:fix":      "eslint src/**/*.ts --fix",
		"test":          "jest",
		"test:watch":    "jest --watch",
		"test:coverage": "jest --coverage",
	}
}

// getKeywords returns keywords based on Express usage
func getKeywords(express bool) []string {
	baseKeywords := []string{"api", "backend", "nodejs", "typescript"}
	if express {
		return append(baseKeywords, "express")
	}
	return baseKeywords
}

// getDependencies returns runtime dependencies
func getDependencies(express bool) map[string]string {
	deps := make(map[string]string)

	if express {
		deps["express"] = "^4.18.2"
		deps["class-validator"] = "^0.14.1"
		deps["class-transformer"] = "^0.5.1"
	}

	return deps
}

// getDevDependencies returns development dependencies for TypeScript
func getDevDependencies(express bool) map[string]string {
	devDeps := map[string]string{
		"nodemon":                          "^3.0.2",
		"@types/validator":                 "^13.12.2",
		"typescript":                       "^5.3.3",
		"@types/node":                      "^20.10.5",
		"eslint":                           "^8.56.0",
		"@typescript-eslint/eslint-plugin": "^6.15.0",
		"@typescript-eslint/parser":        "^6.15.0",
		"@types/jest":                      "^29.5.14",
		"jest":                             "^29.7.0",
		"ts-jest":                          "^29.1.1",
		"@types/supertest":                 "^2.0.12",
		"supertest":                        "^6.3.3",
	}

	if express {
		devDeps["@types/express"] = "^4.17.21"
	}

	return devDeps
}
