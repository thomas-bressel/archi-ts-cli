package files

import (
	"archi-ts-cli/internal/models"
	"fmt"
)

// GetTypeORMDataSourceTemplate return template for data-source.ts file
func GetTypeORMDataSourceTemplate(architecture string) string {
	migrationPath := ""

	switch architecture {
	case string(models.LayeredArchitecture):
		migrationPath = "src/data"
	case string(models.CleanArchitecture):
		migrationPath = "src/infrastructure"
	case string(models.HexagonalArchitecture):
		migrationPath = "storage"
	}

	return fmt.Sprintf(`import "reflect-metadata";
import { DataSource, DataSourceOptions } from "typeorm";

// node.js imports
import path from "path";
import dotenv from "dotenv";
dotenv.config();

// ORM Entities imports
// exemple for layered architecture : import Product from "@datamodels/product.model";
// exemple for hexagonal architecture : import ProductModel from "@models/product.model";

const isSQLite = process.env.DB_TYPE === "sqlite";

// Use process.cwd() to get the project root directory
const STORAGE_DATABASE_PATH = path.join(process.cwd(), "storage/database");

// Function to get the SQLite database path
const getSQLitePath = (): string => {
  if (process.env.SQLITE_FILE) {

    // If SQLITE_FILE starts with './', it's relative to project root
    if (process.env.SQLITE_FILE.startsWith("./")) {
      return path.join(process.cwd(), process.env.SQLITE_FILE.substring(2));
    }

    // If it's an absolute path, use it as is
    if (path.isAbsolute(process.env.SQLITE_FILE)) {
      return process.env.SQLITE_FILE;
    }

    // Otherwise, treat it as relative to project root
    return path.join(process.cwd(), process.env.SQLITE_FILE);
  }

  // Default path
  return path.join(STORAGE_DATABASE_PATH, "database.sqlite");
};

const databaseConfig: DataSourceOptions = isSQLite
  ? {
      type: "sqlite",
      database: getSQLitePath(),
      // Add entities ORM here, exemple :
      // entities for clean architecture: [Product],
      // ORM models for hexagonal architecture: [ProductModel],
      migrations: [
          path.join(process.cwd(), "%s/database/migrations/*.ts")
      ],
      subscribers: [
          path.join(process.cwd(), "%s/subscribers/*.ts")
      ],
      synchronize: process.env.TYPEORM_SYNCHRONIZE === "false",
      dropSchema: process.env.TYPEORM_DROP_SCHEMA === "true",
      logging: process.env.TYPEORM_LOGGING === "true" || false,
    }
  : {
      type: "mysql",
      host: process.env.DB_HOST || "localhost",
      port: parseInt(process.env.DB_PORT || "3306", 10),
      username: process.env.DB_USERNAME || "",
      password: process.env.DB_PASSWORD || "",
      database: process.env.DB_NAME || "",
      // Add entities ORM here, exemple :
      // entities: [Product],
      migrations: [
        path.join(process.cwd(), "%s/database/migrations/*.ts")
      ],
      subscribers: [
        path.join(process.cwd(), "%s/subscribers/*.ts")
      ],
      synchronize: process.env.TYPEORM_SYNCHRONIZE === "false",
      dropSchema: process.env.TYPEORM_DROP_SCHEMA === "true",
      logging: process.env.TYPEORM_LOGGING === "true" || false,
    };

export const AppDataSource = new DataSource(databaseConfig);`, migrationPath, migrationPath, migrationPath, migrationPath)
}

// GetTypeORMCreateDatabaseTemplate return template for create-database.ts file
func GetTypeORMCreateDatabaseTemplate() string {
	return `// mysql imports
import mysql from "mysql2/promise";

// node.js imports
import fs from "fs";
import path from "path";
    

/**
* Create the database if it does not exist.
* @returns Promise<void>
*/
export async function createDatabaseIfNotExists() {

  // Create sqlite file if it doesn't exist
  if (process.env.DB_TYPE === "sqlite") {

    // Use process.cwd() to get the project root directory
    const sqlitePath = path.join(process.cwd(), "storage/database");
    const sqliteFile = path.join(sqlitePath, "database.sqlite");
    
    // Check if the directory exists, if not create it
    if (!fs.existsSync(sqlitePath)) {
      fs.mkdirSync(sqlitePath, { recursive: true });
      console.log("Created directory: " + sqlitePath);
    }
        
    console.log("SQLite database will be created at: " + sqliteFile);
    return;
  }
    
  if (process.env.DB_TYPE !== "mysql") {
    return;
  }
    
  try {
    const connection = await mysql.createConnection({
      host: process.env.DB_HOST,
      user: process.env.DB_USERNAME,
      password: process.env.DB_PASSWORD,
    });
    
      // Create database if it doesn't exist
    await connection.query(
      ` + "`CREATE DATABASE IF NOT EXISTS \\`${process.env.DB_NAME}\\`;`" + `
    );
      console.log("Database \"" + process.env.DB_NAME + "\" checked/created!");
      await connection.end();
    } catch (error) {
      console.error("Error creating MySQL database:", error);
      throw error;
    }
}`
}

// GetHeloperORMScriptTemplate return template for generate-migration.ts file
func GetHelperORMScriptTemplate(architecture string) string {
	migrationPath := ""
	datasourcePath := ""

	switch architecture {
	case string(models.LayeredArchitecture):
		migrationPath = "src/data/database/migrations"
		datasourcePath = "src/data/database/connection"
	case string(models.CleanArchitecture):
		migrationPath = "src/infrastructure/database/migrations"
		datasourcePath = "src/infrastructure/database/config"
	case string(models.HexagonalArchitecture):
		migrationPath = "storage/database/migrations"
		datasourcePath = "src/adapters/secondary/persistence/orm"
	}

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
