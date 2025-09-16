package generate

import (
	"archi-ts-cli/internal/models"
	"archi-ts-cli/internal/templates/generate/exports"
	"fmt"
	"strings"
)

// GetServiceTemplate generate service template
func GetServiceTemplate(cfg models.EntityConfig, architecture string) string {
	lowerName := strings.ToLower(cfg.Name)
	upperName := cfg.Name
	layerImport := ""

	switch architecture {
	case string(models.CleanArchitecture):
		layerImport = exports.GetCleanServiceLayerImports(upperName, lowerName)
	case string(models.LayeredArchitecture):
		layerImport = exports.GetLayeredServiceLayerImports(upperName, lowerName)
	}

	return fmt.Sprintf(`
%s

export class %sService {
  private %sRepository: %sRepository;

  constructor(%sRepository: %sRepository) {
    this.%sRepository = %sRepository;
  }

  /**
   * Get the list of all %s
   * * @returns Promise<%s[]> - Object response from the database with raw entity data
   */
  public async findAll(): Promise<%s[]> {
    try {
      return await this.%sRepository.findAll();
    } catch (error) {
      throw new Error("Failed to fetch %s: " + (error instanceof Error ? error.message : String(error)));
    }
  }
}
`,
		// Arguments for fmt.Sprintf, in order of appearance of %s
		layerImport,
		cfg.Name, lowerName, cfg.Name, // class name and private repository declaration
		lowerName, cfg.Name, lowerName, lowerName, // constructor
		lowerName+"s", // JSDoc comment
		cfg.Name,      // function return type
		cfg.Name,      // function return type
		lowerName,     // this.repository call
		lowerName+"s", // error message
	)
}
