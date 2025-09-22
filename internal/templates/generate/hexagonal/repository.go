package hexagonal

import (
	"archi-ts-cli/internal/models"
	"fmt"
	"strings"
)

// GetRepositoryTemplate gégenerate the repository template
func GetRepositoryTemplate(cfg models.EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	upperName := cfg.Name

	return fmt.Sprintf(`import { %s } from '@entities/%s.entity';
import { %sID } from '@objects/%s-id.value-object';

export interface %sRepository {
    save(%s: %s): Promise<void>;
    findById(id: %sID): Promise<%s | null>;
    // findAll(): Promise<%s[]>;
    // update(%s: %s): Promise<void>;
    // delete(id: %sID): Promise<void>;
    // exists(id: %sID): Promise<boolean>;
}`,
		upperName, lowerName, upperName, lowerName,
		upperName, lowerName, upperName,
		upperName, upperName, upperName,
		lowerName, upperName, upperName, upperName,
	)

}
