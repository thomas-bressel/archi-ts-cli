package hexagonal

import (
	"archi-ts-cli/internal/models"
	"fmt"
)

// GetDtoTemplate gégenerate the repository template
func GetDtoTemplate(cfg models.EntityConfig) string {
	upperName := cfg.Name

	return fmt.Sprintf(`export interface %sResponse {
    id: string;
}`, upperName)
}
