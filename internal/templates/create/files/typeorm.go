package files

// GetTypeORMDataSourceTemplate return template for data-source.ts file
func GetTypeORMDataSourceTemplate() string {
	return `
    import "reflect-metadata";
    import { DataSource, DataSourceOptions } from "typeorm";
    import path from "path";
    import dotenv from "dotenv";

    dotenv.config();

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
          // entities: [
          // ],
          synchronize: process.env.TYPEORM_SYNCHRONIZE === "true",
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
          // entities: [
          // ],
          synchronize: process.env.TYPEORM_SYNCHRONIZE === "true",
          dropSchema: process.env.TYPEORM_DROP_SCHEMA === "true",
          logging: process.env.TYPEORM_LOGGING === "true" || false,
        };

    export const AppDataSource = new DataSource(databaseConfig);`
}

// GetTypeORMCreateDatabaseTemplate return template for create-database.ts file
func GetTypeORMCreateDatabaseTemplate() string {
	return `
    import mysql from "mysql2/promise";
    import fs from "fs";
    import path from "path";
    
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
