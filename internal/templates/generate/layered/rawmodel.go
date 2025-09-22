package layered

import (
	"archi-ts-cli/internal/models"
	"fmt"
	"strings"
)

// getRawModelTemplate génère le template du model de base de données
func GetRawModelTemplate(cfg models.EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	ormName := cfg.Orm

	if ormName == "typeorm" {
		return fmt.Sprintf(`
import { Entity, PrimaryGeneratedColumn, Column } from "typeorm";

@Entity('%s') // Table name
export default class %s {

	constructor(id_%s?: number) {
		this.id_%s = 0;
	}
    @PrimaryGeneratedColumn()
    id_%s: number;
    

	// Exemples : 
    // @Column({ type: 'varchar', length: 255 })
    // name: string;
    
    // @Column({ type: 'decimal', precision: 10, scale: 2 })
    // price: number;
    
    // @Column({ type: 'int', default: 0 })
    // stock: number;
    
    // @Column({ type: 'timestamp', default: () => 'CURRENT_TIMESTAMP' })
    // created_at: Date;
}
`, lowerName, cfg.Name, lowerName, lowerName, lowerName)
	}
	return fmt.Sprintf(`
export default class %s {
	constructor(
		public readonly id_%s: number | null
		// Add other database entity fields here
	) {}
}
`, cfg.Name, lowerName)

}
