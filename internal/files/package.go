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

// generatePackageJson generate the package.json file content based on the project name, language and express usage
func GeneratePackageJson(name, language string, useExpress bool) (string, error) {
	pkg := PackageJSON{
		Name:        name,
		Version:     "1.0.0",
		Description: "Backend API generated with Archi CLI",
		Main:        getMainFile(language),
		Scripts:     getScripts(language),
		Keywords:    getKeywords(useExpress),
		Author:      "",
		License:     "ISC",
	}

	// Add dependencies if Express is needed
	if useExpress {
		pkg.Dependencies = getDependencies(language, useExpress)
	}

	// Always add devDependencies
	pkg.DevDependencies = getDevDependencies(language, useExpress)

	jsonData, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

// getMainFile return the main file based on the language
func getMainFile(language string) string {
	if language == "typescript" {
		return "dist/index.js"
	}
	return "src/index.js"
}

// getScripts returns npm scripts based on the language
func getScripts(language string) map[string]string {
	if language == "typescript" {
		return map[string]string{
			"start":    "node dist/src/index.js",
			"dev":      "nodemon src/index.ts",
			"build":    "tsc",
			"watch":    "tsc --watch",
			"lint":     "eslint src/**/*.ts",
			"lint:fix": "eslint src/**/*.ts --fix",
		}
	}

	return map[string]string{
		"start":    "node src/index.js",
		"dev":      "nodemon src/index.js",
		"lint":     "eslint src/**/*.js",
		"lint:fix": "eslint src/**/*.js --fix",
	}
}

// getKeywords returns keywords based on Express usage
func getKeywords(useExpress bool) []string {
	baseKeywords := []string{"api", "backend", "nodejs"}
	if useExpress {
		return append(baseKeywords, "express")
	}
	return baseKeywords
}

// getDependencies returns runtime dependencies
func getDependencies(language string, useExpress bool) map[string]string {
	deps := make(map[string]string)

	if useExpress {
		deps["express"] = "^4.18.2"
	}

	return deps
}

// getDevDependencies returns development dependencies
func getDevDependencies(language string, useExpress bool) map[string]string {
	devDeps := map[string]string{
		"nodemon": "^3.0.2",
	}

	if language == "typescript" {
		devDeps["typescript"] = "^5.3.3"
		devDeps["@types/node"] = "^20.10.5"
		devDeps["eslint"] = "^8.56.0"
		devDeps["@typescript-eslint/eslint-plugin"] = "^6.15.0"
		devDeps["@typescript-eslint/parser"] = "^6.15.0"
		devDeps["@types/jest"] = "^29.5.14"
		devDeps["jest"] = "^29.7.0"

		if useExpress {
			devDeps["@types/express"] = "^4.17.21"
		}
	} else {
		devDeps["eslint"] = "^8.56.0"
	}

	return devDeps
}
