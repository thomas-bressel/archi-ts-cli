package generate

import (
	"archi-ts-cli/internal/models"
	"fmt"
)

// GetValueObjectTemplate generate entity fields
func GetValueObjectTemplate(cfg models.EntityConfig) string {

	return fmt.Sprintf(`import { v4 as uuidv4 } from 'uuid';

export class %sID {
    constructor(private readonly _value: string) {
        if (!_value?.trim()) {
            throw new Error('%sID cannot be empty');
        }
    }

    /** Returns the underlying string value */
    get value(): string { return this._value; }

    /** Generates a new %sID with a fresh UUID */
    static generate(): %sID {
        return new %sID(uuidv4());
    }

    /** Creates a %sID from an existing string */
    static fromString(value: string): %sID {
        return new %sID(value);
    }

    /** Checks if this %sID is equal to another */
    equals(other: %sID): boolean {
        return this._value === other._value;
    }
}`,
		cfg.Name, cfg.Name, cfg.Name, cfg.Name, cfg.Name, cfg.Name,
		cfg.Name, cfg.Name, cfg.Name, cfg.Name,
	)
}
