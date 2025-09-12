package generate

import (
	"archi-ts-cli/internal/models"
	"fmt"
)

// getModelTemplate génère le template du model métier
func GetModelTemplate(cfg models.EntityConfig) string {
	return fmt.Sprintf(`
export interface %s {
  // Add business model properties here
}
`, cfg.Name)
}
