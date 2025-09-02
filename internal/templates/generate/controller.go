package generate

import (
	"fmt"
	"strings"
)

// GetControllerTemplate génère le template du contrôleur
func GetControllerTemplate(cfg EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	return fmt.Sprintf(`
// Express importation
import { Request, Response } from 'express';

// Services importation
import { %sService } from '../services/%s.service';

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
   * 
   * @returns Promise<Response> - Express response object with appropriate status code and data
   * @throws {500} When an internal server error occurs during retrieval
   */
  public async getAll(req: Request, res: Response): Promise<Response> {
    try {
      const %s = await this.%sService.findAll();
      if (!%s) throw new Error("Empty result");
      return res.status(200).json({ success: true, data: %s });
    } catch (error) {
      return res.status(500).json(error instanceof Error ? error.message : "Internal server error");
    }
  }
}
`, cfg.Name, lowerName,
		cfg.Name, lowerName, cfg.Name, lowerName, cfg.Name, lowerName, lowerName,
		lowerName,
		lowerName+"s", lowerName, lowerName+"s", lowerName+"s")
}
