import "reflect-metadata";
import { DataSource, DataSourceOptions } from "typeorm";

// node.js imports
import path from "path";
import dotenv from "dotenv";
dotenv.config();

// ORM Entities imports
// exemple for layered architecture : import Product from "@datamodels/product.model";

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
      // entities: [Product],
      migrations: [
          path.join(process.cwd(), "src/data/database/migrations/*.ts")
      ],
      subscribers: [
          path.join(process.cwd(), "src/data/subscribers/*.ts")
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
        path.join(process.cwd(), "src/data/database/migrations/*.ts")
      ],
      subscribers: [
        path.join(process.cwd(), "src/data/subscribers/*.ts")
      ],
      synchronize: process.env.TYPEORM_SYNCHRONIZE === "false",
      dropSchema: process.env.TYPEORM_DROP_SCHEMA === "true",
      logging: process.env.TYPEORM_LOGGING === "true" || false,
    };

export const AppDataSource = new DataSource(databaseConfig);