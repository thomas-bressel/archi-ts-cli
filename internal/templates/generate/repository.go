package generate

import (
	"archi-ts-cli/internal/models"
	"archi-ts-cli/internal/templates/generate/exports"
	"fmt"
	"strings"
)

// GetRepositoryTemplate gégenerate the repository template
func GetRepositoryTemplate(cfg models.EntityConfig, architecture string) string {
	lowerName := strings.ToLower(cfg.Name)
	upperName := cfg.Name
	ormName := cfg.Orm
	layerImport := ""

	switch architecture {
	case string(models.CleanArchitecture):
		layerImport = exports.GetCleanRepositoryLayerImports(upperName, lowerName)
	case string(models.LayeredArchitecture):
		layerImport = exports.GetLayeredRepositoryLayerImports(upperName, lowerName)
	}

	if ormName == models.TypeOrm {
		return fmt.Sprintf(`
// Layer importations
import { AppDataSource } from "@connection/data-source";
%s

export class %sRepository {
  private repository = AppDataSource.getRepository(%s);
  // private %sQueries: %sQueries(); if you have a query class

  constructor() {
    // Initialize your database connection/ORM here
    // this.%sQueries = new %sQueries(); if you have a query class
  }
  
  async findAll(): Promise<%s[]> {
     return await this.repository.find();
  }
}
`, layerImport,
			cfg.Name, cfg.Name, lowerName, cfg.Name,
			lowerName, cfg.Name,
			cfg.Name)

	}

	return fmt.Sprintf(`
// Layer importations
%s

export class %sRepository {
  // TODO: Add your database connection/ORM here
  // private %sQueries: %sQueries(); if you have a query class

  constructor() {
    // Initialize your database connection/ORM here
    // this.%sQueries = new %sQueries(); if you have a query class
  }
  
  async findAll(): Promise<%s[]> {
    // TODO: Implement database query
    return [];
  }
}
`, layerImport,
		cfg.Name, lowerName, cfg.Name,
		lowerName, cfg.Name,
		cfg.Name)
}
