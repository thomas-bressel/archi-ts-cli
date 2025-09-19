/**
 * The main server file
 * @module index
**/

// External dependencies
import 'reflect-metadata';
import express, { Request, Response } from "express";
import dotenv from "dotenv";

// Internal dependencies
import { AppDataSource } from "@orm/data-source";
import { createDatabaseIfNotExists } from "@orm/create-database";

// import routes here
// exemple : import { userRouter } from "@routes/user.routes";

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

    // Route paths here
    // exemple : server.use('/api', userRouter);

    // Start server
    const PORT = process.env.LISTEN_PORT ?? 3000;
    server.listen(PORT, () => {
      console.log(`🤘 Server running on http://localhost:${PORT}`);
    });
  } catch (err) {
    console.error("❌ Failed to start server:", err);
    process.exit(1);
  }
})();

export default server;