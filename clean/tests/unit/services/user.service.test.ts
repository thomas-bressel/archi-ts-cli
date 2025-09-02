import { UserService } from '@services/user.service';
import { UserRepository } from '@repositories/user.repository';
import User from '@entities/user.entity';

// Mock the repository module
jest.mock('@repositories/user.repository');

describe('UserService', () => {
  let service: UserService;
  let mockRepository: jest.Mocked<UserRepository>;

  beforeEach(() => {
    // Clear all mocks before each test
    jest.clearAllMocks();

    // Create mock repository instance
    mockRepository = new UserRepository() as jest.Mocked<UserRepository>;
    
    // Initialize service with mock repository
    service = new UserService(mockRepository);
  });

  describe('findAll', () => {
    it('should return all users from repository', async () => {
      // Arrange
      const mockUserData: User[] = [
        new User(1, /* add constructor params */),
        new User(2, /* add constructor params */),
      ];
      mockRepository.findAll.mockResolvedValue(mockUserData);

      // Act
      const result = await service.findAll();

      // Assert
      expect(mockRepository.findAll).toHaveBeenCalledTimes(1);
      expect(result).toEqual(mockUserData);
      expect(result).toHaveLength(2);
    });

    it('should return empty array when repository returns no users', async () => {
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
      await expect(service.findAll()).rejects.toThrow('Failed to fetch users: Database error');
      expect(mockRepository.findAll).toHaveBeenCalledTimes(1);
    });
  });

  // Add more test suites for other methods
  describe('findById', () => {
    it.todo('should return a single user by id');
    it.todo('should return null when user not found');
  });
});
