package hexagonal

import (
	"archi-ts-cli/internal/models"
	"fmt"
	"strings"
)

// GetControllerTemplate generates the controller template.
func GetControllerTemplate(cfg models.EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	upperName := cfg.Name
	titleName := strings.ToTitle(cfg.Name)

	return fmt.Sprintf(`import { injectable, inject } from 'inversify';
import { Request, Response } from 'express';
import { Get%sUseCase } from '@usecases/get-%s.use-case';
import { %sTYPES } from '@src/config/dependencies/%s-types'; 

@injectable()
export class %sController {
    constructor(
        @inject(%sTYPES.Get%sUseCase)
        private readonly get%sUseCase: Get%sUseCase
    ) {}

    async getById(req: Request, res: Response): Promise<void> {
        try {
            const %s = await this.get%sUseCase.execute(req.params.id);
            res.status(200).json(%s);
        } catch (error) {
            if (error instanceof Error) {
                const statusCode = error.name === '%sNotFound' ? 404 : 400;
                res.status(statusCode).json({ error: error.message });
            } else {
                res.status(500).json({ error: 'An unknown error occurred' });
            }
        }
    }
}`, upperName, lowerName, titleName, lowerName,
		upperName, titleName, upperName, upperName, upperName,
		lowerName, upperName, lowerName, upperName,
	)
}
