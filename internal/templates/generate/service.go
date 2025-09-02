package generate

import (
	"fmt"
	"strings"
)

// GetServiceTemplate génère le template du service
func GetServiceTemplate(cfg EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	if cfg.Variant == "typescript" {
		return fmt.Sprintf(`
// Layer imports
import { %sRepository } from '../repositories/%s.repository';
import %s from '../entities/%s.entity';

export class %sService {
  private %sRepository: %sRepository;

  constructor(%sRepository: %sRepository) {
    this.%sRepository = %sRepository;
  }

  async findAll(): Promise<%s[]> {
    return await this.%sRepository.findAll();
  }

  
}
`, cfg.Name, lowerName, cfg.Name, lowerName,
			cfg.Name, lowerName, cfg.Name, lowerName, cfg.Name, lowerName, lowerName,
			cfg.Name, lowerName)
	}

	return fmt.Sprintf(`const %sRepository = require('../repositories/%s.repository');
const %s = require('../entities/%s.entity');

class %sService {
  constructor() {
    this.%sRepository = new %sRepository();
  }

  async findAll() {
    return await this.%sRepository.findAll();
  }

  async findById(id) {
    return await this.%sRepository.findById(id);
  }

  async create(data) {
    // TODO: Add business logic validation here
    const entity = new %s(data);
    return await this.%sRepository.create(entity);
  }

  async update(id, data) {
    // TODO: Add business logic validation here
    const existing = await this.%sRepository.findById(id);
    if (!existing) {
      return null;
    }
    return await this.%sRepository.update(id, data);
  }

  async delete(id) {
    return await this.%sRepository.delete(id);
  }
}

module.exports = %sService;
`, cfg.Name, lowerName, cfg.Name, lowerName, cfg.Name, lowerName, cfg.Name, lowerName, lowerName, cfg.Name, lowerName, lowerName, lowerName, lowerName, cfg.Name)
}
