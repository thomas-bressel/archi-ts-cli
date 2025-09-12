package generate

import (
	"archi-ts-cli/internal/models"
	"fmt"
	"strings"
)

// GetServiceTemplate generate service template
func GetServiceTemplate(cfg models.EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	return fmt.Sprintf(`
// Layer imports
import { %sRepository } from '@repositories/%s.repository';
import { %s } from '@businessmodels/%s.model';

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
		cfg.Name, lowerName, // imports
		cfg.Name, lowerName, // imports
		cfg.Name, lowerName, cfg.Name, // class name and private repository declaration
		lowerName, cfg.Name, lowerName, lowerName, // constructor
		lowerName+"s", // JSDoc comment
		cfg.Name,      // function return type
		cfg.Name,      // function return type
		lowerName,     // this.repository call
		lowerName+"s", // error message
	)
}
