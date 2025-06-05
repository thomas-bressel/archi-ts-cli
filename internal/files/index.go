package files

/*
GetIndexTemplate generate the template for index.ts/js
*/
func GetIndexTemplate(language string, express bool) string {
	if language == "typescript" {
		return getTypeScriptIndexTemplate(express)
	}
	return getJavaScriptIndexTemplate(express)
}

/*
Return a template in the case of typescript depending of ExpressJS or not
*/
func getTypeScriptIndexTemplate(express bool) string {

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
			    stack: 'NodeJS, Typescript',
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
  				  stack: 'NodeJS, Typescript',
  				  library: "no library, I'm a badass developer"
  				}));
			});

			// Start server
			server.listen(PORT, () => {
			  console.log('Server running on http://' + 'localhost:' + PORT);
			});`
	}

}

/*
Return a template in the case of javascript depending of ExpressJS or not
*/
func getJavaScriptIndexTemplate(express bool) string {
	if express {
		return `
			const express = require('express');
			const server = express();
			const PORT = 3000;

			// Routes
			server.get('/', (req, res) => {
			  res.json({
			    message: 'Welcome to Archi API',
			    version: '1.0.0',
			    status: 'running',
			    stack: 'NodeJS, JavaScript',
			    library: 'ExpressJS'
			  });
			});

			// Start server
			server.listen(PORT, () => {
			  console.log('Server running on http://' + 'localhost:' + PORT);
			});

			module.exports = server;`
	} else {
		return `
			const http = require('http');
			const PORT = 3000;

			// Routes
			const server = http.createServer((req, res) => {
				res.setHeader('Content-Type', 'application/json');
  				res.end(JSON.stringify({
  				  message: 'Welcome to Archi API',
  				  version: '1.0.0',
  				  status: 'running',
  				  stack: 'NodeJS, JavaScript',
  				  library: "no library without typescript, you're an old school developer"
  				}));
			});

			// Start server
			server.listen(PORT, () => {
			  console.log('Server running on http://' + 'localhost:' + PORT);
			});

			module.exports = server;`
	}
}
