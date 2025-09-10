package files

import "archi-ts-cli/internal/models"

// GetIndexTemplate generate the template for index.ts
func GetIndexTemplate(architecture models.Architecture, express bool, orm models.Orm) string {

	// For other architectures, use the default template
	if express {
		return getDefaultIndexWithExpress()
	}

	// Only with native Node.js
	return getDefaultIndexPlain()

}

// Default Index file without any library or Orm
func getDefaultIndexPlain() string {
	return `
// Node.js built-in modules
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
func getDefaultIndexWithExpress() string {
	return `
/**
 * The main server file for the React_Node application.
 * @module index
**/

import express, { Request, Response } from 'express';
const server = express();

// Default route
server.get('/', (req: Request, res: Response) => {
  res.json({
    message: 'Welcome to ArchiTS API',
    version: '1.0.0',
    status: 'running',
    stack: 'NodeJS, TypeScript',
    library: 'ExpressJS'
  });
});

// Start server
const PORT = 3000;
server.listen(PORT, () => {
  console.log('Server running on http://' + 'localhost:' + PORT);
});

export default server;`
}
