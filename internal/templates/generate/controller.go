package generate

import (
	"archi-ts-cli/internal/models"
	"archi-ts-cli/internal/templates/generate/exports"
	"fmt"
	"strings"
)

// GetControllerTemplate generates the controller template.
func GetControllerTemplate(cfg models.EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	libraryImport := ""
	controller := ""

	// is Express used ?
	if cfg.Express {
		libraryImport = exports.ExpressImportation()
		controller = exports.ExpressController(lowerName)
	} else {
		libraryImport = exports.NodeImportation()
		controller = exports.NodeController(lowerName, cfg.Name)
	}

	return fmt.Sprintf(`
%s

// Services importation
import { %sService } from '@services/%s.service';

// Libraries importation
import { validate } from 'class-validator';
import { plainToInstance } from 'class-transformer';

export class %sController {
  private %sService: %sService;

  constructor(%sService: %sService) {
    this.%sService = %sService;
  }

%s
}`,
		libraryImport,
		// Arguments for fmt.Sprintf, in order of appearance of %s
		cfg.Name, lowerName, // Service import and path
		cfg.Name,            // Controller class name
		lowerName, cfg.Name, // private service declaration
		lowerName, cfg.Name, // constructor arguments
		lowerName, lowerName, // constructor body
		controller,
	)
}
