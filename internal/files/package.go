package files

import (
	"encoding/json"
)

// generatePackageJson generate the package.json file content based on the project name and language
func GeneratePackageJson(name, language string) (string, error) {
	packageData := map[string]interface{}{
		"name":        name,
		"version":     "1.0.0",
		"description": "Backend API generated with Archi CLI",
		"main":        getMainFile(language),
		"scripts":     getScripts(language),
		"keywords":    []string{"api", "backend", "nodejs", "express"},
		"author":      "",
		"license":     "ISC",
	}

	jsonData, err := json.MarshalIndent(packageData, "", "  ")
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
			"start":    "node dist/index.js",
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
