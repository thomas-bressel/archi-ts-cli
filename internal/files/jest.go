package files

// GetjestConfigTemplate return a template for jest
func GetJestConfigTemplate(extention string) string {
	if extention == "js" {

		return `
		module.exports = {
			preset: 'js-jest',
  			testEnvironment: 'node',
  			testMatch: ['**/tests/**/*.test.js'],
  			collectCoverage: true,
  			coverageDirectory: 'tests/coverage',
  			coverageReporters: ['text', 'lcov'],
		  };`
	}
	return `
		import type { Config } from '@jest/types';
		
		const config: Config.InitialOptions = {
	  	preset: 'ts-jest',
	  	testEnvironment: 'node',
	  	testMatch: ['**/tests/**/*.test.ts', '**/tests/**/*.test.js'],
	  	collectCoverage: true,
	  	coverageDirectory: 'tests/coverage',
	  	coverageReporters: ['text', 'lcov'],
	  	transform: {
	  	  '^.+\\.ts$': 'ts-jest',
	  	},
	  	moduleFileExtensions: ['ts', 'js', 'json'],
	  	};
	  
	  export default config;`

}
