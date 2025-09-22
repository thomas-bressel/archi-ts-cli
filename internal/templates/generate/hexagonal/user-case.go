package hexagonal

import (
	"archi-ts-cli/internal/models"
	"fmt"
	"strings"
)

// GetUseCaseTemplate gégenerate the repository template
func GetUseCaseTemplate(cfg models.EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	upperName := cfg.Name
	titleName := strings.ToTitle(cfg.Name)

	return fmt.Sprintf(`import { injectable, inject } from 'inversify';
import { %sRepository } from '@ports/out/%s.repository';
import { %sResponse } from '@dtos/%s.dto';
import { %sID } from '@objects/%s-id.value-object';
import { %sNotFound } from '@exceptions/%s-not-found.exception';
import { %sTYPES } from '@config/dependencies/%s-types'; 

@injectable()
export class Get%sUseCase {
    constructor(
        @inject(%sTYPES.%sRepository)
        private readonly %sRepository: %sRepository
    ) {}

    async execute(%sId: string): Promise<%sResponse> {
        const %s = await this.%sRepository.findById(%sID.fromString(%sId));
        
        if (!%s) {
            throw new %sNotFound(%sId);
        }

        return {
            id: %s.id.value,
        };
    }
}`,
		upperName, lowerName,
		upperName, lowerName,
		upperName, lowerName,
		upperName, lowerName,
		titleName, lowerName,
		upperName, titleName, upperName, lowerName, upperName,
		lowerName, upperName, lowerName, lowerName,
		upperName, lowerName,
		lowerName, upperName, lowerName, lowerName,
	)

}
