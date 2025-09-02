package generate

import (
	"fmt"
	"strings"
)

// GetRouteTemplate génère le template des routes

// GetRouteTemplate génère le template des routes
func GetRouteTemplate(cfg EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	return fmt.Sprintf(`import { Router } from 'express';
import { %sController } from '../controllers/%s.controller';

const router = Router();
const %sController = new %sController();

// GET /%s
router.get('/', %sController.getAll.bind(%sController));

// GET /%s/:id
router.get('/:id', %sController.getById.bind(%sController));

// POST /%s
router.post('/', %sController.create.bind(%sController));

// PUT /%s/:id
router.put('/:id', %sController.update.bind(%sController));

// DELETE /%s/:id
router.delete('/:id', %sController.delete.bind(%sController));

export { router as %sRoutes };
`, cfg.Name, lowerName, lowerName, cfg.Name, lowerName+"s", lowerName, lowerName, lowerName+"s",
		lowerName, lowerName, lowerName+"s", lowerName, lowerName, lowerName+"s",
		lowerName, lowerName, lowerName+"s", lowerName, lowerName, cfg.Name)
}
