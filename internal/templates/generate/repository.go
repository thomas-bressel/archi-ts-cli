package generate

import (
	"archi-ts-cli/internal/models"
	"fmt"
	"strings"
)

// GetRepositoryTemplate gégenerate the repository template
func GetRepositoryTemplate(cfg models.EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	ormName := cfg.Orm

	if ormName == models.TypeOrm {
		return fmt.Sprintf(`
// Layer importations
import { AppDataSource } from "@connection/data-source";
import %s from '@datamodels/%s.model';

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
`, cfg.Name, lowerName,
			cfg.Name, cfg.Name, lowerName, cfg.Name,
			lowerName, cfg.Name,
			cfg.Name)

	}

	return fmt.Sprintf(`
// Layer importations
import %s from '@datamodels/%s.model';

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
`, cfg.Name, lowerName,
		cfg.Name, lowerName, cfg.Name,
		lowerName, cfg.Name,
		cfg.Name)
}
