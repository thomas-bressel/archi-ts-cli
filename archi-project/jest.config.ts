
		import type { Config } from '@jest/types';
		
		const config: Config.InitialOptions = {
	  	preset: 'ts-jest',
	  	testEnvironment: 'node',
	  	testMatch: ['**/__tests__/**/*.test.ts', '**/__tests__/**/*.test.js'],
	  	collectCoverage: true,
	  	coverageDirectory: '__tests__/coverage',
	  	coverageReporters: ['text', 'lcov'],
	  	transform: {
	  	  '^.+\\.ts$': 'ts-jest',
	  	},
	  	moduleFileExtensions: ['ts', 'js', 'json'],
	  	};
	  
	  export default config;