package files

import "archi-ts-cli/internal/models"

// GetTypeORMDataSourceTemplate return template for data-source.ts file
func GetTypeORMDataSourceTemplate() string {
	return `import "reflect-metadata";
import { DataSource, DataSourceOptions } from "typeorm";

import dotenv from "dotenv";

dotenv.config();
const isSQLite = process.env.DB_TYPE === "sqlite";

const databaseConfig: DataSourceOptions = isSQLite
  ? {
      type: "sqlite",
      database: process.env.SQLITE_FILE || "database.sqlite",
      // entities: [],
      synchronize: process.env.TYPEORM_SYNCHRONIZE === "true",
      dropSchema: process.env.TYPEORM_DROP_SCHEMA === "true",
      logging: false,
    }
  : {
      type: "mysql",
      host: process.env.DB_HOST || "localhost",
      port: parseInt(process.env.DB_PORT || "3306", 10),
      username: process.env.DB_USERNAME || "",
      password: process.env.DB_PASSWORD || "",
      database: process.env.DB_NAME || "",
      // entities: [],
      synchronize: process.env.TYPEORM_SYNCHRONIZE === "true",
      dropSchema: process.env.TYPEORM_DROP_SCHEMA === "true",
      logging: false,
    };

export const AppDataSource = new DataSource(databaseConfig);`
}

// GetTypeORMCreateDatabaseTemplate return template for create-database.ts file
func GetTypeORMCreateDatabaseTemplate() string {
	return `import mysql from "mysql2/promise";

export async function createDatabaseIfNotExists() {
  if (process.env.DB_TYPE !== "mysql") {
    // On ne touche pas à SQLite
    return;
  }

  const connection = await mysql.createConnection({
    host: process.env.DB_HOST,
    user: process.env.DB_USERNAME,
    password: process.env.DB_PASSWORD,
  });

  await connection.query(
    ` + "`CREATE DATABASE IF NOT EXISTS \\`${process.env.DB_NAME}\\`;`" + `
  );
  console.log(` + "`Database \"${process.env.DB_NAME}\" checked/created!`" + `);
  await connection.end();
}`
}

// GetHexagonalIndexTemplate generate  template index.ts for hexagonal architecture
func GetHexagonalIndexTemplate(express bool, orm models.Orm) string {
	if orm == models.TypeOrm {
		if express {
			return getHexagonalIndexWithExpressAndTypeORM()
		}
		return getHexagonalIndexWithTypeORMOnly()
	}

	// Sans ORM
	if express {
		return getHexagonalIndexWithExpressOnly()
	}
	return getHexagonalIndexPlain()
}

func getHexagonalIndexWithExpressAndTypeORM() string {
	return `/**
 * The main server file
 * @module index
**/

// External dependencies
import express, { Request, Response } from "express";
import dotenv from "dotenv";

// Internal dependencies
import { AppDataSource } from "@adapters-typeorm/data-source";
import { createDatabaseIfNotExists } from "@adapters-typeorm/create-database";


/* @Archi::imports-start -- DO NOT REMOVE */
/* @Archi::imports-end */


// Load environment variables from .env file
dotenv.config();

// Initialize Express app
const server = express();
server.use(express.json());

// Default route (health check)
server.get("/", (req: Request, res: Response) => {
  res.json({
    message: "Welcome to ArchiTS API",
    version: "1.0.0",
    status: "running",
    stack: "NodeJS, TypeScript",
    library: "ExpressJS + TypeORM",
  });
});

// Start server with DB
(async () => {
  try {
    
    // Initialize database connection
    await createDatabaseIfNotExists();
    
    // Initialize TypeORM data source
    await AppDataSource.initialize();
    console.log("✅ Data source initialized");

/* @Archi::routes-path-start -- DO NOT REMOVE */
/* @Archi::routes-path-end */

    // Start server
    const PORT = process.env.PORT ?? 3000;
    server.listen(PORT, () => {
      console.log(` + "`🚀 Server running on http://localhost:${PORT}`" + `);
    });
  } catch (err) {
    console.error("❌ Failed to start server:", err);
    process.exit(1);
  }
})();

export default server;`
}

func getHexagonalIndexWithTypeORMOnly() string {
	return `/**
 * The main server file
 * @module index
**/

import * as http from 'http';
import dotenv from "dotenv";

// Internal dependencies
import { AppDataSource } from "./adapters/typeorm/data-source";
import { createDatabaseIfNotExists } from "./adapters/typeorm/create-database";

/* @Archi::imports-start -- DO NOT REMOVE */
/* @Archi::imports-end */

// Load environment variables from .env file
dotenv.config();

const PORT = process.env.PORT ?? 3000;

// Start server with DB
(async () => {
  try {
    // Initialize database connection
    await createDatabaseIfNotExists();
    
    // Initialize TypeORM data source
    await AppDataSource.initialize();
    console.log("✅ Data source initialized");

    // Routes
    const server = http.createServer((req: http.IncomingMessage, res: http.ServerResponse) => {
      res.setHeader('Content-Type', 'application/json');
      res.end(JSON.stringify({
        message: 'Welcome to ArchiTS API',
        version: '1.0.0',
        status: 'running',
        stack: 'NodeJS, TypeScript',
        library: "TypeORM (no Express)"
      }));
    });

    // Start server
    server.listen(PORT, () => {
      console.log(` + "`🚀 Server running on http://localhost:${PORT}`" + `);
    });
  } catch (err) {
    console.error("❌ Failed to start server:", err);
    process.exit(1);
  }
})();`
}

func getHexagonalIndexWithExpressOnly() string {
	return `/**
 * The main server file for Hexagonal Architecture
 * @module index
**/

import express, { Request, Response } from 'express';
import dotenv from "dotenv";

// Load environment variables
dotenv.config();

const server = express();
server.use(express.json());

/* @Archi::imports-start -- DO NOT REMOVE */
/* @Archi::imports-end */

// Default route (health check)
server.get('/', (req: Request, res: Response) => {
  res.json({
    message: 'Welcome to ArchiTS API',
    version: '1.0.0',
    status: 'running',
    architecture: 'Hexagonal',
    stack: 'NodeJS, TypeScript',
    library: 'ExpressJS'
  });
});

/* @Archi::routes-path-start -- DO NOT REMOVE */
/* @Archi::routes-path-end */

// Start server
const PORT = process.env.PORT ?? 3000;
server.listen(PORT, () => {
  console.log(` + "`🚀 Server running on http://localhost:${PORT}`" + `);
});

export default server;`
}

func getHexagonalIndexPlain() string {
	return `/**
 * The main server file for Hexagonal Architecture
 * @module index
**/

import * as http from 'http';
import dotenv from "dotenv";

// Load environment variables
dotenv.config();

const PORT = process.env.PORT ?? 3000;

/* @Archi::imports-start -- DO NOT REMOVE */
/* @Archi::imports-end */

// Routes
const server = http.createServer((req: http.IncomingMessage, res: http.ServerResponse) => {
	res.setHeader('Content-Type', 'application/json');
	res.end(JSON.stringify({
	  message: 'Welcome to ArchiTS API',
	  version: '1.0.0',
	  status: 'running',
	  architecture: 'Hexagonal',
	  stack: 'NodeJS, TypeScript',
	  library: "no library, pure Node.js"
	}));
});

// Start server
server.listen(PORT, () => {
  console.log(` + "`🚀 Server running on http://localhost:${PORT}`" + `);
});`
}
