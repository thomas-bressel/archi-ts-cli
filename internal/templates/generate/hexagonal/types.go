package hexagonal

import (
	"archi-ts-cli/internal/models"
	"fmt"
	"strings"
)

// GetRouteTemplate generates the controller template.
func GetDependenciesTypesTemplate(cfg models.EntityConfig) string {
	upperName := strings.ToTitle(cfg.Name)

	return fmt.Sprintf(`const %sTYPES = {
    %sRepository: Symbol.for('%sRepository'),
    %sController: Symbol.for('%sController'),
    Get%sUseCase: Symbol.for('Get%sUseCase'),
};

export { %sTYPES };`,
		upperName,
		cfg.Name, cfg.Name, cfg.Name, cfg.Name,
		cfg.Name, cfg.Name,
		upperName,
	)
}
