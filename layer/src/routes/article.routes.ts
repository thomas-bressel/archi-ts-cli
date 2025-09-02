import { Router } from 'express';
import { ArticleController } from '../controllers/article.controller';

const router = Router();
const articleController = new ArticleController();

// GET /articles
router.get('/', articleController.getAll.bind(articleController));

// GET /articles/:id
router.get('/:id', articleController.getById.bind(articleController));

// POST /articles
router.post('/', articleController.create.bind(articleController));

// PUT /articles/:id
router.put('/:id', articleController.update.bind(articleController));

// DELETE /articles/:id
router.delete('/:id', articleController.delete.bind(articleController));

export { router as ArticleRoutes };
