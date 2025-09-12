package generate

import (
	"fmt"
	"strings"
)

// getRawModelTemplate génère le template du model de base de données
func GetRawModelTemplate(cfg EntityConfig) string {
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
