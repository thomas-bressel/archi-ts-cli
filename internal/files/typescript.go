package files

import "archi-ts-cli/internal/models"

// GetTsconfigTemplate generates template for tsconfig.json based on architecture
func GetTsconfigTemplate(architecture models.Architecture) string {
	paths := getTsconfigPaths(architecture)

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
    "paths": ` + paths + `
  },
  "include": [
    "src/**/*",
    "tests/**/*"
  ],
  "exclude": [
    "node_modules",
    "dist"
  ]
}
`
}

// getTsconfigPaths returns the paths configuration based on architecture
func getTsconfigPaths(architecture models.Architecture) string {
	switch architecture {
	case models.CleanArchitecture:
		return `{
      "@src/*": ["src/*"],
      "@entities/*": ["src/domain/entities/*"],
      "@controllers/*": ["src/presentation/controllers/*"],
      "@services/*": ["src/data/services/*"],
      "@repositories/*": ["src/data/repositories/*"],
      "@routes/*": ["src/presentation/routes/*"],
      "@middlewares/*": ["src/presentation/middlewares/*"],
      "@utils/*": ["src/shared/utils/*"]
    }`

	case models.HexagonalArchitecture:
		return `{
      "@core/*": ["src/core/*"],
      "@core/domain/*": ["src/core/domain/*"],
      "@core/application/*": ["src/core/application/*"],
      "@ports/*": ["src/ports/*"],
      "@adapters/*": ["src/adapters/*"],
      "@adapters/inbound/*": ["src/adapters/inbound/*"],
      "@adapters/outbound/*": ["src/adapters/outbound/*"],
      "@config/*": ["src/config/*"],
      "@main/*": ["src/main/*"],
      "@src/*": ["src/*"]
    }`

	default: // LayeredArchitecture
		return `{
      "@src/*": ["src/*"],
      "@entities/*": ["src/entities/*"],
      "@controllers/*": ["src/controllers/*"],
      "@services/*": ["src/services/*"],
      "@repositories/*": ["src/repositories/*"],
      "@routes/*": ["src/routes/*"],
      "@middlewares/*": ["src/middlewares/*"],
      "@utils/*": ["src/utils/*"],
      "@config/*": ["src/config/*"]
    }`
	}
}
