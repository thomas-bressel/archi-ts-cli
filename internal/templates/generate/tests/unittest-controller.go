package tests

import (
	"archi-ts-cli/internal/models"
	"archi-ts-cli/internal/templates/generate/exports"
	"fmt"
	"strings"
)

// GetControllerTestTemplate generates unit test template for controllers
func GetControllerUnitTestTemplate(cfg models.EntityConfig) string {

	lowerName := strings.ToLower(cfg.Name)
	libraryImport := ""
	testReqResMock := ""
	testReqResMockSetup := ""
	testAct := ""
	testAssert200 := ""
	testAssert500 := ""

	// is Express used ?
	if cfg.Express {
		libraryImport = exports.ExpressImportation()
		testReqResMock = exports.ExpressReqResMock()
		testReqResMockSetup = exports.ExpressReqResMockSetup()
		testAct = exports.ExpressActController()
		testAssert200 = exports.ExpressAssertController(200)
		testAssert500 = exports.ExpressAssertController(500)
	} else {
		libraryImport = exports.NodeImportation()
		testReqResMock = exports.NodeReqResMock()
		testReqResMockSetup = exports.NodeReqResMockSetup()
		testAct = exports.NodeActController()
		testAssert200 = exports.NodeAssertController(200)
		testAssert500 = exports.NodeAssertController(500)
	}

	return fmt.Sprintf(`%s


// Layers importation
import { %sController } from '@controllers/%s.controller';
import { %sService } from '@services/%s.service';
import { %sRepository } from '@repositories/%s.repository'; // Note: Only needed if you instantiate the service, which is not ideal for a pure mock.

// Mock the service module
jest.mock('@services/%s.service');

describe('%sController', () => {
  let controller: %sController;
  let mockService: jest.Mocked<%sService>;

%s

  beforeEach(() => {
    // Clear all mocks before each test
    jest.clearAllMocks();

    // Create mock service instance. Since the module is mocked, 'new' will return a mocked constructor.
    mockService = new %sService(new %sRepository()) as jest.Mocked<%sService>;
    
    // Initialize controller with mock service (dependency injection)
    controller = new %sController(mockService);

%s

  });

  describe('getAll', () => {
    it('should return all %ss successfully with status 200', async () => {
      // Arrange
      const mock%sData = [
        { id_%s: 1 },
        { id_%s: 2 },
      ];
      mockService.findAll.mockResolvedValue(mock%sData);

      // Act
     %s

      // Assert
      expect(mockService.findAll).toHaveBeenCalledTimes(1);
     %s
      
    });

    it('should return 200 with empty array when service returns empty array', async () => {
      // Arrange
      mockService.findAll.mockResolvedValue([]);

      // Act
      %s

      // Assert
      expect(mockService.findAll).toHaveBeenCalledTimes(1);
      %s
    });

    it('should handle errors when fetching %s fails', async () => {
      // Arrange
      const errorMessage = 'Database connection error';
      mockService.findAll.mockRejectedValue(new Error(errorMessage));

      // Act
      %s

      // Assert
      expect(mockService.findAll).toHaveBeenCalledTimes(1);
      %s
  });

    it('should return "Internal server error" for non-Error exceptions', async () => {
      // Arrange
      mockService.findAll.mockRejectedValue('Unknown error');

      // Act
      %s
      
      // Assert
      expect(mockService.findAll).toHaveBeenCalledTimes(1);
      %s
  });
  });
});`, libraryImport,
		cfg.Name, lowerName, cfg.Name, lowerName, cfg.Name, lowerName,
		lowerName, cfg.Name, cfg.Name, cfg.Name,
		testReqResMock,
		cfg.Name, cfg.Name, cfg.Name, cfg.Name,
		testReqResMockSetup,
		lowerName, cfg.Name, lowerName, lowerName, cfg.Name,
		testAct, testAssert200,
		testAct, testAssert200,
		lowerName, testAct, testAssert500,
		testAct, testAssert500,
	)
}
