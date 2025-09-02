package generate

import (
	"fmt"
	"strings"
)

// GetRouteTemplate génère le template des routes
func GetRouteTemplate(cfg EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	return fmt.Sprintf(`// Express importation
import express, {Request, Response} from "express";

// Layers importation
import { %sController }  from '../controllers/%s.controller'; 
import { %sService } from '../services/%s.service';
import { %sRepository } from "../repositories/%s.repository";


const router = express.Router();

const %sRepository = new %sRepository()
const %sService = new %sService(%sRepository)
const %sController = new %sController(%sService);

// GET /api/v1/admin/%s
router.get("/api/v1/admin/%s", async (req: Request, res: Response) => { 
    %sController.getAll(req, res)
});


export default router;
`,
		// Arguments for fmt.Sprintf, in order of appearance of %s
		cfg.Name, lowerName, // Controller import
		cfg.Name, lowerName, // Service import
		cfg.Name, lowerName, // Repository import
		lowerName, cfg.Name, // Repository variable declaration
		lowerName, cfg.Name, lowerName, // Service variable declaration with dependency injection
		lowerName, cfg.Name, lowerName, // Controller variable declaration with dependency injection
		lowerName+"s", // Route path
		lowerName+"s", // Route path
		lowerName,     // Controller call
	)
}
