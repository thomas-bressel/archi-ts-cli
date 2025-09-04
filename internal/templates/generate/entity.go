package generate

import (
	"fmt"
	"strings"
)

// getEntityTemplate génère le template de l'entité
func GetEntityTemplate(cfg EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	return fmt.Sprintf(`
export default class %s {
  constructor(
    public readonly id_%s: number | null,
    // Add other fields here
  ){}
}
`, cfg.Name, lowerName)
}
