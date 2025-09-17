package generate

import (
	"archi-ts-cli/internal/models"
	"fmt"
	"strings"
)

// getModelTemplate génère le template du model métier
func GetModelTemplate(cfg models.EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	return fmt.Sprintf(`
export class %s {
	constructor(
		public readonly id_%s: number
		// Add business model properties here
	){}
}
`, cfg.Name, lowerName)
}
