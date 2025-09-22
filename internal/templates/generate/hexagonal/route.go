package hexagonal

import (
	"archi-ts-cli/internal/models"
	"fmt"
	"strings"
)

// GetRouteTemplate generates the controller template.
func GetRouteTemplate(cfg models.EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	upperName := strings.ToTitle(cfg.Name)

	return fmt.Sprintf(`import { Router } from 'express';
import { %sController } from '@controllers/%s.controller';
import { container } from '@config/dependencies/%s-containers';
import { %sTYPES } from '@config/dependencies/%s-types'; 

const %sRouter = Router();
const %sController = container.get<%sController>(%sTYPES.%sController);

%sRouter.get('/%ss/:id', (req, res) => %sController.getById(req, res));

export { %sRouter };`,
		cfg.Name, lowerName, lowerName, upperName, lowerName,
		lowerName, lowerName,
		cfg.Name, upperName, cfg.Name,
		lowerName, lowerName, lowerName, lowerName,
	)
}
