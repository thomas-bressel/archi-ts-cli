package generate

import (
	"fmt"
	"strings"
)

// GetControllerTestTemplate generates unit test template for controllers
func GetControllerTestTemplate(cfg EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	return fmt.Sprintf(`import { Request, Response } from 'express';
import { %sController } from '../../../src/controllers/%s.controller';
import { %sService } from '../../../src/services/%s.service';

// Mock the service module
jest.mock('../../../src/services/%s.service');

describe('%sController', () => {
  let controller: %sController;
  let mockService: jest.Mocked<%sService>;
  let mockRequest: Partial<Request>;
  let mockResponse: Partial<Response>;
  let mockNext: jest.Mock;

  beforeEach(() => {
    // Clear all mocks before each test
    jest.clearAllMocks();

    // Create mock service instance
    mockService = new %sService() as jest.Mocked<%sService>;
    
    // Initialize controller with mock service
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

    // Setup mock next function
    mockNext = jest.fn();
  });

  describe('getAll', () => {
    it('should return all %s successfully', async () => {
      // Arrange
      const mock%sData = [
        { id_%s: 1, /* add your fields */ },
        { id_%s: 2, /* add your fields */ },
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
  });

  // Add more test suites for other methods
  describe('getById', () => {
    it.todo('should return a single %s by id');
    it.todo('should return 404 when %s not found');
  });
});
`, cfg.Name, lowerName, cfg.Name, lowerName,
		lowerName,
		cfg.Name, cfg.Name, cfg.Name,
		cfg.Name, cfg.Name,
		cfg.Name,
		lowerName,
		cfg.Name, lowerName, lowerName, cfg.Name,
		cfg.Name,
		lowerName,
		lowerName, lowerName,
	)
}

// GetServiceTestTemplate generates unit test template for services
func GetServiceTestTemplate(cfg EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
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

    it('should propagate repository errors', async () => {
      // Arrange
      const error = new Error('Database error');
      mockRepository.findAll.mockRejectedValue(error);

      // Act & Assert
      await expect(service.findAll()).rejects.toThrow('Database error');
      expect(mockRepository.findAll).toHaveBeenCalledTimes(1);
    });
  });

  // Add more test suites for other methods
  describe('findById', () => {
    it.todo('should return a single %s by id');
    it.todo('should return null when %s not found');
  });
});
`, cfg.Name, lowerName, cfg.Name, lowerName, cfg.Name, lowerName,
		lowerName,
		cfg.Name, cfg.Name, cfg.Name,
		cfg.Name, cfg.Name,
		cfg.Name,
		lowerName,
		cfg.Name, cfg.Name,
		cfg.Name, cfg.Name, cfg.Name,
		cfg.Name,
		lowerName, lowerName,
	)
}

// GetRepositoryTestTemplate generates unit test template for repositories
func GetRepositoryTestTemplate(cfg EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
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
      const expected%sData: %s[] = [
        new %s(1, /* add constructor params */),
        new %s(2, /* add constructor params */),
      ];

      // Act
      const result = await repository.findAll();

      // Assert
      expect(result).toBeDefined();
      expect(Array.isArray(result)).toBe(true);
    });

    it('should return empty array when no %s found', async () => {
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
`, cfg.Name, lowerName, cfg.Name, lowerName,
		cfg.Name, cfg.Name,
		cfg.Name,
		lowerName,
		cfg.Name, cfg.Name,
		cfg.Name, cfg.Name,
		lowerName,
		lowerName, lowerName,
	)
}
