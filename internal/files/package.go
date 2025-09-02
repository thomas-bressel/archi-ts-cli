package files

import (
	"encoding/json"
)

// PackageJSON represents the structure of package.json with ordered fields
type PackageJSON struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Description     string            `json:"description"`
	Main            string            `json:"main"`
	Scripts         map[string]string `json:"scripts"`
	Keywords        []string          `json:"keywords"`
	Author          string            `json:"author"`
	License         string            `json:"license"`
	Dependencies    map[string]string `json:"dependencies,omitempty"`
	DevDependencies map[string]string `json:"devDependencies"`
}

// GeneratePackageJson generate the package.json file content based on the project name and express usage
func GeneratePackageJson(name string, useExpress bool) (string, error) {
	pkg := PackageJSON{
		Name:        name,
		Version:     "1.0.0",
		Description: "Backend API generated with Archi CLI",
		Main:        "dist/index.js",
		Scripts:     getScripts(),
		Keywords:    getKeywords(useExpress),
		Author:      "",
		License:     "MIT",
	}

	// Add dependencies if Express is needed
	if useExpress {
		pkg.Dependencies = getDependencies(useExpress)
	}

	// Always add devDependencies
	pkg.DevDependencies = getDevDependencies(useExpress)

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
func getKeywords(useExpress bool) []string {
	baseKeywords := []string{"api", "backend", "nodejs", "typescript"}
	if useExpress {
		return append(baseKeywords, "express")
	}
	return baseKeywords
}

// getDependencies returns runtime dependencies
func getDependencies(useExpress bool) map[string]string {
	deps := make(map[string]string)

	if useExpress {
		deps["express"] = "^4.18.2"
		deps["class-validator"] = "^0.14.1"
		deps["class-transformer"] = "^0.5.1"
	}

	return deps
}

// getDevDependencies returns development dependencies for TypeScript
func getDevDependencies(useExpress bool) map[string]string {
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

	if useExpress {
		devDeps["@types/express"] = "^4.17.21"
	}

	return devDeps
}
