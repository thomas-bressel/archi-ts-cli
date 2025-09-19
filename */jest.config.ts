import type { Config } from '@jest/types';

const config: Config.InitialOptions = {
  preset: 'ts-jest',
  testEnvironment: 'node',
  moduleNameMapper: {
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
      "^@config/(.*)$": "<rootDir>/src/config/*$1", 
      "^@storage/(.*)$": "<rootDir>/storage/$1",
    },
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

export default config;