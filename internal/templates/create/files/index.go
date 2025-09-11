package files

import "archi-ts-cli/internal/models"

// GetIndexTemplate generate the template for index.ts
// If an ORM exists then return the Express + ORM template
// If any ORM exists then only return Express Template
func GetIndexTemplate(architecture models.Architecture, express bool, orm models.Orm) string {

	// If an Orm is chosen then Express too
	if orm == models.TypeOrm {
		return getDefaultIndexWithExpressAndTypeORM(architecture)
	}

	// Generate Express template with its appropriate architecture
	if express {
		return getDefaultIndexWithExpress()
	}

	// Generate native Node.js template with its appropriate architecture
	return getDefaultIndexPlain()

}

// Default Index file without any library or Orm
// [X] Layered, Clean or Hexagonale Architecture
func getDefaultIndexPlain() string {

	return `
    /**
     * The main server file
     * @module index
    **/
    import * as http from "http";
    import dotenv from "dotenv";

    // Layers importation HERE
    // exemple : import { handleUserRoutes } from "@routes/user.routes";

    // Load environment variables
    dotenv.config();
    const PORT: number = process.env.LISTEN_PORT ? parseInt(process.env.LISTEN_PORT) : 3000;

    /**
     * Default route handler
     */
    function defaultRoute(req: http.IncomingMessage, res: http.ServerResponse) {
        res.statusCode = 404;
        res.setHeader("Content-Type", "application/json");
        res.end(JSON.stringify({
            message: 'Route '+ req.url + ' not found'
        }));
    }

    /**
     * Static routes
     */
    const staticRoutes: Record<string, (req: http.IncomingMessage, res: http.ServerResponse) => void> = {
        "/": (req, res) => {
            res.statusCode = 200;
            res.setHeader("Content-Type", "application/json");
            res.end(JSON.stringify({
                message: 'Welcome to Archi API',
                version: '1.10.0',
                status: 'running',
                stack: 'NodeJS, TypeScript',
                library: "no library, I'm a badass developer"
            }));
        },
        "/ping": (req, res) => {
            res.statusCode = 200;
            res.end("pong");
        }
    };

    /**
     * Create HTTP server
     */
    const server = http.createServer(async (req, res) => {
        try {
            if (req.url && staticRoutes[req.url]) {
                staticRoutes[req.url](req, res);
                return;
            }

    		 // Handle routes HERE
            // exemple : if (await handleUserRoutes(req, res)) return;

            defaultRoute(req, res);

        } catch (error) {
            res.statusCode = 500;
            res.setHeader("Content-Type", "application/json");
            res.end(JSON.stringify({
                message: error instanceof Error ? error.message : "Internal server error"
            }));
        }
    });

    /**
     * Start server
     */
    server.listen(PORT, () => {
        console.log('Server running on http://localhost:' + PORT);
    });
    `
}

// Default Index file with Express but no Orm
// [X] Layered Architecture
// [O] Clean Architecture
// [X] Hexagonal Architecture
// [X] Express.js
// [X] TypeORM
func getDefaultIndexWithExpressAndTypeORM(architecture models.Architecture) string {
	var imports string = ""
	switch architecture {
	case models.LayeredArchitecture:
		imports = `
        import { AppDataSource } from "@connection/data-source";
        import { createDatabaseIfNotExists } from "@connection/create-database";`
	case models.CleanArchitecture:
		imports = `
        import { AppDataSource } from "@connection/data-source";
        import { createDatabaseIfNotExists } from "@connection/create-database";`
	case models.HexagonalArchitecture:
		imports = `
        import { AppDataSource } from "@adapters-typeorm/data-source";
        import { createDatabaseIfNotExists } from "@adapters-typeorm/create-database";`
	default:
		return ""
	}

	return `
    /**
     * The main server file
     * @module index
    **/

    // External dependencies
    import express, { Request, Response } from "express";
    import dotenv from "dotenv";

    // Internal dependencies
    ` + imports + `

    // import routes here

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

        // Start server
        const PORT = process.env.LISTEN_PORT ?? 3000;
        server.listen(PORT, () => {
          console.log(` + "`🤘 Server running on http://localhost:${PORT}`" + `);
        });
      } catch (err) {
        console.error("❌ Failed to start server:", err);
        process.exit(1);
      }
    })();

    export default server;`
}

// Default Index file with Express but no Orm
// [X] Layered, Clean or Hexagonale Architecture
// [X] Express.js
func getDefaultIndexWithExpress() string {
	var imports string = ""

	return `
    /**
     * The main server file
     * @module index
    **/

    // External dependencies
    import express, { Request, Response } from "express";
    import dotenv from "dotenv";

    // Internal dependencies
    ` + imports + `

    // import routes here

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

        // Start server
        const PORT = process.env.LISTEN_PORT ?? 3000;
        server.listen(PORT, () => {
          console.log(` + "`🤘 Server running on http://localhost:${PORT}`" + `);
        });
      } catch (err) {
        console.error("❌ Failed to start server:", err);
        process.exit(1);
      }
    })();

    export default server;`
}
