package hexagonal

import (
	"archi-ts-cli/internal/models"
	"fmt"
	"strings"
)

// GetRepositoryTemplate gégenerate the repository template
func GetOrmRepositoryTemplate(cfg models.EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	upperName := cfg.Name

	return fmt.Sprintf(`import { injectable } from 'inversify'; 
import { Repository } from 'typeorm';
import { %s } from '@entities/%s.entity';
import { %sID } from '@objects/%s-id.value-object';
import { %sRepository } from '@ports/out/%s.repository';
import { %sModel } from '@models/%s.model';

@injectable()
export class Orm%sRepository implements %sRepository {
    constructor(private readonly ormRepository: Repository<%sModel>) {}

    async save(%s: %s): Promise<void> {
        const model = new %sModel();
        model.id = %s.id.value;
        
        await this.ormRepository.save(model);
    }

    async findById(id: %sID): Promise<%s | null> {
        const model = await this.ormRepository.findOne({
            where: { id: id.value }
        });
        
        if (!model) return null;
        
        return %s.reconstitute(model.id);
    }

    // Exemples : 
    // async findAll(): Promise<%s[]> {
    //     const models = await this.ormRepository.find();
    //     return models.map(model => 
    //         %s.reconstitute(model.id)
    //     );
    // }

    // async update(%s: %s): Promise<void> {
    //     await this.ormRepository.update(%s.id.value, {
    //     });
    // }

    // async delete(id: %sID): Promise<void> {
    //     await this.ormRepository.delete({ id: id.value });
    // }

    // async exists(id: %sID): Promise<boolean> {
    //     const count = await this.ormRepository.count({
    //         where: { id: id.value }
    //     });
    //     return count > 0;
    // }
}`,
		upperName, lowerName, upperName, lowerName,
		upperName, lowerName, upperName, lowerName,
		upperName, upperName, upperName,
		lowerName, upperName, upperName, lowerName,
		upperName, upperName, upperName,
		upperName, upperName,
		lowerName, upperName, lowerName,
		upperName,
		upperName,
	)

}
