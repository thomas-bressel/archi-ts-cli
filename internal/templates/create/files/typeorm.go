package files

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
    return;
  }

  const connection = await mysql.createConnection({
    host: process.env.DB_HOST,
    user: process.env.DB_USERNAME,
    password: process.env.DB_PASSWORD,
  });

  // Create database if it doesn't exist
  await connection.query(
    ` + "`CREATE DATABASE IF NOT EXISTS \\`${process.env.DB_NAME}\\`;`" + `
  );
  console.log(` + "`Database \"${process.env.DB_NAME}\" checked/created!`" + `);
  await connection.end();
}`
}
