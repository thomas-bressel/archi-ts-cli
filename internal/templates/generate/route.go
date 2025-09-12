package generate

import (
	"archi-ts-cli/internal/models"
	"archi-ts-cli/internal/templates/generate/exports"
	"fmt"
	"strings"
)

// GetRouteTemplate generate route template
func GetRouteTemplate(cfg models.EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	libraryImport := ""
	expressExport := ""
	expressRouter := ""
	expressRouteExemple := ""
	// is Express used ?
	if cfg.Express {
		libraryImport = exports.ExpressImportation()
		expressRouter = exports.ExpressRouter()
		expressRouteExemple = exports.ExpressRouteExemple(lowerName)
		expressExport = exports.ExpressExportation()
	} else {
		libraryImport = exports.NodeImportation()
		expressRouteExemple = exports.NodeRouteExemple(lowerName, cfg.Name)
	}

	return fmt.Sprintf(`	
%s

// Layers importation
import { %sController }  from '@controllers/%s.controller'; 
import { %sService } from '@services/%s.service';
import { %sRepository } from "@repositories/%s.repository";

%s

// Dependencies injections
const %sRepository = new %sRepository()
const %sService = new %sService(%sRepository)
const %sController = new %sController(%sService);

%s

%s
`,
		// Header
		libraryImport,
		// Arguments for fmt.Sprintf, in order of appearance of %s
		cfg.Name, lowerName, // Controller import
		cfg.Name, lowerName, // Service import
		cfg.Name, lowerName, // Repository import
		expressRouter,
		lowerName, cfg.Name, // Repository variable declaration
		lowerName, cfg.Name, lowerName, // Service variable declaration with dependency injection
		lowerName, cfg.Name, lowerName, // Controller variable declaration with dependency injection
		expressRouteExemple, expressExport,
	)
}
