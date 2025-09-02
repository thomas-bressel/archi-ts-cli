package generate

import (
	"fmt"
	"strings"
)

// getEntityTemplate génère le template de l'entité
func GetEntityTemplate(cfg EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	if cfg.Variant == "typescript" {
		return fmt.Sprintf(`
  export default class %s {
    constructor(
    public readonly id_%s: number | null,
    // Add other fields here
    ){}
  }
`, cfg.Name, lowerName)
	}

	return fmt.Sprintf(`class %s {
  constructor(data = {}) {
    this.id = data.id || 0;
    this.name = data.name || '';
    this.description = data.description || '';
    this.isActive = data.isActive !== undefined ? data.isActive : true;
    this.createdAt = data.createdAt || new Date();
    this.updatedAt = data.updatedAt || new Date();
  }

  toJSON() {
    return {
      id: this.id,
      name: this.name,
      description: this.description,
      isActive: this.isActive,
      createdAt: this.createdAt,
      updatedAt: this.updatedAt,
    };
  }
}

module.exports = %s;
`, cfg.Name, cfg.Name)
}
