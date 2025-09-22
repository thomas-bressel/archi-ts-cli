package files

import "archi-ts-cli/internal/models"

// GetJestConfigTemplate returns a template for jest based on architecture and ORM
func GetJestConfigTemplate(architecture models.Architecture, orm models.Orm) string {
	paths := getJestConfigPaths(architecture)

	return `import type { Config } from '@jest/types';

const config: Config.InitialOptions = {
  preset: 'ts-jest',
  testEnvironment: 'node',
  moduleNameMapper: ` + paths + `,
  testMatch: ['**/tests/**/*.test.ts'],
  collectCoverage: true,
  coverageDirectory: 'tests/coverage',
  coverageReporters: ['text', 'lcov'],
  transform: {
    '^.+\\.ts$': ['ts-jest', {
      tsconfig: 'tsconfig.json'
    }],
  },
  moduleFileExtensions: ['ts', 'js', 'json'],
};

export default config;`
}

// getJestConfigPaths returns the paths configuration based on architecture and ORM
// [X] Layered Architecture
// [X] Clean Architecture
// [X] Hexagonal Architecture
func getJestConfigPaths(architecture models.Architecture) string {
	switch architecture {
	case models.LayeredArchitecture:
		return `{
    "^@src/(.*)$": "<rootDir>/src/$1",
    "^@config/(.*)$": "<rootDir>/src/common/config/$1",
    "^@constants/(.*)$": "<rootDir>/src/common/constants/$1",
    "^@errors/(.*)$": "<rootDir>/src/common/errors/$1",
    "^@logging/(.*)$": "<rootDir>/src/common/logging/$1",
    "^@utils/(.*)$": "<rootDir>/src/common/utils/$1",
    "^@controllers/(.*)$": "<rootDir>/src/presentation/controllers/$1",
    "^@middlewares/(.*)$": "<rootDir>/src/presentation/middlewares/$1",
    "^@routes/(.*)$": "<rootDir>/src/presentation/routes/$1",
    "^@services/(.*)$": "<rootDir>/src/business/services/$1",
    "^@interfaces/(.*)$": "<rootDir>/src/business/interfaces/$1",
    "^@businessmodels/(.*)$": "<rootDir>/src/business/models/$1",
    "^@repositories/(.*)$": "<rootDir>/src/data/repositories/$1",
    "^@datamodels/(.*)$": "<rootDir>/src/data/models/$1",
    "^@connection/(.*)$": "<rootDir>/src/data/database/connection/$1",
    "^@migration/(.*)$": "<rootDir>/src/data/database/migration/$1",
    "^@seeds/(.*)$": "<rootDir>/src/data/database/seeds/$1",
    "^@storage/(.*)$": "<rootDir>/storage/$1",
	}`

	case models.CleanArchitecture:
		return `{
      "^@src/(.*)$": "<rootDir>/src/$1",
      "^@entities/(.*)$": "<rootDir>/src/domain/entities/$1",
      "^@errors/(.*)$": "<rootDir>/src/domain/errors/$1",
      "^@usecases/(.*)$": "<rootDir>/src/application/use-cases/$1",
      "^@services/(.*)$": "<rootDir>/src/application/use-cases/$1",
      "^@interfaces/(.*)$": "<rootDir>/src/application/interfaces/$1",
      "^@dtos/(.*)$": "<rootDir>/src/application/dtos/$1",
      "^@controllers/(.*)$": "<rootDir>/src/presentation/controllers/$1",
      "^@routes/(.*)$": "<rootDir>/src/presentation/routes/$1",
      "^@middlewares/(.*)$": "<rootDir>/src/presentation/middlewares/$1",
      "^@validators/(.*)$": "<rootDir>/src/presentation/validators/$1",
      "^@repositories/(.*)$": "<rootDir>/src/infrastructure/repositories/$1",
      "^@datamodels/(.*)$": "<rootDir>/src/infrastructure/models/$1",
      "^@cache/(.*)$": "<rootDir>/src/infrastructure/cache/$1",
      "^@email/(.*)$": "<rootDir>/src/infrastructure/email/$1",
      "^@mappers/(.*)$": "<rootDir>/src/infrastructure/mappers/$1",
      "^@config/(.*)$": "<rootDir>/src/infrastructure/database/config/$1",
      "^@utils/(.*)$": "<rootDir>/src/shared/utils/$1",
      "^@storage/(.*)$": "<rootDir>/storage/$1",
    }`

	case models.HexagonalArchitecture:
		return `{
      "^@src/(.*)$": "<rootDir>/src/$1",
      "^@entities/(.*)$": "<rootDir>/src/domain/entities/$1",
      "^@objects/(.*)$": "<rootDir>/src/domain/value-objects/$1",
      "^@events/(.*)$": "<rootDir>/src/domain/events/$1",
      "^@exceptions/(.*)$": "<rootDir>/src/domain/exceptions/$1",
      "^@usecases/(.*)$": "<rootDir>/src/application/use-cases/$1",
      "^@services/(.*)$": "<rootDir>/src/application/use-cases/$1",
      "^@ports/(.*)$": "<rootDir>/src/application/ports/$1",
      "^@dtos/(.*)$": "<rootDir>/src/application/dtos/$1",
      "^@http/(.*)$": "<rootDir>/src/adapters/primary/http/$1",
      "^@controllers/(.*)$": "<rootDir>/src/adapters/primary/http/controllers/$1",
      "^@routes/(.*)$": "<rootDir>/src/adapters/primary/http/routes/*$1",
      "^@middlewares/(.*)$": "<rootDir>/src/adapters/primary/http/middlewares/$1",
      "^@cli/(.*)$": "<rootDir>/src/adapters/primary/cli/$1",
      "^@orm/(.*)$": "<rootDir>/src/adapters/secondary/persistence/orm/$1",
      "^@models/(.*)$": "<rootDir>/src/adapters/secondary/persistence/models/$1",
      "^@repositories/(.*)$": "<rootDir>/src/adapters/secondary/persistence/repositories/$1",
      "^@email/(.*)$": "<rootDir>/src/adapters/secondary/email/$1",
      "^@cache/(.*)$": "<rootDir>/src/adapters/secondary/cache/$1",
      "^@config/(.*)$": "<rootDir>/src/config/$1", 
      "^@storage/(.*)$": "<rootDir>/storage/$1",
    }`
	default:
		return ""
	}
}
