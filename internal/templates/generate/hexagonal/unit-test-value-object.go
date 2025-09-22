package hexagonal

import (
	"archi-ts-cli/internal/models"
	"fmt"
	"strings"
)

// GetControllerTestTemplate generates unit test template for controllers
func GetValueObjectUnitTestTemplate(cfg models.EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	upperName := cfg.Name

	return fmt.Sprintf(`import { %sID } from '@objects/%s-id.value-object';

describe('%sID Value Object', () => {
    describe('Constructor', () => {
        it('should create a %sID with valid string', () => {
            // Arrange
            const validId = 'valid-%s-id-123';

            // Act
            const %sId = new %sID(validId);

            // Assert
            expect(%sId.value).toBe(validId);
        });

        it('should throw error when value is empty string', () => {
            // Act & Assert
            expect(() => new %sID('')).toThrow('%sID cannot be empty');
        });

        it('should throw error when value is only whitespace', () => {
            // Act & Assert
            expect(() => new %sID('   ')).toThrow('%sID cannot be empty');
        });

        it('should throw error when value is null', () => {
            // Act & Assert
            expect(() => new %sID(null as any)).toThrow('%sID cannot be empty');
        });

        it('should throw error when value is undefined', () => {
            // Act & Assert
            expect(() => new %sID(undefined as any)).toThrow('%sID cannot be empty');
        });
    });

    describe('value getter', () => {
        it('should return the original value', () => {
            // Arrange
            const expectedValue = 'test-%s-id';
            const %sId = new %sID(expectedValue);

            // Act
            const actualValue = %sId.value;

            // Assert
            expect(actualValue).toBe(expectedValue);
        });

        it('should return trimmed value when created with whitespace', () => {
            // Arrange
            const valueWithSpaces = '  valid-id  ';
            const %sId = new %sID(valueWithSpaces);

            // Act
            const actualValue = %sId.value;

            // Assert
            expect(actualValue).toBe(valueWithSpaces);
        });
    });

    describe('generate static method', () => {
        it('should generate a new %sID with UUID format', () => {
            // Act
            const %sId = %sID.generate();

            // Assert
            expect(%sId).toBeInstanceOf(%sID);
            expect(%sId.value).toMatch(/^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$/i);
        });

        it('should generate different UUIDs on multiple calls', () => {
            // Act
            const %sId1 = %sID.generate();
            const %sId2 = %sID.generate();

            // Assert
            expect(%sId1.value).not.toBe(%sId2.value);
        });

        it('should generate valid %sID instances', () => {
            // Act
            const %sId = %sID.generate();

            // Assert
            expect(%sId).toBeInstanceOf(%sID);
            expect(typeof %sId.value).toBe('string');
            expect(%sId.value.length).toBeGreaterThan(0);
        });
    });

    describe('fromString static method', () => {
        it('should create %sID from valid string', () => {
            // Arrange
            const testString = 'custom-%s-id';

            // Act
            const %sId = %sID.fromString(testString);

            // Assert
            expect(%sId).toBeInstanceOf(%sID);
            expect(%sId.value).toBe(testString);
        });

        it('should throw error when creating from empty string', () => {
            // Act & Assert
            expect(() => %sID.fromString('')).toThrow('%sID cannot be empty');
        });

        it('should throw error when creating from whitespace-only string', () => {
            // Act & Assert
            expect(() => %sID.fromString('   ')).toThrow('%sID cannot be empty');
        });

        it('should handle special characters', () => {
            // Arrange
            const specialString = '%s-id_123@domain.com';

            // Act
            const %sId = %sID.fromString(specialString);

            // Assert
            expect(%sId.value).toBe(specialString);
        });
    });

    describe('equals method', () => {
        it('should return true when comparing identical %sIDs', () => {
            // Arrange
            const id = 'same-id';
            const %sId1 = new %sID(id);
            const %sId2 = new %sID(id);

            // Act
            const result = %sId1.equals(%sId2);

            // Assert
            expect(result).toBe(true);
        });

        it('should return false when comparing different %sIDs', () => {
            // Arrange
            const %sId1 = new %sID('id-1');
            const %sId2 = new %sID('id-2');

            // Act
            const result = %sId1.equals(%sId2);

            // Assert
            expect(result).toBe(false);
        });

        it('should be symmetric', () => {
            // Arrange
            const %sId1 = new %sID('test-id');
            const %sId2 = new %sID('test-id');

            // Act & Assert
            expect(%sId1.equals(%sId2)).toBe(%sId2.equals(%sId1));
        });

        it('should be reflexive', () => {
            // Arrange
            const %sId = new %sID('test-id');

            // Act & Assert
            expect(%sId.equals(%sId)).toBe(true);
        });

        it('should handle case sensitivity', () => {
            // Arrange
            const %sId1 = new %sID('Test-ID');
            const %sId2 = new %sID('test-id');

            // Act
            const result = %sId1.equals(%sId2);

            // Assert
            expect(result).toBe(false);
        });
    });

    describe('Edge cases and integration', () => {
        it('should work with generated IDs in equals method', () => {
            // Arrange
            const generatedId = %sID.generate();
            const sameId = %sID.fromString(generatedId.value);

            // Act
            const result = generatedId.equals(sameId);

            // Assert
            expect(result).toBe(true);
        });

        it('should handle very long strings', () => {
            // Arrange
            const longString = 'a'.repeat(1000);

            // Act
            const %sId = new %sID(longString);

            // Assert
            expect(%sId.value).toBe(longString);
            expect(%sId.value.length).toBe(1000);
        });

        it('should handle unicode characters', () => {
            // Arrange
            const unicodeString = '%s-id-ëàü-测试-🚀';

            // Act
            const %sId = new %sID(unicodeString);

            // Assert
            expect(%sId.value).toBe(unicodeString);
        });
    });
});`,
		upperName, lowerName, //import
		upperName,                                             // describe
		upperName, lowerName, lowerName, upperName, lowerName, // it
		upperName, upperName, upperName, upperName, upperName, upperName, upperName, upperName, // it
		lowerName, lowerName, upperName, lowerName, lowerName, upperName, lowerName, // describe
		upperName, lowerName, upperName, lowerName, upperName, lowerName, lowerName, upperName, lowerName, upperName, lowerName, lowerName, upperName, lowerName, upperName, lowerName, upperName, lowerName, lowerName,
		upperName, lowerName, lowerName, upperName, lowerName, upperName, lowerName, upperName, upperName, upperName, upperName, lowerName, lowerName, upperName, lowerName,
		upperName, lowerName, upperName, lowerName, upperName, lowerName, lowerName, upperName, lowerName, upperName, lowerName, upperName, lowerName, lowerName, lowerName, upperName, lowerName, upperName, lowerName, lowerName, lowerName, lowerName, lowerName, upperName, lowerName, lowerName, lowerName, upperName, lowerName, upperName, lowerName, lowerName,
		upperName, upperName, lowerName, upperName, lowerName, lowerName, lowerName, lowerName, upperName, lowerName,
	)
}
