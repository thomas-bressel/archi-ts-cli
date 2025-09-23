package typeorm

import (
	"fmt"
)

// GetHeloperORMScriptTemplate return template for generate-migration.ts file
func GetHelperScriptTemplate(migrationPath string, datasourcePath string) string {
	return fmt.Sprintf(`#!/usr/bin/env node

const { execSync } = require('child_process');
const path = require('path');

// Get the migration name from command line arguments
const migrationName = process.argv[2];

if (!migrationName) {
  console.error('❌ Erreur: It is necessary to provide a name for the migration.');
  console.log('Usage: npm run migration:generate -- NameOfTheMigration');
  process.exit(1);
}

// Build the full path for the migration file
const migrationPath = path.join('%s', migrationName);

// Build the TypeORM CLI command
const command = 'npx typeorm-ts-node-commonjs migration:generate ' + migrationPath + ' -d ./%s/data-source.ts -p';

console.log('📝 Generation of the migration: ', migrationName);
console.log('📂 Into: ', migrationPath);

try {
  // Launch the command
  execSync(command, { stdio: 'inherit' });
  console.log('✅ Migration successfull generated!');
} catch (error) {
  console.error('❌ Error during migration generation:', error);
  process.exit(1);
}`, migrationPath, datasourcePath)

}
