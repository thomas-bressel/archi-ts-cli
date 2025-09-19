package generate

import (
	"archi-ts-cli/internal/models"
	"fmt"
	"strings"
)

// getEntityTemplate génère le template de l'entité
func GetEntityTemplate(cfg models.EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)

	switch cfg.Architecture {
	case models.HexagonalArchitecture:
		return fmt.Sprintf(`import { %sID } from '@objects/%s-id.value-object';

export class %s {
    constructor(
        private readonly _id: %sID
    ) {}

    // Getters
    get id(): %sID { return this._id; }

    // Business methods
    
    // exemple : 
    // updateName(name: %sName): void {
    //     this._name = name;
    // }


    // Factory

    // exemple : 
    // static create(name: string, age: number): %s {
    //     return new %s(
    //         %sID.generate(),
    //         %sName.create(name)
    //     );
    // }

    static reconstitute(id: string): %s {
        return new %s(
            %sID.fromString(id)
        );
    }
}`,
			cfg.Name, lowerName, // Import
			cfg.Name, cfg.Name, // constructor
			cfg.Name, cfg.Name, cfg.Name, cfg.Name, cfg.Name, cfg.Name, cfg.Name, cfg.Name, cfg.Name,
		)
	default:
		return fmt.Sprintf(`export default class %s {
  constructor(
    public readonly id_%s: number | null,
    // Add other fields here
  ){}
}`, cfg.Name, lowerName)
	}

}
