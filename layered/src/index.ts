/**
 * The main server file for the React_Node application.
 * @module index
**/

import express, { Request, Response } from 'express';
const server = express();


/* @Archi::imports-start -- DO NOT REMOVE */
/* @Archi::imports-end */

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

export default server;