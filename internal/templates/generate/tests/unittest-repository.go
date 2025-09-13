package tests

import (
	"archi-ts-cli/internal/models"
	"fmt"
	"strings"
)

// GetRepositoryTestTemplate generates unit test template for repositories
func GetRepositoryUnitTestTemplate(cfg models.EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	lowerPlural := lowerName + "s"
	return fmt.Sprintf(`import { %sRepository } from '@repositories/%s.repository';
import %s from '@datamodels/%s.model';

// Mock database/ORM if needed
// jest.mock('@config/database');

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
