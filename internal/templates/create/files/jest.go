package files

import "archi-ts-cli/internal/models"

// GetJestConfigTemplate returns a template for jest based on architecture
func GetJestConfigTemplate(architecture models.Architecture) string {
	moduleNameMapper := getJestModuleNameMapper(architecture)

	return `import type { Config } from '@jest/types';

const config: Config.InitialOptions = {
  preset: 'ts-jest',
  testEnvironment: 'node',
  moduleNameMapper: ` + moduleNameMapper + `,
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

// getJestModuleNameMapper returns the module name mapper based on architecture
func getJestModuleNameMapper(architecture models.Architecture) string {
	switch architecture {
	case models.CleanArchitecture:
		return `{
    "^@src/(.*)$": "<rootDir>/src/$1",
    "^@entities/(.*)$": "<rootDir>/src/domain/entities/$1",
    "^@controllers/(.*)$": "<rootDir>/src/presentation/controllers/$1",
    "^@services/(.*)$": "<rootDir>/src/data/services/$1",
    "^@repositories/(.*)$": "<rootDir>/src/data/repositories/$1",
    "^@middlewares/(.*)$": "<rootDir>/src/presentation/middleware/$1",
    "^@utils/(.*)$": "<rootDir>/src/shared/utils/$1"
  }`

	case models.HexagonalArchitecture:
		return `{
    "^@core/(.*)$": "<rootDir>/src/core/$1",
    "^@ports/(.*)$": "<rootDir>/src/ports/$1",
    "^@adapters/(.*)$": "<rootDir>/src/adapters/$1",
    "^@config/(.*)$": "<rootDir>/src/config/$1",
    "^@main/(.*)$": "<rootDir>/src/main/$1",
    "^@src/(.*)$": "<rootDir>/src/$1"
  }`

	default: // LayeredArchitecture
		return `{
    "^@src/(.*)$": "<rootDir>/src/$1",
    "^@entities/(.*)$": "<rootDir>/src/entities/$1",
    "^@controllers/(.*)$": "<rootDir>/src/controllers/$1",
    "^@services/(.*)$": "<rootDir>/src/services/$1",
    "^@repositories/(.*)$": "<rootDir>/src/repositories/$1",
    "^@routes/(.*)$": "<rootDir>/src/routes/$1",
    "^@middlewares/(.*)$": "<rootDir>/src/middlewares/$1",
    "^@utils/(.*)$": "<rootDir>/src/utils/$1"
  }`
	}
}
