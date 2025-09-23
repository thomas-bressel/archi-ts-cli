package files

import (
	"archi-ts-cli/internal/models"
	"archi-ts-cli/internal/templates/create/files/clean"
	"archi-ts-cli/internal/templates/create/files/hexagonal"
	"archi-ts-cli/internal/templates/create/files/layered"
)

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
func getJestConfigPaths(architecture models.Architecture) string {
	switch architecture {
	case models.LayeredArchitecture:
		return layered.GetJestConfigPaths()
	case models.CleanArchitecture:
		return clean.GetJestConfigPaths()
	case models.HexagonalArchitecture:
		return hexagonal.GetJestConfigPaths()
	default:
		return ""
	}
}
