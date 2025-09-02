// Express importation
import { Request, Response } from 'express';

// Layers importation
import { UserController } from '@controllers/user.controller';
import { UserService } from '@services/user.service';
import { UserRepository } from '@repositories/user.repository'; // Note: Only needed if you instantiate the service, which is not ideal for a pure mock.

// Mock the service module
jest.mock('@services/user.service');

describe('UserController', () => {
  let controller: UserController;
  let mockService: jest.Mocked<UserService>;
  let mockRequest: Partial<Request>;
  let mockResponse: Partial<Response>;

  beforeEach(() => {
    // Clear all mocks before each test
    jest.clearAllMocks();

    // Create mock service instance. Since the module is mocked, 'new' will return a mocked constructor.
    mockService = new UserService(new UserRepository()) as jest.Mocked<UserService>;
    
    // Initialize controller with mock service (dependency injection)
    controller = new UserController(mockService);

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
    it('should return all users successfully with status 200', async () => {
      // Arrange
      const mockUserData = [
        { id_user: 1 },
        { id_user: 2 },
      ];
      mockService.findAll.mockResolvedValue(mockUserData);

      // Act
      await controller.getAll(mockRequest as Request, mockResponse as Response);

      // Assert
      expect(mockService.findAll).toHaveBeenCalledTimes(1);
      expect(mockResponse.status).toHaveBeenCalledWith(200);
      expect(mockResponse.json).toHaveBeenCalledWith({
        success: true,
        data: mockUserData,
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

    it('should handle errors when fetching users fails', async () => {
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
