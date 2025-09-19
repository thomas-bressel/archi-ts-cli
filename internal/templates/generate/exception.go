package generate

import (
	"archi-ts-cli/internal/models"
	"fmt"
	"strings"
)

// GetExceptionTemplate génère le template de l'entité
func GetExceptionTemplate(cfg models.EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)

	return fmt.Sprintf(`export class %sNotFound extends Error {
    constructor(%sId: string) {
        super('%s with ID ${%sId} not found');
        this.name = '%sNotFound';
    }
}`,
		cfg.Name,
		lowerName, // constructor
		cfg.Name, lowerName, cfg.Name,
	)
}
