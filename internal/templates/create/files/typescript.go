package files

import (
	"archi-ts-cli/internal/models"
	"archi-ts-cli/internal/templates/create/files/clean"
	"archi-ts-cli/internal/templates/create/files/hexagonal"
	"archi-ts-cli/internal/templates/create/files/layered"
)

// GetTsconfigTemplate generates template for tsconfig.json based on architecture
func GetTsconfigTemplate(architecture models.Architecture) string {

	// define variable paths depending of architecture
	paths := getTsconfigPaths(architecture)

	// define types for hexagonal architecture
	types := ""
	if architecture == models.HexagonalArchitecture {
		types = `"types": ["reflect-metadata", "jest"],`
	}

	return getTsconfigBaseTemplate(types, paths)
}

// getTsconfigPaths returns the paths configuration based on architecture and ORM
func getTsconfigPaths(architecture models.Architecture) string {
	switch architecture {
	case models.CleanArchitecture:
		return clean.GetTsconfigPaths()
	case models.HexagonalArchitecture:
		return hexagonal.GetTsconfigPaths()
	case models.LayeredArchitecture:
		return layered.GetTsconfigPaths()
	default:
		return ``
	}
}

// getTsconfigBaseTemplate returns default base template
func getTsconfigBaseTemplate(types string, paths string) string {
	return `{
  "compilerOptions": {
    "baseUrl": "./",
    "target": "es2017",
    "module": "commonjs",
    "moduleResolution": "node",
    "strict": true,
    "esModuleInterop": true,
    "skipLibCheck": true,
    "outDir": "./dist",
    "rootDir": "./",
    "forceConsistentCasingInFileNames": true,
    "experimentalDecorators": true,
    "emitDecoratorMetadata": true,
    ` + types + `
    "paths": ` + paths + `
  },
    "ts-node": {
    "require": ["tsconfig-paths/register"],
    "transpileOnly": true,
    "files": true
  },
  "include": [
    "src/**/*",
    "tests/**/*"
  ],
  "exclude": [
    "node_modules",
    "dist"
  ]
}`
}
