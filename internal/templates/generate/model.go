package generate

import (
	"fmt"
	"strings"
)

// getModelTemplate génère le template du model métier
func GetModelTemplate(cfg EntityConfig) string {
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
