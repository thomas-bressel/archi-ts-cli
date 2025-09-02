package generate

import (
	"fmt"
	"strings"
)

// GetRepositoryTemplate génère le template du repository
func GetRepositoryTemplate(cfg EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	return fmt.Sprintf(`
// Layer importations
import %s from '@entities/%s.entity';

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
