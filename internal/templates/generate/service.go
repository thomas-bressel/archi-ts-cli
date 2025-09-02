package generate

import (
	"fmt"
	"strings"
)

// GetServiceTemplate génère le template du service
func GetServiceTemplate(cfg EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
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
