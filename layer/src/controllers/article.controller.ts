
// Express importation
import { Request, Response } from 'express';

// Services importation
import { ArticleService } from '../services/article.service';

// Libraries importation
import { validate } from 'class-validator';
import { plainToInstance } from 'class-transformer';

export class ArticleController {
  private articleService: ArticleService;

  constructor(articleService: ArticleService) {
    this.articleService = articleService;
  }


  /**
   * Get the list of all article
   * @param req 
   * @param res 
   * 
   * @returns Promise<Response> - Express response object with appropriate status code and data
   * @throws {500} When an internal server error occurs during user retrieval
   */
  public async getAll(req: Request, res: Response): Promise<Response> {
    try {
      const articles = await this.articleService.findAll();
      if (!articles) throw new Error("Résultat vide dans users");
      return res.status(200).json({ success: true, data: articles });
    } catch (error) {
      return res.status(500).json(error instanceof Error ? error.message : "Erreur interne du serveur");
    }
  }
}
