package tests

import (
	"archi-ts-cli/internal/models"
	"archi-ts-cli/internal/templates/generate/exports"
	"fmt"
	"strings"
)

// GetControllerTestTemplate generates unit test template for controllers
func GetApiIntegrationTestTemplate(cfg models.EntityConfig, architecture string) string {
	lowerName := strings.ToLower(cfg.Name)
	upperName := cfg.Name
	layerImport := ""

	switch architecture {
	case string(models.CleanArchitecture):
		layerImport = exports.GetTestCleanImports(upperName, lowerName)
	case string(models.LayeredArchitecture):
		layerImport = exports.GetTestLayeredImports(upperName, lowerName)
	}

	// Adapte template depending if there is an orm of not
	if cfg.Orm == models.TypeOrm {
		return fmt.Sprintf(`// Imports
import express, { Application } from 'express';
import request from 'supertest';

import { %sController } from '@controllers/%s.controller';
import { %sService } from '@services/%s.service';
import { %sRepository } from '@repositories/%s.repository';

%s

// Mock the repository for integration tests
jest.mock('@repositories/%s.repository');

describe('%s Integration Tests - Express', () => {
  let app: Application;
  let mockRepository: jest.Mocked<%sRepository>;
  let %sService: %sService;
  let %sController: %sController;

  beforeAll(() => {
    // Configure Express application for tests
    app = express();
    app.use(express.json());
  });

  beforeEach(() => {
    jest.clearAllMocks();
    
    // Setup mock repository
    mockRepository = new %sRepository() as jest.Mocked<%sRepository>;
    
    // Create instances with dependency injection
    %sService = new %sService(mockRepository);
    %sController = new %sController(%sService);
    
    // Setup Express routes
    setupRoutes();
  });

  const setupRoutes = () => {
    // Use the actual controller method with proper binding
    app.get('/api/v1/admin/%ss', async (req, res) => {
      await %sController.getAll(req, res);
    });

    // Route for unsupported methods
    app.all('/api/v1/admin/%ss', (req, res) => {
      if (req.method !== 'GET') {
        res.status(405).json({
          success: false,
          message: 'Method not allowed'
        });
      }
    });
  };

  describe('GET /api/v1/admin/%ss', () => {
    it('should return all %ss with status 200', async () => {
      // Arrange
      const mock%ss = [
        Object.assign(new %s(1), { id_%s: 1 }),
        Object.assign(new %s(2), { id_%s: 2 }),
        Object.assign(new %s(3), { id_%s: 3 })
      ];
      
      mockRepository.findAll = jest.fn().mockResolvedValue(mock%ss);

      // Act
      const response = await request(app)
        .get('/api/v1/admin/%ss')
        .expect('Content-Type', /json/);

      // Assert
      expect(response.status).toBe(200);
      expect(response.body).toEqual({
        success: true,
        data: mock%ss
      });
      expect(mockRepository.findAll).toHaveBeenCalledTimes(1);
    });

    it('should return empty array when no %ss found with status 200', async () => {
      // Arrange
      mockRepository.findAll = jest.fn().mockResolvedValue([]);

      // Act
      const response = await request(app)
        .get('/api/v1/admin/%ss')
        .expect('Content-Type', /json/);

      // Assert
      expect(response.status).toBe(200);
      expect(response.body).toEqual({
        success: true,
        data: []
      });
      expect(mockRepository.findAll).toHaveBeenCalledTimes(1);
    });

    it('should handle database errors and return 500', async () => {
      // Arrange
      const errorMessage = 'Database connection failed';
      mockRepository.findAll = jest.fn().mockRejectedValue(new Error(errorMessage));

      // Act
      const response = await request(app)
        .get('/api/v1/admin/%ss')
        .expect('Content-Type', /json/);

      // Assert
      expect(response.status).toBe(500);
      expect(response.body).toBe('Failed to fetch %ss: ' + errorMessage);
      expect(mockRepository.findAll).toHaveBeenCalledTimes(1);
    });

    it('should handle non-Error exceptions gracefully', async () => {
      // Arrange
      const errorString = 'Unexpected error occurred';
      mockRepository.findAll = jest.fn().mockRejectedValue(errorString);

      // Act
      const response = await request(app)
        .get('/api/v1/admin/%ss')
        .expect('Content-Type', /json/);

      // Assert
      expect(response.status).toBe(500);
      expect(response.body).toBe('Failed to fetch %ss: ' + errorString);
      expect(mockRepository.findAll).toHaveBeenCalledTimes(1);
    });
  });

  describe('HTTP Methods', () => {
    it('should return 405 for POST method', async () => {
      // Act
      const response = await request(app)
        .post('/api/v1/admin/%ss')
        .send({ name: 'Test %s' });

      // Assert
      expect(response.status).toBe(405);
      expect(response.body).toEqual({
        success: false,
        message: 'Method not allowed'
      });
    });

    it('should return 405 for PUT method', async () => {
      // Act
      const response = await request(app)
        .put('/api/v1/admin/%ss')
        .send({ name: 'Updated %s' });

      // Assert
      expect(response.status).toBe(405);
      expect(response.body).toEqual({
        success: false,
        message: 'Method not allowed'
      });
    });

    it('should return 405 for DELETE method', async () => {
      // Act
      const response = await request(app)
        .delete('/api/v1/admin/%ss');

      // Assert
      expect(response.status).toBe(405);
      expect(response.body).toEqual({
        success: false,
        message: 'Method not allowed'
      });
    });

    it('should return 405 for PATCH method', async () => {
      // Act
      const response = await request(app)
        .patch('/api/v1/admin/%ss')
        .send({ name: 'Patched %s' });

      // Assert
      expect(response.status).toBe(405);
      expect(response.body).toEqual({
        success: false,
        message: 'Method not allowed'
      });
    });
  });

  describe('Response Headers and Structure', () => {
    it('should return correct content-type header', async () => {
      // Arrange
      mockRepository.findAll = jest.fn().mockResolvedValue([]);

      // Act
      const response = await request(app)
        .get('/api/v1/admin/%ss');

      // Assert
      expect(response.headers['content-type']).toMatch(/application\/json/);
    });

    it('should return consistent response structure for success', async () => {
      // Arrange
      const mock%ss = [
        Object.assign(new %s(1), { id_%s: 1 })
      ];
      mockRepository.findAll = jest.fn().mockResolvedValue(mock%ss);

      // Act
      const response = await request(app)
        .get('/api/v1/admin/%ss');

      // Assert
      expect(response.body).toHaveProperty('success');
      expect(response.body).toHaveProperty('data');
      expect(response.body.success).toBe(true);
      expect(Array.isArray(response.body.data)).toBe(true);
    });

    it('should return consistent response structure for error', async () => {
      // Arrange
      mockRepository.findAll = jest.fn().mockRejectedValue(new Error('Test error'));

      // Act
      const response = await request(app)
        .get('/api/v1/admin/%ss');

      // Assert
      expect(response.status).toBe(500);
      expect(typeof response.body).toBe('string');
      expect(response.body).toContain('Failed to fetch %ss');
    });
  });

  describe('Edge Cases', () => {
    it('should handle null %ss response', async () => {
      // Arrange
      mockRepository.findAll = jest.fn().mockResolvedValue(null as any);

      // Act
      const response = await request(app)
        .get('/api/v1/admin/%ss');

      // Assert
      expect(response.status).toBe(200);
      expect(response.body).toEqual({
        success: true,
        data: []
      });
    });

    it('should handle undefined %ss response', async () => {
      // Arrange
      mockRepository.findAll = jest.fn().mockResolvedValue(undefined as any);

      // Act
      const response = await request(app)
        .get('/api/v1/admin/%ss');

      // Assert
      expect(response.status).toBe(200);
      expect(response.body).toEqual({
        success: true,
        data: []
      });
    });

    it('should handle very long error messages', async () => {
      // Arrange
      const longErrorMessage = 'A'.repeat(1000);
      mockRepository.findAll = jest.fn().mockRejectedValue(new Error(longErrorMessage));

      // Act
      const response = await request(app)
        .get('/api/v1/admin/%ss');

      // Assert
      expect(response.status).toBe(500);
      expect(response.body).toContain(longErrorMessage);
    });
  });

  describe('Performance and Load', () => {
    it('should handle multiple concurrent requests', async () => {
      // Arrange
      const mock%ss = [
        Object.assign(new %s(1), { id_%s: 1 }),
        Object.assign(new %s(2), { id_%s: 2 })
      ];
      mockRepository.findAll = jest.fn().mockResolvedValue(mock%ss);

      // Act - Make 5 concurrent requests
      const promises = Array(5).fill(null).map(() => 
        request(app).get('/api/v1/admin/%ss')
      );
      
      const responses = await Promise.all(promises);

      // Assert
      responses.forEach(response => {
        expect(response.status).toBe(200);
        expect(response.body).toEqual({
          success: true,
          data: mock%ss
        });
      });
      expect(mockRepository.findAll).toHaveBeenCalledTimes(5);
    });

    it('should handle large dataset response', async () => {
      // Arrange - Create 1000 %ss
      const large%sSet = Array.from({ length: 1000 }, (_, i) => 
        Object.assign(new %s(1), { id_%s: i + 1 })
      );
      mockRepository.findAll = jest.fn().mockResolvedValue(large%sSet);

      // Act
      const response = await request(app)
        .get('/api/v1/admin/%ss');

      // Assert
      expect(response.status).toBe(200);
      expect(response.body.success).toBe(true);
      expect(response.body.data).toHaveLength(1000);
      expect(response.body.data[0].id_%s).toBe(1);
      expect(response.body.data[999].id_%s).toBe(1000);
    });

    it('should handle slow database response', async () => {
      // Arrange - Simulate slow database with 2 second delay
      mockRepository.findAll = jest.fn().mockImplementation(() => 
        new Promise(resolve => 
          setTimeout(() => resolve([Object.assign(new %s(1), { id_%s: 1 })]), 100)
        )
      );

      // Act
      const startTime = Date.now();
      const response = await request(app)
        .get('/api/v1/admin/%ss');
      const endTime = Date.now();

      // Assert
      expect(response.status).toBe(200);
      expect(response.body.success).toBe(true);
      expect(endTime - startTime).toBeGreaterThanOrEqual(100);
    });
  });

  describe('Query Parameters', () => {
    it('should ignore query parameters for this endpoint', async () => {
      // Arrange
      mockRepository.findAll = jest.fn().mockResolvedValue([]);

      // Act
      const response = await request(app)
        .get('/api/v1/admin/%ss?page=1&limit=10&sort=name');

      // Assert
      expect(response.status).toBe(200);
      expect(response.body).toEqual({
        success: true,
        data: []
      });
      // Verify findAll was called without parameters
      expect(mockRepository.findAll).toHaveBeenCalledWith();
    });
  });

  describe('Request Headers', () => {
    it('should handle requests with various headers', async () => {
      // Arrange
      mockRepository.findAll = jest.fn().mockResolvedValue([]);

      // Act
      const response = await request(app)
        .get('/api/v1/admin/%ss')
        .set('Accept', 'application/json')
        .set('User-Agent', 'Jest Test')
        .set('X-Custom-Header', 'test-value');

      // Assert
      expect(response.status).toBe(200);
      expect(response.body).toEqual({
        success: true,
        data: []
      });
    });
  });
});`, cfg.Name, lowerName, cfg.Name, lowerName, cfg.Name, lowerName, layerImport,
			lowerName,
			cfg.Name, cfg.Name, lowerName, cfg.Name, lowerName, cfg.Name,
			cfg.Name, cfg.Name, lowerName, cfg.Name, lowerName, cfg.Name, lowerName,
			lowerName, lowerName, lowerName, lowerName, lowerName,
			cfg.Name, cfg.Name, lowerName, cfg.Name, lowerName, cfg.Name, lowerName, cfg.Name,
			lowerName, cfg.Name,
			lowerName, lowerName, lowerName, lowerName, lowerName, lowerName, lowerName,
			cfg.Name, lowerName, cfg.Name, lowerName, lowerName, cfg.Name, lowerName,
			cfg.Name, cfg.Name, lowerName, cfg.Name, lowerName,
			lowerName, lowerName, lowerName, lowerName, lowerName, lowerName, lowerName,
			cfg.Name, cfg.Name, lowerName, cfg.Name, lowerName, cfg.Name, lowerName, cfg.Name,
			lowerName, cfg.Name, cfg.Name, lowerName, cfg.Name, lowerName, lowerName, lowerName,
			cfg.Name, lowerName, lowerName, lowerName, lowerName,
		)
	}

	if cfg.Express {
		return fmt.Sprintf(`// Imports
import express, { Application } from 'express';
import request from 'supertest';

// Layer imports
import { %sController } from '@controllers/%s.controller';
import { %sService } from '@services/%s.service';
import { %sRepository } from '@repositories/%s.repository';


%s

// Mock the repository for integration tests
jest.mock('@repositories/%s.repository');

describe('%s Integration Tests - Express', () => {
  let app: Application;
  let mockRepository: jest.Mocked<%sRepository>;
  let %sController: %sController;

  beforeAll(() => {
    // Configure Express application for tests
    app = express();
    app.use(express.json());
  });

  beforeEach(() => {
    jest.clearAllMocks();
    
    // Setup mock repository
    mockRepository = new %sRepository() as jest.Mocked<%sRepository>;
    
    // Mock the constructor to return our mock
    (%sRepository as jest.MockedClass<typeof %sRepository>).mockImplementation(() => mockRepository);
    
    // Create instances with dependency injection
    const %sService = new %sService(mockRepository);
    %sController = new %sController(%sService);
    
    // Setup Express routes
    setupRoutes();
  });

  const setupRoutes = () => {
    // Adapt controller method for Express
    app.get('/api/v1/admin/%s', async (req, res) => {
      try {
        const %ss = await %sController['%sService'].findAll();

        if (!%ss || %ss.length === 0) {
          return res.status(200).json({
            success: true,
            data: []
          });
        }

        res.status(200).json({
          success: true,
          data: %ss
        });
      } catch (error) {
        res.status(500).json({
          success: false,
          message: error instanceof Error ? error.message : "Internal server error"
        });
      }
    });

    // Route for unsupported methods
    app.all('/api/v1/admin/%s', (req, res) => {
      if (req.method !== 'GET') {
        res.status(405).json({
          success: false,
          message: 'Method not allowed'
        });
      }
    });
  };

  describe('GET /api/v1/admin/%s', () => {
    it('should return all %ss with status 200', async () => {
      // Arrange
      const mock%ss: %s[] = [
        new %s(1),
        new %s(2),
        new %s(3)
      ];
      mockRepository.findAll.mockResolvedValue(mock%ss);

      // Act & Assert
      const response = await request(app)
        .get('/api/v1/admin/%s')
        .expect(200)
        .expect('Content-Type', /json/);

      expect(response.body.success).toBe(true);
      expect(response.body.data).toHaveLength(3);
      expect(response.body.data[0]).toHaveProperty('id_%s', 1);
      expect(response.body.data[1]).toHaveProperty('id_%s', 2);
      expect(response.body.data[2]).toHaveProperty('id_%s', 3);
    });

    it('should return empty array when no %ss found with status 200', async () => {
      // Arrange
      mockRepository.findAll.mockResolvedValue([]);

      // Act & Assert
      const response = await request(app)
        .get('/api/v1/admin/%s')
        .expect(200)
        .expect('Content-Type', /json/);

      expect(response.body.success).toBe(true);
      expect(response.body.data).toEqual([]);
      expect(response.body.data).toHaveLength(0);
    });

    it('should handle database errors and return 500', async () => {
      // Arrange
      mockRepository.findAll.mockRejectedValue(new Error('Database connection failed'));

      // Act & Assert
      const response = await request(app)
        .get('/api/v1/admin/%s')
        .expect(500)
        .expect('Content-Type', /json/);

      expect(response.body.success).toBe(false);
      expect(response.body.message).toContain('Failed to fetch %ss');
    });

    it('should handle non-Error exceptions gracefully', async () => {
      // Arrange
      mockRepository.findAll.mockRejectedValue('Unknown error string');

      // Act & Assert
      const response = await request(app)
        .get('/api/v1/admin/%s')
        .expect(500)
        .expect('Content-Type', /json/);

      expect(response.body.success).toBe(false);
      expect(response.body.message).toContain('Failed to fetch %ss');
    });
  });

  describe('HTTP Methods', () => {
    it('should return 405 for POST method', async () => {
      // Act & Assert
      const response = await request(app)
        .post('/api/v1/admin/%s')
        .send({})
        .expect(405)
        .expect('Content-Type', /json/);

      expect(response.body.success).toBe(false);
      expect(response.body.message).toBe('Method not allowed');
    });

    it('should return 405 for PUT method', async () => {
      // Act & Assert
      const response = await request(app)
        .put('/api/v1/admin/%s')
        .send({})
        .expect(405)
        .expect('Content-Type', /json/);

      expect(response.body.success).toBe(false);
      expect(response.body.message).toBe('Method not allowed');
    });

    it('should return 405 for DELETE method', async () => {
      // Act & Assert
      const response = await request(app)
        .delete('/api/v1/admin/%s')
        .expect(405)
        .expect('Content-Type', /json/);

      expect(response.body.success).toBe(false);
      expect(response.body.message).toBe('Method not allowed');
    });
  });

  describe('Response Headers and Structure', () => {
    it('should return correct content-type header', async () => {
      // Arrange
      mockRepository.findAll.mockResolvedValue([]);

      // Act & Assert
      await request(app)
        .get('/api/v1/admin/%s')
        .expect('Content-Type', /json/)
        .expect(200);
    });

    it('should return consistent response structure for success', async () => {
      // Arrange
      const mock%ss: %s[] = [new %s(1)];
      mockRepository.findAll.mockResolvedValue(mock%ss);

      // Act & Assert
      const response = await request(app)
        .get('/api/v1/admin/%s')
        .expect(200);

      expect(response.body).toHaveProperty('success');
      expect(response.body).toHaveProperty('data');
      expect(response.body.success).toBe(true);
      expect(Array.isArray(response.body.data)).toBe(true);
    });

    it('should return consistent response structure for error', async () => {
      // Arrange
      mockRepository.findAll.mockRejectedValue(new Error('Test error'));

      // Act & Assert
      const response = await request(app)
        .get('/api/v1/admin/%s')
        .expect(500);

      expect(response.body).toHaveProperty('success');
      expect(response.body).toHaveProperty('message');
      expect(response.body.success).toBe(false);
      expect(typeof response.body.message).toBe('string');
    });
  });

  describe('Edge Cases', () => {
    it('should handle null %ss response', async () => {
      // Arrange
      mockRepository.findAll.mockResolvedValue(null as any);

      // Act & Assert
      const response = await request(app)
        .get('/api/v1/admin/%s')
        .expect(200);

      expect(response.body.success).toBe(true);
      expect(response.body.data).toEqual([]);
    });

    it('should handle undefined %ss response', async () => {
      // Arrange
      mockRepository.findAll.mockResolvedValue(undefined as any);

      // Act & Assert
      const response = await request(app)
        .get('/api/v1/admin/%s')
        .expect(200);

      expect(response.body.success).toBe(true);
      expect(response.body.data).toEqual([]);
    });
  });

  describe('Performance and Load', () => {
    it('should handle multiple concurrent requests', async () => {
      // Arrange
      const mock%ss: %s[] = [new %s(1), new %s(2)];
      mockRepository.findAll.mockResolvedValue(mock%ss);

      // Act - Make 5 concurrent requests
      const requests = Array(5).fill(null).map(() => 
        request(app).get('/api/v1/admin/%s')
      );

      const responses = await Promise.all(requests);

      // Assert
      responses.forEach(response => {
        expect(response.status).toBe(200);
        expect(response.body.success).toBe(true);
        expect(response.body.data).toHaveLength(2);
      });

      expect(mockRepository.findAll).toHaveBeenCalledTimes(5);
    });

    it('should handle large dataset response', async () => {
      // Arrange
      const largeDataset: %s[] = Array(1000).fill(null).map((_, index) => new %s(index + 1));
      mockRepository.findAll.mockResolvedValue(largeDataset);

      // Act & Assert
      const response = await request(app)
        .get('/api/v1/admin/%s')
        .expect(200);

      expect(response.body.success).toBe(true);
      expect(response.body.data).toHaveLength(1000);
      expect(response.body.data[0].id_%s).toBe(1);
      expect(response.body.data[999].id_%s).toBe(1000);
    });
  });
});`, cfg.Name, lowerName, cfg.Name, lowerName, cfg.Name, lowerName, layerImport,
			lowerName,
			cfg.Name, cfg.Name, lowerName, cfg.Name,
			cfg.Name, cfg.Name, cfg.Name, cfg.Name,
			lowerName, cfg.Name, lowerName, cfg.Name, lowerName,
			lowerName, lowerName, lowerName, lowerName, lowerName, lowerName, lowerName, lowerName,
			lowerName, lowerName, cfg.Name, cfg.Name, cfg.Name, cfg.Name, cfg.Name, cfg.Name,
			lowerName, lowerName, lowerName, lowerName,
			lowerName, lowerName,
			lowerName, lowerName, lowerName, lowerName,
			lowerName, lowerName, lowerName, lowerName,
			cfg.Name, cfg.Name, cfg.Name, cfg.Name,
			lowerName, lowerName, lowerName, lowerName,
			lowerName, lowerName,
			cfg.Name, cfg.Name, cfg.Name, cfg.Name, cfg.Name,
			lowerName,
			cfg.Name, cfg.Name, lowerName, lowerName, lowerName,
		)
	} else {

		return fmt.Sprintf(`// Imports
import * as http from "http";
import { %sRepository } from '@repositories/%s.repository';
import { %sService } from '@services/%s.service';
import { %sController } from '@controllers/%s.controller';
import dotenv from "dotenv";
dotenv.config();

%s

// Mock the repository for integration tests
jest.mock('@repositories/%s.repository');

describe('%s Integration Tests - Node.js Native', () => {
  let mockRepository: jest.Mocked<%sRepository>;
  let %sController: %sController;
  const basePort = Number(process.env.LISTEN_PORT_TEST); // Base port for tests

  // Utility function to create a server for each test
  const createTestServer = (): http.Server => {
    return http.createServer(async (req, res) => {
      if (!req.url) {
        res.statusCode = 400;
        res.end(JSON.stringify({ message: 'Bad request' }));
        return;
      }

      const url = new URL(req.url, 'http://localhost:' + basePort);
      const pathname = url.pathname;
      const method = req.method?.toUpperCase();

      try {
        // GET /api/v1/admin/%s
        if (pathname === "/api/v1/admin/%s" && method === "GET") {
          await %sController.getAll(req, res);
          return;
        }
        
        // Route not found
        res.statusCode = 404;
        res.setHeader("Content-Type", "application/json");
        res.end(JSON.stringify({ message: 'Route not found' }));
      } catch (error) {
        res.statusCode = 500;
        res.setHeader("Content-Type", "application/json");
        res.end(JSON.stringify({
          message: error instanceof Error ? error.message : "Internal server error"
        }));
      }
    });
  };

  beforeEach(() => {
    jest.clearAllMocks();
    
    // Setup mock repository with all necessary mocks
    const MockedRepository = %sRepository as jest.MockedClass<typeof %sRepository>;
    mockRepository = {
      findAll: jest.fn()
    } as any;
    
    // Mock the constructor to return our mocked instance
    MockedRepository.mockImplementation(() => mockRepository);
    
    // Create instances with dependency injection
    const %sService = new %sService(mockRepository);
    %sController = new %sController(%sService);
  });

  describe('GET /api/v1/admin/%s', () => {
    it('should return all %ss with status 200', (done) => {
      // Arrange
      const mock%ss: %s[] = [
        new %s(1),
        new %s(2)
      ];
      mockRepository.findAll.mockResolvedValue(mock%ss);

      const server = createTestServer();
      const port = basePort;

      server.listen(port, () => {
        // Act
        const req = http.request({
          hostname: 'localhost',
          port: port,
          path: '/api/v1/admin/%s',
          method: 'GET',
          headers: {
            'Content-Type': 'application/json'
          }
        }, (res) => {
          let data = '';
          
          res.on('data', (chunk) => {
            data += chunk;
          });

          res.on('end', () => {
            try {
              // Assert
              expect(res.statusCode).toBe(200);
              expect(res.headers['content-type']).toContain('application/json');
              
              const response = JSON.parse(data);
              expect(response.success).toBe(true);
              expect(response.data).toHaveLength(2);
              expect(response.data[0]).toHaveProperty('id_%s', 1);
              expect(response.data[1]).toHaveProperty('id_%s', 2);
              
              server.close(done);
            } catch (error) {
              server.close();
              done(error);
            }
          });
        });

        req.on('error', (error) => {
          server.close();
          done(error);
        });
        req.end();
      });

      server.on('error', (error) => {
        done(error);
      });
    }, 10000); // 10 seconds timeout

    it('should return empty array when no %ss found with status 200', (done) => {
      // Arrange
      mockRepository.findAll.mockResolvedValue([]);

      const server = createTestServer();
      const port = basePort + 1;

      server.listen(port, () => {
        // Act
        const req = http.request({
          hostname: 'localhost',
          port: port,
          path: '/api/v1/admin/%s',
          method: 'GET'
        }, (res) => {
          let data = '';
          
          res.on('data', (chunk) => {
            data += chunk;
          });

          res.on('end', () => {
            try {
              // Assert
              expect(res.statusCode).toBe(200);
              
              const response = JSON.parse(data);
              expect(response.success).toBe(true);
              expect(response.data).toEqual([]);
              expect(response.data).toHaveLength(0);
              
              server.close(done);
            } catch (error) {
              server.close();
              done(error);
            }
          });
        });

        req.on('error', (error) => {
          server.close();
          done(error);
        });
        req.end();
      });

      server.on('error', (error) => {
        done(error);
      });
    }, 10000);

    it('should handle database errors and return 500', (done) => {
      // Arrange
      mockRepository.findAll.mockRejectedValue(new Error('Database connection failed'));

      const server = createTestServer();
      const port = basePort + 2;

      server.listen(port, () => {
        // Act
        const req = http.request({
          hostname: 'localhost',
          port: port,
          path: '/api/v1/admin/%s',
          method: 'GET'
        }, (res) => {
          let data = '';
          
          res.on('data', (chunk) => {
            data += chunk;
          });

          res.on('end', () => {
            try {
              // Assert
              expect(res.statusCode).toBe(500);
              
              const response = JSON.parse(data);
              expect(response.success).toBe(false);
              expect(response.message).toContain('Failed to fetch %ss');
              
              server.close(done);
            } catch (error) {
              server.close();
              done(error);
            }
          });
        });

        req.on('error', (error) => {
          server.close();
          done(error);
        });
        req.end();
      });

      server.on('error', (error) => {
        done(error);
      });
    }, 10000);

    it('should return 404 for unknown routes', (done) => {
      const server = createTestServer();
      const port = basePort + 3;

      server.listen(port, () => {
        // Act
        const req = http.request({
          hostname: 'localhost',
          port: port,
          path: '/api/v1/unknown',
          method: 'GET'
        }, (res) => {
          let data = '';
          
          res.on('data', (chunk) => {
            data += chunk;
          });

          res.on('end', () => {
            try {
              // Assert
              expect(res.statusCode).toBe(404);
              
              const response = JSON.parse(data);
              expect(response.message).toBe('Route not found');
              
              server.close(done);
            } catch (error) {
              server.close();
              done(error);
            }
          });
        });

        req.on('error', (error) => {
          server.close();
          done(error);
        });
        req.end();
      });

      server.on('error', (error) => {
        done(error);
      });
    }, 10000);

    it('should handle POST method correctly (method not allowed)', (done) => {
      const server = createTestServer();
      const port = basePort + 4;

      server.listen(port, () => {
        // Act
        const req = http.request({
          hostname: 'localhost',
          port: port,
          path: '/api/v1/admin/%s',
          method: 'POST'
        }, (res) => {
          let data = '';
          
          res.on('data', (chunk) => {
            data += chunk;
          });

          res.on('end', () => {
            try {
              // Assert
              expect(res.statusCode).toBe(404); // Route not found since method is not supported
              
              const response = JSON.parse(data);
              expect(response.message).toBe('Route not found');
              
              server.close(done);
            } catch (error) {
              server.close();
              done(error);
            }
          });
        });

        req.on('error', (error) => {
          server.close();
          done(error);
        });
        req.end();
      });

      server.on('error', (error) => {
        done(error);
      });
    }, 10000);
  });

  describe('Response Headers', () => {
    it('should set correct content-type header', (done) => {
      // Arrange
      mockRepository.findAll.mockResolvedValue([]);

      const server = createTestServer();
      const port = basePort + 5;

      server.listen(port, () => {
        // Act
        const req = http.request({
          hostname: 'localhost',
          port: port,
          path: '/api/v1/admin/%s',
          method: 'GET'
        }, (res) => {
          try {
            // Assert
            expect(res.headers['content-type']).toContain('application/json');
            
            server.close(done);
          } catch (error) {
            server.close();
            done(error);
          }
        });

        req.on('error', (error) => {
          server.close();
          done(error);
        });
        req.end();
      });

      server.on('error', (error) => {
        done(error);
      });
    }, 10000);
  });

  describe('Error Handling', () => {
    it('should handle non-Error exceptions gracefully', (done) => {
      // Arrange
      mockRepository.findAll.mockRejectedValue('Unknown error string');

      const server = createTestServer();
      const port = basePort + 6;

      server.listen(port, () => {
        // Act
        const req = http.request({
          hostname: 'localhost',
          port: port,
          path: '/api/v1/admin/%s',
          method: 'GET'
        }, (res) => {
          let data = '';
          
          res.on('data', (chunk) => {
            data += chunk;
          });

          res.on('end', () => {
            try {
              // Assert
              expect(res.statusCode).toBe(500);
              
              const response = JSON.parse(data);
              expect(response.success).toBe(false);
              expect(response.message).toContain('Failed to fetch %ss');
              
              server.close(done);
            } catch (error) {
              server.close();
              done(error);
            }
          });
        });

        req.on('error', (error) => {
          server.close();
          done(error);
        });
        req.end();
      });

      server.on('error', (error) => {
        done(error);
      });
    }, 10000);
  });
});`, cfg.Name, lowerName, cfg.Name, lowerName, cfg.Name, lowerName, layerImport,
			lowerName,
			cfg.Name, cfg.Name, lowerName, cfg.Name,
			lowerName, lowerName, lowerName,
			cfg.Name, cfg.Name,
			lowerName, cfg.Name, lowerName, cfg.Name, lowerName,
			lowerName, lowerName, cfg.Name, cfg.Name, cfg.Name, cfg.Name, cfg.Name, lowerName, lowerName, lowerName,
			lowerName, lowerName, lowerName, lowerName, lowerName, lowerName, lowerName, lowerName,
		)
	}

}
