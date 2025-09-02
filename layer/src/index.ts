
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

			export default server;