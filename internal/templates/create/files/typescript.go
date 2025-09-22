package files

import "archi-ts-cli/internal/models"

// GetTsconfigTemplate generates template for tsconfig.json based on architecture
// [X] Common template for all type of architecture
func GetTsconfigTemplate(architecture models.Architecture, orm models.Orm) string {
	paths := getTsconfigPaths(architecture)
	types := ""
	if architecture == models.HexagonalArchitecture {
		types = `"types": ["reflect-metadata", "jest"],`
	}

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
}
`
}

// getTsconfigPaths returns the paths configuration based on architecture and ORM
// [X] Layered Architecture
// [X] Clean Architecture
// [X] Hexagonal Architecture
func getTsconfigPaths(architecture models.Architecture) string {
	switch architecture {
	case models.CleanArchitecture:
		return `{
      "@src/*": ["src/*"],
      "@entities/*": ["src/domain/entities/*"],
      "@errors/*": ["src/domain/errors/*"],
      "@usecases/*": ["src/application/use-cases/*"],
      "@services/*": ["src/application/use-cases/*"],
      "@interfaces/*": ["src/application/interfaces/*"],
      "@dtos/*": ["src/application/dtos/*"],
      "@controllers/*": ["src/presentation/controllers/*"],
      "@routes/*": ["src/presentation/routes/*"],
      "@middlewares/*": ["src/presentation/middlewares/*"],
      "@validators/*": ["src/presentation/validators/*"],
      "@datamodels/*": ["src/infrastructure/models/*"],
      "@repositories/*": ["src/infrastructure/repositories/*"],
      "@cache/*": ["src/infrastructure/cache/*"],
      "@email/*": ["src/infrastructure/email/*"],
      "@mappers/*": ["src/infrastructure/mappers/*"],
      "@config/*": ["src/infrastructure/database/config/*"],
      "@utils/*": ["src/shared/utils/*"],
      "@storage/*": ["storage/*"],
    }`

	case models.HexagonalArchitecture:
		return `{
      "@src/*": ["src/*"],
      "@entities/*": ["src/domain/entities/*"], 
      "@objects/*": ["src/domain/value-objects/*"], 
      "@events/*": ["src/domain/events/*"],
      "@exceptions/*": ["src/domain/exceptions/*"],
      "@usecases/*": ["src/application/use-cases/*"], 
      "@services/*": ["src/application/use-cases/*"],
      "@ports/*": ["src/application/ports/*"], 
      "@dtos/*": ["src/application/dtos/*"],
      "@http/*": ["src/adapters/primary/http/*"],
      "@controllers/*": ["src/adapters/primary/http/controllers/*"], 
      "@routes/*": ["src/adapters/primary/http/routes/*"],  
      "@middlewares/*": ["src/adapters/primary/http/middlewares/*"],
      "@cli/*": ["src/adapters/primary/cli/*"],
      "@orm/*": ["src/adapters/secondary/persistence/orm/*"], 
      "@models/*": ["src/adapters/secondary/persistence/models/*"], 
      "@repositories/*": ["src/adapters/secondary/persistence/repositories/*"],
      "@email/*": ["src/adapters/secondary/email/*"],
      "@cache/*": ["src/adapters/secondary/cache/*"],
      "@config/*": ["src/config/*"],  
      "@storage/*": ["storage/*"],
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
      "@storage/*": ["storage/*"],
    }`

	default:
		return ``
	}
}
