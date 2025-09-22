package tests

import (
	"archi-ts-cli/internal/models"
	"archi-ts-cli/internal/templates/generate/exports"
	"fmt"
	"strings"
)

// GetRepositoryTestTemplate generates unit test template for repositories
func GetRepositoryUnitTestTemplate(cfg models.EntityConfig, architecture string) string {
	lowerName := strings.ToLower(cfg.Name)
	lowerPlural := lowerName + "s"

	upperName := cfg.Name
	layerImport := ""
	aliasPath := ""

	switch architecture {
	case string(models.CleanArchitecture):
		layerImport = exports.GetTestCleanImports(upperName, lowerName)
		aliasPath = "@config"
	case string(models.LayeredArchitecture):
		layerImport = exports.GetTestLayeredImports(upperName, lowerName)
		aliasPath = "@connection"
	}

	if cfg.Orm == models.TypeOrm {
		return fmt.Sprintf(`
import { %sRepository } from '@repositories/%s.repository';
%s
import { AppDataSource } from '%s/data-source';
import { Repository } from 'typeorm';

// Mock the data source module
jest.mock('%s/data-source', () => ({
  AppDataSource: {
    getRepository: jest.fn()
  }
}));

describe('%sRepository', () => {
  let %sRepository: %sRepository;
  let mockTypeOrmRepository: jest.Mocked<Repository<%s>>;

  beforeEach(() => {
    // Clear all mocks before each test
    jest.clearAllMocks();
    
    // Create a mock TypeORM repository with all necessary methods
    mockTypeOrmRepository = {
      find: jest.fn(),
      findOne: jest.fn(),
      save: jest.fn(),
      remove: jest.fn(),
      delete: jest.fn(),
      create: jest.fn(),
      update: jest.fn(),
      count: jest.fn(),
      findAndCount: jest.fn(),
      // Add other Repository methods as needed
    } as unknown as jest.Mocked<Repository<%s>>;

    // Mock getRepository to return our mock repository
    (AppDataSource.getRepository as jest.Mock).mockReturnValue(mockTypeOrmRepository);
    
    // Initialize repository - this will now use the mocked AppDataSource
    %sRepository = new %sRepository();
  });

  afterEach(() => {
    jest.restoreAllMocks();
  });

  describe('findAll', () => {
    it('should return all %ss from database', async () => {
      // Arrange
      const mock%ss: %s[] = [
        Object.assign(new %s(1), { 
          id_%s: 1
          // Add other properties as needed based on your %s model
        }),
        Object.assign(new %s(2), { 
          id_%s: 2
          // Add other properties as needed
        }),
        Object.assign(new %s(3), { 
          id_%s: 3
          // Add other properties as needed
        })
      ];

      // Mock the find method to return our test data
      mockTypeOrmRepository.find.mockResolvedValue(mock%ss);

      // Act
      const result = await %sRepository.findAll();

      // Assert
      expect(result).toBeDefined();
      expect(Array.isArray(result)).toBe(true);
      expect(result).toEqual(mock%ss);
      expect(result).toHaveLength(3);
      expect(mockTypeOrmRepository.find).toHaveBeenCalledTimes(1);
      expect(mockTypeOrmRepository.find).toHaveBeenCalledWith();
    });

    it('should return empty array when no %ss found', async () => {
      // Arrange
      mockTypeOrmRepository.find.mockResolvedValue([]);

      // Act
      const result = await %sRepository.findAll();

      // Assert
      expect(result).toEqual([]);
      expect(result).toHaveLength(0);
      expect(mockTypeOrmRepository.find).toHaveBeenCalledTimes(1);
    });

    it('should handle database errors properly', async () => {
      // Arrange
      const dbError = new Error('Database connection failed');
      mockTypeOrmRepository.find.mockRejectedValue(dbError);

      // Act & Assert
      await expect(%sRepository.findAll()).rejects.toThrow('Database connection failed');
      expect(mockTypeOrmRepository.find).toHaveBeenCalledTimes(1);
    });

    it('should handle null response from database', async () => {
      // Arrange
      mockTypeOrmRepository.find.mockResolvedValue(null as any);

      // Act
      const result = await %sRepository.findAll();

      // Assert
      expect(result).toBeNull();
      expect(mockTypeOrmRepository.find).toHaveBeenCalledTimes(1);
    });

    it('should call repository.find with no parameters', async () => {
      // Arrange
      mockTypeOrmRepository.find.mockResolvedValue([]);

      // Act
      await %sRepository.findAll();

      // Assert
      expect(mockTypeOrmRepository.find).toHaveBeenCalledWith();
      // Verify no parameters were passed
      expect(mockTypeOrmRepository.find).not.toHaveBeenCalledWith(expect.anything());
    });
  });

  describe('Repository Initialization', () => {
    it('should get %s repository from AppDataSource', () => {
      // Assert
      expect(AppDataSource.getRepository).toHaveBeenCalledTimes(1);
    });

    it('should create repository instance successfully', () => {
      // Assert
      expect(%sRepository).toBeDefined();
      expect(%sRepository).toBeInstanceOf(%sRepository);
    });
  });
});`, cfg.Name, lowerName, layerImport, aliasPath, aliasPath,
			cfg.Name, lowerName, cfg.Name, cfg.Name, cfg.Name, lowerName, cfg.Name,
			lowerName, cfg.Name, cfg.Name, cfg.Name, lowerName, cfg.Name, cfg.Name, lowerName, cfg.Name, lowerName, cfg.Name, lowerName, cfg.Name,
			lowerName, lowerName, lowerName, lowerName, lowerName,
			cfg.Name, lowerName, lowerName, cfg.Name,
		)
	}

	return fmt.Sprintf(`// Layer imports
import { %sRepository } from '@repositories/%s.repository';

%s

describe('%sRepository', () => {
  let repository: %sRepository;

  beforeEach(() => {    
    // Initialize repository
    repository = new %sRepository();
  });

  describe('findAll', () => {
    it('should return all %s from database', async () => {
      // Arrange

      // Act

      // Assert
    });

    it('should return empty array when no %s found', async () => {
      // Arrange

      // Act

      // Assert
    });
  });
});
`,
		// Arguments for fmt.Sprintf
		cfg.Name, lowerName, layerImport, // imports
		cfg.Name,    // describe
		cfg.Name,    // let declaration
		cfg.Name,    // new Repository
		lowerPlural, // it should return all...
		lowerPlural, // it should return empty...
	)
}
