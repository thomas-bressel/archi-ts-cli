package files

// GetJestConfigTemplate return a template for jest (TypeScript only)
func GetJestConfigTemplate() string {
	return `
import type { Config } from '@jest/types';

const config: Config.InitialOptions = {
  preset: 'ts-jest',
  testEnvironment: 'node',
  testMatch: ['**/tests/**/*.test.ts'],
  collectCoverage: true,
  coverageDirectory: 'tests/coverage',
  coverageReporters: ['text', 'lcov'],
  transform: {
    '^.+\\.ts$': 'ts-jest',
  },
  moduleFileExtensions: ['ts', 'json'],
};

export default config;`
}
