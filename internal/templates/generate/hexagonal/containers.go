package hexagonal

import (
	"archi-ts-cli/internal/models"
	"fmt"
	"strings"
)

// GetRouteTemplate generates the controller template.
func GetDependenciesContainersTemplate(cfg models.EntityConfig) string {
	upperName := cfg.Name
	lowerName := strings.ToLower(cfg.Name)
	titleName := strings.ToTitle(cfg.Name)

	return fmt.Sprintf(`import 'reflect-metadata';
import { Container } from 'inversify';
import { AppDataSource } from '@orm/data-source';
import { %sTYPES } from '@src/config/dependencies/%s-types'; 

// Interfaces
import { %sRepository } from '@ports/out/%s.repository';

// Implémentations  
import { Orm%sRepository } from '@repositories/orm-%s.repository';
import { Get%sUseCase } from '@usecases/get-%s.use-case';
import { %sController } from '@controllers/%s.controller';

// Model
import { %sModel } from '@models/%s.model';

const container = new Container();

// Configuration du Repository
container.bind<%sRepository>(%sTYPES.%sRepository).toConstantValue(
    new Orm%sRepository(AppDataSource.getRepository(%sModel))
);

// Configuration des Use Cases
container.bind<Get%sUseCase>(%sTYPES.Get%sUseCase).to(Get%sUseCase);

// Configuration des Controllers
container.bind<%sController>(%sTYPES.%sController).to(%sController);

export { container };`,
		titleName, lowerName,
		upperName, lowerName,
		upperName, lowerName, upperName, lowerName, upperName, lowerName,
		upperName, lowerName,
		upperName, titleName, upperName, upperName, upperName,
		upperName, titleName, upperName, upperName,
		upperName, titleName, upperName, upperName,
	)
}
