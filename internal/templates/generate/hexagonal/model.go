package hexagonal

import (
	"archi-ts-cli/internal/models"
	"fmt"
	"strings"
)

// GetUseCaseTemplate gégenerate the repository template
func GetModelTemplate(cfg models.EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	upperName := cfg.Name

	return fmt.Sprintf(`import { Entity, PrimaryColumn, Column } from 'typeorm';

@Entity('%ss')
export class %sModel {
    @PrimaryColumn('uuid')
    id!: string;


    // Exemples
    // @Column({ type: 'varchar', length: 100 })
    // name!: string;

    // @Column({ type: 'int' })
    // age!: number;
}`,
		lowerName,
		upperName)

}
