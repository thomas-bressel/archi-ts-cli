package tests

import (
	"archi-ts-cli/internal/models"
	"archi-ts-cli/internal/templates/generate/exports"
	"fmt"
	"strings"
)

// GetServiceTestTemplate generates unit test template for services
func GetServiceUnitTestTemplate(cfg models.EntityConfig, architecture string) string {
	lowerName := strings.ToLower(cfg.Name)
	lowerPlural := lowerName + "s"
	upperName := cfg.Name
	layerImport := ""

	switch architecture {
	case string(models.CleanArchitecture):
		layerImport = exports.GetTestCleanImports(upperName, lowerName)
	case string(models.LayeredArchitecture):
		layerImport = exports.GetTestLayeredImports(upperName, lowerName)
	}

	return fmt.Sprintf(`// Layers importation
import { %sService } from '@services/%s.service';
import { %sRepository } from '@repositories/%s.repository';
%s

// Mock the repository module
jest.mock('@repositories/%s.repository');

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
		cfg.Name, lowerName, cfg.Name, lowerName, layerImport, // imports
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
