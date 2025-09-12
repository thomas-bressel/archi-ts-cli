package generate

import (
	"archi-ts-cli/internal/models"
	"fmt"
	"strings"
)

// getRawModelTemplate génère le template du model de base de données
func GetRawModelTemplate(cfg models.EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	return fmt.Sprintf(`
export default interface %s {
    id_%s: number | null;
    // Add other database entity fields here
}
`, cfg.Name, lowerName)
}
