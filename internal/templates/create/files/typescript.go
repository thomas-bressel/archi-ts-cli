package files

import "archi-ts-cli/internal/models"

// GetTsconfigTemplate generates template for tsconfig.json based on architecture
// [X] Common template for all type of architecture
func GetTsconfigTemplate(architecture models.Architecture, orm models.Orm) string {
	paths := getTsconfigPaths(architecture, orm)

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

// getTsconfigPaths returns the paths configuration based on architecture and ORM
// [X] Layered Architecture
// [X] Clean Architecture
// [X] Hexagonal Architecture
func getTsconfigPaths(architecture models.Architecture, orm models.Orm) string {
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
		// For hexagonale with TypeORM
		if orm == models.TypeOrm {
			return `{
      "@src/*": ["src/*"],
      "@entities/*": ["src/domain/entities/*"],
      "@controllers/*": ["src/interfaces/controllers/*"],
      "@services/*": ["src/application/services/*"],
      "@repositories/*": ["src/adapters/typeorm/repositories/*"],
      "@ports/*": ["src/domain/ports/*"],
      "@routes/*": ["src/interfaces/routes/*"],
      "@adapters-typeorm/*": ["src/adapters/typeorm/*"],
    }`
		}
		// for hexagonale whitout ORM
		return `{
      "@src/*": ["src/*"],
      "@entities/*": ["src/core/domain/entities/*"],
      "@controllers/*": ["src/adapters/inbound/http/controllers/*"],
      "@services/*": ["src/core/application/services/*"],
      "@repositories/*": ["src/adapters/outbound/repositories/*"],
      "@ports/*": ["src/ports/*"],
      "@routes/*": ["src/adapters/inbound/http/routes/*"]
    }`

	case models.LayeredArchitecture:
		return `{
      "@src/*": ["src/*"],
      "@config/*": ["src/common/config/*"],
      "@constants/*": ["src/common/constants/*"],
      "@errors/*": ["src/common/errors/*"],
      "@logging/*": ["src/common/logging/*"],
      "@utils/*": ["src/common/utils/*"],
      "@controllers/*": ["src/presentation/controllers/*"],
      "@middlewares/*": ["src/presentation/middlewares/*"],
      "@routes/*": ["src/presentation/routes/*"],
      "@services/*": ["src/business/services/*"],
      "@interfaces/*": ["src/business/interfaces/*"],
      "@businessmodels/*": ["src/business/models/*"],
      "@repositories/*": ["src/data/repositories/*"],
      "@datamodels/*": ["src/data/models/*"],
      "@connection/*": ["src/data/database/connection/*"],
      "@migration/*": ["src/data/database/migration/*"],
      "@seeds/*": ["src/data/database/seeds/*"],
    }`

	default:
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
