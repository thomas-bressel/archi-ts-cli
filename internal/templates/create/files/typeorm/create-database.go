package typeorm

func GetCreateDatabaseTemplate() string {
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
