package hexagonal

import (
	"archi-ts-cli/internal/models"
	"fmt"
	"strings"
)

// GetControllerTestTemplate generates unit test template for controllers
func GetUseCaseUnitTestTemplate(cfg models.EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	upperName := cfg.Name
	titleName := strings.ToTitle(cfg.Name)

	return fmt.Sprintf(`// Layer imports 
import { Get%sUseCase } from '@usecases/get-%s.use-case';
import { %sRepository } from '@ports/out/%s.repository';
import { %sResponse } from '@dtos/%s.dto';
import { %sID } from '@objects/%s-id.value-object';
import { %sNotFound } from '@exceptions/%s-not-found.exception';
import { %s } from '@entities/%s.entity';

// Mock %sTYPES module
jest.mock('@config/dependencies/%s-types', () => ({
    %sTYPES: {
        %sRepository: Symbol.for('%sRepository'),
    },
}));

// Repository mock
const mock%sRepository = {
    findById: jest.fn(),
    save: jest.fn(),
    delete: jest.fn(),
    findAll: jest.fn(),
} as jest.Mocked<%sRepository>;

// Create mock %s using real entity factory method
const createMock%s = (id: string): %s => {
    return %s.reconstitute(id);
};

const mock%s = createMock%s('test-%s-id-123');

describe('Get%sUseCase', () => {
    let get%sUseCase: Get%sUseCase;

    beforeEach(() => {
        // Reset all mocks before each test
        jest.clearAllMocks();
        
        // Create new use case instance with mocked repository
        get%sUseCase = new Get%sUseCase(mock%sRepository);
    });

    describe('execute', () => {
        describe('when %s exists', () => {
            it('should return %sResponse when %s is found', async () => {
                // Arrange
                const %sId = 'test-%s-id-123';
                const expected%sId = %sID.fromString(%sId);
                
                mock%sRepository.findById.mockResolvedValue(mock%s);

                // Act
                const result = await get%sUseCase.execute(%sId);

                // Assert
                expect(mock%sRepository.findById).toHaveBeenCalledTimes(1);
                expect(mock%sRepository.findById).toHaveBeenCalledWith(expected%sId);
                
                expect(result).toEqual<%sResponse>({
                    id: %sId,
                });
            });

            it('should call repository with correct %sID value object', async () => {
                // Arrange
                const %sId = 'another-%s-id';
                mock%sRepository.findById.mockResolvedValue(mock%s);

                // Act
                await get%sUseCase.execute(%sId);

                // Assert
                const expected%sIdCall = mock%sRepository.findById.mock.calls[0][0];
                expect(expected%sIdCall).toBeInstanceOf(%sID);
                expect(expected%sIdCall.value).toBe(%sId);
            });

            it('should handle special characters in %s ID', async () => {
                // Arrange
                const %sId = '%s-id_with@special.chars-123';
                const %sWithSpecialId = createMock%s(%sId);
                
                mock%sRepository.findById.mockResolvedValue(%sWithSpecialId);

                // Act
                const result = await get%sUseCase.execute(%sId);

                // Assert
                expect(result.id).toBe(%sId);
            });

            it('should handle UUID format %s ID', async () => {
                // Arrange
                const uuid%sId = '123e4567-e89b-12d3-a456-426614174000';
                const %sWithUuid = createMock%s(uuid%sId);
                
                mock%sRepository.findById.mockResolvedValue(%sWithUuid);

                // Act
                const result = await get%sUseCase.execute(uuid%sId);

                // Assert
                expect(result.id).toBe(uuid%sId);
                expect(mock%sRepository.findById).toHaveBeenCalledWith(
                    expect.objectContaining({ value: uuid%sId })
                );
            });
        });

        describe('when %s does not exist', () => {
            it('should throw %sNotFound exception when %s is not found', async () => {
                // Arrange
                const nonExistent%sId = 'non-existent-%s-id';
                mock%sRepository.findById.mockResolvedValue(null);

                // Act & Assert
                await expect(
                    get%sUseCase.execute(nonExistent%sId)
                ).rejects.toThrow(%sNotFound);

                expect(mock%sRepository.findById).toHaveBeenCalledTimes(1);
                expect(mock%sRepository.findById).toHaveBeenCalledWith(
                    expect.objectContaining({ value: nonExistent%sId })
                );
            });

            it('should throw %sNotFound with correct %s ID in exception', async () => {
                // Arrange
                const nonExistent%sId = 'missing-%s-123';
                mock%sRepository.findById.mockResolvedValue(null);

                // Act & Assert
                try {
                    await get%sUseCase.execute(nonExistent%sId);
                    fail('Should have thrown %sNotFound exception');
                } catch (error: unknown) {
                    expect(error).toBeInstanceOf(%sNotFound);
                    if (error instanceof %sNotFound) {
                        // Exception uses single quotes, so no string interpolation occurs
                        expect(error.message).toBe('%s with ID ${%sId} not found');
                    }
                }
            });

            it('should throw %sNotFound when repository returns null', async () => {
                // Arrange
                const %sId = 'null-%s';
                mock%sRepository.findById.mockResolvedValue(null);

                // Act & Assert
                await expect(
                    get%sUseCase.execute(%sId)
                ).rejects.toThrow(%sNotFound);
            });
        });

        describe('input validation', () => {
            it('should handle empty string %s ID', async () => {
                // Arrange
                const empty%sId = '';

                // Act & Assert
                // %sID.fromString should throw error for empty strings
                await expect(
                    get%sUseCase.execute(empty%sId)
                ).rejects.toThrow();
            });

            it('should handle whitespace-only %s ID', async () => {
                // Arrange
                const whitespace%sId = '   ';

                // Act & Assert
                await expect(
                    get%sUseCase.execute(whitespace%sId)
                ).rejects.toThrow();
            });

            it('should handle very long %s ID', async () => {
                // Arrange
                const long%sId = 'a'.repeat(1000);
                const %sWithLongId = createMock%s(long%sId);
                
                mock%sRepository.findById.mockResolvedValue(%sWithLongId);

                // Act
                const result = await get%sUseCase.execute(long%sId);

                // Assert
                expect(result.id).toBe(long%sId);
            });
        });

        describe('repository error handling', () => {
            it('should propagate repository errors', async () => {
                // Arrange
                const %sId = 'error-%s-id';
                const repositoryError = new Error('Database connection failed');
                
                mock%sRepository.findById.mockRejectedValue(repositoryError);

                // Act & Assert
                await expect(
                    get%sUseCase.execute(%sId)
                ).rejects.toThrow('Database connection failed');

                expect(mock%sRepository.findById).toHaveBeenCalledTimes(1);
            });

            it('should handle repository timeout', async () => {
                // Arrange
                const %sId = 'timeout-%s-id';
                const timeoutError = new Error('Query timeout');
                
                mock%sRepository.findById.mockRejectedValue(timeoutError);

                // Act & Assert
                await expect(
                    get%sUseCase.execute(%sId)
                ).rejects.toThrow('Query timeout');
            });
        });

        describe('integration with %sID value object', () => {
            it('should create %sID with same value as input', async () => {
                // Arrange
                const %sId = 'integration-test-id';
                mock%sRepository.findById.mockResolvedValue(mock%s);

                // Act
                await get%sUseCase.execute(%sId);

                // Assert
                const calledWith = mock%sRepository.findById.mock.calls[0][0];
                expect(calledWith.value).toBe(%sId);
                expect(calledWith).toBeInstanceOf(%sID);
            });

            it('should handle %sID equality correctly', async () => {
                // Arrange
                const %sId = 'equality-test-id';
                const expected%sId = %sID.fromString(%sId);
                
                mock%sRepository.findById.mockResolvedValue(mock%s);

                // Act
                await get%sUseCase.execute(%sId);

                // Assert
                const actual%sId = mock%sRepository.findById.mock.calls[0][0];
                expect(actual%sId.equals(expected%sId)).toBe(true);
            });
        });
    });

    describe('dependency injection', () => {
        it('should be instantiable with mocked repository', () => {
            // Act
            const useCase = new Get%sUseCase(mock%sRepository);

            // Assert
            expect(useCase).toBeInstanceOf(Get%sUseCase);
        });
    });
});`,
		upperName, lowerName, upperName, lowerName, upperName, lowerName, upperName, lowerName, upperName, lowerName, upperName, lowerName, // Imports
		titleName, lowerName, titleName, upperName, upperName, upperName, upperName, lowerName, upperName, upperName, upperName, upperName, upperName, lowerName, // Mock
		upperName, upperName, upperName, upperName, upperName, upperName, // describe
		lowerName, upperName, lowerName, lowerName, lowerName, upperName, upperName, lowerName, upperName, upperName, upperName, lowerName, upperName, upperName, upperName, upperName, lowerName,
		upperName, lowerName, lowerName, upperName, upperName, upperName, lowerName, upperName, upperName, upperName, upperName, upperName, lowerName,
		lowerName, lowerName, lowerName, lowerName, upperName, lowerName, upperName, lowerName, upperName, lowerName, lowerName,
		lowerName, upperName, lowerName, upperName, upperName, upperName, lowerName, upperName, upperName, upperName, upperName, upperName,
		lowerName, upperName, lowerName, upperName, lowerName, upperName, upperName, upperName, upperName, upperName, upperName, upperName,
		upperName, lowerName, upperName, lowerName, upperName, upperName, upperName, upperName, upperName, upperName, upperName, lowerName,
		upperName, lowerName, lowerName, upperName, upperName, lowerName, upperName,
		lowerName, upperName, upperName, upperName, upperName, lowerName, upperName, upperName, upperName, lowerName, upperName, lowerName, upperName, upperName, upperName, lowerName, upperName, upperName, upperName,
		lowerName, lowerName, upperName, upperName, lowerName, upperName, lowerName, lowerName, upperName, upperName, lowerName,
		upperName, upperName, lowerName, upperName, upperName, upperName, lowerName, upperName, lowerName, upperName, upperName, lowerName, upperName, upperName, lowerName, upperName, upperName, upperName, lowerName, upperName, upperName, upperName, upperName,
		upperName, upperName, upperName,
	)
}
