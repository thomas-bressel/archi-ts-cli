package generate

import (
	"fmt"
	"strings"
)

// GetControllerTestTemplate generates unit test template for controllers
func GetControllerTestTemplate(cfg EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	lowerPlural := lowerName + "s"
	return fmt.Sprintf(`// Express importation
import { Request, Response } from 'express';

// Layers importation
import { %sController } from '../../../src/controllers/%s.controller';
import { %sService } from '../../../src/services/%s.service';
import { %sRepository } from '../../../src/repositories/%s.repository'; // Note: Only needed if you instantiate the service, which is not ideal for a pure mock.

// Mock the service module
jest.mock('../../../src/services/%s.service');

describe('%sController', () => {
  let controller: %sController;
  let mockService: jest.Mocked<%sService>;
  let mockRequest: Partial<Request>;
  let mockResponse: Partial<Response>;

  beforeEach(() => {
    // Clear all mocks before each test
    jest.clearAllMocks();

    // Create mock service instance. Since the module is mocked, 'new' will return a mocked constructor.
    mockService = new %sService(new %sRepository()) as jest.Mocked<%sService>;
    
    // Initialize controller with mock service (dependency injection)
    controller = new %sController(mockService);

    // Setup mock request
    mockRequest = {
      params: {},
      body: {},
      query: {},
      headers: {},
    };

    // Setup mock response with chainable methods
    mockResponse = {
      status: jest.fn().mockReturnThis(),
      json: jest.fn().mockReturnThis(),
      send: jest.fn().mockReturnThis(),
    };
  });

  describe('getAll', () => {
    it('should return all %s successfully with status 200', async () => {
      // Arrange
      const mock%sData = [
        { id_%s: 1 },
        { id_%s: 2 },
      ];
      mockService.findAll.mockResolvedValue(mock%sData);

      // Act
      await controller.getAll(mockRequest as Request, mockResponse as Response);

      // Assert
      expect(mockService.findAll).toHaveBeenCalledTimes(1);
      expect(mockResponse.status).toHaveBeenCalledWith(200);
      expect(mockResponse.json).toHaveBeenCalledWith({
        success: true,
        data: mock%sData,
      });
    });

    it('should return 200 with empty array when service returns empty array', async () => {
      // Arrange
      mockService.findAll.mockResolvedValue([]);

      // Act
      await controller.getAll(mockRequest as Request, mockResponse as Response);

      // Assert
      expect(mockService.findAll).toHaveBeenCalledTimes(1);
      expect(mockResponse.status).toHaveBeenCalledWith(200);
      expect(mockResponse.json).toHaveBeenCalledWith({
        success: true,
        data: [],
      });
    });

    it('should handle errors when fetching %s fails', async () => {
      // Arrange
      const errorMessage = 'Database connection error';
      mockService.findAll.mockRejectedValue(new Error(errorMessage));

      // Act
      await controller.getAll(mockRequest as Request, mockResponse as Response);

      // Assert
      expect(mockService.findAll).toHaveBeenCalledTimes(1);
      expect(mockResponse.status).toHaveBeenCalledWith(500);
      expect(mockResponse.json).toHaveBeenCalledWith(errorMessage);
    });

    it('should return "Internal server error" for non-Error exceptions', async () => {
      // Arrange
      mockService.findAll.mockRejectedValue('Unknown error');

      // Act
      await controller.getAll(mockRequest as Request, mockResponse as Response);

      // Assert
      expect(mockService.findAll).toHaveBeenCalledTimes(1);
      expect(mockResponse.status).toHaveBeenCalledWith(500);
      expect(mockResponse.json).toHaveBeenCalledWith("Internal server error");
    });
  });
});
`,
		// Arguments for fmt.Sprintf, in order of appearance of %s
		cfg.Name, lowerName, cfg.Name, lowerName, cfg.Name, lowerName, // Imports
		lowerName,          // jest.mock
		cfg.Name,           // describe block
		cfg.Name, cfg.Name, // let declarations
		cfg.Name, cfg.Name, cfg.Name, // beforeEach: new Service()
		cfg.Name,    // beforeEach: new Controller()
		lowerPlural, // it should return all...
		cfg.Name,    // mock%sData
		lowerName,   // id_%s: 1
		lowerName,   // id_%s: 2
		cfg.Name,    // mockResolvedValue
		cfg.Name,    // expect mockData
		lowerPlural, // it should handle errors...
	)
}

