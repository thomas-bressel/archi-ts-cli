package files

// GetIndexTemplate generate the template for index.ts
func GetIndexTemplate(express bool) string {
	if express {
		return `
import express, { Express, Request, Response } from 'express';
const server = express();
const PORT = 3000;

// Routes
server.get('/', (req: Request, res: Response) => {
  res.json({
    message: 'Welcome to Archi API',
    version: '1.0.0',
    status: 'running',
    stack: 'NodeJS, TypeScript',
    library: 'ExpressJS'
  });
});

// Start server
server.listen(PORT, () => {
  console.log('Server running on http://' + 'localhost:' + PORT);
});

export default server;`
	} else {
		return `
import * as http from 'http';
const PORT = 3000;

// Routes
const server = http.createServer((req: http.IncomingMessage, res: http.ServerResponse) => {
	res.setHeader('Content-Type', 'application/json');
	res.end(JSON.stringify({
	  message: 'Welcome to Archi API',
	  version: '1.0.0',
	  status: 'running',
	  stack: 'NodeJS, TypeScript',
	  library: "no library, I'm a badass developer"
	}));
});

// Start server
server.listen(PORT, () => {
  console.log('Server running on http://' + 'localhost:' + PORT);
});`
	}
}
