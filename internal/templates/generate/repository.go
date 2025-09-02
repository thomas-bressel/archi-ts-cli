package generate

import (
	"fmt"
	"strings"
)

// GetRepositoryTemplate génère le template du repository
func GetRepositoryTemplate(cfg EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	if cfg.Variant == "typescript" {
		return fmt.Sprintf(`
  // Layer importations
  import  %s  from '../entities/%s.entity';

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

	return fmt.Sprintf(`const %sQuery = require('../database/queries/%s.query');
const %s = require('../entities/%s.entity');

class %sRepository {
  constructor() {
    this.%sQuery = new %sQuery();
  }

  async findAll() {
    // TODO: Implement database query
    const results = await this.%sQuery.selectAll();
    return results.map(row => new %s(row));
  }

  async findById(id) {
    // TODO: Implement database query
    const result = await this.%sQuery.selectById(id);
    return result ? new %s(result) : null;
  }

  async create(data) {
    // TODO: Implement database query
    const result = await this.%sQuery.insert(data);
    return new %s(result);
  }

  async update(id, data) {
    // TODO: Implement database query
    const result = await this.%sQuery.update(id, data);
    return result ? new %s(result) : null;
  }

  async delete(id) {
    // TODO: Implement database query
    return await this.%sQuery.delete(id);
  }
}

module.exports = %sRepository;
`, cfg.Name, lowerName, cfg.Name, lowerName, cfg.Name, lowerName, cfg.Name, lowerName, cfg.Name, lowerName, cfg.Name, lowerName, cfg.Name, lowerName, cfg.Name, lowerName, cfg.Name)
}