// GetServiceTestTemplate generates unit test template for services
func GetServiceTestTemplate(cfg EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	lowerPlural := lowerName + "s"
	return fmt.Sprintf(`import { %sService } from '../../../src/services/%s.service';
import { %sRepository } from '../../../src/repositories/%s.repository';
import %s from '../../../src/entities/%s.entity';

// Mock the repository module
jest.mock('../../../src/repositories/%s.repository');

describe('%sService', () => {
  let service: %sService;
  let mockRepository: jest.Mocked<%sRepository>;

  beforeEach(() => {
    // Clear all mocks before each test
    jest.clearAllMocks();

    // Create mock repository instance
    mockRepository = new %sRepository() as jest.Mocked<%sRepository>;
    
    // Initialize service with mock repository
    service = new %sService(mockRepository);
  });

  describe('findAll', () => {
    it('should return all %s from repository', async () => {
      // Arrange
      const mock%sData: %s[] = [
        new %s(1, /* add constructor params */),
        new %s(2, /* add constructor params */),
      ];
      mockRepository.findAll.mockResolvedValue(mock%sData);

      // Act
      const result = await service.findAll();

      // Assert
      expect(mockRepository.findAll).toHaveBeenCalledTimes(1);
      expect(result).toEqual(mock%sData);
      expect(result).toHaveLength(2);
    });

    it('should return empty array when repository returns no %s', async () => {
      // Arrange
      mockRepository.findAll.mockResolvedValue([]);

      // Act
      const result = await service.findAll();

      // Assert
      expect(mockRepository.findAll).toHaveBeenCalledTimes(1);
      expect(result).toEqual([]);
      expect(result).toHaveLength(0);
    });

    it('should propagate repository errors', async () => {
      // Arrange
      const error = new Error('Database error');
      mockRepository.findAll.mockRejectedValue(error);
    
      // Act & Assert
      await expect(service.findAll()).rejects.toThrow('Failed to fetch %s: Database error');
      expect(mockRepository.findAll).toHaveBeenCalledTimes(1);
    });
  });

  // Add more test suites for other methods
  describe('findById', () => {
    it.todo('should return a single %s by id');
    it.todo('should return null when %s not found');
  });
});
`,
		// Arguments for fmt.Sprintf
		cfg.Name, lowerName, cfg.Name, lowerName, cfg.Name, lowerName, // imports
		lowerName,          // jest.mock
		cfg.Name,           // describe
		cfg.Name, cfg.Name, // let declarations
		cfg.Name, cfg.Name, // beforeEach: mockRepository
		cfg.Name,                               // beforeEach: new Service()
		lowerPlural,                            // it should return all...
		cfg.Name, cfg.Name, cfg.Name, cfg.Name, // mockData
		cfg.Name,             // mockResolvedValue
		cfg.Name,             // expect(result).toEqual
		lowerPlural,          // it should return empty...
		lowerPlural,          // propagate errors
		lowerName, lowerName, // it.todo
	)
}

// GetRepositoryTestTemplate generates unit test template for repositories
func GetRepositoryTestTemplate(cfg EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	lowerPlural := lowerName + "s"
	return fmt.Sprintf(`import { %sRepository } from '../../../src/repositories/%s.repository';
import %s from '../../../src/entities/%s.entity';

// Mock database/ORM if needed
// jest.mock('../../../src/config/database');

describe('%sRepository', () => {
  let repository: %sRepository;

  beforeEach(() => {
    // Clear all mocks before each test
    jest.clearAllMocks();
    
    // Initialize repository
    repository = new %sRepository();
  });

  describe('findAll', () => {
    it('should return all %s from database', async () => {
      // Arrange
      // This test might need more complex mocking depending on your data source
      const expected%sData: %s[] = [
        new %s(1, /* add constructor params */),
        new %s(2, /* add constructor params */),
      ];
      // Example of mocking a database call if needed:
      // YourDataSource.find.mockResolvedValue(expected%sData);

      // Act
      const result = await repository.findAll();

      // Assert
      expect(result).toBeDefined();
      expect(Array.isArray(result)).toBe(true);
      // expect(result).toEqual(expected%sData);
    });

    it('should return empty array when no %s found', async () => {
      // Arrange
      // Mock the data source to return an empty array
      // YourDataSource.find.mockResolvedValue([]);

      // Act
      const result = await repository.findAll();

      // Assert
      expect(result).toEqual([]);
      expect(result).toHaveLength(0);
    });
  });

  // Add more test suites for other methods
  describe('findById', () => {
    it.todo('should return a single %s by id');
    it.todo('should return null when %s not found');
  });
});
`,
		// Arguments for fmt.Sprintf
		cfg.Name, lowerName, cfg.Name, lowerName, // imports
		cfg.Name,                               // describe
		cfg.Name,                               // let declaration
		cfg.Name,                               // new Repository
		lowerPlural,                            // it should return all...
		cfg.Name, cfg.Name, cfg.Name, cfg.Name, // mockData
		cfg.Name,             // mockResolvedValue (commented out)
		cfg.Name,             // expect(result).toEqual (commented out)
		lowerPlural,          // it should return empty...
		lowerName, lowerName, // it.todo
	)
}
