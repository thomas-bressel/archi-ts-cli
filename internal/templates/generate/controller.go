package generate

import (
	"fmt"
	"strings"
)

// GetControllerTemplate generates the controller template.
func GetControllerTemplate(cfg EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	return fmt.Sprintf(`
// Express importation
import { Request, Response } from 'express';

// Services importation
import { %sService } from '@services/%s.service';

// Libraries importation
import { validate } from 'class-validator';
import { plainToInstance } from 'class-transformer';

export class %sController {
  private %sService: %sService;

  constructor(%sService: %sService) {
    this.%sService = %sService;
  }

  /**
   * Get the list of all %s
   * @param req 
   * @param res 
   * * @returns Promise<Response> - Express response object with appropriate status code and data
   * @throws {500} When an internal server error occurs during retrieval
   */
  public async getAll(req: Request, res: Response): Promise<Response> {
    try {
      const %s = await this.%sService.findAll();
  
      if (!%s || %s.length === 0) return res.status(200).json({ success: true, data: [] });
  
      return res.status(200).json({ success: true, data: %s });
    } catch (error) {
      return res.status(500).json(error instanceof Error ? error.message : "Internal server error");
    }
  }
}
`,
		// Arguments for fmt.Sprintf, in order of appearance of %s
		cfg.Name, lowerName, // Service import and path
		cfg.Name,            // Controller class name
		lowerName, cfg.Name, // private service declaration
		lowerName, cfg.Name, // constructor arguments
		lowerName, lowerName, // constructor body
		lowerName+"s",                // JSDoc comment for "all entities"
		lowerName+"s",                // `const entities` variable
		lowerName,                    // `this.service.findAll()` call
		lowerName+"s", lowerName+"s", // `if (!entities || entities.length === 0)`
		lowerName+"s", // `return res.status(200).json({..., data: entities })`
	)
}
